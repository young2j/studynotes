# 切片的解析

当我们的代码敲下`[]`时，便会被go编译器解析为抽象语法树上的切片节点, 被初始化为切片表达式`SliceType`:

```go
// go/src/cmd/compile/internal/syntax/parser.go

// TypeSpec = identifier [ TypeParams ] [ "=" ] Type .
func (p *parser) typeDecl(group *Group) Decl {
	...
	if p.tok == _Lbrack {
		// d.Name "[" ...
		// array/slice type or type parameter list
		pos := p.pos()
		p.next()
		switch p.tok {
		...
		case _Rbrack:
			// d.Name "[" "]" ...
			p.next()
			d.Type = p.sliceType(pos)
		...
		}
	} 
	...
}

func (p *parser) sliceType(pos Pos) Expr {
	t := new(SliceType)
	t.pos = pos
	t.Elem = p.type_()
	return t
}

// go/src/cmd/compile/internal/syntax/nodes.go
type (
	...
  // []Elem
	SliceType struct {
		Elem Expr
		expr
	}
  ...
)
```

编译时切片定义为`Slice`结构体，属性只包含同一类型的元素`Elem`，编译时通过`NewSlice()`函数进行创建：

```go
// go/src/cmd/compile/internal/types/type.go

type Slice struct {
	Elem *Type // element type
}


func NewSlice(elem *Type) *Type {
	if t := elem.cache.slice; t != nil {
		if t.Elem() != elem {
			base.Fatalf("elem mismatch")
		}
		if elem.HasTParam() != t.HasTParam() || elem.HasShape() != t.HasShape() {
			base.Fatalf("Incorrect HasTParam/HasShape flag for cached slice type")
		}
		return t
	}

	t := newType(TSLICE)
	t.extra = Slice{Elem: elem}
	elem.cache.slice = t
	if elem.HasTParam() {
		t.SetHasTParam(true)
	}
	if elem.HasShape() {
		t.SetHasShape(true)
	}
	return t
}
```

# 切片的初始化

切片有两种初始化方式，一种声明即初始化称为字面量初始化，一种称为`make`初始化，例如:

```go
litSlic := []int{1,2,3,4}  // 字面量初始化
makeSlic := make([]int,0)  // make初始化
```

## 字面量初始化

切片字面量的初始化是在生成抽象语法树后进行遍历的`walk`阶段完成的。通过`walkComplit`方法，首先会进行类型检查，此时会计算出切片元素的个数`length`，然后通过`slicelit`方法完成具体的初始化工作。整个过程会先创建一个数组存储于静态区`(static array)`，并在堆区创建一个新的切片`(auto array)`，然后将静态区的数据复制到堆区`(copy the static array to the auto array)`，对于切片中的元素会按索引位置一个一个的进行赋值。 在程序启动时这一过程会加快切片的初始化。

```go
// go/src/cmd/compile/internal/walk/complit.go

// walkCompLit walks a composite literal node:
// OARRAYLIT, OSLICELIT, OMAPLIT, OSTRUCTLIT (all CompLitExpr), or OPTRLIT (AddrExpr).
func walkCompLit(n ir.Node, init *ir.Nodes) ir.Node {
	if isStaticCompositeLiteral(n) && !ssagen.TypeOK(n.Type()) {
		n := n.(*ir.CompLitExpr) // not OPTRLIT
		// n can be directly represented in the read-only data section.
		// Make direct reference to the static data. See issue 12841.
		vstat := readonlystaticname(n.Type())
		fixedlit(inInitFunction, initKindStatic, n, vstat, init)
		return typecheck.Expr(vstat)
	}
	var_ := typecheck.Temp(n.Type())
	anylit(n, var_, init)
	return var_
}
```

类型检查时，计算出切片长度的过程为:

```go
// go/src/cmd/compile/internal/typecheck/expr.go

func tcCompLit(n *ir.CompLitExpr) (res ir.Node) {
	...
	t := n.Type()
	base.AssertfAt(t != nil, n.Pos(), "missing type in composite literal")

	switch t.Kind() {
	...
	case types.TSLICE:
		length := typecheckarraylit(t.Elem(), -1, n.List, "slice literal")
		n.SetOp(ir.OSLICELIT)
		n.Len = length
	...
  }

	return n
}
```

切片的具体初始化过程为:

1. 在静态存储区创建一个数组；

2. 将数组赋值给一个常量部分；

3. 创建一个自动指针即切片分配到堆区，并指向数组；

4. 将数组中的数据从静态区拷贝到切片的堆区；

5. 对每一个切片元素按索引位置分别进行赋值；

6. 最后将分配到堆区的切片赋值给定义的变量；

   > 源代码通过注释也写明了整个过程。

```go
// go/src/cmd/compile/internal/walk/complit.go

func anylit(n ir.Node, var_ ir.Node, init *ir.Nodes) {
	t := n.Type()
	switch n.Op() {
  ...
  case ir.OSLICELIT:
		n := n.(*ir.CompLitExpr)
		slicelit(inInitFunction, n, var_, init)
  ...
  }
}

func slicelit(ctxt initContext, n *ir.CompLitExpr, var_ ir.Node, init *ir.Nodes) {
	// make an array type corresponding the number of elements we have
	t := types.NewArray(n.Type().Elem(), n.Len)
	types.CalcSize(t)

	if ctxt == inNonInitFunction {
		// put everything into static array
		vstat := staticinit.StaticName(t)

		fixedlit(ctxt, initKindStatic, n, vstat, init)
		fixedlit(ctxt, initKindDynamic, n, vstat, init)

		// copy static to slice
		var_ = typecheck.AssignExpr(var_)
		name, offset, ok := staticinit.StaticLoc(var_)
		if !ok || name.Class != ir.PEXTERN {
			base.Fatalf("slicelit: %v", var_)
		}
		staticdata.InitSlice(name, offset, vstat.Linksym(), t.NumElem())
		return
	}

	// recipe for var = []t{...}
	// 1. make a static array
	//	var vstat [...]t
	// 2. assign (data statements) the constant part
	//	vstat = constpart{}
	// 3. make an auto pointer to array and allocate heap to it
	//	var vauto *[...]t = new([...]t)
	// 4. copy the static array to the auto array
	//	*vauto = vstat
	// 5. for each dynamic part assign to the array
	//	vauto[i] = dynamic part
	// 6. assign slice of allocated heap to var
	//	var = vauto[:]
	//
	// an optimization is done if there is no constant part
	//	3. var vauto *[...]t = new([...]t)
	//	5. vauto[i] = dynamic part
	//	6. var = vauto[:]

	// if the literal contains constants,
	// make static initialized array (1),(2)
	var vstat ir.Node

	mode := getdyn(n, true)
	if mode&initConst != 0 && !isSmallSliceLit(n) {
		if ctxt == inInitFunction {
			vstat = readonlystaticname(t)
		} else {
			vstat = staticinit.StaticName(t)
		}
		fixedlit(ctxt, initKindStatic, n, vstat, init)
	}

	// make new auto *array (3 declare)
	vauto := typecheck.Temp(types.NewPtr(t))

	// set auto to point at new temp or heap (3 assign)
	var a ir.Node
	if x := n.Prealloc; x != nil {
		// temp allocated during order.go for dddarg
		if !types.Identical(t, x.Type()) {
			panic("dotdotdot base type does not match order's assigned type")
		}
		a = initStackTemp(init, x, vstat)
	} else if n.Esc() == ir.EscNone {
		a = initStackTemp(init, typecheck.Temp(t), vstat)
	} else {
		a = ir.NewUnaryExpr(base.Pos, ir.ONEW, ir.TypeNode(t))
	}
	appendWalkStmt(init, ir.NewAssignStmt(base.Pos, vauto, a))

	if vstat != nil && n.Prealloc == nil && n.Esc() != ir.EscNone {
		// If we allocated on the heap with ONEW, copy the static to the
		// heap (4). We skip this for stack temporaries, because
		// initStackTemp already handled the copy.
		a = ir.NewStarExpr(base.Pos, vauto)
		appendWalkStmt(init, ir.NewAssignStmt(base.Pos, a, vstat))
	}

	// put dynamics into array (5)
	var index int64
	for _, value := range n.List {
		if value.Op() == ir.OKEY {
			kv := value.(*ir.KeyExpr)
			index = typecheck.IndexConst(kv.Key)
			if index < 0 {
				base.Fatalf("slicelit: invalid index %v", kv.Key)
			}
			value = kv.Value
		}
		a := ir.NewIndexExpr(base.Pos, vauto, ir.NewInt(index))
		a.SetBounded(true)
		index++

		// TODO need to check bounds?

		switch value.Op() {
		case ir.OSLICELIT:
			break

		case ir.OARRAYLIT, ir.OSTRUCTLIT:
			value := value.(*ir.CompLitExpr)
			k := initKindDynamic
			if vstat == nil {
				// Generate both static and dynamic initializations.
				// See issue #31987.
				k = initKindLocalCode
			}
			fixedlit(ctxt, k, value, a, init)
			continue
		}

		if vstat != nil && ir.IsConstNode(value) { // already set by copy from static value
			continue
		}

		// build list of vauto[c] = expr
		ir.SetPos(value)
		as := ir.NewAssignStmt(base.Pos, a, value)
		appendWalkStmt(init, orderStmtInPlace(typecheck.Stmt(as), map[string][]*ir.Name{}))
	}

	// make slice out of heap (6)
	a = ir.NewAssignStmt(base.Pos, var_, ir.NewSliceExpr(base.Pos, ir.OSLICE, vauto, nil, nil, nil))
	appendWalkStmt(init, orderStmtInPlace(typecheck.Stmt(a), map[string][]*ir.Name{}))
}
```



## make初始化

当使用`make`初始化一个切片时，会被编译器解析为一个`OMAKESLICE`操作:

```go
// go/src/cmd/compile/internal/walk/expr.go

func walkExpr1(n ir.Node, init *ir.Nodes) ir.Node {
	switch n.Op() {
	...
	case ir.OMAKESLICE:
		n := n.(*ir.MakeExpr)
		return walkMakeSlice(n, init)
	...
}
```

如果`make`初始化一个较大的切片则会逃逸到堆中，如果分配了一个较小的切片则直接在栈中分配。

* 在`walkMakeSlice`函数中，如果未指定切片的容量`Cap`，则初始容量等于切片的长度。
* 如果切片的初始化未发生内存逃逸`n.Esc() == ir.EscNone`，则会先在内存中创建一个同样容量大小的数组`NewArray()`, 然后按切片长度将数组中的值`arr[:l]`赋予切片。
* 如果发生了内存逃逸，切片会调用运行时函数`makeslice`和`makeslice64`在堆中完成对切片的初始化。

```go
// go/src/cmd/compile/internal/walk/builtin.go

func walkMakeSlice(n *ir.MakeExpr, init *ir.Nodes) ir.Node {
	l := n.Len
	r := n.Cap
	if r == nil {
		r = safeExpr(l, init)
		l = r
	}
	...
	if n.Esc() == ir.EscNone {
		if why := escape.HeapAllocReason(n); why != "" {
			base.Fatalf("%v has EscNone, but %v", n, why)
		}
		// var arr [r]T
		// n = arr[:l]
		i := typecheck.IndexConst(r)
		if i < 0 {
			base.Fatalf("walkExpr: invalid index %v", r)
		}
		...
		t = types.NewArray(t.Elem(), i) // [r]T
		var_ := typecheck.Temp(t)
		appendWalkStmt(init, ir.NewAssignStmt(base.Pos, var_, nil))  // zero temp
		r := ir.NewSliceExpr(base.Pos, ir.OSLICE, var_, nil, l, nil) // arr[:l]
		// The conv is necessary in case n.Type is named.
		return walkExpr(typecheck.Expr(typecheck.Conv(r, n.Type())), init)
	}

	// n escapes; set up a call to makeslice.
	// When len and cap can fit into int, use makeslice instead of
	// makeslice64, which is faster and shorter on 32 bit platforms.

	len, cap := l, r

	fnname := "makeslice64"
	argtype := types.Types[types.TINT64]

	// Type checking guarantees that TIDEAL len/cap are positive and fit in an int.
	// The case of len or cap overflow when converting TUINT or TUINTPTR to TINT
	// will be handled by the negative range checks in makeslice during runtime.
	if (len.Type().IsKind(types.TIDEAL) || len.Type().Size() <= types.Types[types.TUINT].Size()) &&
		(cap.Type().IsKind(types.TIDEAL) || cap.Type().Size() <= types.Types[types.TUINT].Size()) {
		fnname = "makeslice"
		argtype = types.Types[types.TINT]
	}
	fn := typecheck.LookupRuntime(fnname)
	ptr := mkcall1(fn, types.Types[types.TUNSAFEPTR], init, reflectdata.TypePtr(t.Elem()), typecheck.Conv(len, argtype), typecheck.Conv(cap, argtype))
	ptr.MarkNonNil()
	len = typecheck.Conv(len, types.Types[types.TINT])
	cap = typecheck.Conv(cap, types.Types[types.TINT])
	sh := ir.NewSliceHeaderExpr(base.Pos, t, ptr, len, cap)
	return walkExpr(typecheck.Expr(sh), init)
}

```

切片在栈中初始化还是在堆中初始化，存在一个临界值进行判断。临界值`maxImplicitStackVarSize`默认为64kb。从下面的源代码可以看到，显式变量声明`explicit variable declarations `和隐式变量`implicit variables`逃逸的临界值并不一样。

* 当我们使用`var变量声明`以及`:=赋值操作`时，内存逃逸的临界值为`10M`, 小于该值的对象会分配在栈中。

* 当我们使用如下操作时，内存逃逸的临界值为`64kb`，小于该值的对象会分配在栈中。

  ```go
  p := new(T)          
  p := &T{}           
  s := make([]T, n)    
  s := []byte("...") 
  ```

```go
// go/src/cmd/compile/internal/ir/cfg.go

var (
	// maximum size variable which we will allocate on the stack.
	// This limit is for explicit variable declarations like "var x T" or "x := ...".
	// Note: the flag smallframes can update this value.
	MaxStackVarSize = int64(10 * 1024 * 1024)

	// maximum size of implicit variables that we will allocate on the stack.
	//   p := new(T)          allocating T on the stack
	//   p := &T{}            allocating T on the stack
	//   s := make([]T, n)    allocating [n]T on the stack
	//   s := []byte("...")   allocating [n]byte on the stack
	// Note: the flag smallframes can update this value.
	MaxImplicitStackVarSize = int64(64 * 1024)

	// MaxSmallArraySize is the maximum size of an array which is considered small.
	// Small arrays will be initialized directly with a sequence of constant stores.
	// Large arrays will be initialized by copying from a static temp.
	// 256 bytes was chosen to minimize generated code + statictmp size.
	MaxSmallArraySize = int64(256)
)
```

切片的make初始化就属于`s := make([]T, n)`操作，当切片元素分配的内存大小大于`64kb`时, 切片会逃逸到堆中进行初始化。此时会调用运行时函数`makeslice`来完成这一个过程:

```go
// go/src/runtime/slice.go

func makeslice(et *_type, len, cap int) unsafe.Pointer {
	mem, overflow := math.MulUintptr(et.size, uintptr(cap))
	if overflow || mem > maxAlloc || len < 0 || len > cap {
		// NOTE: Produce a 'len out of range' error instead of a
		// 'cap out of range' error when someone does make([]T, bignumber).
		// 'cap out of range' is true too, but since the cap is only being
		// supplied implicitly, saying len is clearer.
		// See golang.org/issue/4085.
		mem, overflow := math.MulUintptr(et.size, uintptr(len))
		if overflow || mem > maxAlloc || len < 0 {
			panicmakeslicelen()
		}
		panicmakeslicecap()
	}

	return mallocgc(mem, et, true)
}
```

根据切片的运行时结构定义，运行时切片结构底层维护着切片的长度`len`、容量`cap`以及指向数组数据的指针`array`:

```go
// go/src/runtime/slice.go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

// 或者
// go/src/reflect/value.go
// SliceHeader is the runtime representation of a slice.
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
```

# 切片的截取

从切片的运行时结构已经知道，切片底层数据是一个数组，切片本身只是持有一个指向改数组数据的指针。因此，当我们对切片进行截取操作时，新的切片仍然指向原切片的底层数据，当对原切片数据进行更新时，意味着新切片相同索引位置的数据也发生了变化:

```go
slic := []int{1, 2, 3, 4, 5}
slic1 := slic[:2]
fmt.Printf("slic1: %v\n", slic1)

slic[0] = 0

fmt.Printf("slic: %v\n", slic)
fmt.Printf("slic1: %v\n", slic1)

// slic1: [1 2]
// slic: [0 2 3 4 5]
// slic1: [0 2]
```

切片截取后，虽然底层数据没有发生变化，但指向的数据范围发生了变化，表现为截取后的切片长度、容量会相应发生变化:

* 长度为截取的范围
* 容量为截取起始位置到原切片末尾的范围

```go
slic := []int{1, 2, 3, 4, 5}
slic1 := slic[:2]
slic2 := slic[2:]

fmt.Printf("len(slic): %v\n", len(slic))
fmt.Printf("cap(slic): %v\n", cap(slic))

fmt.Printf("len(slic1): %v\n", len(slic1))
fmt.Printf("cap(slic1): %v\n", cap(slic1))

fmt.Printf("len(slic2): %v\n", len(slic2))
fmt.Printf("cap(slic2): %v\n", cap(slic2))

// len(slic): 5
// cap(slic): 5

// len(slic1): 2
// cap(slic1): 5

// len(slic2): 3
// cap(slic2): 3
```

所以，切片截取变化的是底层data指针、长度以及容量，data指针指向的数组数据本身没有变化。切片的赋值拷贝就等价于于全切片，底层`data`指针仍然指向相同的数组地址，长度和容量保持不变:

```go
slic := []int{1, 2, 3, 4, 5}
s := slic  // 等价于 s := slic[:]
```

当切片作为参数传递时，即使切片中包含大量的数据，也只是切片数据地址的拷贝，拷贝的成本是较低的。

# 切片的复制

当我们想要完整拷贝一个切片时，可以使用内置的`copy`函数，效果类似于"深拷贝"。

```go
slic := []int{1, 2, 3, 4, 5}
var slic1 []int
copy(slic1, slic)

fmt.Printf("slic: %p\n", slic)
fmt.Printf("slic1: %p\n", slic1)

// slic: 0xc0000aa030
// slic1: 0x0
```

完整复制后，新的切片指向了新的内存地址。切片的复制在运行时会调用`slicecopy()`函数，通过`memmove`移动数据到新的内存地址:

```go
// go/src/runtime/slice.go

func slicecopy(toPtr unsafe.Pointer, toLen int, fromPtr unsafe.Pointer, fromLen int, width uintptr) int {
	if fromLen == 0 || toLen == 0 {
		return 0
	}

	n := fromLen
	if toLen < n {
		n = toLen
	}
	...
	if size == 1 { // common case worth about 2x to do here
		// TODO: is this still worth it with new memmove impl?
		*(*byte)(toPtr) = *(*byte)(fromPtr) // known to be a byte pointer
	} else {
		memmove(toPtr, fromPtr, size)
	}
	return n
}
```

# 切片的扩容

切片元素个数可以动态变化，切片初始化后会确定一个初始化容量，当容量不足时会在运行时通过`growslice`进行扩容:

```go

func growslice(et *_type, old slice, cap int) slice {
	...
	newcap := old.cap
	doublecap := newcap + newcap
	if cap > doublecap {
		newcap = cap
	} else {
		const threshold = 256
		if old.cap < threshold {
			newcap = doublecap
		} else {
			// Check 0 < newcap to detect overflow
			// and prevent an infinite loop.
			for 0 < newcap && newcap < cap {
				// Transition from growing 2x for small slices
				// to growing 1.25x for large slices. This formula
				// gives a smooth-ish transition between the two.
				newcap += (newcap + 3*threshold) / 4
			}
			// Set newcap to the requested cap when
			// the newcap calculation overflowed.
			if newcap <= 0 {
				newcap = cap
			}
		}
	}
	...
	memmove(p, old.array, lenmem)

	return slice{p, old.len, newcap}
}
```

从`growslice`的代码可以看出：

* 当新申请的容量(`cap`)大于二倍旧容量(`old.cap`)时，最终容量(`newcap`)是新申请的容量；
* 当新申请的容量(`cap`)小于二倍旧容量(`old.cap`)时，

  * 如果旧容量小于256，最终容量为旧容量的2倍；

  * 如果旧容量大于等于256，则会按照公式`newcap += (newcap + 3*threshold) / 4`来确定最终容量。实际的表现为:

    - 当切片长度小于等于1024时，最终容量是旧容量的2倍；

    - 当切片长度大于1024时，最终容量是旧容量的1.25倍，随着长度的增长，大于1.25倍；
* 扩容后，会通过`memmove()`函数将旧的数组移动到新的地址，因此扩容后新的切片一般和原来的地址不同。

示例:

```go
var slic []int
oldCap := cap(slic)
for i := 0; i < 2048; i++ {
  slic = append(slic, i)
  newCap := cap(slic)
  grow := float32(newCap) / float32(oldCap)
  if newCap != oldCap {
    fmt.Printf("len(slic):%v cap(slic):%v grow:%v %p\n", len(slic), cap(slic), grow, slic)
  }

  oldCap = newCap
}

// len(slic):1     cap(slic):1     grow:+Inf       0xc0000140c0
// len(slic):2     cap(slic):2     grow:2          0xc0000140e0
// len(slic):3     cap(slic):4     grow:2          0xc000020100
// len(slic):5     cap(slic):8     grow:2          0xc00001e340
// len(slic):9     cap(slic):16    grow:2          0xc000026080
// len(slic):17    cap(slic):32    grow:2          0xc00007e000
// len(slic):33    cap(slic):64    grow:2          0xc000100000
// len(slic):65    cap(slic):128   grow:2          0xc000102000
// len(slic):129   cap(slic):256   grow:2          0xc000104000
// len(slic):257   cap(slic):512   grow:2          0xc000106000
// len(slic):513   cap(slic):1024  grow:2          0xc000108000
// len(slic):1025  cap(slic):1280  grow:1.25       0xc00010a000
// len(slic):1281  cap(slic):1696  grow:1.325      0xc000114000
// len(slic):1697  cap(slic):2304  grow:1.3584906  0xc00011e000
```

# 总结

* 切片在编译时定义为`Slice`结构体，并通过`NewSlice()`函数进行创建；

  ```go
  type Slice struct {
  	Elem *Type // element type
  }
  ```

* 切片的运行时定义为`slice`结构体, 底层维护着指向数组数据的指针，切片长度以及容量；

  ```go
  type slice struct {
  	array unsafe.Pointer
  	len   int
  	cap   int
  }
  ```

* 切片字面量初始化时，会在编译时的类型检查阶段计算出切片的长度，然后在walk遍历语法树时创建底层数组，并将切片中的每个字面量元素按索引赋值给数组，切片的数据指针指向该数组；

* 切片make初始化时，会调用运行时`makeslice`函数进行内存分配，当内存占用大于64kb时会逃逸到堆中；

* 切片截取后，底层数组数据没有发生变化，但指向的数据范围发生了变化，表现为截取后的切片长度、容量会相应发生变化:

  * 长度为截取的范围
  * 容量为截取起始位置到原切片末尾的范围

* 使用`copy`复制切片时，会在运行时会调用`slicecopy()`函数，通过`memmove`移动数据到了新的内存地址；

* 切片扩容是通过运行时`growslice`函数完成的，一般表现为：

  * 当切片长度小于等于1024时，最终容量是旧容量的2倍；

  - 当切片长度大于1024时，最终容量是旧容量的1.25倍，并随着长度的增长，缓慢大于1.25倍；

  * 扩容时会通过`memmove()`函数将旧的数组移动到新的地址，因此扩容后地址会发生变化。

