# Go Agenda

Agenda is a meeting management system.

[![Build Status](https://travis-ci.org/MegaShow/goagenda.svg?branch=master)](https://travis-ci.org/MegaShow/goagenda)
[![Coverage Status](https://coveralls.io/repos/github/MegaShow/goagenda/badge.svg)](https://coveralls.io/github/MegaShow/goagenda)
[![CodeFactor](https://www.codefactor.io/repository/github/megashow/goagenda/badge)](https://www.codefactor.io/repository/github/megashow/goagenda)
[![Teambition](https://img.shields.io/badge/teambition-tasks-ff69b4.svg)](https://www.teambition.com/project/5bc6ffbaf10ae90018184bd0/)
[![GoDoc](https://godoc.org/github.com/MegaShow/goagenda?status.svg)](https://godoc.org/github.com/MegaShow/goagenda)

\[English\]  \[[中文](README_zh.md)\]  \[[Contributing](CONTRIBUTING.md)\] \[[Test Case](test/TEST_CASE.md)\]

## Quick Start

Install agenda.

```sh
$ go install github.com/MegaShow/goagenda
```

Create soft link.

```sh
$ ln -S /usr/bin/agenda $GOPATH/bin/goagenda
```

Or rename the binary file.

```sh
$ mv $GOPATH/bin/goagenda $GOPATH/bin/agenda
```

Initial the config file while first running.

```sh
$ agenda
```

Test.

```sh
$ agenda
```

## User manual

### Register

```
$ agenda register -u <user> -p <password> [-e <email>] [-t <telephone>]
```

* You can use `r`  or `reg` as an alias for `register`.
* The length of user's name must be between `1` and `32`, and name must be start with letter, including only letters, digits and underline.
* The length of password must be between `6` and `64`, and password has no other limits.

### Login

```
$ agenda login -u <user> -p <password>
```

* You can use `l` or `li` as an alias for `login`.
* You should be in the not-logged state. 

### Logout

```
$ agenda logout
```

* You can use `lo` as an alias for `logout`.
* You should login.

### Show status

```
$ agenda status
```

* You can use `s` as an alias for `status`.
* Show if you are logged or not.

### Set information of user

```
$ agenda user set [-p <password>] [-e <email>] [-t <telephone>]
```

- You can use `u` as an alias for `user`.
- You can use `s` as an alias for `set`.
- The password cannot be empty.
- The email address and telephone can be set empty. Please input  `-e ""` or `-t ""` to indicate it.

### Delete user

```
$ agenda user delete -u <user> -p <password>
```

- You can use `d` as an alias for `delete`.
- You should verify your user's name and password.
- Your account will be deleted，and you will logout.
- The meetings you initiate will be deleted, and the meetings you participate will remove you in the list of participators.
- If this meeting has no participator after you remove participators, it will be deleted.

### List users

```
$ agenda user list [-u <user>]
```

 - You can use `l` as an alias for `list`.
 - List the details of the user you input.
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

### Add participators

```sh
$ agenda meeting add -t <title> participators
```

- You can use `a` as an alias for `add`.
- If you want to add more than one participator, please input like `p1 p2 p3 ...` .
- You must be the initiator of this meeting.
- If you add the participator who has been added into the meeting, you will not get the error message but this participator will not be add to the meeting twice.

### Remove participators

```
$ agenda meeting remove -t <title> <participators list>
```

* You can use `r` as an alias for `remove`.
* You must be the initiator of this meeting.
* Once list includes invalid participator name (non-exist or no participator of this meeting), remove operating will be cancelled.
* If this meeting has no participator after you remove participators, it will be deleted.

### Delete a meeting

```
$ agenda meeting delete [-t <title> | -a]
```

* You can use `d` as an alias for `delete`.
* You must be the initiator of this meeting.
* When set `-a`, all the meetings you initiate will be deleted.

### Quit a meeting

```
$ agenda meeting quit -t <title>
```

* You can use `q` as an alias for `quit`.
* You can't be the initiator of this meeting.
* If this meeting has no participator after you quit, it will be deleted.

### List meetings

```
$ agenda meeting list [-t <title>] [-s <startTime>] [-e <endTime>]
```

* You can use `l` as an alias for `list`.
* It will only list the meetings you initiate or participate.
* The list will be sorted by start time.

### Print log

```
$ agenda log
```

* Print log information while `Log.IsOpen` is configured as `true`.
* **It is a bad command, and it's considered whether to be deprecated.**

## License

Under Apache License 2.0.

