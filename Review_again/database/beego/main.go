package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq" // 导入数据库驱动
)

// Model Struct
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	// 设置默认数据库
	orm.RegisterDataBase("default", "postgres", "user=root password=zzh15937 dbname=postgres host=127.0.0.1 port=5433 sslmode=disable")

	// 注册定义的 model
	orm.RegisterModel(new(User))
	//RegisterModel 也可以同时注册多个 model
	//orm.RegisterModel(new(User), new(Profile), new(Post))

	// 创建 table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	user := User{Name: "slene"}

	// 插入表
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// 更新表
	user.Name = "赵忠鹤"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// 读取 one
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// 删除表
	// num, err = o.Delete(&u)
	// fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
