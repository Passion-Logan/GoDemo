package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type TableName struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3307)/test")
	if err != nil {
		fmt.Println("数据库打开失败:", err)
		return
	}
	Db = database
}

func insertData() {
	r, err := Db.Exec("insert into table_name(name) values (?)", "test")
	if err != nil {
		fmt.Println("数据插入失败", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("SQL执行失败", err)
		return
	}
	fmt.Println("成功的id", id)
}

func queryData(id int) []TableName {
	var tables []TableName
	err := Db.Select(&tables, "select id, name from table_name where id = ?", id)
	if err != nil {
		fmt.Println("SQL执行失败", err)
		return tables
	}
	return tables
}

func updateData(old TableName, new TableName) int64 {
	res, err := Db.Exec("update table_name set name = ? where id = ?", new.Name, old.Id)
	if err != nil {
		fmt.Println("SQL执行失败", err)
		return 0
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("失败行, ", err)
	}

	return row
}

func deleteData(id int) int64 {
	res, err := Db.Exec("delete from table_name where id = ?", id)
	if err != nil {
		fmt.Println("SQL执行失败", err)
		return 0
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("失败行, ", err)
	}

	return row
}

func main() {
	fmt.Println("删除成功行数:", deleteData(8))

	defer Db.Close()
}
