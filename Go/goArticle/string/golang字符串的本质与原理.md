# 字符串的本质

## 字符串的定义

`golang`中的字符(`character`)串指的是所有8比特位字节字符串的集合，通常(非必须)是`UTF-8` 编码的文本。 字符串可以为空，但不能是`nil`。 字符串在编译时即确定了长度，值是不可变的。

```go
// go/src/builtin/builtin.go

// string is the set of all strings of 8-bit bytes, conventionally but not
// necessarily representing UTF-8-encoded text. A string may be empty, but
// not nil. Values of string type are immutable.
type string string
```

字符串在本质上是一串字符数组，每个字符在存储时都对应了一个或多个整数，整数是多少取决于字符集的编码方式。

```go
s := "golang"
for i := 0; i < len(s); i++ {
  fmt.Printf("s[%v]: %v\n",i, s[i])
}

// s[0]: 103
// s[1]: 111
// s[2]: 108
// s[3]: 97
// s[4]: 110
// s[5]: 103
```

字符串在编译时类型为`string`,在运行时其类型定义为一个结构体，位于`reflect`包中:

```go
// go/src/reflect/value.go

// StringHeader is the runtime representation of a string.
// ...
type StringHeader struct {
	Data uintptr
	Len  int
}
```

根据运行时字符串的定义可知，在程序运行的过程中，字符串存储了长度(`Len`)及指向实际数据的指针(`Data`)。

## 字符串的长度

`golang`中所有文件都采用`utf8`编码，字符常量也使用`utf8`编码字符集。1个英文字母占1个字节长度，一个中文占3个字节长度。go中对字符串取长度`len(s)`指的是字节长度，而不是字符个数，这与动态语言如`python`中的表现有所差别。如:

```python
print(len("go语言")) 
# 4
```

```go
s := "go语言"
fmt.Printf("len(s): %v\n", len(s)) 
// len(s): 8
```

## 字符与符文

`go`中存在一个特殊类型——符文类型(`rune`)，用来表示和区分字符串中的字符。`rune`的本质是`int32`。字符串符文的个数往往才比较符合我们直观感受上的字符串长度。要计算字符串符文长度，可以先将字符串转为`[]rune`类型，或者利用标准库中的`utf8.RuneCountInString()`函数。

```go
s := "go语言"
fmt.Println(len([]rune(s)))
// 4

count := utf8.RuneCountInString(s)
fmt.Println(count)
// 4
```

当用`range`遍历字符串时，遍历的就不再是单字节，而是单个符文`rune`。

```go
s := "go语言"
for _, r := range s {
    fmt.Printf("rune: %v  string: %#U\n", r, r)
}

// rune: 103  unicode: U+0067 'g'
// rune: 111  unicode: U+006F 'o'
// rune: 35821  unicode: U+8BED '语'
// rune: 35328  unicode: U+8A00 '言'
```

# 字符串的原理

## 字符串的解析

`golang`在词法解析阶段，通过扫描源代码，将双引号和反引号开头的内容分别识别为标准字符串和原始字符串:

```go
// go/src/cmd/compile/internal/syntax/scanner.go

func (s *scanner) next() {
	...
	switch s.ch {
	...
	case '"':
		s.stdString()

	case '`':
		s.rawString()
  ...
```

然后，不断的扫描下一个字符，直到遇到另一个双引号和反引号即结束扫描。并通过`string(s.segment())`将解析到的字节转换为字符串，同时通过`setLlit()`方法将扫描到的内容类型(`kind`)标记为`StringLit`。

```go
func (s *scanner) stdString() {
	ok := true
	s.nextch()

	for {
    if s.ch == '"' {
			s.nextch()
			break
		}
    ...
		s.nextch()
	}

	s.setLit(StringLit, ok)
}

func (s *scanner) rawString() {
	ok := true
	s.nextch()

	for {
    if s.ch == '`' {
			s.nextch()
			break
		}
		...
		s.nextch()
	}
  
	s.setLit(StringLit, ok)
}

// setLit sets the scanner state for a recognized _Literal token.
func (s *scanner) setLit(kind LitKind, ok bool) {
	s.nlsemi = true
	s.tok = _Literal
	s.lit = string(s.segment())
	s.bad = !ok
	s.kind = kind
}
```

## 字符串的拼接

字符串可以通过`+`进行拼接:

```go
s := "go" + "lang"
```

在编译阶段构建抽象语法树时，等号右边的`"go"+"lang"`会被解析为一个字符串相加的表达式(`AddStringExpr`)节点，该表达式的操作`op`为`OADDSTR`。相加的各部分字符串被解析为节点`Node`列表，并赋给表达式的`List`字段：

```go
// go/src/cmd/compile/internal/ir/expr.go

// An AddStringExpr is a string concatenation Expr[0] + Exprs[1] + ... + Expr[len(Expr)-1].
type AddStringExpr struct {
	miniExpr
	List     Nodes
	Prealloc *Name
}

func NewAddStringExpr(pos src.XPos, list []Node) *AddStringExpr {
	n := &AddStringExpr{}
	n.pos = pos
	n.op = OADDSTR
	n.List = list
	return n
}
```

在构建抽象语法树时，会遍历整个语法树的表达式，在遍历的过程中，识别到操作`Op`的类型为`OADDSTR`，则会调用`walkAddString`对字符串加法表达式进行进一步处理:	

```go
// go/src/cmd/compile/internal/walk/expr.go
func walkExpr(n ir.Node, init *ir.Nodes) ir.Node {
	...
	n = walkExpr1(n, init)
	...
	return n
}

func walkExpr1(n ir.Node, init *ir.Nodes) ir.Node {
	switch n.Op() {
	...
	case ir.OADDSTR:
		return walkAddString(n.(*ir.AddStringExpr), init)
  ...
	}
  ...
}

```

`walkAddString`首先计算相加的字符串的个数`c`,如果相加的字符串个数小于2，则会报错。接下来会对相加的字符串字节长度求和，如果字符串总字节长度小于32，则会通过`stackBufAddr()`在栈空间开辟一块32字节的缓存空间。否则会在堆区开辟一个足够大的内存空间，用于存储多个字符串。

```go
// go/src/cmd/compile/internal/walk/walk.go
const tmpstringbufsize = 32


// go/src/cmd/compile/internal/walk/expr.go
func walkAddString(n *ir.AddStringExpr, init *ir.Nodes) ir.Node {
	c := len(n.List)

	if c < 2 {
		base.Fatalf("walkAddString count %d too small", c)
	}

	buf := typecheck.NodNil()
	if n.Esc() == ir.EscNone {
		sz := int64(0)
		for _, n1 := range n.List {
			if n1.Op() == ir.OLITERAL {
				sz += int64(len(ir.StringVal(n1)))
			}
		}

		// Don't allocate the buffer if the result won't fit.
		if sz < tmpstringbufsize {
			// Create temporary buffer for result string on stack.
			buf = stackBufAddr(tmpstringbufsize, types.Types[types.TUINT8])
		}
	}

	// build list of string arguments
	args := []ir.Node{buf}
	for _, n2 := range n.List {
		args = append(args, typecheck.Conv(n2, types.Types[types.TSTRING]))
	}

	var fn string
	if c <= 5 {
		// small numbers of strings use direct runtime helpers.
		// note: order.expr knows this cutoff too.
		fn = fmt.Sprintf("concatstring%d", c)
	} else {
		// large numbers of strings are passed to the runtime as a slice.
		fn = "concatstrings"

		t := types.NewSlice(types.Types[types.TSTRING])
		// args[1:] to skip buf arg
		slice := ir.NewCompLitExpr(base.Pos, ir.OCOMPLIT, t, args[1:])
		slice.Prealloc = n.Prealloc
		args = []ir.Node{buf, slice}
		slice.SetEsc(ir.EscNone)
	}

	cat := typecheck.LookupRuntime(fn)
	r := ir.NewCallExpr(base.Pos, ir.OCALL, cat, nil)
	r.Args = args
	r1 := typecheck.Expr(r)
	r1 = walkExpr(r1, init)
	r1.SetType(n.Type())

	return r1
}
```

如果用于相加的字符串个数小于等于5个，则会调用运行时的字符串拼接`concatstring1-concatstring5`函数。否则调用运行时的`concatstrings`函数，并将字符串通过切片`slice`的形式传入。类型检查中的`typecheck.LookupRuntime(fn)`方法查找到运行时的字符串拼接函数后，将其构建为一个调用表达式，操作`Op`为`OCALL`，最后遍历调用表达式完成调用。`concatstring1-concatstring5`中的每一个调用最终都会调用`concatstrings`函数。

```go
// go/src/runtime/string.go

const tmpStringBufSize = 32
type tmpBuf [tmpStringBufSize]byte

func concatstring2(buf *tmpBuf, a0, a1 string) string {
	return concatstrings(buf, []string{a0, a1})
}

func concatstring3(buf *tmpBuf, a0, a1, a2 string) string {
	return concatstrings(buf, []string{a0, a1, a2})
}

func concatstring4(buf *tmpBuf, a0, a1, a2, a3 string) string {
	return concatstrings(buf, []string{a0, a1, a2, a3})
}

func concatstring5(buf *tmpBuf, a0, a1, a2, a3, a4 string) string {
	return concatstrings(buf, []string{a0, a1, a2, a3, a4})
}
```

`concatstring1-concatstring5`已经存在一个32字节的临时缓存空间供其使用, 并通过`slicebytetostringtmp`函数将该缓存空间的首地址作为字符串的地址，字节长度作为字符串的长度。如果待拼接字符串的长度大于32字节，则会调用`rawstring`函数，该函数会在堆区为字符串分配存储空间, 并且将该存储空间的地址指向字符串。由此可以看出，字符串的底层是字节切片，且指向同一片内存区域。在分配好存储空间、完成指针指向等工作后，待拼接的字符串切片会被一个一个地通过内存拷贝`copy(b,x)`到分配好的存储空间`b`上。

```go
// concatstrings implements a Go string concatenation x+y+z+...
func concatstrings(buf *tmpBuf, a []string) string {
	...
  l := 0

	for i, x := range a {
    ...
    n := len(x)
		...
		l += n
		...
	}
	s, b := rawstringtmp(buf, l)
	for _, x := range a {
		copy(b, x)
		b = b[len(x):]
	}
	return s
}

func rawstringtmp(buf *tmpBuf, l int) (s string, b []byte) {
	if buf != nil && l <= len(buf) {
		b = buf[:l]
		s = slicebytetostringtmp(&b[0], len(b))
	} else {
		s, b = rawstring(l)
	}
	return
}

func slicebytetostringtmp(ptr *byte, n int) (str string) {
	...
	stringStructOf(&str).str = unsafe.Pointer(ptr)
	stringStructOf(&str).len = n
	return
}

// rawstring allocates storage for a new string. The returned
// string and byte slice both refer to the same storage.
func rawstring(size int) (s string, b []byte) {
	p := mallocgc(uintptr(size), nil, false)

	stringStructOf(&s).str = p
	stringStructOf(&s).len = size

	*(*slice)(unsafe.Pointer(&b)) = slice{p, size, size}

	return
}

type stringStruct struct {
	str unsafe.Pointer
	len int
}
func stringStructOf(sp *string) *stringStruct {
	return (*stringStruct)(unsafe.Pointer(sp))
}
```

## 字符串的转换

尽管字符串的底层是字节数组， 但字节数组与字符串的相互转换并不是简单的指针引用，而是涉及了内存复制。当字符串大于32字节时，还需要申请堆内存。

```go
s := "go语言"
b := []byte(s) // stringtoslicebyte
ss := string(b) // slicebytetostring
```

当字符串转换为字节切片时，需要调用`stringtoslicebyte`函数，当字符串小于32字节时，可以直接使用缓存`buf`，但是当字节长度大于等于32时,`rawbyteslice`函数需要向堆区申请足够的内存空间，然后通过内存复制将字符串拷贝到目标地址。

```go
// go/src/runtime/string.go
func stringtoslicebyte(buf *tmpBuf, s string) []byte {
	var b []byte
	if buf != nil && len(s) <= len(buf) {
		*buf = tmpBuf{}
		b = buf[:len(s)]
	} else {
		b = rawbyteslice(len(s))
	}
	copy(b, s)
	return b
}

func rawbyteslice(size int) (b []byte) {
	cap := roundupsize(uintptr(size))
	p := mallocgc(cap, nil, false)
	if cap != uintptr(size) {
		memclrNoHeapPointers(add(p, uintptr(size)), cap-uintptr(size))
	}

	*(*slice)(unsafe.Pointer(&b)) = slice{p, size, int(cap)}
	return
}


func slicebytetostring(buf *tmpBuf, ptr *byte, n int) (str string) {
	...
	var p unsafe.Pointer
	if buf != nil && n <= len(buf) {
		p = unsafe.Pointer(buf)
	} else {
		p = mallocgc(uintptr(n), nil, false)
	}
	stringStructOf(&str).str = p
	stringStructOf(&str).len = n
	memmove(p, unsafe.Pointer(ptr), uintptr(n))
	return
}
```

字节切片转换为字符串时，原理同上。因此字符串和切片的转换涉及内存拷贝，在一些密集转换的场景中，需要评估转换带来的性能损耗。

# 总结

1. 字符串常量存储在静态存储区，其内容不可以被改变。
2. 字符串的本质是字符数组，底层是字节数组，且与字符串指向同一个内存地址。
3. 字符串的长度是字节长度，要获取直观长度，需要先转换为符文数组，或者通过`utf8`标准库的方法进行处理。
4. 字符串通过扫描源代码的双引号和反引号进行解析。
5. 字符串常量的拼接发生在编译时，且根据拼接字符串的个数调用了对应的运行时拼接函数。
6. 字符串变量的拼接发生在运行时。
7. 无论是字符串的拼接还是转换，当字符串长度小于32字节时，可以直接使用栈区32字节的缓存，反之，需要向堆区申请足够的存储空间。
8. 字符串与字节数组的相互转换并不是无损的指针引用，涉及到了内存复制。在转换密集的场景需要考虑转换的性能和空间损耗。