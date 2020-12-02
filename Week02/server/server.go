package server

import (
	"Week02/dao"
	"fmt"
)

func ServerHandle() {
	db, err := dao.InitDB("mysql", "root:root@tcp(localhost:6033)/mysql?charset=utf8")
	if err != nil {
		fmt.Printf("stack trace : \n%+v\n", err)
	}

	sql := "select studentId, studentName from mysql.class order by studentId limit 1"
	value, err := dao.DoWork(db, sql)
	if err != nil {
		fmt.Printf("stack trace : \n%+v\n", err)
	}

	fmt.Println("value is %d",  value)
}
