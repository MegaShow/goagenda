# Go Agenda

Agenda is a meeting management system.

[![Build Status](https://travis-ci.org/MegaShow/goagenda.svg?branch=master)](https://travis-ci.org/MegaShow/goagenda)
[![Coverage Status](https://coveralls.io/repos/github/MegaShow/goagenda/badge.svg)](https://coveralls.io/github/MegaShow/goagenda)
[![CodeFactor](https://www.codefactor.io/repository/github/megashow/goagenda/badge)](https://www.codefactor.io/repository/github/megashow/goagenda)
[![Teambition](https://img.shields.io/badge/teambition-tasks-ff69b4.svg)](https://www.teambition.com/project/5bc6ffbaf10ae90018184bd0/)
[![GoDoc](https://godoc.org/github.com/MegaShow/goagenda?status.svg)](https://godoc.org/github.com/MegaShow/goagenda)

\[English\]  \[[中文](README_zh.md)\]  \[[Contributing](CONTRIBUTING.md)\]

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

### List users

### Create a meeting

```
$ agenda meeting create -t <title> -p <participator> -s <startTime> -e <endTime>
```

- If you want to add more than one participator, please input like `-p p1,p2,...,pN` .
- Input time like `year-month-day-hour-minute` .

### Set information of meeting

```
$ agenda meeting set -t <title> [-p <participator>] [-s <startTime>] [-e <endTime>]
```

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

