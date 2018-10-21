# Go Agenda

Agenda是一个基于命令行的会议管理系统。

[![Build Status](https://camo.githubusercontent.com/d80b55e7e52b05b91016002052d2de3f144efaf2/68747470733a2f2f7472617669732d63692e6f72672f4d65676153686f772f676f6167656e64612e7376673f6272616e63683d6d6173746572)](https://travis-ci.org/MegaShow/goagenda) [![Coverage Status](https://camo.githubusercontent.com/fc80330179322c05a7040c0c66ac2767f7014bc9/68747470733a2f2f636f766572616c6c732e696f2f7265706f732f6769746875622f4d65676153686f772f676f6167656e64612f62616467652e737667)](https://coveralls.io/github/MegaShow/goagenda) [![CodeFactor](https://camo.githubusercontent.com/4a0e71f2baf8df0404196f19a33af516a8237233/68747470733a2f2f7777772e636f6465666163746f722e696f2f7265706f7369746f72792f6769746875622f6d65676173686f772f676f6167656e64612f6261646765)](https://www.codefactor.io/repository/github/megashow/goagenda) [![Teambition](https://camo.githubusercontent.com/0e1dfb705a24405737069292a8d44a4476443c88/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f7465616d626974696f6e2d7461736b732d6666363962342e737667)](https://www.teambition.com/project/5bc6ffbaf10ae90018184bd0/) [![GoDoc](https://camo.githubusercontent.com/6cf0ccd061f613abf1a477893b7050b634576704/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f4d65676153686f772f676f6167656e64613f7374617475732e737667)](https://godoc.org/github.com/MegaShow/goagenda)

[[English](https://github.com/yuyuforest/goagenda/blob/master/README.md)] [中文] [[Contributing](https://github.com/yuyuforest/goagenda/blob/master/CONTRIBUTING.md)]

## 快速入门

Coming soon.

## 用户手册

### 注册

### 登入/登出

### 显示当前状态

### 设置用户信息

```
$ agenda user set [-p <password>] [-e <email>] [-t <telephone>]
```

- 密码不能为空。
- 邮件地址和电话可以为空。请输入""以表示空字符串。

### 删除用户

```
$ agenda user delete
```

- 必须处于登陆状态才能删除
- 你将会删除你自己的账户
- 删除用户后，你将登出改账户

### 显示所有用户

### 创建所有会议

### 修改会议信息

### 添加或移除与会者

### 删除或退出会议

### 显示所有会议

### 打印日志

## 协议

使用Apache License 2.0协议。