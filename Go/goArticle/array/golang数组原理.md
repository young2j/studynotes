

# 编译时数组类型解析

## ArrayType

数组是内存中一片连续的区域，在声明时需要指定长度，数组的声明有如下三种方式，`[...]`的方式在编译时会自动推断长度。

```go
var arr1 [3]int
var arr2 = [3]int{1,2,3}
arr3 := [...]int{1,2,3}
```

在词法及语法解析时，上述三种方式声明的数组会被解析为`ArrayType`, 当遇到`[...]`的声明时，其长度会被标记为`nil`，将在后续阶段进行自动推断。

```go
// go/src/cmd/compile/internal/syntax/parser.go

func (p *parser) typeOrNil() Expr {
  ...
	pos := p.pos()
	switch p.tok {
	...
	case _Lbrack:
		// '[' oexpr ']' ntype
		// '[' _DotDotDot ']' ntype
		p.next()
		if p.got(_Rbrack) {
			return p.sliceType(pos)
		}
		return p.arrayType(pos, nil)
  ...
}
  
// "[" has already been consumed, and pos is its position.
// If len != nil it is the already consumed array length.
func (p *parser) arrayType(pos Pos, len Expr) Expr {
	...
	if len == nil && !p.got(_DotDotDot) {
		p.xnest++
		len = p.expr()
		p.xnest--
	}
	...
	p.want(_Rbrack)
	t := new(ArrayType)
	t.pos = pos
	t.Len = len
	t.Elem = p.type_()
	return t
}
```

```go
// go/src/cmd/compile/internal/syntax/nodes.go

type (
  ...
	// [Len]Elem
	ArrayType struct {
		Len  Expr // nil means Len is ...
		Elem Expr
		expr
	}
  ...
)
```

## types2.Array

在对生成的表达式进行类型检查时，如果是`ArrayType`类型，且其长度`Len`为`nil`时，会初始化一个`types2.Array`并将其长度标记为`-1`，然后通过`check.indexedElts(e.ElemList, utyp.elem, utyp.len)`返回数组长度`n`并赋值给`Len`,完成自动推断。

```go
// go/src/cmd/compile/internal/types2/array.go

// An Array represents an array type.
type Array struct {
	len  int64
	elem Type
}
```

```go
// go/src/cmd/compile/internal/types2/expr.go

// exprInternal contains the core of type checking of expressions.
// Must only be called by rawExpr.
func (check *Checker) exprInternal(x *operand, e syntax.Expr, hint Type) exprKind {
	...
	switch e := e.(type) {
	...
	case *syntax.CompositeLit:
		var typ, base Type

		switch {
		case e.Type != nil:
			// composite literal type present - use it
			// [...]T array types may only appear with composite literals.
			// Check for them here so we don't have to handle ... in general.
			if atyp, _ := e.Type.(*syntax.ArrayType); atyp != nil && atyp.Len == nil {
				// We have an "open" [...]T array type.
				// Create a new ArrayType with unknown length (-1)
				// and finish setting it up after analyzing the literal.
				typ = &Array{len: -1, elem: check.varType(atyp.Elem)}
				base = typ
				break
			}
			typ = check.typ(e.Type)
			base = typ
      ...
		}

		switch utyp := coreType(base).(type) {
		...
		case *Array:
			if utyp.elem == nil {
				check.error(e, "illegal cycle in type declaration")
				goto Error
			}
			n := check.indexedElts(e.ElemList, utyp.elem, utyp.len)
			// If we have an array of unknown length (usually [...]T arrays, but also
			// arrays [n]T where n is invalid) set the length now that we know it and
			// record the type for the array (usually done by check.typ which is not
			// called for [...]T). We handle [...]T arrays and arrays with invalid
			// length the same here because it makes sense to "guess" the length for
			// the latter if we have a composite literal; e.g. for [n]int{1, 2, 3}
			// where n is invalid for some reason, it seems fair to assume it should
			// be 3 (see also Checked.arrayLength and issue #27346).
			if utyp.len < 0 {
				utyp.len = n
				// e.Type is missing if we have a composite literal element
				// that is itself a composite literal with omitted type. In
				// that case there is nothing to record (there is no type in
				// the source at that point).
				if e.Type != nil {
					check.recordTypeAndValue(e.Type, typexpr, utyp, nil)
				}
			}
		...
		}
	...
}
```

## types.Array

在生成中间结果时，`types2.Array`最终会通过`types.NewArray()`转换成`types.Array`类型。

```go
// go/src/cmd/compile/internal/noder/types.go

// typ0 converts a types2.Type to a types.Type, but doesn't do the caching check
// at the top level.
func (g *irgen) typ0(typ types2.Type) *types.Type {
	switch typ := typ.(type) {
	...
	case *types2.Array:
		return types.NewArray(g.typ1(typ.Elem()), typ.Len())
	...
}
```

```go
// go/src/cmd/compile/internal/types/type.go

// Array contains Type fields specific to array types.
type Array struct {
	Elem  *Type // element type
	Bound int64 // number of elements; <0 if unknown yet
}

// NewArray returns a new fixed-length array Type.
func NewArray(elem *Type, bound int64) *Type {
	if bound < 0 {
		base.Fatalf("NewArray: invalid bound %v", bound)
	}
	t := newType(TARRAY)
	t.extra = &Array{Elem: elem, Bound: bound}
	t.SetNotInHeap(elem.NotInHeap())
	if elem.HasTParam() {
		t.SetHasTParam(true)
	}
	if elem.HasShape() {
		t.SetHasShape(true)
	}
	return t
}
```

# 编译时数组字面量初始化

数组类型解析可以得到数组元素的类型`Elem`以及数组长度`Bound`,而数组字面量的初始化是在编译时类型检查阶段完成的，通过函数`tcComplit -> typecheckarraylit`循环字面量分别进行赋值。

```go
// go/src/cmd/compile/internal/typecheck/expr.go

func tcCompLit(n *ir.CompLitExpr) (res ir.Node) {
	...
	t := n.Type()
	base.AssertfAt(t != nil, n.Pos(), "missing type in composite literal")

	switch t.Kind() {
	...
	case types.TARRAY:
		typecheckarraylit(t.Elem(), t.NumElem(), n.List, "array literal")
		n.SetOp(ir.OARRAYLIT)
	...

	return n
}
```

```go
// go/src/cmd/compile/internal/typecheck/typecheck.go

// typecheckarraylit type-checks a sequence of slice/array literal elements.
func typecheckarraylit(elemType *types.Type, bound int64, elts []ir.Node, ctx string) int64 {
	...
	for i, elt := range elts {
		ir.SetPos(elt)
		r := elts[i]
		...
		r = Expr(r)
		r = AssignConv(r, elemType, ctx)
		...
}
```

# 编译时数组索引越界检查

在对数组进行索引访问时，如果访问越界在编译时就无法通过检查。例如:

```go
arr := [...]string{"s1", "s2", "s3"}
e3 := arr[3]

// invalid array index 3 (out of bounds for 3-element array)
```

数组在类型检查阶段会对访问数组的索引进行验证:

```go
// go/src/cmd/compile/internal/typecheck/typecheck.go
func typecheck1(n ir.Node, top int) ir.Node {
  ...
	switch n.Op() {
  ...
  case ir.OINDEX:
		n := n.(*ir.IndexExpr)
		return tcIndex(n)
  ...
  }
}


// go/src/cmd/compile/internal/typecheck/expr.go
func tcIndex(n *ir.IndexExpr) ir.Node {
	...
	l := n.X
	n.Index = Expr(n.Index)
	r := n.Index
	t := l.Type()
	...
	switch t.Kind() {
	...
	case types.TSTRING, types.TARRAY, types.TSLICE:
		n.Index = indexlit(n.Index)
		if t.IsString() {
			n.SetType(types.ByteType)
		} else {
			n.SetType(t.Elem())
		}
		why := "string"
		if t.IsArray() {
			why = "array"
		} else if t.IsSlice() {
			why = "slice"
		}

		if n.Index.Type() != nil && !n.Index.Type().IsInteger() {
			base.Errorf("non-integer %s index %v", why, n.Index)
			return n
		}

		if !n.Bounded() && ir.IsConst(n.Index, constant.Int) {
			x := n.Index.Val()
			if constant.Sign(x) < 0 {
				base.Errorf("invalid %s index %v (index must be non-negative)", why, n.Index)
			} else if t.IsArray() && constant.Compare(x, token.GEQ, constant.MakeInt64(t.NumElem())) {
				base.Errorf("invalid array index %v (out of bounds for %d-element array)", n.Index, t.NumElem())
			} else if ir.IsConst(n.X, constant.String) && constant.Compare(x, token.GEQ, constant.MakeInt64(int64(len(ir.StringVal(n.X))))) {
				base.Errorf("invalid string index %v (out of bounds for %d-byte string)", n.Index, len(ir.StringVal(n.X)))
			} else if ir.ConstOverflow(x, types.Types[types.TINT]) {
				base.Errorf("invalid %s index %v (index too large)", why, n.Index)
			}
		}
	...
	}
	return n
}
```

# 运行时数组内存分配

数组是内存区域一块连续的存储空间。在运行时会通过`mallocgc`给数组分配具体的存储空间。`newarray`中如果数组元素刚好只有一个，则空间大小为元素类型的大小`typ.size`， 如果有多个元素则内存大小为`n*typ.size`。但这并不是实际分配的内存大小，实际分配多少内存，取决于`mallocgc`，涉及到`golang`的内存分配原理。但可以看到如果待分配的对象不超过`32kb`,`mallocgc`会直接将其分配在缓存空间中，如果大于`32kb`则直接从堆区分配内存空间。

```go
// go/src/runtime/malloc.go

// newarray allocates an array of n elements of type typ.
func newarray(typ *_type, n int) unsafe.Pointer {
	if n == 1 {
		return mallocgc(typ.size, typ, true)
	}
	mem, overflow := math.MulUintptr(typ.size, uintptr(n))
	if overflow || mem > maxAlloc || n < 0 {
		panic(plainError("runtime: allocation size out of range"))
	}
	return mallocgc(mem, typ, true)
}

// Allocate an object of size bytes.
// Small objects are allocated from the per-P cache's free lists.
// Large objects (> 32 kB) are allocated straight from the heap.
func mallocgc(size uintptr, typ *_type, needzero bool) unsafe.Pointer {
	...
}

```



# 总结

* 数组在编译阶段最终被解析为`types.Array`类型，包含元素类型`Elem`和数组长度`Bound`

  ```go
  type Array struct {
  	Elem  *Type // element type
  	Bound int64 // number of elements; <0 if unknown yet
  }
  ```

* 如果数组长度未指定，例如使用了语法糖`[...]`，则会在表达式类型检查时计算出数组长度。

* 数组字面量初始化以及索引越界检查都是在编译时类型检查阶段完成的。

* 在运行时通过`newarray()`函数对数组内存进行分配，如果数组大小超过`32kb`则会直接分配到堆区内存。