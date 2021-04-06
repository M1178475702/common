package common

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
	"time"
)

func ifFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type User struct {
	Id int `json:"id"`
}

func TestTrxLock(t *testing.T) {
	defer func() {
		err := recover()
		fmt.Println(err)
		time.Sleep(60 * time.Second)
	}()
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	ifFatal(err)
	tx, err := db.Begin()
	ifFatal(err)

	//row := tx.QueryRow("select * from user")
	//ifFatal(err)

	//user := &User{}

	_, err = tx.Exec("update user set name='666' where id = 9")
	ifFatal(err)
	_, err = tx.Exec("insert into user (name) values ('zx')")
	ifFatal(err)

	d := make([]int, 0, 1)
	d[0] = 1 //panic

	err = tx.Commit()
	ifFatal(err)
}
