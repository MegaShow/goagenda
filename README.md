# Go Agenda

Agenda is a meeting management system.

[![Build Status](https://travis-ci.org/MegaShow/goagenda.svg?branch=master)](https://travis-ci.org/MegaShow/goagenda)
[![Coverage Status](https://coveralls.io/repos/github/MegaShow/goagenda/badge.svg)](https://coveralls.io/github/MegaShow/goagenda)
[![CodeFactor](https://www.codefactor.io/repository/github/megashow/goagenda/badge)](https://www.codefactor.io/repository/github/megashow/goagenda)
[![Teambition](https://img.shields.io/badge/teambition-tasks-ff69b4.svg)](https://www.teambition.com/project/5bc6ffbaf10ae90018184bd0/)
[![GoDoc](https://godoc.org/github.com/MegaShow/goagenda?status.svg)](https://godoc.org/github.com/MegaShow/goagenda)

\[English\]  \[[中文](README_zh.md)\]  \[[Contributing](CONTRIBUTING.md)\] \[[Test Case](test/TEST_CASE.md)\]

## Quick Start

Coming soon.

## User manual

### Register

### Login/Logout

### Show status

### Set information of user

```
$ agenda user set [-p <password>] [-e <email>] [-t <telephone>]
```

- The password cannot be empty.
- The email address and telephone can be set empty. Please input  `-e ""` or `-t ""` to indicate it.

### Delete user

```
$ agenda user delete -u <user> -p <password>
```

- You should login.
- Your account will be deleted，and you will logout.

### List users

```
$ agenda user list [-u <user>]
```

 - List the details of the user you input
 - If you don't input the username, you will get all users' information.

### Create a meeting

```
$ agenda meeting create -t <title> -s <startTime> -e <endTime> -p <participators> 
```

- You can use `meet` or `m` as aliases for `meeting`.
- You can use `c` as an alias for `create`.
- If you want to add more than one participator, please input like `-p p1,p2,...,pN` .
- Input time like `YYYY-MM-DD/hh:mm` or `YYYY-M-D/h:m` , using 24-hour.

### Set information of meeting

```
$ agenda meeting set -t <title> [-s <startTime>] [-e <endTime>] [-p <participators>]
```

- You can use `s` as an alias for `set`.

### Add or remove participator

### Delete a meeting

### Quit a meeting

```
$ agenda meeting quit -t <title>
```

### List meetings

### Print log

## License

Under Apache License 2.0.

