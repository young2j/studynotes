package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var dsn = "root:rootroot@tcp(%s:3306)/hello?parseTime=true"
var logger *log.Logger

func init() {
	dir, _ := os.Getwd()
	fmt.Println("当前路径：", dir)

	logFile := dir+"/logs/log.log"
	fileOut, err := os.Create(logFile)
	if os.IsExist(err){
		fileOut, err = os.OpenFile(logFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	outs := io.MultiWriter(os.Stdout, fileOut)
	logger = log.New(outs, "", log.Lshortfile|log.Ldate|log.Ltime)
	env := os.Getenv("TEST_CICD_ENV")
	switch env {
	case "ci":
		dsn = fmt.Sprintf(dsn, "mysql")
	case "dev":
		dsn = fmt.Sprintf(dsn, "127.0.0.1")
	default:
		dsn = fmt.Sprintf(dsn, "127.0.0.1")
	}
	logger.Printf("current running env: %s dsn: %s\n", env, dsn)
}

func queryUser() string {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Fatal(err)
	}
	var userName string
	row := db.QueryRow("select name from users where id=?", 3)
	if err = row.Scan(&userName); err != nil {
		logger.Fatal(err)
	}

	return userName
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		userName := queryUser()
		rw.Write([]byte(userName))
		logger.Printf("%s %s %s \n", r.URL.User, r.Method, r.RequestURI)
	})
	http.ListenAndServe("0.0.0.0:8888", nil)
}
