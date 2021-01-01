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
