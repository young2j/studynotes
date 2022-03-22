# go-zero开发辅助工具

## 辅助工具命令行参数

```shell
cd templates/tools
go run . --help
  -f string
        传入goctl生成的模型文件,如 usersCustomerModel.go
        传入补全rpc代码时的.proto文件,如 access.proto
  -s string
        模块的名称，默认提取模型文件名第一个单词
  -t string
        生成的目标, tag-补全模型tag, api-生成.api文件, rpc-生成.proto文件, logic-补全rpc逻辑代码
```

## 第一步: 生成模型代码

根据项目下的 `templates/mongo`利用 `goctl`工具生成mongo模型的增删改查代码:

```shell
# cd services/users/endpoint/api
goctl model mongo -type UsersStaff -c -style goZero -dir . -home ../../../templates/
```

## 第二步: 设计模型字段

在生产的代码中设计 `mongo`模型字段(字段名+类型+注释):

```go
// cd services/users/model/usersStaffModel.go
// 后台用户模型
type UsersStaff struct {
	Id          bson.ObjectId  // ObjectId
	Uid         string         // uuid
	Name        string         // 姓名
	Phone       string         // 用户电话
	Email       string         // 邮箱
	DateJoined  time.Time      // 入职时间
	OnJob       int32          // 是否在职
	EmployeeId  string         // 工号
	Role        string         // 角色名,角色需要进行分配
	CreatedTime time.Time      // 创建时间
	UpdatedTime time.Time      // 更新时间
}
```

## 第三步：补全模型tag

```shell
# cd templates/tools
# go run . --help
go run . -t tag -f ../../services/users/model/usersStaffModel.go 
```

控制台输出：复制粘贴即可

```shell
type UsersStaff struct {
  Id bson.ObjectId `json:"id" bson:"_id,omitempty" description:"ObjectId"`
  Uid string `json:"uid" bson:"uid" description:"uuid"`
  Name string `json:"name" bson:"name" description:"姓名"`
  Phone string `json:"phone" bson:"phone" description:"用户电话"`
  Email string `json:"email" bson:"email" description:"邮箱"`
  DateJoined time.Time `json:"dateJoined" bson:"dateJoined" description:"入职时间"`
  OnJob int32 `json:"onJob" bson:"onJob" description:"是否在职"`
  EmployeeId string `json:"employeeId" bson:"employeeId" description:"工号"`
  Role string `json:"role" bson:"role" description:"角色名,角色需要进行分配"`
  CreatedTime time.Time `json:"createdTime" bson:"createdTime" description:"创建时间"`
  UpdatedTime time.Time `json:"updatedTime" bson:"updatedTime" description:"更新时间"`

}
```

## 第四步：生成api文件

```shell
go run . -t api -f ../../services/users/model/usersStaffModel.go [-s users]
# ls services/users/model
# users.api
# ...
```

## 第五步: 生成rpc文件

```shell
go run . -t rpc -f ../../services/users/model/usersStaffModel.go [-s users]
# ls services/users/model
# users.proto
# ...
```

## 第六步: 补全rpc-crud代码

> 注意此操作会覆盖文件内容，原则上只运行一次

```
go run . -t logic -f ../../services/users/endpoint/rpc/users.proto
```
