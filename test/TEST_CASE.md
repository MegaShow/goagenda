# 测试样例

## Register

## Login

## Logout

## Status

## User/Set

## User/List

## User/Delete

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

**controller/meeting**

- 检测是否登录
- 检测议题参数是否合法
- 检测时间参数是否合法
  - 检查重设的时间的格式
  - 如果重设了0个或1个时间，则不检查顺序
  - 如果重设了2个时间，需要检查顺序

**service/meeting**

- 检测会议是否已存在（不存在则非法）
- 检测该会议的发起人是否当前用户
- 检测新的时间段的顺序是否正确
  - 前提：重设了1个时间
- 检测某些与会者是否存在
  - 前提：重设了与会者
- 检测重叠会议（检测的与会者包括发起人），找到的重叠会议不包括当前会议
- 去掉重复的与会者
  - 前提：重设了与会者

### 测试命令

```shell
rm data/*.json
./agenda register -u Amy -p 123456
./agenda r -u Bob -p 654321
./agenda r -u Cici -p 123456
./agenda r -u Duke -p 654321
./agenda r -u Ella -p 123456
./agenda r -u Frank -p 654321
./agenda meeting s -t me1 -s 2018-11-1/09:00 -e 2018-11-1/11:00 -p Bob
./agenda login -u Amy -p 123456
./agenda meeting set -t me1
./agenda m s -t "" -p Bob
./agenda m s -t me1 -p ""
./agenda m s -t me1 -s 2018-11-1-09-00 -p Bob
./agenda m s -t me1 -s 2018-11-1/12:00 -e 2018-11-1/11:00 -p Bob
./agenda m s -t me1 -s 2018-11-1/09:00 -e 2018-11-1/11:00 -p Bob,Cici
./agenda m create -t me1 -s 2018-11-1/09:00 -e 2018-11-1/11:00 -p Bob,Cici
./agenda m s -t me1 -s 2018-11-1/13:00 -p Bob
./agenda m s -t me1 -e 2018-11-1/8:30 -p Bob
./agenda m s -t me1 -p Bob,Gavin
./agenda m s -t me1 -s 2018-11-1/09:30 -e 2018-11-1/11:30 -p Ella,Ella,Bob,Amy
./agenda m create -t me2 -s 2018-11-1/15:00 -e 2018-11-1/17:30 -p Bob,Duke
./agenda m s -t me2 -s 2018-11-1/10:00 -e 2018-11-1/12:30
./agenda m s -t me2 -s 2018-11-1/8:0 -e 2018-11-1/9:0 -p Bob,Frank
./agenda logout
./agenda login -u Cici -p 123456
./agenda m create -t me3 -s 2018-11-1/7:30 -e 2018-11-1/8:45 -p Ella,Duke
./agenda m s -t me1 -s 2018-11-1/8:00
./agenda m s -t me3 -p Frank,Duke
./agenda m s -t me3 -p Duke,Amy
./agenda m s -t me3 -e 2018-11-1/7:45
./agenda m s -t me3 -s 2018-11-1/6:0
./agenda m s -t me3 -p Amy,Ella,Bob,Duke,Frank
./agenda logout

```

### 测试结果

```shell
./agenda register -u Amy -p 123456

./agenda r -u Bob -p 654321

./agenda r -u Cici -p 123456

./agenda r -u Duke -p 654321

./agenda r -u Ella -p 123456

./agenda r -u Frank -p 654321

./agenda meeting s -t me1 -s 2018-11-1/09:00 -e 2018-11-1/11:00 -p Bob
# 不成功，因为未登录

./agenda login -u Amy -p 123456

./agenda meeting set -t me1
# set nothing

./agenda m s -t "" -p Bob
# 不成功，因为议题参数不合法

./agenda m s -t me1 -p ""
# 不成功，因为与会者参数不合法

./agenda m s -t me1 -s 2018-11-1-09-00 -p Bob
# 不成功，因为时间参数格式不对

./agenda m s -t me1 -s 2018-11-1/12:00 -e 2018-11-1/11:00 -p Bob
# 不成功，因为时间参数顺序不对

./agenda m s -t me1 -s 2018-11-1/09:00 -e 2018-11-1/11:00 -p Bob,Cici
# 不成功，因为会议me1不存在

./agenda m create -t me1 -s 2018-11-1/09:00 -e 2018-11-1/11:00 -p Bob,Cici
# 创建会议me1
# me1 Amy 2018-11-1/09:00 2018-11-1/11:00 Bob,Cici

./agenda m s -t me1 -s 2018-11-1/13:00 -p Bob
# 不成功，因为重设的startTime不在原来的endTime之前

./agenda m s -t me1 -e 2018-11-1/8:30 -p Bob
# 不成功，因为重设的endTime不在原来的startTime之后

./agenda m s -t me1 -p Bob,Gavin
# 不成功，因为重设的与会者中Gavin不存在

./agenda m s -t me1 -s 2018-11-1/09:30 -e 2018-11-1/11:30 -p Ella,Ella,Bob,Amy
# 成功，对me1，重设时间段，重设与会者，并且去掉了重复的与会者
# 在data/meeting.json里查看结果
# me1 Amy 2018-11-1/09:30 2018-11-1/11:30 Bob,Ella

./agenda m create -t me2 -s 2018-11-1/15:00 -e 2018-11-1/17:30 -p Bob,Duke
# 创建会议me2
# me1 Amy 2018-11-1/09:30 2018-11-1/11:30 Bob,Ella
# me2 Amy 2018-11-1/15:00 2018-11-1/17:30 Bob,Duke

./agenda m s -t me2 -s 2018-11-1/10:00 -e 2018-11-1/12:30
# 不成功，因为Amy或Bob有重叠的会议me1
# （测试：重设了时间后检测重叠会议）

./agenda m s -t me2 -s 2018-11-1/8:0 -e 2018-11-1/9:0 -p Bob,Frank
# 成功
# （测试：一场会议结束时间与另一场会议开始时间重叠的情况）
# me1 Amy 2018-11-1/09:30 2018-11-1/11:30 Bob,Ella
# me2 Amy 2018-11-1/08:00 2018-11-1/09:00 Bob,Frank

./agenda logout

./agenda login -u Cici -p 123456

./agenda m create -t me3 -s 2018-11-1/7:30 -e 2018-11-1/8:45 -p Ella,Duke
# 创建会议me3
# me1 Amy 2018-11-1/09:30 2018-11-1/11:30 Bob,Ella
# me2 Amy 2018-11-1/08:00 2018-11-1/09:00 Bob,Frank
# me3 Cici 2018-11-1/07:30 2018-11-1/08:45 Duke,Ella

./agenda m s -t me1 -s 2018-11-1/8:00
# 不成功，因为me1的发起人不是当前用户Cici

./agenda m s -t me3 -p Frank,Duke
# 不成功，因为Frank有重叠的会议me2
# （测试：重设了与会者后检测重叠会议）

./agenda m s -t me3 -p Duke,Amy
# 不成功，因为Amy有重叠的会议me2
# （测试：检测的其他重叠会议的与会者，应包括这些会议的发起人）

./agenda m s -t me3 -e 2018-11-1/7:45
# 成功

./agenda m s -t me3 -s 2018-11-1/6:00
# 成功

./agenda m s -t me3 -p Amy,Ella,Bob,Duke,Frank
# 成功
# me1 Amy 2018-11-1/09:30 2018-11-1/11:30 Bob,Ella
# me2 Amy 2018-11-1/08:00 2018-11-1/09:00 Bob,Frank
# me3 Cici 2018-11-1/06:00 2018-11-1/07:45 Amy,Bob,Duke,Ella,Frank

./agenda logout

```
