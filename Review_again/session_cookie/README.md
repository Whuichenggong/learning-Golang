## session和cookie

### 前言： 

在浏览需要认证的服务页面以及页面统计中相当关键

需要了解“登录”过程中到底发生了什么。

当用户来到微博登录页面，输入用户名和密码之后点击“登录”后浏览器将认证信息POST给远端的服务器，服务器执行验证逻辑，如果验证通过，则浏览器会跳转到登录用户的微博首页，在登录成功后，服务器如何验证我们对其他受限制页面的访问呢？因为**HTTP协议是无状态的**，所以很显然服务器不可能知道我们已经在上一次的HTTP请求中通过了验证。当然，**最简单的解决方案就是所有的请求里面都带上用户名和密码**，这样虽然可行，但大大加重了服务器的负担（对于每个request都需要到数据库验证），也大大降低了用户体验(每个页面都需要重新输入用户名密码，每个页面都带有登录表单)。既然直接在请求中带上用户名与密码不可行，那么就只有在服务器或客户端**保存一些类似的可以代表身份的信息了**，所以就有了**cookie与session。**

## cookie：

简而言之就是在**本地计算机**保存一些用户操作的历史信息（当然包括登录信息），并在用户**再次访问该站点时浏览器通过HTTP协议将本地cookie内容发送给服务器**，从而完成验证，或继续上一步操作。

浏览器端保存的cookie信息 

cookie是有时间限制的，根据生命期不同分成两种：会话cookie和持久cookie；

### Go 设置 Cookie:

```go
http.SetCookie(w ResponseWriter, cookie *Cookie)
```

```go
expiration := time.Now()
expiration = expiration.AddDate(1, 0, 0)
cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
http.SetCookie(w, &cookie)
``` 

### Go 获取 Cookie:

1.

```go
cookie, _ := r.Cookie("username")
fmt.Fprint(w, cookie)
```

2.
```go
for _, cookie := range r.Cookies() {
	fmt.Fprint(w, cookie.Name)
}
```


## Session：

在**服务器**上保存用户操作的历史信息。服务器使用session id来标识session，session id由服务器负责产生，保证**随机性与唯一性**，相当于一个随机密钥，避免在握手或传输中暴露用户真实密码。但该方式下，仍然需要将发送请求的客户端与session进行对应，所以可以借助cookie机制来获取客户端的标识（即session id），也可以通过GET方式将id提交给服务器。


## 例子：
当你登录一个网站时，服务器为你创建了一个 Session，并在 Cookie 中存储一个 Session ID。每次你访问网站时，浏览器会自动将 Cookie 中的 Session ID 发送给服务器，服务器使用该 Session ID 查找对应的 Session 数据（例如用户信息），从而知道你已经登录，不需要重新输入密码。

目的都是： 都是为了克服http协议无状态的缺陷