package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=astaxie password=astaxie dbname=test sslmode=disable")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username,department,created) VALUES($1,$2,$3) RETURNING uid")
	checkErr(err)

	//执行准备好的语句，将占位符替换为实际值。
	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	//pg不支持这个函数，因为他没有类似MySQL的自增ID
	// id, err := res.LastInsertId()
	// checkErr(err)
	// fmt.Println(id)

	//执行语句并获取返回值
	var lastInsertId int
	err = db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertId)
	checkErr(err)
	fmt.Println("最后插入id =", lastInsertId)

	//更新数据
	stmt, err = db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", 1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("影响行数 =", affect)

	//查询数据 返回多行结果
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	//遍历结果集
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string

		// 将当前行的列值存入变量，映射到 Go 的变量类型。
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)

		//删除数据
		stmt, err = db.Prepare("delete from userinfo where uid=$1")
		checkErr(err)

		res, err = stmt.Exec(1)
		checkErr(err)

		affect, err = res.RowsAffected()
		checkErr(err)

		fmt.Println(affect)

		db.Close()
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
