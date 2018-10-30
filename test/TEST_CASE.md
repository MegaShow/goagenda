# 测试样例

## Register

## Login

## Logout

## Status

## User/Set

## User/List

## User/Delete

**controller/user**

* 要检测是否登录
* 要检测用户、密码是否合法
* 要检测用户是否为当前用户

**service/user**

* 要检测用户、密码是否正确
* 要检测是否需要删除或退出会议

### 测试命令

```sh
./agenda register -u Amy -p 123456
./agenda r -u Bob -p 654321
./agenda r -u Cici -p 123456
./agenda r -u Duke -p 654321
./agenda r -u Ella -p 123456
./agenda login -u Amy -p 123456
./agenda meeting c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
./agenda meeting c -t me2 -s 2018-10-27/09:00 -e 2018-10-27/11:00 -p Duke
./agenda logout
./agenda login -u Bob -p 654321
./agenda meeting c -t me3 -s 2018-10-28/09:00 -e 2018-10-28/11:00 -p Duke
./agenda logout
./agenda user delete -u Bob -p 654321
./agenda login -u Bob -p 654321
./agenda user delete -u Bob -p 123456
./agenda user delete -u Bob -p 654321
./agenda login -u Ella -p 123456
./agenda u d -u Ella -p 123456
./agenda login -u Cici -p 123456
./agenda u d -u Cici -p 123456
./agenda login -u Amy -p 123456
./agenda u d -u Amy -p 123456
```

### 测试结果

```sh
./agenda register -u Amy -p 123456
./agenda r -u Bob -p 654321
./agenda r -u Cici -p 123456
./agenda r -u Duke -p 654321
./agenda r -u Ella -p 123456
./agenda login -u Amy -p 123456
./agenda meeting c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
./agenda meeting c -t me2 -s 2018-10-27/09:00 -e 2018-10-27/11:00 -p Duke
./agenda logout
./agenda login -u Bob -p 654321
./agenda meeting c -t me3 -s 2018-10-28/09:00 -e 2018-10-28/11:00 -p Duke
./agenda logout

./agenda user delete -u Bob -p 654321
# 失败，未登录

./agenda login -u Bob -p 654321

./agenda user delete -u Bob -p 123456
# 失败，用户、密码不正确

./agenda user delete -u Bob -p 654321
# 成功，删除me3、修改me1

./agenda login -u Ella -p 123456

./agenda u d -u Ella -p 123456
# 成功

./agenda login -u Cici -p 123456

./agenda u d -u Cici -p 123456
# 成功，删除me1

./agenda login -u Amy -p 123456

./agenda u d -u Amy -p 123456
# 成功，删除me3
```

## Meeting/Create

**controller/meeting**

- 要检测是否登录
- 要检测时间、议题是否合法

**service/meeting**

- 要检测议题是否已存在
- 要检测某些与会者是否存在
- 要检测与会者是否有别的重叠会议（检测的与会者也包括会议发起人）
- 要去掉重复的与会者

### 测试命令

```sh
./agenda register -u Amy -p 123456
./agenda r -u Bob -p 654321
./agenda r -u Cici -p 123456
./agenda r -u Duke -p 654321
./agenda r -u Ella -p 123456
./agenda meeting c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/08:00 -p Bob,Cici
./agenda login -u Amy -p 123456
./agenda meeting c -t me0 -s "" -e 2018-10-26/08:00 -p Bob,Cici
./agenda meeting c -t me0 -s 2018-10-26/09:00 -e 2018-10-26-08-00 -p Bob,Cici
./agenda meeting c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/08:00 -p Bob,Cici
./agenda meet c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/09:00 -p Bob,Cici
./agenda m c -t "" -s 2018-10-26/09:00 -e 2018-10-26/10:00 -p Bob,Cici
./agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/10:00 -p Bob,Frank
./agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
./agenda m c -t me1 -s 2018-10-26/13:00 -e 2018-10-26/15:00 -p Bob,Cici
./agenda m c -t me2 -s 2018-10-26/08:00 -e 2018-10-26/9:30 -p Bob,Duke
./agenda m c -t me2 -s 2018-10-26/08:00 -e 2018-10-26/9:00 -p Bob,Duke,Duke,Amy
./agenda logout
./agenda login -u Duke -p 654321
./agenda m c -t me3 -s 2018-10-26/10:5 -e 2018-10-26/12:00 -p Amy,Ella
./agenda m c -t me3 -s 2018-10-26/8:7 -e 2018-10-26/12:00 -p Ella
./agenda m c -t me3 -s 2018-10-27/16:00 -e 2018-10-27/17:00 -p Cici,Ella
./agenda m c -t me4 -s 2018-10-26/9:30 -e 2018-10-26/10:30 -p Ella
./agenda logout
```

### 测试结果

```sh
./agenda register -u Amy -p 123456

./agenda r -u Bob -p 654321

./agenda r -u Cici -p 123456

./agenda r -u Duke -p 654321

./agenda r -u Ella -p 123456

./agenda meeting c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/08:00 -p Bob,Cici
# 不成功，因为未登陆

./agenda login -u Amy -p 123456

./agenda meeting c -t me0 -s "" -e 2018-10-26/08:00 -p Bob,Cici
# 不成功，因为输入的时间为空

./agenda meeting c -t me0 -s 2018-10-26/09:00 -e 2018-10-26-08-00 -p Bob,Cici
# 不成功，因为输入的时间格式非法

./agenda meeting c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/08:00 -p Bob,Cici
# 不成功，因为时间顺序不对（开始时间小于结束时间）

./agenda meet c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/09:00 -p Bob,Cici
# 不成功，因为时间顺序不对（开始时间和结束时间相等）

./agenda m c -t "" -s 2018-10-26/09:00 -e 2018-10-26/10:00 -p Bob,Cici
# 不成功，因为议题不合法

./agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/10:00 -p Bob,Frank
# 不成功，因为用户Frank不存在

./agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
# 成功

./agenda m c -t me1 -s 2018-10-26/13:00 -e 2018-10-26/15:00 -p Bob,Cici
# 不成功，因为议题me1已存在

./agenda m c -t me2 -s 2018-10-26/08:00 -e 2018-10-26/9:30 -p Bob,Duke
# 不成功，因为Amy和Bob有重叠的会议me1

./agenda m c -t me2 -s 2018-10-26/08:00 -e 2018-10-26/9:00 -p Bob,Duke,Duke,Amy
# 成功
# （检测一场会议的endTime与另一场会议的startTime重叠的情况：me2结束时间与me1开始时间重叠）
# （同时检测与会者去重的情况，在data/meeting.json里查看，结果应该是Bob,Duke）

./agenda logout

./agenda login -u Duke -p 654321

./agenda m c -t me3 -s 2018-10-26/10:5 -e 2018-10-26/12:00 -p Amy,Ella
# 不成功，因为Amy正在主持另一会议me1

./agenda m c -t me3 -s 2018-10-26/8:7 -e 2018-10-26/12:00 -p Ella
# 不成功，因为当前用户Duke正在开另一会议me2

./agenda m c -t me3 -s 2018-10-27/16:00 -e 2018-10-27/17:00 -p Cici,Ella
# 成功

./agenda m c -t me4 -s 2018-10-26/9:30 -e 2018-10-26/10:30 -p Ella
# 成功
# （检测有时间段重叠、但人员不相干的会议的情况）

./agenda logout
```

## Meeting/Set

## Meeting/Add

## Meeting/Remove

**controller/meeting**

- 要检测是否登录
- 要检测议题是否合法

**service/meeting**

- 要检测会议是否存在
- 要检测当前用户是否为会议发起者
- 要检测用户列表是否合法

### 测试命令

```sh
./agenda register -u Amy -p 123456
./agenda r -u Bob -p 654321
./agenda r -u Cici -p 123456
./agenda r -u Duke -p 654321
./agenda r -u Ella -p 123456
./agenda login -u Amy -p 123456
./agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici,Ella
./agenda logout
./agenda m remove -t me1 Bob
./agenda login -u Bob -p 654321
./agenda m r -t me1 Bob
./agenda logout
./agenda login -u Amy -p 123456
./agenda m r -t me2 Bob
./agenda m r -t me1 Bob Duke
./agenda m r -t me1 Bob Frank
./agenda m r -t me1 Bob Ella
./agenda m r -t me1 Cici
```

### 测试结果

```sh
./agenda register -u Amy -p 123456
./agenda r -u Bob -p 654321
./agenda r -u Cici -p 123456
./agenda r -u Duke -p 654321
./agenda r -u Ella -p 123456
./agenda login -u Amy -p 123456
./agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici,Ella
./agenda logout

./agenda m remove -t me1 Bob
# 失败，未登录

./agenda login -u Bob -p 654321

./agenda m r -t me1 Bob
# 失败，不是发起者

./agenda logout
./agenda login -u Amy -p 123456

./agenda m r -t me2 Bob
# 失败，会议不存在

./agenda m r -t me1 Bob Duke
# 失败，Duke不是与会者

./agenda m r -t me1 Bob Frank
# 失败，Frank不存在

./agenda m r -t me1 Bob Ella
# 成功

./agenda m r -t me1 Cici
# 成功，并删除会议
```

## Meeting/Delete

**controller/meeting**

- 要检测是否登录
- 要检测议题是否合法

**service/meeting**

* 要检测会议是否存在
* 要检测当前用户是否为会议发起者

### 测试命令

```sh
./agenda register -u Amy -p 123456
./agenda r -u Bob -p 654321
./agenda r -u Cici -p 123456
./agenda r -u Duke -p 654321
./agenda r -u Ella -p 123456
./agenda login -u Amy -p 123456
./agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
./agenda logout
./agenda m delete -t me1
./agenda login -u Bob -p 654321
./agenda m delete -t me1
./agenda m d -t me2
./agenda logout
./agenda login -u Amy -p 123456
./agenda m d -t me1
./agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
./agenda m c -t me2 -s 2018-10-27/09:00 -e 2018-10-27/11:00 -p Bob,Cici
./agenda m d -a -t me1
./agenda m d -a
```

### 测试结果

```sh
./agenda register -u Amy -p 123456
./agenda r -u Bob -p 654321
./agenda r -u Cici -p 123456
./agenda r -u Duke -p 654321
./agenda r -u Ella -p 123456
./agenda login -u Amy -p 123456
./agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
./agenda logout

./agenda m delete -t me1
# 失败，用户未登录

./agenda login -u Bob -p 654321

./agenda m delete -t me1
# 失败，用户不是会议发起者

./agenda m d -t me2
# 失败，找不到该会议

./agenda logout
./agenda login -u Amy -p 123456

./agenda m d -t me1
# 成功

./agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
./agenda m c -t me2 -s 2018-10-27/09:00 -e 2018-10-27/11:00 -p Bob,Cici

./agenda m d -a -t me1
# 失败，不能同时指定-a和-t

./agenda m d -a
# 成功，删除两个会议
```

## Meeting/Quit

**controller/meeting**

- 要检测是否登录
- 要检测议题是否合法

**service/meeting**

- 要检测会议是否存在
- 要检测当前用户是否为会议发起者，发起者不能退出会议
- 要检测当前用户是否为会议参与者
- 要检测当前用户退出会议之后会议参与人数是否为0，是否需要删除会议

### 测试命令

```sh
./agenda register -u Amy -p 123456
./agenda r -u Bob -p 654321
./agenda r -u Cici -p 123456
./agenda r -u Duke -p 654321
./agenda r -u Ella -p 123456
./agenda login -u Amy -p 123456
./agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
./agenda logout
./agenda m quit -t me1
./agenda login -u Bob -p 654321
./agenda m quit -t me1
./agenda m q -t me1
./agenda m q -t me2
./agenda logout
./agenda login -u Amy -p 123456
./agenda m q -t me1
./agenda logout
./agenda login -u Cici -p 123456
./agenda m q -t me1
```

### 测试结果

```sh
./agenda register -u Amy -p 123456
./agenda r -u Bob -p 654321
./agenda r -u Cici -p 123456
./agenda r -u Duke -p 654321
./agenda r -u Ella -p 123456
./agenda login -u Amy -p 123456
./agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
./agenda logout

./agenda m quit -t me1
# 失败，未登录

./agenda login -u Bob -p 654321

./agenda m quit -t me1
# 成功

./agenda m q -t me1
# 失败，不是会议参与者

./agenda m q -t me2
# 失败，不存在该会议

./agenda logout
./agenda login -u Amy -p 123456

./agenda m q -t me1
# 失败，为会议发起者

./agenda logout
./agenda login -u Cici -p 123456

./agenda m q -t me1
# 成功，并删除会议
```

## Meeting/List

**controller/meeting**

- 要检测是否登录
- 要检测议题是否合法
- 要检测时间是否合法

### 测试命令

```sh
./agenda register -u Amy -p 123456
./agenda r -u Bob -p 654321
./agenda r -u Cici -p 123456
./agenda r -u Duke -p 654321
./agenda r -u Ella -p 123456
./agenda login -u Amy -p 123456
./agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
./agenda logout
./agenda login -u Bob -p 654321
./agenda m c -t me2 -s 2018-10-24/09:00 -e 2018-10-24/11:00 -p Duke
./agenda m l
./agenda m l -t me3
./agenda m l -t me1
./agenda m l -s 2018-10-25/09:00 -e 2019-10-25/09:00
./agenda m l -s 2018-10-25/09:00
./agenda m l -e 2019-10-25/09:00
./agenda m l -t me1 -s 2018-10-25/09:00 -e 2019-10-25/09:00
./agenda m l -s 2018-10-25/09:00 -e 2018-10-25/11:00
./agenda logout
./agenda login -u Amy -p 123456
./agenda m l
./agenda logout
./agenda login -u Ella -p 123456
./agenda m l
./agenda logout
```

### 测试结果

```sh
./agenda register -u Amy -p 123456
./agenda r -u Bob -p 654321
./agenda r -u Cici -p 123456
./agenda r -u Duke -p 654321
./agenda r -u Ella -p 123456
./agenda login -u Amy -p 123456
./agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
./agenda logout
./agenda login -u Bob -p 654321
./agenda m c -t me2 -s 2018-10-24/09:00 -e 2018-10-24/11:00 -p Duke

./agenda m l
# me1、me2

./agenda m l -t me3
# nil

./agenda m l -t me1
# me1

./agenda m l -s 2018-10-25/09:00 -e 2019-10-25/09:00
# me1

./agenda m l -s 2018-10-25/09:00
# me1

./agenda m l -e 2019-10-25/09:00
# me1, me2

./agenda m l -t me1 -s 2018-10-25/09:00 -e 2019-10-25/09:00
# me1

./agenda m l -s 2018-10-25/09:00 -e 2018-10-25/11:00
# nil

./agenda logout
./agenda login -u Amy -p 123456

./agenda m l
# me1

./agenda logout
./agenda login -u Ella -p 123456

./agenda m l
# nil

./agenda logout
```

