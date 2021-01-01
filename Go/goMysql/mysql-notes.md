# 概述

`go`语言的标准库`database/sql`提供了操作关系型数据库的通用方法，它是对数据库接口的抽象和封装，必须配合数据库驱动使用。因此要实现数据库操作，代码中需要引入两个包——一个是面向用户操作层面的`database/sql`，一个是数据库驱动层面的`go-sql-driver/mysql`。不同的数据库需要相应引入不同的驱动库。此外，第三方库`sqlx`在`database/sql`的基础上零侵入地实现了更为便捷的`API`，因此实际编写代码的时候，通常需要同时引入上述三个包。本文就这三个包的常用方法及示例进行归纳整理。

# 客户端连接

## 基本连接与参数

```GO
package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
  // sql.Open返回一个*sql.DB对象, 它的作用是：
  // - 通过驱动程序打开和关闭与基础数据库的连接。
  // - 根据需要管理连接池。
  // sql.Open不会立即建立与数据库的任何连接，也不会验证驱动程序连接的参数(但会检查格式)。这个过程是懒加载的，也就是第一次用到的时候才会真正建立连接。
  db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hello")
  if err != nil {
      log.Fatal(err)
  }

  defer db.Close()

  // 设置连接的最大生存时间，以确保连接可以被驱动安全关闭。官方建议小于5分钟。
  db.SetConnMaxLifetime(time.Minute * 3)
  // 设置打开的最大连接数，取决于mysql服务器和具体应用程序
  db.SetMaxOpenConns(10)
  // 设置最大闲置连接数，这个连接数应大于等于打开的最大连接，否则需要额外连接时会频繁进行打开关闭。
  // 最好与最大连接数保持相同，当大于最大连接数时，内部自动会减少到与最大连接数相同。
  db.SetMaxIdleConns(10)

  // 设置闲置连接的最大存在时间, support>=go1.15
  db.SetConnMaxIdleTime(time.Minute * 3)
}
```

## DSN(data source name)

`DSN`以字符串的形式具体定义了数据库地址以及连接参数。连接参数以`key=value`的形式跟在具体数据库名称后，类似于url参数格式。具体的params可在`mysql.Config`结构体中查看。下面为参考示例：

```go
// dsn完整格式
dsn := "username:password@protocol(host:port)/dbname?param1=value1&...&paramN=valueN"

// loc=Local 设置时间为系统所在时间，并将数据库中的时间字段解析为go中的time.Time类型
username:password@unix(/tmp/mysql.sock)/dbname?loc=Local&parseTime=true

// 跳过tsl验证、事务自动提交
username:password@tcp(localhost:3306)/dbname?tls=skip-verify&autocommit=true

// ipv6连接
username:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname?timeout=90s&collation=utf8mb4_unicode_ci

// remote host
username:password@tcp(your-remote-uri.com:3306)/dbname

// 默认localhost:3306
user:password@tcp/dbname?charset=utf8mb4,utf8

// 默认tcp(localhost:3306)
username:password@/dbname
```

# 客户端操作

## 关于context api

go在1.7版本中引入 了`context` 包，此后标准库中很多接口都加上了 `context` 参数，其中就包含`database/sql`。因此你会看到两种`API`， 如：

* `DB.Query()` 与 `DB.QueryContext()`

* `DB.QueryRow()` 与 `DB.QueryRowContext()`

* `DB.Exec()` 与 `DB.ExecContext()`

* `DB.Begin()` 与 `DB.BeginTx()`

* ...

  Context API 在使用时，第一个参数需要传递ctx。
  除此以外，两种API在使用上没有任何差别。本文主要使用也推荐使用ContextAPI。

## QueryContext—查询多行

```GO
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


func main() {
  // 准备一个ctx
  ctx := context.Background()
  // sql.Open不会建立与数据库的任何连接，也不会验证驱动程序连接的参数(但会检查格式)。这个过程是懒加载的，也就是第一次用到的时候才会真正建立连接。
  db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hello")
  if err != nil {
      log.Fatal(err)
  }
  defer db.Close()

  var (
      id   int
      name string
  )

  rows, err := db.QueryContext(ctx, "SELECT id, name FROM users WHERE id<=?", 2)
  if err != nil {
      log.Fatal(err)
  }
  // 应该总是调用rows.Close(), 这个方法是无害的，甚至可以调用多次。
  defer rows.Close()
  for rows.Next() {
      err := rows.Scan(&id, &name)
      if err != nil {
          log.Fatal(err)
      }
      fmt.Printf("id: %v  name:%v\n", id, name)
      // id: 1  name:a君
      // id: 2  name:b君
  }
// 处理Rows.Scan可能遇到的任何错误
  err = rows.Err()
  if err != nil {
      log.Fatal(err)
  }
}
```

## QueryRowContext—查询单行

```GO
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
  // 准备一个ctx
  ctx := context.Background()

  // sql.Open不会建立与数据库的任何连接，也不会验证驱动程序连接的参数(但会检查格式)。这个过程是懒加载的，也就是第一次用到的时候才会真正建立连接。
  db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hello")
  if err != nil {
      log.Fatal(err)
  }
  defer db.Close()

  var (
      id   int
      name string
  )

  err = db.QueryRowContext(ctx, "SELECT id, name FROM users WHERE id=?", 1).Scan(&id, &name)

  switch {
      case err == sql.ErrNoRows:
          fmt.Println("return no result")
      case err != nil:
          log.Fatal(err)
      default:
          fmt.Printf("id: %v  name:%v\n", id, name)
          // id: 1  name:a君
  }
}
```

## Prepare—语句预处理

下面一行命令声明了一条预处理的sql语句。

```GO
stmt, err := db.Prepare("INSERT INTO  users(name) VALUES(?)")
```

预处理语句通常具有**安全、高效、便捷**的好处。在数据库层面，它**与一个数据库连接相绑定**。典型的流程是，客户端发送一个带有占位符的sql描述语句到服务器进行预处理，服务器返回该描述语句的ID，然后客户端通过发送这个ID和相应的参数至服务器执行该语句。

**在go中，预处理的工作机制是：**

1. 准备好一个描述语句，然后会在连接池中的一个连接上进行预处理；
2. 每一个描述语句都会记住使用的是池中的哪一个连接；
3. 在具体执行语句的时候，就会对应使用这个记住的连接，如果这个连接被关闭或者被占用，那么就会重新获取另一个连接进行预处理并执行。

**使用预处理语句需要注意的是： 在高并发场景下，可能会预先准备或者重新准备大量的处理语句，这样将导致大量连接处于忙碌状态，甚至超过服务器限制。**

## ExecContext—执行修改

```GO
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
  // ctx
  ctx := context.Background()
  // sql.Open不会建立与数据库的任何连接，也不会验证驱动程序连接的参数(但会检查格式)。这个过程是懒加载的，也就是第一次用到的时候才会真正建立连接。
  db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hello")
  if err != nil {
      log.Fatal(err)
  }

  defer db.Close()

  names := []string{"c君", "d君", "e君"}

  stmt, err := db.Prepare("INSERT INTO  users(name) VALUES(?)")
  if err != nil {
      log.Fatal(err)
  }
  defer stmt.Close()

  for _, name := range names {
      result, err := stmt.ExecContext(ctx, name)
      if err != nil {
          log.Fatal(err)
      }
      lastInserID, _ := result.LastInsertId()
      rowsAffected, _ := result.RowsAffected()
      fmt.Printf("lastInsetId: %v, rowsAffected: %v\n", lastInserID, rowsAffected)
      // lastInsetId: 6, rowsAffected: 1
      // lastInsetId: 7, rowsAffected: 1
      // lastInsetId: 8, rowsAffected: 1
  }
}
```

## 使用Query 还是 Exec ？

考虑如下两个操作:

```go
result, err := db.Exec("DELETE FROM users")
result, err := db.Query("DELETE FROM users")
// handle result...
```

当我们需要对返回结果`result`做进一步处理时，这两种方式都可以使用。因为处理完毕后连接就会关闭或者放回连接池。

然而，对于修改操作我们往往会忽略操作结果，比如：

```go
_, err := db.Exec("DELETE FROM users")  // OK
_, err := db.Query("DELETE FROM users") // BAD
```

**如果你想忽略查询结果，那永远不要像上面那样使用Query()，因为这个操作会保持数据库连接直到连接关闭。如果存在未读的数据行，这个连接就不会得到释放，其他操作也无法再使用这个连接。而等待垃圾回收这个连接也要花相当长一段时间。所以任何时候记住查询操作使用`Query`，修改操作使用`Exec`就对了。**

## Tx—事务

### 事务开启、执行、回滚与提交

```GO
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
  // ctx
  ctx := context.Background()
  // sql.Open不会建立与数据库的任何连接，也不会验证驱动程序连接的参数(但会检查格式)。这个过程是懒加载的，也就是第一次用到的时候才会真正建立连接。
  db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hello")
  if err != nil {
      log.Fatal(err)
  }

  defer db.Close()

  // 开启事务
  tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
  if err != nil {
      log.Fatal(err)
  }
  // 执行事务一
  _, err = tx.ExecContext(ctx, "UPDATE users SET name=? WHERE id=?", "字母君", 1)
  if err != nil {
      // 回滚
      if rollBackErr := tx.Rollback(); rollBackErr != nil {
          log.Fatalf("transaction rollback err: %v", rollBackErr.Error())
      }
      log.Fatal(err)
  }

  // 执行事务二
  var row struct {
      id   int
      name string
      age []byte // null field
  }
  err = tx.QueryRowContext(ctx, "SELECT * FROM users WHERE id=?", 1).Scan(&row.id, &row.name, &row.age)
  switch {
      case err == sql.ErrNoRows:
          fmt.Println("no data")
      case err != nil:
          // 回滚
          if rollBackErr := tx.Rollback(); rollBackErr != nil {
              log.Fatalf("transaction rollback err: %v", rollBackErr.Error())
          }
          log.Fatal(err)
      default:
          fmt.Println("row:", row)
          // row: {1 字母君 []}
  }

  // 提交事务
  if err = tx.Commit(); err != nil {
      log.Fatal(err)
  }
}
```

### 在事务中使用预处理语句

在事务中使用预处理语句时，一定不要在事务开启前就做好语句准备，也就是说**不要在事务中使用事务外准备的描述语句**。因为预处理语句特定于某一个连接，而事务会开启一个新的连接，**在事务中使用外层语句实质上会在事务的连接中重新准备sql语句**。

```GO
package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
  //ctx
  ctx := context.Background()
  // sql.Open不会建立与数据库的任何连接，也不会验证驱动程序连接的参数(但会检查格式)。这个过程是懒加载的，也就是第一次用到的时候才会真正建立连接。
  db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hello")
  if err != nil {
      log.Fatal(err)
  }
  defer db.Close()

// 错误：不要在事务开启前做语句准备
// stmt, err := db.Prepare("UPDATE users SET name=? WHERE id=?")

// 开启事务
  tx, err := db.Begin()
  if err != nil {
      log.Fatal(err)
  }
  // 准备sql语句
  stmt, err := tx.Prepare("UPDATE users SET name=? WHERE id=?")
  if err != nil {
      log.Fatal(err)
  }
  // defer stmt.Close()

  // 执行语句
  _, err = stmt.ExecContext(ctx, "字母酱", 2)
  if err != nil {
      tx.Rollback()
      log.Fatal(err)
  }

  // 提交事务
  if err = tx.Commit(); err != nil {
      log.Fatal(err)
  }
  // go1.4及之前版本，一定要将这句放在事务tx.Commit()之后，否则会造成并发访问连接状态不一致的问题。
  defer stmt.Close()
}
```

## Null—蛋疼的空值处理

假如数据库中`users`表有三个字段`id<INT>, name<VARCHAR>, age<INT>`，其中`age`字段都为`NULL`.  现在查询一行数据：

```GO
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
  //ctx
  ctx := context.Background()
  // db
  db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hello")
  if err != nil {
      log.Fatal(err)
  }
  defer db.Close()

  var (
      id   int
      name string
      age  int
  )

  err = db.QueryRowContext(ctx, "SELECT * FROM users WHERE id=?", 1).Scan(&id, &name, &age)
  if err != nil {
      if err == sql.ErrNoRows {
          fmt.Println("return no result")
      } else {
          log.Fatal(err)
    // 2021/01/01 19:42:32 sql: Scan error on column index 2, name "age": converting NULL to int is unsupported
      }
  }

  fmt.Printf("id: %v  name:%v age: %v \n", id, name, age)
}
```

可以看到，go无法完成从`NULL`到`int`的转换，因为go中所有类型都具有零值，没有空值的说法。但是，在实际运用中，`NULL`值是非常普遍的，有时候我们甚至不知道哪些字段存在`NULL`，是否存在`NULL`.所以一个未知的`NULL`就可能导致程序挂掉。解决方案是声明可能为空的字段为`[]byte、sql.RawBytes、sql.Null*`类型。

### 声明空值为[]byte

当声明空值为`[]byte`时，会返回一个空切片`[]`, 如下：

```GO
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
  //ctx
  ctx := context.Background()
  // db
  db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hello")
  if err != nil {
      log.Fatal(err)
  }
  defer db.Close()

  var (
      id   int
      name string
      age  []byte //null
  )

  err = db.QueryRowContext(ctx, "SELECT * FROM users WHERE id=?", 1).Scan(&id, &name, &age)
  if err != nil {
      if err == sql.ErrNoRows {
          fmt.Println("return no result")
      } else {
          log.Fatal(err)
      }
  }

  fmt.Printf("id: %v  name:%v age: %v \n", id, name, age)
  // id: 1  name:字母君 age: []
}
```

### 声明空值为`sql.RawBytes`

最蛋疼的操作，甚至不能直接声明空值字段类型为`sql.RawBytes`，要先声明为空接口，再`new(sql.RawBytes)`一个指针引用到该空接口（经过实践，这一步都可以省了，直接空接口），才能支持`Scan`操作。最终会返回`<nil>`.

```go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
  //ctx
  ctx := context.Background()
  // db
  db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hello")
  if err != nil {
      log.Fatal(err)
  }
  defer db.Close()

  var (
      id   int
      name string
      age  interface{}
  )
  // age = new(sql.RawBytes)

  err = db.QueryRowContext(ctx, "SELECT * FROM users WHERE id=?", 1).Scan(&id, &name, &age)
  if err != nil {
      if err == sql.ErrNoRows {
          fmt.Println("return no result")
      } else {
          log.Fatal(err)
      }
  }

  fmt.Printf("id: %v  name:%v age: %v \n", id, name, age)
  // id: 1  name:字母君 age: <nil>
}
```

### 声明为`sql.Null*`

可以根据不同的字段类型，声明对应的空值为`sql.Nullint64、sql.NullString`等。这种方式最为严谨，虽然较为繁琐，但是依然推荐尽量使用这种处理方式。

```GO
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
  //ctx
  ctx := context.Background()
  // db
  db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hello")
  if err != nil {
      log.Fatal(err)
  }
  defer db.Close()

  var (
      id   int
      name string
      age  sql.NullInt64 // 整型空值类型
  )

  err = db.QueryRowContext(ctx, "SELECT * FROM users WHERE id=?", 1).Scan(&id, &name, &age)
  if err != nil {
      if err == sql.ErrNoRows {
          fmt.Println("return no result")
      } else {
          log.Fatal(err)
      }
  }

  if age.Valid{ //如果age不是空值，而是一个有效值
      fmt.Printf("id: %v  name:%v age: %v \n", id, name, age.Int64)
  } else {
      fmt.Printf("id: %v  name:%v age: %v \n", id, name, "NULL")
      // id: 1  name:字母君 age: NULL
  }
}
```

# 更易用的sqlx

`sqlx`在 `database/sql`的基础上扩展了更多便捷的功能，使用更加方便。`sqlx`并没有对`database/sql`底层接口做任何修改，而是通过新的方法扩展原有的功能，因此可以和`database/sql`接口方法一起使用。`sqlx`主要包含如下特征：

1. 可以直接将行数据序列化到结构体struct、映射map以及切片slice里；
2. 支持为预处理语句提供命名参数；
3. `Get`和`Select` api可以快速地获取查询结果到结构体或者切片中；
4. 更加简洁地错误处理, 通过`Must*API`实现了错误的自动panic。

## 安装

```go
go get github.com/jmoiron/sqlx
```

## QueryRowxContext—查询单行

```go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type userStruct struct {
	ID   int           `db:"id"`
	Name string        `db:"name"`
	Age  sql.NullInt64 `db:"age"`
}


func main() {
  //ctx
  ctx := context.Background()
  // db
  db := sqlx.MustOpen("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hello")
  defer db.Close()

  // 序列化到结构体
  var user userStruct
  row := db.QueryRowxContext(ctx, "SELECT * FROM users WHERE id=?", 1)
  err := row.StructScan(&user)
  if err != nil {
      log.Fatal(err)
  }
  // 序列化到map
  var userMap = make(map[string]interface{}, 3)
  row = db.QueryRowxContext(ctx, "SELECT * FROM users WHERE id=?", 2)
  err = row.MapScan(userMap)
  if err != nil {
      log.Fatal(err)
  }

  // 序列化到slice
  row = db.QueryRowxContext(ctx, "SELECT * FROM users WHERE id=?", 3)
  userSlice, err := row.SliceScan()
  if err != nil {
      log.Fatal(err)
  }

  fmt.Printf("%#v\n", user)
  // main.userStruct{ID:1, Name:"a君", Age:sql.NullInt64{Int64:0, Valid:false}}
  fmt.Printf("%#v\n", userMap)
  // map[string]interface{}{"age":interface {}(nil), "id":2, "name":[]uint8{0x62, 0xe5, 0x90, 0x9b}}
  fmt.Printf("%#v\n", userSlice)
  // []interface{}{3, []uint8{0x63, 0xe5, 0x90, 0x9b}, interface {}(nil)}
}
```

## GetContext—查询单行(便捷API)

省去了手动Scan操作。

```go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type userStruct struct {
  ID   int           `db:"id"`
  Name string        `db:"name"`
  Age  sql.NullInt64 `db:"age"`
}


func main() {
  //ctx
  ctx := context.Background()
  // db
  db := sqlx.MustOpen("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hello")
  defer db.Close()

  // 序列化到结构体
  var user userStruct
  err := db.GetContext(ctx, &user, "SELECT * FROM users WHERE age=?", 20)
  if err != nil {
      log.Fatal(err)
  }

  fmt.Printf("%#v\n", user)
  // main.userStruct{ID:7, Name:"d君", Age:sql.NullInt64{Int64:20, Valid:true}}
}
```

## QueryxContext—查询多行

```go
package main

import (
  "context"
  "database/sql"
  "fmt"
  "log"

  _ "github.com/go-sql-driver/mysql"
  "github.com/jmoiron/sqlx"
)

type userStruct struct {
  ID   int           `db:"id"`
  Name string        `db:"name"`
  Age  sql.NullInt64 `db:"age"`
}

func main() {
  //ctx
  ctx := context.Background()
  // db
  db := sqlx.MustOpen("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hello")
  defer db.Close()

  // 序列化到结构体
  var user userStruct
  rows, err := db.QueryxContext(ctx, "SELECT * FROM users")
  if err != nil {
      log.Fatal(err)
  }

  for rows.Next() {
      err = rows.StructScan(&user)
      if err != nil {
          log.Fatal(err)
      }
      fmt.Printf("%#v\n", user)
  }
}
  // main.userStruct{ID: 1, Name: "a君", Age: sql.NullInt64{Int64: 0, Valid: false}}
  // main.userStruct{ID: 2, Name: "b君", Age: sql.NullInt64{Int64: 0, Valid: false}}
  // main.userStruct{ID: 3, Name: "c君", Age: sql.NullInt64{Int64: 0, Valid: false}}
  // main.userStruct{ID: 7, Name: "d君", Age: sql.NullInt64{Int64: 18, Valid: true}}
  // main.userStruct{ID: 8, Name: "e君", Age: sql.NullInt64{Int64: 23, Valid: true}}
```

## SelectContext—查询多行(便捷API)

```go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type userStruct struct {
  ID   int           `db:"id"`
  Name string        `db:"name"`
  Age  sql.NullInt64 `db:"age"`
}

func main() {
  //ctx
  ctx := context.Background()
  // db
  db := sqlx.MustOpen("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hello")
  defer db.Close()

  // 序列化到结构体切片
  var users []userStruct
  err := db.SelectContext(ctx, &users, "SELECT * FROM users")
  if err != nil {
      log.Fatal(err)
  }

  for _, user := range users {
      fmt.Printf("%#v\n", user)
  }
  // main.userStruct{ID: 1, Name: "a君", Age: sql.NullInt64{Int64: 0, Valid: false}}
  // main.userStruct{ID: 2, Name: "b君", Age: sql.NullInt64{Int64: 0, Valid: false}}
  // main.userStruct{ID: 3, Name: "c君", Age: sql.NullInt64{Int64: 0, Valid: false}}
  // main.userStruct{ID: 7, Name: "d君", Age: sql.NullInt64{Int64: 20, Valid: true}}
  // main.userStruct{ID: 8, Name: "e君", Age: sql.NullInt64{Int64: 23, Valid: true}}
  // main.userStruct{ID: 9, Name: "f君", Age: sql.NullInt64{Int64: 2, Valid: true}}
}
```

## NamedQueryContext—命名参数查询

```go
package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type userStruct struct {
  ID   int    `db:"id"`
  Name string `db:"name"`
  Age  int    `db:"age"`
}

func main() {
  //ctx
  ctx := context.Background()
  // db
  db := sqlx.MustOpen("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hello")
  defer db.Close()

  // 序列化到结构体
  var user userStruct
  rows, err := db.NamedQueryContext(ctx, "SELECT * FROM users WHERE age < :adult", 
                                  map[string]interface{}{"adult": 18})
  if err != nil {
      log.Fatal(err)
  }

  for rows.Next() {
      err = rows.StructScan(&user)
      if err != nil {
          log.Fatal(err)
      }
      fmt.Printf("%#v\n", user)
  }
  // main.userStruct{ID:9, Name:"f君", Age:2}
}
```

## MustExecContext—执行修改(通过占位符)

```go
package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


func main() {
  //ctx
  ctx := context.Background()
  // db
  db := sqlx.MustOpen("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hello")
  defer db.Close()

  result := db.MustExecContext(ctx, "UPDATE users SET age=? WHERE id=?", 20, 7)

  lastInsertID, _ := result.LastInsertId()
  rowsAffected, _ := result.RowsAffected()
  fmt.Println(lastInsertID, rowsAffected)
  // 0 1
}
```

## NamedExecContext—执行修改(通过命名参数)

```GO
package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type user struct {
  Name string `db:"name"`
  Age  int    `db:"age"`
}

func main() {
  //ctx
  ctx := context.Background()
  // db
  db := sqlx.MustOpen("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hello")
  defer db.Close()

  result, err := db.NamedExecContext(ctx, "INSERT INTO users(name,age) VALUES(:name, :age)", &user{Name: "f君", Age: 2})
  // 使用map结构
  // result, err = db.NamedExecContext(
  // 	ctx,
  // 	"INSERT INTO users(name,age) VALUES(:name,:age)",
  // 	map[string]interface{}{
  // 		"name": "f君",
  // 		"age":  2,
  // 	})

  if err != nil {
      log.Fatal(err)
  }

  lastInsertID, _ := result.LastInsertId()
  rowsAffected, _ := result.RowsAffected()
  fmt.Println(lastInsertID, rowsAffected)
  // 9 1 
}
```

## Tx—使用事务

```go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type userStruct struct {
  ID   int           `db:"id"`
  Name string        `db:"name"`
  Age  sql.NullInt64 `db:"age"`
}

func main() {
  //ctx
  ctx := context.Background()
  // db
  db := sqlx.MustOpen("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hello")
  defer db.Close()

  // 开启事务
  tx := db.MustBegin()
  // 事务1
  tx.MustExecContext(ctx, "UPDATE users SET age=? WHERE id=?", 18, 3)

  // 事务2
  stmt, err := tx.PreparexContext(ctx, "SELECT * FROM users WHERE id=?")
  if err != nil {
      log.Fatal(err)
  }
  user1 := userStruct{}
  err = stmt.GetContext(ctx, &user1, 3)
  if err != nil {
      log.Fatal(err)
  }
  fmt.Printf("%#v\n", user1)
  // main.userStruct{ID: 3, Name: "c君", Age: sql.NullInt64{Int64: 18, Valid: true}}

  // 事务3
  namedStmt, err := tx.PrepareNamedContext(ctx, "SELECT * FROM users WHERE id=:id")
  if err != nil {
      log.Fatal(err)
  }
  user2 := userStruct{}
  err = namedStmt.GetContext(ctx, &user2, map[string]interface{}{"id": 8})
  if err != nil {
      log.Fatal(err)
  }
  fmt.Printf("%#v\n", user2)
  // main.userStruct{ID: 8, Name: "e君", Age: sql.NullInt64{Int64: 23, Valid: true}}

  // 事务4
  var users []userStruct
  err = tx.SelectContext(ctx, &users, "SELECT * FROM users")
  if err != nil {
      log.Fatal(err)
  }
  for _, user := range users {
      fmt.Printf("%#v\n", user)
  }
  // main.userStruct{ID: 1, Name: "a君", Age: sql.NullInt64{Int64: 0, Valid: false}}
  // main.userStruct{ID: 2, Name: "b君", Age: sql.NullInt64{Int64: 0, Valid: false}}
  // main.userStruct{ID: 3, Name: "c君", Age: sql.NullInt64{Int64: 18, Valid: true}}
  // main.userStruct{ID: 7, Name: "d君", Age: sql.NullInt64{Int64: 20, Valid: true}}
  // main.userStruct{ID: 8, Name: "e君", Age: sql.NullInt64{Int64: 23, Valid: true}}
  // main.userStruct{ID: 9, Name: "f君", Age: sql.NullInt64{Int64: 2, Valid: true}}

  // 事务提交
  tx.Commit()

  // 关闭语句
  defer stmt.Close()
  defer namedStmt.Close()
}
```

# 参考资料

https://pkg.go.dev/database/sql](https://pkg.go.dev/database/sql)

https://pkg.go.dev/github.com/go-sql-driver/mysql](https://pkg.go.dev/github.com/go-sql-driver/mysql)

[https://github.com/jmoiron/sqlx](https://github.com/jmoiron/sqlx)