# copier

在日常开发中，我们时常需要将一个变量的值快速应用到另一个变量上，动态语言往往都有相应的语法让开发人员快速进行上述操作，例如`python`中将一个`dict`(d1)的`key:value`解包给到另一个`dict`(d2):

```python
d1 = {"key1": 1, "key2": 2}
d2 = {"key3": 3, **d1 } # {'key3': 3, 'key1': 1, 'key2': 2}
```

再如`javascript`中可以使用`...`运算符进行相应的操作:

```javascript
const d1 = { key1: 1, key2: 2}
const d2 = { key3: 3, ...d1 } // {key3: 3, key1: 1, key2: 2}
```

然而，静态语言中各变量有强制的类型限制。在`golang`中当我们需要从一个结构体将值拷贝到另一个结构体时，即使结构体字段相同，仍然需要先声明并初始化一个目标变量，然后一个字段一个字段地进行赋值， 例如有两个字段相同的结构体`Perm1和Perm2`:

```go
type Perm1 struct {
	Action string
	Label  string
}

type Perm2 struct {
	Action     string
	Label      string
}

// 现有Perm1类型的一个变量, 需要转换为一个Perm2类型的变量
perm1 := Perm1{Action: "GET", Label: "rest-get-method"}
perm2 := Perm2{
  Action: perm1.Action,
  Label:  perm1.Label,
}
```

当结构体字段较多，或者类型嵌套较深时，仅仅只是为了拷贝数据值，就需要写大段大段的代码来一个字段一个字段地进行赋值。为此，需要一个包来专门提供数据值拷贝的功能， 例如第三方包[`copier`](https://github.com/jinzhu/copier), 当数据结构较复杂时，我们只需要一行代码，就可以将同名字段的值进行拷贝:

```go
perm1 := Perm1{Action: "GET", Label: "rest-get-method"}
perm2 := Perm2{}
copier.Copy(&perm2, &perm1)
```

最近在业务开发中，需要在http响应结果`httpResp`、rpc响应结果`rpcResp`和数据库查询结果`model`三者之间进行数据转换，于是使用了`copier`来进行数据拷贝。然而，后端数据库使用了`mongodb`, `copier`却不支持`bson.ObjectId`与`string`之间的转换。在rpc服务中会报出如下错误:

```shell
grpc: error while marshaling: string field contains invalid UTF-8
```

原因是`copier`将`ObjectId`转换为了非`UTF-8`编码的值, 同时`string`也无法正常转换为`ObjectId`:

```shell
# primitive.ObjectId -> string
from.Id1: ObjectIdHex("61fd0e4d18ef1dc958a6a796") to.Id1: a�M��X��� 
from.Id2: ObjectID("61fd0e4d5093876fcc4c0990") to.Id2:  
# string -> bson.ObjectId
from.Id1Hex: 61f04828eb37b662c8f3b085 to.Id1Hex: ObjectIdHex("363166303438323865623337623636326338663362303835") 
from.Id2Hex: 61f04828eb37b662c8f3b085 to.Id2Hex: ObjectID("000000000000000000000000") 
```

使用过程中也发现`copier`不支持如下类型的拷贝:

```go
psm1 := map[int]*[]string{1: {"a", "b", "c"}}
psm2 := make(map[int]*[]string) // make(map[int][]string) 类型可以正常拷贝
copier.Copy(&psm2, &psm1)
// panic: reflect.Value.Addr of unaddressable value
```

由于数据结构嵌套较深，需要写大量的数据转换代码，仅仅是为了将`ObejctId`转换为`string`，代码可读性变得极差。`copier`目前也无法应用。

# gocopy

基于业务需要，便自行实现了一个`golang`数据值拷贝的库——[`gocopy`](https://github.com/young2j/gocopy), 原理同`copier`都是利用反射`reflect`来实现。

## Copy slice

`gocopy`支持拷贝`slice`：

```go
s1 := []int{3, 4, 5}
s2 := make([]int, 0)
gocopy.Copy(&s2, &s1)
fmt.Printf("s2: %v\n", s2)
// s2: [3 4 5]
```

## Copy map

`gocopy`支持拷贝`map`:

```go
m1 := map[string]int{"key1": 1, "key2": 2}
m2 := make(map[string]int)
gocopy.Copy(&m2, &m1)
fmt.Printf("m2: %v\n", m2)
// m2: map[key1:1 key2:2]
```

再看看`copier`拷贝会报错的例子:

```go
psm1 := map[int]*[]string{1: {"a", "b", "c"}}
psm2 := make(map[int]*[]string)
copier.Copy(&psm2, &psm1)
fmt.Printf("psm2: %#v\n", psm2)
// psm2: map[int]*[]string{1:(*[]string)(0xc0000a6570)}
```

## Copy struct

`gocopy`支持拷贝`struct`：

```go
roll := 100
	st1 := model.AccessRolePerms{
		Role: "角色",
		Roll: &roll,
		EmbedFields: model.EmbedFields{
			EmbedF1: "embedF1",
		},
		Actions: []string{"GET", "POST"},
		Perms:   []*model.Perm{{Action: "GET", Label: "rest-get-method"}},
		PermMap: map[string]*model.Perm{"perm": {Action: "PUT", Label: "rest-put-method"}},
	}
	st2 := types.AccessRolePerms{}
	gocopy.Copy(&st2, &st1)
	fmt.Println("==============================")
	fmt.Printf("st2.Role: %v\n", *st2.Role)
	fmt.Printf("st2.Roll: %v\n", *st2.Roll)
	fmt.Printf("st2.Actions: %v\n", st2.Actions)

	for _, v := range st2.Perms {
		fmt.Printf("Perms: %#v\n", v)
	}
	for k, v := range st2.PermMap {
		fmt.Printf("PermMap k:%v v:%#v\n", k, v)
	}

// st2.Role: 角色
// st2.Roll: 100
// st2.Actions: [GET POST]
// Perms: &types.Perm{Action:"GET", Label:"rest-get-method"}
// PermMap k:perm v:&types.Perm{Action:"PUT", Label:"rest-put-method"}
```

## Append mode

`gocopy`还支持附加拷贝模式(`append mode`):

### Append slice

```go
opts := gocopy.Option{Append: true}
as1 := []int{3, 4, 5}
as2 := []int{1, 2}
gocopy.CopyWithOption(&as2, &as1, &opts)
fmt.Printf("as2: %v\n", as2)
// as2: [1 2 3 4 5]
```

### Append map

```go
opts := gocopy.Option{Append: true}
am1 := map[string]int{"key1": 1, "key2": 2}
am2 := map[string]int{"key0": 0, "key2": 3}
gocopy.CopyWithOption(&am2, &am1, &opts)
fmt.Printf("am2: %v\n", am2)

ams1 := map[string][]int{"key1": {1}, "key2": {2}}
ams2 := map[string][]int{"key0": {0}, "key2": {3}}
gocopy.CopyWithOption(&ams2, &ams1, &opts)
fmt.Printf("ams2: %v\n", ams2)

// am2: map[key0:0 key1:1 key2:2]
// ams2: map[key0:[0] key1:[1] key2:[3 2]]
```

### Append struct map/slice field

```go
opts := gocopy.Option{Append: true}
ast1 := model.AccessRolePerms{
  Actions: []string{"PUT", "DELETE"},
  Perms:   []*model.Perm{{Action: "PUT", Label: "rest-put-method"}},
  PermMap: map[string]*model.Perm{"delete": {Action: "DELETE", Label: "rest-delete-method"}},
}
ast2 := types.AccessRolePerms{
  Actions: []string{"GET", "POST"},
  Perms:   []*types.Perm{{Action: "GET", Label: "rest-get-method"}},
  PermMap: map[string]*types.Perm{"get": {Action: "GET", Label: "rest-get-method"}},
}
gocopy.CopyWithOption(&ast2, &ast1, &opts)

fmt.Printf("ast2.Actions: %v\n", ast2.Actions)
for i, perm := range ast2.Perms {
  fmt.Printf("ast2.Perms[%v]: %#v\n", i, perm)
}
for i, pm := range ast2.PermMap {
  fmt.Printf("ast2.PermMap[%v]: %#v\n", i, pm)
}

// ast2.Actions: [GET POST PUT DELETE]
// ast2.Perms[0]: &types.Perm{Action:"GET", Label:"rest-get-method"}
// ast2.Perms[1]: &types.Perm{Action:"PUT", Label:"rest-put-method"}
// ast2.PermMap[delete]: &types.Perm{Action:"DELETE", Label:"rest-delete-method"}
// ast2.PermMap[get]: &types.Perm{Action:"GET", Label:"rest-get-method"}
```

## Copy specified field

`gocopy`可以通过字段名指定将某一个字段拷贝至另一个字段:

```go
// from field to another field
ost1 := model.AccessRolePerms{
  From: "fromto",
}
ost2 := types.AccessRolePerms{}
opt := gocopy.Option{
  NameFromTo:       map[string]string{"From": "To"},
}
gocopy.CopyWithOption(&ost2, &ost1, &opt)

fmt.Printf("ost2.To: %v\n", ost2.To)

// ost2.To: fromto
```

## ObjectId and String

`gocopy`支持将`ObjectId`字段转换为`string`类型，反之亦然。

```go
// objectId to string and vice versa
from := model.AccessRolePerms{
  Id1:    bson.NewObjectId(),  // "github.com/globalsign/mgo/bson"
  Id2:    primitive.NewObjectID(), // "go.mongodb.org/mongo-driver/bson/primitive"
  Id1Hex: "61f04828eb37b662c8f3b085",
  Id2Hex: "61f04828eb37b662c8f3b085",
}
to := types.AccessRolePerms{
  Actions: []string{"GET", "POST"},
}
option := &gocopy.Option{
  ObjectIdToString: map[string]string{"Id1": "mgo", "Id2": "official"},
  StringToObjectId: map[string]string{"Id1Hex": "mgo", "Id2Hex": "official"},
  Append:           true,
}
gocopy.CopyWithOption(&to, from, option)

fmt.Printf("from.Id1: %v to.Id1: %v \n", from.Id1, to.Id1)
fmt.Printf("from.Id2: %v to.Id2: %v \n", from.Id2, to.Id2)
fmt.Printf("from.Id1Hex: %v to.Id1Hex: %v \n", from.Id1Hex, to.Id1Hex)
fmt.Printf("from.Id2Hex: %v to.Id2Hex: %v \n", from.Id2Hex, to.Id2Hex)

// from.Id1: ObjectIdHex("61f6cdf318ef1d4366bca973") to.Id1:61f6cdf318ef1d4366bca973
// from.Id2: ObjectID("61f6cdf3cc541c1bc35a41fc") to.Id2:61f6cdf3cc541c1bc35a41fc
// from.Id1Hex: 61f04828eb37b662c8f3b085 to.Id1Hex:ObjectIdHex("61f04828eb37b662c8f3b085")
// from.Id2Hex: 61f04828eb37b662c8f3b085 to.Id2Hex:ObjectID("61f04828eb37b662c8f3b085")
```

## Benchmark

使用`gocopy`和`copier`分别拷贝相同的结构体，做个简单的`benchmark`, 内存分配及占用降低了约50%，运行效率也提升了约50%。

```shell
goos: darwin
goarch: amd64
pkg: github.com/young2j/gocopy
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkCopy
BenchmarkCopy-4     	  122139	      8884 ns/op	    5592 B/op	      81 allocs/op
BenchmarkCopier
BenchmarkCopier-4   	   62940	     18695 ns/op	   14640 B/op	     166 allocs/op
PASS
ok  	github.com/young2j/gocopy	4.999s
```

## Github

https://github.com/young2j/gocopy

欢迎star🌟





