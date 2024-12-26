package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type User struct {
	ID         int `orm:"pk"`
	Name       string
	Departname string
	Created    time.Time
	Age        int16   // 将 Profile.Age 添加到 User 中
	Posts      []*Post `orm:"reverse(many)"`
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`
	Tags  []*Tag `orm:"rel(m2m)"` //设置一对多关系
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	// 设置默认数据库
	orm.RegisterDataBase("default", "postgres", "user=root password=zzh15937 dbname=postgres host=127.0.0.1 port=5433 sslmode=disable")
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Post), new(Tag))
	orm.Debug = true
}

func main() {

	o := orm.NewOrm()
	var user User
	user.Name = "赵忠鹤"
	user.Departname = "后端开发"

	id, err := o.Insert(&user)
	if err == nil {
		fmt.Println(id)
	}
}

// func main() {
// 	db, err := sql.Open("postgres", "user=astaxie password=astaxie dbname=test sslmode=disable")
// 	checkErr(err)

// 	//插入数据
// 	stmt, err := db.Prepare("INSERT INTO userinfo(username,department,created) VALUES($1,$2,$3) RETURNING uid")
// 	checkErr(err)

// 	//执行准备好的语句，将占位符替换为实际值。
// 	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
// 	checkErr(err)

// 	//pg不支持这个函数，因为他没有类似MySQL的自增ID
// 	// id, err := res.LastInsertId()
// 	// checkErr(err)
// 	// fmt.Println(id)

// 	//执行语句并获取返回值
// 	var lastInsertId int
// 	err = db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertId)
// 	checkErr(err)
// 	fmt.Println("最后插入id =", lastInsertId)

// 	//更新数据
// 	stmt, err = db.Prepare("update userinfo set username=$1 where uid=$2")
// 	checkErr(err)

// 	res, err = stmt.Exec("astaxieupdate", 1)
// 	checkErr(err)

// 	affect, err := res.RowsAffected()
// 	checkErr(err)

// 	fmt.Println("影响行数 =", affect)

// 	//查询数据 返回多行结果
// 	rows, err := db.Query("SELECT * FROM userinfo")
// 	checkErr(err)

// 	//遍历结果集
// 	for rows.Next() {
// 		var uid int
// 		var username string
// 		var department string
// 		var created string

// 		// 将当前行的列值存入变量，映射到 Go 的变量类型。
// 		err = rows.Scan(&uid, &username, &department, &created)
// 		checkErr(err)
// 		fmt.Println(uid)
// 		fmt.Println(username)
// 		fmt.Println(department)
// 		fmt.Println(created)

// 		//删除数据
// 		stmt, err = db.Prepare("delete from userinfo where uid=$1")
// 		checkErr(err)

// 		res, err = stmt.Exec(1)
// 		checkErr(err)

// 		affect, err = res.RowsAffected()
// 		checkErr(err)

// 		fmt.Println(affect)

// 		db.Close()
// 	}
// }

// func checkErr(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
