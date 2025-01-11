<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>登录页面</title>
</head>
<body>

<!-- 显示服务器返回的动态消息 -->
<!-- 双括号这部分代码body的什么位置最终就会显示在网页的什么位置 -->
<!--   隐藏的token字段，用于比对 -->
   
    <form action="/login" method="post">
        <label>用户名: <input type="text" name="username"></label><br>
        <label>密码: <input type="password" name="password"></label><br>
      
        <input type="hidden" name="token" value="{{.Token}}">
        <label> <input type="radio" name="gender" value="1">男</label><br>
        <label> <input type="radio" name="gender" value="2">女</label><br>
        <label>选择水果:
            <select name="fruit">
                <option value="apple">Apple</option>
                <option value="pear">Pear</option>
                <option value="banana">Banana</option>
            </select>
        </label><br>
        <input type="submit" value="登录">
    </form>
     {{if .Message}}
    <div style="color: red;">{{.Message}}</div>
    {{end}}
</body>
</html>
