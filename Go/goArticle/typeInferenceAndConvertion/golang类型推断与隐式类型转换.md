# 前言

golang类型推断可以省略类型，像写动态语言代码一样，让编程变得更加简单，同时也保留了静态类型的安全性。
类型推断往往也伴随着类型的隐式转换，二者均与golang的编译器有关。在了解了golang的类型推断与隐式类型转换原理后，将对如下问题信手拈来——下述代码能通过编译吗？b的值是什么类型？

```go
// eg.1
a := 1.1
b := 1 + a

// eg.2
a := 1
b := 1.1 + a

// eg.3
a1 := 1
a2 := 1.1
b := a1 + a2

// eg.4
const b = 3 * 0.333

// eg.5
const a int = 1.0
const b = a * 0.333

// eg.6
const a = 1.0/3
b := &a
```

要弄清楚上述示例，在了解变量类型推断之前，最好先了解常量的隐式类型转换。

# 常量的隐式类型转换

## 常量的声明

* 未命名常量只在编译期间存在，不会存储在内存中；而命名常量存在于内存静态区，不允许修改。

  考虑如下代码:

  ```go
  const k = 5
  ```

  `5`就是未命名常量，而`k`即为命名常量，当编译后`k`的值为5，而等号右边的5不再存在。

* 常量不允许取址。

  ```go
  const k = 5
  addr := &k
  // invalid operation: cannot take address of k (untyped int constant 5)
  ```

## 常量的类型转换

* 兼容的类型可以进行隐式转换。例如

  ```go
  const c int = 123
  const c int = 123.0
  const c int = 123.1 // cannot use 123.1 (untyped float constant) as int value in constant declaration (truncated)
  
  const c float64 = 123.0
  const c float64 = 123
  ```

* 运算中的隐式转换

  * 除位运算、未定义常量外，运算符两边的操作数类型必须相同

  * 如果运算符两边是不同类型的未定义常量(untyped constant)，则隐式转换的优先级为: 

    整数(int) <符文数(rune)<浮点数(float)<复数(Imag)

   例如

  ```go
  const c = 1/2    // 1和2类型相同，无隐式转换发生
  const c = 1/2.0  // 整数优先转换为浮点数1.0， c的结果为0.5(float64)
  
  const a int = 1
  const c = a * 1.1 // *左边的a是已定义类型的常量，因此1.1将被转换为int，但浮点数1.1与整形不兼容，无法进行转换，因此编译器会报错
  //  (untyped float constant) truncated to int 
  ```

基于上述说明，前言中的示例4、5、6均可迎刃而解。

## 隐式转换的原理

常量隐式转换的统一在编译时的类型检查阶段完成。通过`defaultlit2`函数进行处理。其中，`l和r`分别代表运算符左右两边的节点。

```go
// go/src/cmd/compile/internal/typecheck/const.go

func defaultlit2(l ir.Node, r ir.Node, force bool) (ir.Node, ir.Node) {
	if l.Type() == nil || r.Type() == nil {
		return l, r
	}

	if !l.Type().IsInterface() && !r.Type().IsInterface() {
		// Can't mix bool with non-bool, string with non-string.
		if l.Type().IsBoolean() != r.Type().IsBoolean() {
			return l, r
		}
		if l.Type().IsString() != r.Type().IsString() {
			return l, r
		}
	}

	if !l.Type().IsUntyped() {
		r = convlit(r, l.Type())
		return l, r
	}

	if !r.Type().IsUntyped() {
		l = convlit(l, r.Type())
		return l, r
	}

	if !force {
		return l, r
	}

	// Can't mix nil with anything untyped.
	if ir.IsNil(l) || ir.IsNil(r) {
		return l, r
	}
	t := defaultType(mixUntyped(l.Type(), r.Type()))
	l = convlit(l, t)
	r = convlit(r, t)
	return l, r
}
```

从源代码中可以看到，如果左右两边均不是接口类型，那么：

* `bool`型不能与非`bool`型进行转换，即

  ```go
  c := true + 12 // 错误
  ```

* `string`型不能与非`string`型进行转换, 即

  ```go
  c := "123" + 12 // 错误
  ```

* `nil`不能与任意未定义类型进行转换，即

  ```go
  c := nil + 12 // 错误
  ```

* 如果操作符左边的节点有定义类型，则将操作符右边的节点转换为左边节点的类型，即

  ```go
  const a int = 1
  const b int = 1.0
  
  const c = a + 1.0 // 1.0转换为a的类型int
  const c = a + b // b的类型已经在前面转换为int
  ```

* 如果操作符左边的节点为未命名常量，而右边的节点有定义类型，则将左边节点的类型转换为右边节点的类型，即

  ```go
  const a int = 1
  
  const c = 1.0 + a // 1.0转换为a的类型int
  ```

  综上所述，可以得出:

  任何时候，**已定义类型的常量都不会发生类型转换**。换言之，**编译器不允许对变量标识符引用的值进行强制类型转换**。即无关优先级，下述c=xx代码中的`a、b`均不会发生类型转换，只能是为定义类型的常量`1.0`转换为`a、b`的类型。

  ```go
  const a int = 1
  const b int = 1.0
  
  const c = a + 1.0
  const c = a + b 
  const c = 1.0 + b
  ```

# 变量的类型推断

golang使用特殊的操作符":="用于变量的类型推断，且其只能作用于函数或方法体内部。

> 操作符":="在《go语言实战》中有个名字叫“短变量声明操作符”

初识go语言的人总是会有疑问，下面三个语句有啥差别:

```go
a := 123 
var a = 123
var a int = 123.0
```

从结果上来说，上述三个语句是等效的。但编译阶段的执行细节是不同的。

## 类型推断的原理

编译器的执行过程为：`词法(token)解析->语法(syntax)分析->抽象语法树(ast)构建->类型检查->生成中间代码->代码优化->生成机器码`。

类型推断发生于前四个阶段，即`词法(token)解析->语法(syntax)分析->抽象语法树(ast)构建->类型检查`。以`a := 123`为例:

* 在词法解析阶段, 会将赋值语句右边的常量`123`解析为一个未定义的类型，称为未定义常量。编译器会逐个读取该常量的UTF-8字符，首个字符为"的则为字符串，首个字符为'0'-'9'的则为数字, 数字中包含"."号的则为浮点数。

  ```go
  // go/src/cmd/compile/internal/syntax/scanner.go
  func (s *scanner) next() {
    ...
    switch s.ch {
  	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
  		s.number(false)
  
  	case '"':
  		s.stdString()
  
  	case '`':
  		s.rawString()
  
  	case '\'':
  		s.rune()
      ...
    }
  }
  ```

* 在语法分析阶段，会对解析的词进行具体的语法分析。例如上述`s.number(false)`就是依次扫描`123`三个符文(rune)然后按照无小数点的数字来做具体分析。

  当无小数点符号`.`时，如果首字符为'0', 则扫描下一位字符，'x'、'o'、'b'分别代表我们写的代码表示的是十六进制、八进制及二进制数字。当首字符不是'0'时，每一位字符均作为十进制数字进行处理。

  当有小数点时(`seenPoint=true`)，每一位字符均作为十进制浮点数字面类型处理(`FloatLit`)

  ```go
  // go/src/cmd/compile/internal/syntax/scanner.go
  func (s *scanner) number(seenPoint bool) {
    ...
    base := 10        // number base
    ...
  	// integer part
  	if !seenPoint {
  		if s.ch == '0' {
  			s.nextch()
  			switch lower(s.ch) {
  			case 'x':
  				s.nextch()
  				base, prefix = 16, 'x'
  			case 'o':
  				s.nextch()
  				base, prefix = 8, 'o'
  			case 'b':
  				s.nextch()
  				base, prefix = 2, 'b'
  			default:
  				base, prefix = 8, '0'
  				digsep = 1 // leading 0
  			}
        digsep |= s.digits(base, &invalid)
        ...
  		}
  	...
  	}
  
  	// fractional part
  	if seenPoint {
  		kind = FloatLit
  		digsep |= s.digits(base, &invalid)
  	}
   ...
  }
  ```

  最后`a := 123` 整个语句会被解析为一个赋值语句`AssignStmt`，通过如下结构体进行表示:

  ```go
  // go/src/cmd/compile/internal/syntax/nodes.go
  type (
    ...
  	AssignStmt struct {
  		Op       Operator // 0 means no operation
  		Lhs, Rhs Expr     // Rhs == nil means Lhs++ (Op == Add) or Lhs-- (Op == Sub)
  		simpleStmt
  	}
    ...
  )
  ```

* 基于语法分析的结果，整个代码结构会被构建为一颗抽象语法树(ast)。抽象语法树是go编译器的中间结果`ir(intermediate representation)`，赋值语句`AssignStmt`会被构建为`ir.AssignStmt`，`:=`符号两边的字符被构建为节点`ir.Node`。

  ```go
  // go/src/cmd/compile/internal/ir/node.go
  
  // An AssignStmt is a simple assignment statement: X = Y.
  // If Def is true, the assignment is a :=.
  type AssignStmt struct {
  	miniStmt
  	X   Node
  	Def bool
  	Y   Node
  }
  
  // A Node is the abstract interface to an IR node.
  type Node interface {
  	...
  	// Source position.
  	Pos() src.XPos
  	SetPos(x src.XPos)
  	...
  	// Fields specific to certain Ops only.
  	Type() *types.Type
  	SetType(t *types.Type)
  	Val() constant.Value
  	SetVal(v constant.Value)
  	...
  	// Typecheck values:
  	//  0 means the node is not typechecked
  	//  1 means the node is completely typechecked
  	//  2 means typechecking of the node is in progress
  	//  3 means the node has its type from types2, but may need transformation
  	Typecheck() uint8
  	SetTypecheck(x uint8)
  }
  ```

* 最后，编译器会对抽象语法树的节点进行类型检查(typecheck)。检查的过程中，会将右边的节点`rhs`的类型`r.Type()`赋值给左边的节点`lhs`，因此最终变量a的类型(Kind)即为`123`的类型，为整型(types.TINT, go/src/cmd/compile/internal/types/type.go)。

  ```go
  // go/src/cmd/compile/internal/typecheck/stmt.go
  
  // type check assignment.
  // if this assignment is the definition of a var on the left side,
  // fill in the var's type.
  func tcAssign(n *ir.AssignStmt) {
    ...
    lhs, rhs := []ir.Node{n.X}, []ir.Node{n.Y}
    assign(n, lhs, rhs)
    ...
  }
  
  func assign(stmt ir.Node, lhs, rhs []ir.Node) {
    ...
    assignType := func(i int, typ *types.Type) {
      checkLHS(i, typ)
      if typ != nil {
        checkassignto(typ, lhs[i])
      }
    }
    ...
    assignType(0, r.Type())
    ...
  }
  ```

  ```go
  // go/src/cmd/compile/internal/typecheck/typecheck.go
  
  func checkassignto(src *types.Type, dst ir.Node) {
  	...
  	if op, why := Assignop(src, dst.Type()); op == ir.OXXX {
  		base.Errorf("cannot assign %v to %L in multiple assignment%s", src, dst, why)
  		return
  	}
  }
  ```

## 类型推断示例分析

根据上述原理，再看这三个表达式有何编译的执行过程有何不同：

```go
a := 123 
var a = 123
var a int = 123.0
```

`a := 123` 会显式的触发类型推断，编译器解析右边的每一个字符为十进制数字(IntLit)，然后构建为一个整型节点，在类型检查的时候，将其类型赋值给左边的节点变量`a`。

由于`var a = 123`左边的`a`未显式指定其类型，因此仍然会触发类型推断，`ir.AssignStmt.Def=false`，过程同上，依然在类型检查的时候，将`123`的类型赋值给左边的`a`。

对于`var a int = 123.0`， 由于`123.0`包含小数点'.'，编译器解析右边的每一个字符为十进制浮点数(FloatLit)，由于赋值操作符`=`左边显式定义了`a`的类型为`int`, 因此在类型检查阶段，右边的`123.0`会发生隐式类型转换，因为类型兼容，会转换为整型`123`。因此对于显式指定类型的表达式不会发生类型推断。

同理，结合类型转换的原理，前言中的示例1、2、3便可迎刃而解。

# 总结

1. 常量不允许取址。

2. 运算符两边的操作数类型必须相同。

3. 如果运算符两边是不同类型的未定义常量(untyped constant)，则会发生隐式转换，且转换的优先级为: 

   整数(int) <符文数(rune)<浮点数(float)<复数(Imag)。

4. 如果运算符的某一边是已定义类型常量(变量标识符)，则该已定义类型的常量任何时候都不会发生类型转换。因为**编译器不允许对变量标识符引用的值进行强制类型转换**。

5. `:=`会显式的触发类型推断，其只能作用于函数或方法体内。

6. 不指定类型的`var`变量声明，也会触发类型推断，可声明于局部也可声明在全局。

7. 指定类型的`var`变量声明，不会触发类型推断(因为类型已经显式指定了)，但有可能发生类型隐式转换。