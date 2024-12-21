# database/sql接口

为开发数据库驱动定义的一些标准接口，根据定义的接口开发响应数据库驱动，按照标准接口开发的代码，迁移数据库的时候，不用做任何修改

## sql.Register
1. 注册数据库驱动
2. 第三方数据库驱动会实现init函数，在init函数中调用sql.Register注册驱动


~~~go
func init() {
	sql.Register("sqlite3", &SQLiteDriver{})
}
~~~

database/sql内部通过一个map来存储用户定义的相应驱动

~~~go
var drivers = make(map[string]driver.Driver)

drivers[name] = driver
~~~

在项目中引用第三方库时时常会看到这个：

~~~go
import _ "github.com/lib/pq"
~~~

是用来忽略变量赋值的占位符，那么包引入用到这个符号也是相似的作用，这儿使用_的意思是引入后面的包名而不直接使用这个包中定义的函数，变量等资源。


---

# 回顾之前篇章：


# 1.流程和函数

##  main函数和init函数

Go里面有两个保留的函数：init函数（能够应用于所有的package）和main函数（只能应用于package main）
这两个函数定义时不能有任何的参数和返回值

被建议： 用户在一个package中每个文件只写一个init函数！
Go程序会自动调用init()和main() 

[![image.png]](image.png)

## import

fmt是Go语言的标准库，其实是去GOROOT环境变量指定目录下去加载该模块，当然Go的import还支持如下两种方式来**加载自己写的模块**

1. 相对路径

```go
import “./model” //当前文件同一目录的model目录，但是不建议这种方式来import
```

2. 绝对路径

```go
import “/home/user/go/src/myproject/model” //绝对路径
```

3. 点操作

```go
import(
     . "fmt"
 )
```

点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，也就是前面你调用的fmt.Println("hello world")可以省略的写成Println("hello world")


4. 别名

```go
import (
    f "fmt"
)

func main() {
    f.Println("hello world")
}
```

回到本篇章：

5. _操作

```go_ "github.com/ziutek/mymysql/godrv"


_操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。

---


## driver.Driver

Driver是一个数据库驱动的接口，他定义了一个method： Open(name string)，这个方法返回一个数据库的Conn接口

```go
type Driver interface {
	Open(name string) (Conn, error)
}
```

返回的Conn只能用来进行一次goroutine的操作，也就是说不能把这个Conn应用于Go的多个goroutine里面。

错误：

```go
...
go goroutineA (Conn)  //执行查询操作
go goroutineB (Conn)  //执行插入操作
...
```
Go不知道某个操作究竟是由哪个goroutine发起的,从而导致数据混乱

## driver.Conn

数据库连接的接口定义，这个Conn只能应用在一个goroutine里面，不能使用在多个goroutine里面，详情请参考上面的说明。

```go
type Conn interface {
	Prepare(query string) (Stmt, error)
	Close() error
	Begin() (Tx, error)
}
```

Prepare函数返回与当前连接相关的执行Sql语句的准备状态，可以进行查询、删除等操作。

Close函数关闭当前的连接，执行释放连接拥有的资源等清理工作。因为驱动实现了database/sql里面建议的conn pool，所以你不用再去实现缓存conn之类的，这样会容易引起问题。

Begin函数开启一个事务，返回一个Tx接口，用来执行事务相关的操作。

## driver.Stmt

Stmt是一种准备好的状态，和Conn相关联，而且只能应用于一个goroutine中，不能应用于多个goroutine。

```go
type Stmt interface {
	Close() error
	NumInput() int
	Exec(args []Value) (Result, error)
	Query(args []Value) (Rows, error)
}
```

Close函数关闭当前的Stmt，执行释放Stmt拥有的资源等清理工作。

NumInput函数返回当前Stmt需要的参数个数。

Exec函数执行一个修改操作，返回一个Result接口，用来获取执行结果。

Query函数执行一个查询操作，返回一个Rows接口，用来获取查询结果。

## driver.Tx

Tx是事务的接口，用来执行事务相关的操作。

```go
type Tx interface {
	Commit() error
	Rollback() error
	Prepare(query string) (Stmt, error)
}
```

Commit函数提交当前事务，如果提交成功，则返回nil，否则返回错误。

Rollback函数回滚当前事务，如果回滚成功，则返回nil，否则返回错误。

Prepare函数和Conn的Prepare函数类似，用来准备执行Sql语句的状态。

## driver.Tx

事务处理一般就两个过程，递交或者回滚。数据库驱动里面也只需要实现这两个函数就可以

```go
type Tx interface {
	Commit() error
	Rollback() error
}
```
这两个函数一个用来递交一个事务，一个用来回滚事务。

# driver.Execer

Conn可选择实现的接口

```go
type Execer interface {
	Exec(query string, args []Value) (Result, error)
}
```
如果这个接口没有定义，那么在调用DB.Exec,就会首先调用Prepare返回Stmt，然后执行Stmt的Exec，然后关闭Stmt。

 ## driver.Result

这个是执行Update/Insert等操作返回的结果接口定义

 ```go
 type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}
```

LastInsertId函数返回由数据库执行插入操作得到的自增ID号。

RowsAffected函数返回执行Update/Insert等操作影响的数据条目数。

## driver.Rows

```go
type Rows interface {
	Columns() []string
	Close() error
	Next(dest []Value) error
}
```
Columns函数返回查询数据库表的字段信息，这个返回的slice和sql查询的字段一一对应，而不是返回整个表的所有字段。

Close函数用来关闭Rows迭代器。

Next函数用来返回下一条数据，把数据赋值给dest。dest里面的元素必须是driver.Value的值除了string，返回的数据里面所有的string都必须要转换成[]byte。如果最后没数据了，Next函数最后返回io.EOF。

## driver.RowsAffected

RowsAffected其实就是一个int64的别名，但是他实现了Result接口，用来底层实现Result的表示方式

```go
type RowsAffected int64

func (RowsAffected) LastInsertId() (int64, error)

func (v RowsAffected) RowsAffected() (int64, error)
```

## driver.Value

```go
type Value interface{}
```

drive的Value是驱动必须能够操作的Value，Value要么是nil，要么是下面的任意一种

```go
int64
float64
bool
[]byte
string   [*]除了Rows.Next返回的不能是string.
time.Time
```

driver.ValueConverter

```go
type ValueConverter interface {
	ConvertValue(v interface{}) (Value, error)
}
```

这个ValueConverter有很多好处：

转化driver.value到数据库表相应的字段，例如int64的数据如何转化成数据库表uint16字段

把数据库查询结果转化成driver.Value值

在scan函数里面如何把driver.Value值转化成用户定义的值

## driver.Valuer

```go
type Valuer interface {
	Value() (Value, error)
}
```

## database/sql

database/sql在database/sql/driver提供的接口基础上定义了一些更高阶的方法，用以简化数据库操作,同时内部还建议性地实现一个conn pool。

```go
type DB struct {
	driver 	 driver.Driver
	dsn    	 string
	mu       sync.Mutex // protects freeConn and closed
	freeConn []driver.Conn
	closed   bool
}
```
Open函数返回的是DB对象，里面有一个freeConn，它就是那个简易的连接池。它的实现相当简单或者说简陋，就是当执行db.prepare -> db.prepareDC的时候会defer dc.releaseConn，然后调用db.putConn，也就是把这个连接放入连接池，每次调用db.conn的时候会先判断freeConn的长度是否大于0，大于0说明有可以复用的conn，直接拿出来用就是了，如果不大于0，则创建一个conn，然后再返回之。



1. 我的疑惑：

Go语言的database/sql库本身只提供了通用的接口，而具体的数据库驱动需要我们手动导入。 这句话时什么意思 什么是提供通用接口，数据库驱动又是什么，这有什么区别

2. 回答：

什么是通用接口
通用接口是指 database/sql 提供的一组方法和类型，这些方法和类型不依赖于具体的数据库种类，具有通用性。例如：

通用类型：
*sql.DB：表示数据库连接池。
*sql.Rows：表示查询返回的多行结果集。
*sql.Stmt：表示预处理语句。
通用方法：
Query(query string, args ...interface{})：执行查询语句。
Exec(query string, args ...interface{})：执行非查询语句（如 INSERT、UPDATE）。
Begin()：开始一个事务。

可以用统一的方式操作各种数据库，而不需要针对每种数据库写不同的代码。

为了使用具体数据库的功能，必须手动导入相应的驱动，这样 database/sql 才能调用驱动实现的具体功能。









