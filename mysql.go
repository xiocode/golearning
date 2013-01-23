package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	//"time"
)

func main() {
	db, err := sql.Open("mysql", "root:299792458@tcp(127.0.0.1:3306)/xio?charset=utf8")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO test (name) VALUES (?)")
	checkErr(err)

	res, err := stmt.Exec("研发部门")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update test set name=? where id=?")
	checkErr(err)

	res, err = stmt.Exec("xioupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM test")
	checkErr(err)

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(name)
	}

	//删除数据
	stmt, err = db.Prepare("delete from test where id=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
