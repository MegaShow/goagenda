# Go Agenda

Agenda是一个基于命令行的会议管理系统。

[![Build Status](https://travis-ci.org/MegaShow/goagenda.svg?branch=master)](https://travis-ci.org/MegaShow/goagenda)
[![Coverage Status](https://coveralls.io/repos/github/MegaShow/goagenda/badge.svg)](https://coveralls.io/github/MegaShow/goagenda)
[![CodeFactor](https://www.codefactor.io/repository/github/megashow/goagenda/badge)](https://www.codefactor.io/repository/github/megashow/goagenda)
[![Teambition](https://img.shields.io/badge/teambition-tasks-ff69b4.svg)](https://www.teambition.com/project/5bc6ffbaf10ae90018184bd0/)
[![GoDoc](https://godoc.org/github.com/MegaShow/goagenda?status.svg)](https://godoc.org/github.com/MegaShow/goagenda)

\[[English](README.md)\]  [中文]  \[[Contributing](CONTRIBUTING.md)\]

## 快速入门

Coming soon.

## 用户手册

### 注册

### 登入/登出

### 显示当前状态

### 修改用户信息

```
$ agenda user set [-p <password>] [-e <email>] [-t <telephone>]
```

- 密码不能为空。
- 邮件地址和电话可以设为空。请输入 `-e ""` 或 `-t ""` 以表示置空。

### 删除用户

```
$ agenda user delete
```

- 必须处于登陆状态才能删除
- 你将会删除你自己的账户
- 删除用户后，你将登出改账户

### 显示所有用户

### 创建会议

```
$ agenda meeting create -t <title> -p <participator> -s <startTime> -e <endTime>
```

- 如果你想添加多个与会者，请像这样输入：`-p p1,p2,...,pN` 
- 你可以这样输入时间： `year-month-day-hour-minute` 

### 修改会议信息

```
$ agenda meeting set -t <title> [-p <participator>] [-s <startTime>] [-e <endTime>]
```

### 添加或移除与会者

### 删除会议

###  退出会议

```
$ agenda meeting quit -t <title>
```

### 查询会议

### 打印日志

## 协议

使用Apache License 2.0协议。
