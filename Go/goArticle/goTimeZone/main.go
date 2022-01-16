package main

import (
	"database/sql"
	"fmt"
	"time"

	// driver
	_ "github.com/go-sql-driver/mysql"
)

var localMysqlDSN = "root:rootroot@tcp(127.0.0.1:3306)/hello?parseTime=true"
func main() {
	db, _ := sql.Open("mysql", localMysqlDSN)
	defer db.Close()
	
	var (
		id int64
		updatedAt time.Time
	)
	row := db.QueryRow("SELECT * FROM time_test")
	_ = row.Scan(&id, &updatedAt)
	fmt.Printf("id: %v  updatedAt:%v\n", id, updatedAt)

	cstSh, _ := time.LoadLocation("Asia/Shanghai")
	usNY, _ := time.LoadLocation("America/New_York")
	updatedAtCN :=updatedAt.In(cstSh)
	updatedAtNY :=updatedAt.In(usNY)
	fmt.Printf("updatedAtCN: %v\n", updatedAtCN)
	fmt.Printf("updatedAtNY: %v\n", updatedAtNY)

	// 两个时区今日0点
	todayDateStrCN := updatedAtCN.Format("2006-01-02")
	todayDateStrNY := updatedAtNY.Format("2006-01-02")
	todayZeroTimeCN, _ := time.ParseInLocation("2006-01-02", todayDateStrCN, cstSh) //上海时区
	todayZeroTimeNY, _ := time.ParseInLocation("2006-01-02", todayDateStrNY, usNY) // 纽约时区

	fmt.Println("todayZeroTimeCN:", todayZeroTimeCN)
	fmt.Println("todayZeroTimeNY:", todayZeroTimeNY)
	fmt.Println("updatedAt is before China-today  :", updatedAt.Before(todayZeroTimeCN))
	fmt.Println("updatedAt is before NewYork-today:", updatedAt.Before(todayZeroTimeNY))
	fmt.Println("updatedAt is after China-today   :", updatedAt.After(todayZeroTimeCN))
	fmt.Println("updatedAt is after NewYork-today :", updatedAt.After(todayZeroTimeNY))
}