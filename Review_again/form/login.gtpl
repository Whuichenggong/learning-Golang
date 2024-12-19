<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>登录页面</title>
</head>
<body>
    <form action="/login" method="post">
        <label>用户名: <input type="text" name="username"></label><br>
        <label>密码: <input type="password" name="password"></label><br>
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
</body>
</html>
