# Go Agenda

Agenda是一个基于命令行的会议管理系统。

[![Build Status](https://travis-ci.org/MegaShow/goagenda.svg?branch=master)](https://travis-ci.org/MegaShow/goagenda)
[![Coverage Status](https://coveralls.io/repos/github/MegaShow/goagenda/badge.svg)](https://coveralls.io/github/MegaShow/goagenda)
[![CodeFactor](https://www.codefactor.io/repository/github/megashow/goagenda/badge)](https://www.codefactor.io/repository/github/megashow/goagenda)
[![Teambition](https://img.shields.io/badge/teambition-tasks-ff69b4.svg)](https://www.teambition.com/project/5bc6ffbaf10ae90018184bd0/)
[![GoDoc](https://godoc.org/github.com/MegaShow/goagenda?status.svg)](https://godoc.org/github.com/MegaShow/goagenda)

\[[English](README.md)\]  [中文]  \[[Contributing](CONTRIBUTING.md)\] \[[Test Case](test/TEST_CASE.md)\]

## 快速入门

下载源码。

```sh
$ go get -u github.com/MegaShow/goagenda
```

编译源码。

```sh
$ cd $GOPATH/src/github.com/MegaShow/goagenda
$ go build -o agenda
```

复制粘贴二进制文件和配置文件。

```sh
$ mkdir -p /opt/agenda
$ cp agenda /opt/agenda/agenda
$ mkdir -p ~/agenda
$ chmod -R 777 ~/agenda
$ cp .agenda.yaml ~/agenda/.agenda.yaml
```

创建软链接。

```sh
$ ln -S /usr/bin/agenda /opt/agenda/agenda
```

测试。

```sh
$ agenda
```

## 用户手册

### 注册

```
$ agenda register -u <user> -p <password> [-e <email>] [-t <telephone>]
```

- 你可以使用`r`或`reg`代替`register`。
- 用户名不得超过`32`位，并且必须以字母开头，只能包括字母、数字、下划线。
- 密码不得少于`6`位，不得超过`64`位，除此之外无其他要求。

### 登入

```
$ agenda login -u <user> -p <password>
```

- 你可以使用`l`或`li`代替`login`。
- 你应该处于未登录状态。

### 登出

```
$ agenda logout
```

- 你可以使用`lo`代替`logout`。
- 你应该处于登录状态。

### 显示当前状态

```
$ agenda status
```

- 你可以使用`s`代替`status`。
- 显示你登录与否。

### 修改用户信息

```
$ agenda user set [-p <password>] [-e <email>] [-t <telephone>]
```

- 你可以使用`u`代替`user`。
- 你可以使用`s`代替`set`。
- 密码不能为空。
- 邮件地址和电话可以设为空。请输入 `-e ""` 或 `-t ""` 以表示置空。

### 删除用户

```
$ agenda user delete -u <user> -p <password>
```

- 你可以使用`d`代替`delete`。
- 你输入用户名和密码必须是你的现在登陆的用户名和密码。
- 你将会删除你自己的账户，删除用户后，你将登出改账户。
- 你发起的会议将被删除，你参与的会议将会从与会者列表中移除你。
- 如果你退出会议之后，该会议的参与人数为0，那该会议将被删除。

### 显示所有用户

```
$ agenda user list [-u <user>]
```
 - 你可以使用`l`代替`list`。
 - 你将会获得你输入的用户的详情（包括用户名、邮箱和联系电话）。
 - 如果你不输入用户名称，你将会获得所有用户的详情。

### 创建会议

```
$ agenda meeting create -t <title> -s <startTime> -e <endTime> -p <participators>
```

- 你可以使用 `meet` 或 `m` 代替 `meeting` 。
- 你可以使用 `c` 代替 `create` 。
- 如果你想添加多个与会者，请像这样输入：`-p p1,p2,...,pN` 。
- 你可以这样输入时间： `YYYY-MM-DD/hh:mm` 或 `YYYY-M-D/h:m` ，使用24小时制。

### 修改会议信息

```
$ agenda meeting set -t <title> [-s <startTime>] [-e <endTime>] [-p <participators>]
```

- 你可以使用 `s` 代替 `set` 。

### 添加与会者

```sh
$ agenda meeting add -t <title> participators
```

- 你可以用 `a` 代替 `add`。 
- 如果你想要添加多个与会者，请像这样输入： `p1 p2 p3 ...` 。
- 你必须是这个会议的发起者。
- 如果你添加了原来会议就存在的成员，你将不会收到错误信息，但是这个人不会再次加入到会议中。

### 移除与会者

```
$ agenda meeting remove -t <title> <participators list>
```

* 你可以使用`r`代替`remove`。
* 你必须是该会议的发起者。
* 一旦列表中存在非法的用户名(不存在或不是该会议的与会者)，移除操作将被取消。
* 如果你移除与会者之后，该会议的参与人数为0，那该会议将被删除。

### 删除会议

```
$ agenda meeting delete [-t <title> | -a]
```

* 你可以使用`d`代替`delete`。
* 你必须是该会议的发起者。
* 当设置了`-a`时，你所发起的所有会议都会被删除。

###  退出会议

```
$ agenda meeting quit -t <title>
```

* 你可以使用`q`代替`quit`。
* 你不能是该会议的发起者。
* 如果你退出会议之后，该会议的参与人数为0，那该会议将被删除。

### 查询会议

```
$ agenda meeting list [-t <title>] [-s <startTime>] [-e <endTime>]
```

* 你可以使用`l`代替`list`。
* 该命令只会列出你发起或参与的会议。
* 列表将根据会议开始时间排序。

### 打印日志

```
$ agenda log
```

- 打印Log信息，如果`Log.IsOpen`被配置为`true`。
- **这是一个糟糕的命令，我们正在考虑是否弃用它。**

## 协议

使用Apache License 2.0协议。

