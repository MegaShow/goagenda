# 测试样例

建议这样进行测试：

- 使用 `go install` 或 `go get` 安装Agenda后，将 `$GOPATH/bin` 下生成的 `goagenda.exe` 重命名为 `agenda.exe`

- 在进行每次测试前，清除数据文件，在类Unix终端或Windows Powershell下，可以使用 `rm ~/.agenda/data/*` 
- 测试完后，`rm -r ~/.agenda` 清除配置文件

请注意：

- 下面部分测试例子使用了空字符串 `""` ，而Windows Powershell并不识别 `""` 

## Register

**service/admin**

- 检测用户名是否已存在

### 测试例子

```sh
agenda register -u "@" -p 123456
agenda register -u Amy -p 123
agenda register -u Amy -p 123456
agenda register -u Amy -p 123456
```

### 测试结果

```sh
agenda register -u "@" -p 123456
# 失败，用户名含有非法字符

agenda register -u Amy -p 123
# 失败，密码过短

agenda register -u Amy -p 123456

agenda register -u Amy -p 123456
# 失败，用户名已存在
```

## Login

**controller/admin**

- 检测是否已登录

**service/admin**

- 检测用户名、密码是否正确

### 测试例子

```sh
agenda register -u Amy -p 123456
agenda register -u Bob -p 654321
agenda login -u Amy -p 12345678
agenda login -u AA -p 123456
agenda login -u Amy -p 123456
agenda login -u Amy -p 123456
agenda login -u Bob -p 654321
agenda logout
```

### 测试结果

```sh
agenda register -u Amy -p 123456

agenda register -u Bob -p 654321

agenda login -u Amy -p 12345678
# 失败，密码错误

agenda login -u AA -p 123456
# 失败，用户名错误

agenda login -u Amy -p 123456

agenda login -u Amy -p 123456
# 失败，该用户已登录

agenda login -u Bob -p 654321
# 失败，当前已有登录用户Amy

agenda logout
```

## Logout

**controller/admin**

- 检测是否已登录

### 测试例子

```sh
agenda register -u Amy -p 123456
agenda logout
agenda login -u Amy -p 123456
agenda logout
```

### 测试结果

```sh
agenda register -u Amy -p 123456

agenda logout
# 失败，没有登录的用户

agenda login -u Amy -p 123456

agenda logout
```

## Status

### 测试例子

```sh
agenda register -u Amy -p 123456
agenda status
agenda login -u Amy -p 123456
agenda status
agenda logout
agenda status
```

### 测试结果

```sh
agenda register -u Amy -p 123456

agenda status
# 无登录用户

agenda login -u Amy -p 123456

agenda status
# 当前用户：Amy

agenda logout

agenda status
# 无登录用户
```

## User/Set

**controller/user**

- 检测是否登录

**service/user**

- 如果重设了密码，需要进行加密操作

### 测试例子

```sh
agenda register -u Amy -p 123456
agenda user set -e yuyu@qq.com
agenda login -u Amy -p 123456
agenda u s -e y
agenda u s -t 2321
agenda u s -p 12ab
agenda u s -p ""
agenda u s -p "123456@"
agenda u s
agenda u s -e yuyu@qq.com -t 13800138000 -p "123abc"
agenda u list
agenda u s -e "" -t ""
agenda u list
agenda logout
agenda login -u Amy -p 123abc
agenda logout
```

### 测试结果

```sh
agenda register -u Amy -p 123456

agenda user set -e yuyu@qq.com
# 失败，因为未登陆

agenda login -u Amy -p 123456

agenda u s -e y
# 失败，因为邮箱参数格式不正确

agenda u s -t 2321
# 失败，因为电话参数格式不正确

agenda u s -p 12ab
# 失败，因为密码过短

agenda u s -p ""
# 失败，因为密码为空

agenda u s -p "123456@"

agenda u s
# set nothing

agenda u s -e yuyu@qq.com -t 13800138000 -p "123abc"

agenda u list

agenda u s -e "" -t ""
# 设邮箱为空，电话为空

agenda u list

agenda logout

agenda login -u Amy -p 123abc
# 测试新密码是否生效

agenda logout
```

## User/List

controller/meeting**

- 要检测是否登录
- 要检测用户名是否合法

**service/meeting**

- 如果有指定用户，要检测用户是否存在

### 测试命令

```sh
agenda register -u Amy -p 123456
agenda r -u Bob -p 654321
agenda r -u Cici -p 123456
agenda r -u Duke -p 654321
agenda r -u Ella -p 123456
agenda user list
agenda login -u Amy -p 123456
agenda user list
agenda user list -u Amy
agenda user list -u amy
```

### 测试结果

```sh
agenda register -u Amy -p 123456
agenda r -u Bob -p 654321
agenda r -u Cici -p 123456
agenda r -u Duke -p 654321
agenda r -u Ella -p 123456
agenda user list
# 失败，未登录

agenda login -u Amy -p 123456
agenda user list
# 成功，列出所有用户信息

agenda user list -u Amy
# 成功，列出Amy信息

agenda user list -u amy
# 失败，找不到该用户
```

## User/Delete

**controller/user**

* 检测是否登录
* 检测用户、密码是否合法
* 检测用户是否为当前用户

**service/user**

* 检测用户、密码是否正确
* 检测是否需要删除或退出会议

### 测试命令

```sh
agenda register -u Amy -p 123456
agenda r -u Bob -p 654321
agenda r -u Cici -p 123456
agenda r -u Duke -p 654321
agenda r -u Ella -p 123456
agenda login -u Amy -p 123456
agenda meeting c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
agenda meeting c -t me2 -s 2018-10-27/09:00 -e 2018-10-27/11:00 -p Duke
agenda logout
agenda login -u Bob -p 654321
agenda meeting c -t me3 -s 2018-10-28/09:00 -e 2018-10-28/11:00 -p Duke
agenda logout
agenda user delete -u Bob -p 654321
agenda login -u Bob -p 654321
agenda user delete -u Bob -p 123456
agenda user delete -u Bob -p 654321
agenda login -u Ella -p 123456
agenda u d -u Ella -p 123456
agenda login -u Cici -p 123456
agenda u d -u Cici -p 123456
agenda login -u Amy -p 123456
agenda u d -u Amy -p 123456
```

### 测试结果

```sh
agenda register -u Amy -p 123456
agenda r -u Bob -p 654321
agenda r -u Cici -p 123456
agenda r -u Duke -p 654321
agenda r -u Ella -p 123456
agenda login -u Amy -p 123456
agenda meeting c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
agenda meeting c -t me2 -s 2018-10-27/09:00 -e 2018-10-27/11:00 -p Duke
agenda logout
agenda login -u Bob -p 654321
agenda meeting c -t me3 -s 2018-10-28/09:00 -e 2018-10-28/11:00 -p Duke
agenda logout

agenda user delete -u Bob -p 654321
# 失败，未登录

agenda login -u Bob -p 654321

agenda user delete -u Bob -p 123456
# 失败，用户、密码不正确

agenda user delete -u Bob -p 654321
# 成功，删除me3、修改me1

agenda login -u Ella -p 123456

agenda u d -u Ella -p 123456
# 成功

agenda login -u Cici -p 123456

agenda u d -u Cici -p 123456
# 成功，删除me1

agenda login -u Amy -p 123456

agenda u d -u Amy -p 123456
# 成功，删除me2
```

## Meeting/Create

**controller/meeting**

- 检测是否登录
- 检测时间、议题是否合法

**service/meeting**

- 检测议题是否已存在
- 检测某些与会者是否存在
- 检测与会者是否有别的重叠会议（检测的与会者也包括会议发起人）
- 去掉重复的与会者

### 测试命令

```sh
agenda register -u Amy -p 123456
agenda r -u Bob -p 654321
agenda r -u Cici -p 123456
agenda r -u Duke -p 654321
agenda r -u Ella -p 123456
agenda meeting c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/08:00 -p Bob,Cici
agenda login -u Amy -p 123456
agenda meeting c
agenda meeting c -t me0 -s "" -e 2018-10-26/08:00 -p Bob,Cici
agenda meeting c -t me0 -s 2018-10-26/09:00 -e 2018-10-26-08-00 -p Bob,Cici
agenda meeting c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/08:00 -p Bob,Cici
agenda meet c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/09:00 -p Bob,Cici
agenda m c -t "" -s 2018-10-26/09:00 -e 2018-10-26/10:00 -p Bob,Cici
agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/10:00 -p Bob,Frank
agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
agenda m c -t me1 -s 2018-10-26/13:00 -e 2018-10-26/15:00 -p Bob,Cici
agenda m c -t me2 -s 2018-10-26/08:00 -e 2018-10-26/9:30 -p Bob,Duke
agenda m c -t me2 -s 2018-10-26/08:00 -e 2018-10-26/9:00 -p Bob,Duke,Duke,Amy
agenda logout
agenda login -u Duke -p 654321
agenda m c -t me3 -s 2018-10-26/10:5 -e 2018-10-26/12:00 -p Amy,Ella
agenda m c -t me3 -s 2018-10-26/8:7 -e 2018-10-26/12:00 -p Ella
agenda m c -t me3 -s 2018-10-27/16:00 -e 2018-10-27/17:00 -p Cici,Ella
agenda m c -t me4 -s 2018-10-26/9:30 -e 2018-10-26/10:30 -p Ella
agenda logout
```

### 测试结果

```sh
agenda register -u Amy -p 123456

agenda r -u Bob -p 654321

agenda r -u Cici -p 123456

agenda r -u Duke -p 654321

agenda r -u Ella -p 123456

agenda meeting c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/08:00 -p Bob,Cici
# 失败，因为未登陆

agenda login -u Amy -p 123456

agenda meeting c
# 失败，因为无参数

agenda meeting c -t me0 -s "" -e 2018-10-26/08:00 -p Bob,Cici
# 失败，因为输入的时间为空

agenda meeting c -t me0 -s 2018-10-26/09:00 -e 2018-10-26-08-00 -p Bob,Cici
# 失败，因为输入的时间格式非法

agenda meeting c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/08:00 -p Bob,Cici
# 失败，因为时间顺序不对（开始时间小于结束时间）

agenda meet c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/09:00 -p Bob,Cici
# 失败，因为时间顺序不对（开始时间和结束时间相等）

agenda m c -t "" -s 2018-10-26/09:00 -e 2018-10-26/10:00 -p Bob,Cici
# 失败，因为议题不合法

agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/10:00 -p Bob,Frank
# 失败，因为用户Frank不存在

agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
# 成功
# me1 2018-10-26/09:00 2018-10-26/11:00 Amy Bob,Cici

agenda m c -t me1 -s 2018-10-26/13:00 -e 2018-10-26/15:00 -p Bob,Cici
# 失败，因为议题me1已存在

agenda m c -t me2 -s 2018-10-26/08:00 -e 2018-10-26/9:30 -p Bob,Duke
# 失败，因为Amy和Bob有重叠的会议me1

agenda m c -t me2 -s 2018-10-26/08:00 -e 2018-10-26/9:00 -p Bob,Duke,Duke,Amy
# 成功
# （检测一场会议的endTime与另一场会议的startTime重叠的情况：me2结束时间与me1开始时间重叠）
# （同时检测与会者去重的情况，在data/meeting.json里查看，结果应该是Bob,Duke）
# me1 2018-10-26/09:00 2018-10-26/11:00 Amy Bob,Cici
# me2 2018-10-26/08:00 2018-10-26/09:00 Amy Bob,Duke

agenda logout

agenda login -u Duke -p 654321

agenda m c -t me3 -s 2018-10-26/10:5 -e 2018-10-26/12:00 -p Amy,Ella
# 失败，因为Amy正在主持另一会议me1

agenda m c -t me3 -s 2018-10-26/8:7 -e 2018-10-26/12:00 -p Ella
# 失败，因为当前用户Duke正在开另一会议me2

agenda m c -t me3 -s 2018-10-27/16:00 -e 2018-10-27/17:00 -p Cici,Ella
# 成功
# me1 2018-10-26/09:00 2018-10-26/11:00 Amy Bob,Cici
# me2 2018-10-26/08:00 2018-10-26/09:00 Amy Bob,Duke
# me3 2018-10-27/16:00 2018-10-27/17:00 Duke Cici,Ella

agenda m c -t me4 -s 2018-10-26/9:30 -e 2018-10-26/10:30 -p Ella
# 成功
# （检测有时间段重叠、但人员不相干的会议的情况）
# me1 2018-10-26/09:00 2018-10-26/11:00 Amy Bob,Cici
# me2 2018-10-26/08:00 2018-10-26/09:00 Amy Bob,Duke
# me3 2018-10-27/16:00 2018-10-27/17:00 Duke Cici,Ella
# me4 2018-10-26/09:30 2018-10-26/10:30 Duke Ella

agenda logout
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

```sh
agenda register -u Amy -p 123456
agenda r -u Bob -p 654321
agenda r -u Cici -p 123456
agenda r -u Duke -p 654321
agenda r -u Ella -p 123456
agenda r -u Frank -p 654321
agenda meeting s -t me1 -s 2018-11-1/09:00 -e 2018-11-1/11:00 -p Bob
agenda login -u Amy -p 123456
agenda meeting set -t me1
agenda m s -t "" -p Bob
agenda m s -t me1 -p ""
agenda m s -t me1 -s 2018-11-1-09-00 -p Bob
agenda m s -t me1 -s 2018-11-1/12:00 -e 2018-11-1/11:00 -p Bob
agenda m s -t me1 -s 2018-11-1/09:00 -e 2018-11-1/11:00 -p Bob,Cici
agenda m create -t me1 -s 2018-11-1/09:00 -e 2018-11-1/11:00 -p Bob,Cici
agenda m s -t me1 -s 2018-11-1/13:00 -p Bob
agenda m s -t me1 -e 2018-11-1/8:30 -p Bob
agenda m s -t me1 -p Bob,Gavin
agenda m s -t me1 -s 2018-11-1/09:30 -e 2018-11-1/11:30 -p Ella,Ella,Bob,Amy
agenda m create -t me2 -s 2018-11-1/15:00 -e 2018-11-1/17:30 -p Bob,Duke
agenda m s -t me2 -s 2018-11-1/10:00 -e 2018-11-1/12:30
agenda m s -t me2 -s 2018-11-1/8:0 -e 2018-11-1/9:0 -p Bob,Frank
agenda logout
agenda login -u Cici -p 123456
agenda m create -t me3 -s 2018-11-1/7:30 -e 2018-11-1/8:45 -p Ella,Duke
agenda m s -t me1 -s 2018-11-1/8:00
agenda m s -t me3 -p Frank,Duke
agenda m s -t me3 -p Duke,Amy
agenda m s -t me3 -e 2018-11-1/7:45
agenda m s -t me3 -s 2018-11-1/6:0
agenda m s -t me3 -p Amy,Ella,Duke,Bob
agenda m c -t me4 -s 2018-11-1/10:00 -e 2018-11-1/10:30 -p Frank
agenda m s -t me4 -s 2018-11-1/7:00 -e 2018-11-1/7:30
agenda logout
```

### 测试结果

```sh
agenda register -u Amy -p 123456

agenda r -u Bob -p 654321

agenda r -u Cici -p 123456

agenda r -u Duke -p 654321

agenda r -u Ella -p 123456

agenda r -u Frank -p 654321

agenda meeting s -t me1 -s 2018-11-1/09:00 -e 2018-11-1/11:00 -p Bob
# 失败，因为未登录

agenda login -u Amy -p 123456

agenda meeting set -t me1
# set nothing

agenda m s -t "" -p Bob
# 失败，因为议题参数不合法

agenda m s -t me1 -p ""
# 失败，因为与会者参数不合法

agenda m s -t me1 -s 2018-11-1-09-00 -p Bob
# 失败，因为时间参数格式不对

agenda m s -t me1 -s 2018-11-1/12:00 -e 2018-11-1/11:00 -p Bob
# 失败，因为时间参数顺序不对

agenda m s -t me1 -s 2018-11-1/09:00 -e 2018-11-1/11:00 -p Bob,Cici
# 失败，因为会议me1不存在

agenda m create -t me1 -s 2018-11-1/09:00 -e 2018-11-1/11:00 -p Bob,Cici
# 创建会议me1
# me1 2018-11-1/09:00 2018-11-1/11:00 Amy Bob,Cici

agenda m s -t me1 -s 2018-11-1/13:00 -p Bob
# 失败，因为重设的startTime不在原来的endTime之前

agenda m s -t me1 -e 2018-11-1/8:30 -p Bob
# 失败，因为重设的endTime不在原来的startTime之后

agenda m s -t me1 -p Bob,Gavin
# 失败，因为重设的与会者中Gavin不存在

agenda m s -t me1 -s 2018-11-1/09:30 -e 2018-11-1/11:30 -p Ella,Ella,Bob,Amy
# 成功，对me1，重设时间段，重设与会者，并且去掉了重复的与会者
# 在data/meeting.json里查看结果
# me1 2018-11-1/09:30 2018-11-1/11:30 Amy Bob,Ella

agenda m create -t me2 -s 2018-11-1/15:00 -e 2018-11-1/17:30 -p Bob,Duke
# 创建会议me2
# me1 2018-11-1/09:30 2018-11-1/11:30 Amy Bob,Ella
# me2 2018-11-1/15:00 2018-11-1/17:30 Amy Bob,Duke

agenda m s -t me2 -s 2018-11-1/10:00 -e 2018-11-1/12:30
# 失败，因为Amy或Bob有重叠的会议me1
# （测试：重设了时间后检测重叠会议）

agenda m s -t me2 -s 2018-11-1/8:0 -e 2018-11-1/9:0 -p Bob,Frank
# 成功
# （测试：一场会议结束时间与另一场会议开始时间重叠的情况）
# me1 2018-11-1/09:30 2018-11-1/11:30 Amy Bob,Ella
# me2 2018-11-1/08:00 2018-11-1/09:00 Amy Bob,Frank

agenda logout

agenda login -u Cici -p 123456

agenda m create -t me3 -s 2018-11-1/7:30 -e 2018-11-1/8:45 -p Ella,Duke
# 创建会议me3
# me1 2018-11-1/09:30 2018-11-1/11:30 Amy Bob,Ella
# me2 2018-11-1/08:00 2018-11-1/09:00 Amy Bob,Frank
# me3 2018-11-1/07:30 2018-11-1/08:45 Cici Duke,Ella

agenda m s -t me1 -s 2018-11-1/8:00
# 失败，因为me1的发起人不是当前用户Cici

agenda m s -t me3 -p Frank,Duke
# 失败，因为Frank有重叠的会议me2
# （测试：重设了与会者后检测重叠会议）

agenda m s -t me3 -p Duke,Amy
# 失败，因为Amy有重叠的会议me2
# （测试：检测的其他重叠会议的与会者，应包括这些会议的发起人）

agenda m s -t me3 -e 2018-11-1/7:45
# 成功

agenda m s -t me3 -s 2018-11-1/6:00
# 成功

agenda m s -t me3 -p Amy,Ella,Duke,Bob
# 成功
# me1 2018-11-1/09:30 2018-11-1/11:30 Amy Bob,Ella
```

## Meeting/Add

**controller/meeting**

- 要检测是否登录
- 要检测议题是否合法

**service/meeting**

- 要检测会议是否存在
- 要检测当前用户是否为会议发起者
- 要检测用户列表是否合法

### 测试命令

```sh
agenda register -u Amy -p 123456
agenda r -u Bob -p 654321
agenda r -u Cici -p 123456
agenda r -u Duke -p 654321
agenda r -u Ella -p 123456
agenda login -u Amy -p 123456
agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici,Ella
agenda logout
agenda m remove -t me1 Bob
agenda login -u Bob -p 654321
agenda m r -t me1 Bob
agenda logout
agenda login -u Amy -p 123456
agenda m r -t me2 Bob
agenda m r -t me1 Bob Duke
agenda m r -t me1 Bob Frank
agenda m r -t me1 Bob Ella
agenda m r -t me1 Cici
```

### 测试结果

```sh
.agenda register -u Amy -p 123456
agenda r -u Bob -p 654321
agenda r -u Cici -p 123456
agenda r -u Duke -p 654321
agenda r -u Ella -p 123456
agenda login -u Amy -p 123456
agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici,Ella
agenda logout

agenda m remove -t me1 Bob
# 失败，未登录

agenda login -u Bob -p 654321

agenda m r -t me1 Bob
# 失败，不是发起者

agenda logout
agenda login -u Amy -p 123456

agenda m r -t me2 Bob
# 失败，会议不存在

agenda m r -t me1 Bob Duke
# 失败，Duke不是与会者

agenda m r -t me1 Bob Frank
# 失败，Frank不存在

agenda m r -t me1 Bob Ella
# 成功

agenda m r -t me1 Cici
# 成功，并删除会议
```


## Meeting/Remove

**controller/meeting**

- 检测是否登录
- 检测议题是否合法

**service/meeting**

- 检测会议是否存在
- 检测当前用户是否为会议发起者
- 检测用户列表是否合法

### 测试命令

```sh
agenda register -u Amy -p 123456
agenda r -u Bob -p 654321
agenda r -u Cici -p 123456
agenda r -u Duke -p 654321
agenda r -u Ella -p 123456
agenda login -u Amy -p 123456
agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici,Ella
agenda logout
agenda m remove -t me1 Bob
agenda login -u Bob -p 654321
agenda m r -t me1 Bob
agenda logout
agenda login -u Amy -p 123456
agenda m r -t me2 Bob
agenda m r -t me1 Bob Duke
agenda m r -t me1 Bob Frank
agenda m r -t me1 Bob Ella
agenda m r -t me1 Cici
```

### 测试结果

```sh
agenda register -u Amy -p 123456
agenda r -u Bob -p 654321
agenda r -u Cici -p 123456
agenda r -u Duke -p 654321
agenda r -u Ella -p 123456
agenda login -u Amy -p 123456
agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici,Ella
agenda logout

agenda m remove -t me1 Bob
# 失败，未登录

agenda login -u Bob -p 654321

agenda m r -t me1 Bob
# 失败，不是发起者

agenda logout
agenda login -u Amy -p 123456

agenda m r -t me2 Bob
# 失败，会议不存在

agenda m r -t me1 Bob Duke
# 失败，Duke不是与会者

agenda m r -t me1 Bob Frank
# 失败，Frank不存在

agenda m r -t me1 Bob Ella
# 成功

agenda m r -t me1 Cici
# 成功，并删除会议
```

## Meeting/Delete

**controller/meeting**

- 检测是否登录
- 检测议题是否合法

**service/meeting**

* 检测会议是否存在
* 检测当前用户是否为会议发起者

### 测试命令

```sh
agenda register -u Amy -p 123456
agenda r -u Bob -p 654321
agenda r -u Cici -p 123456
agenda r -u Duke -p 654321
agenda r -u Ella -p 123456
agenda login -u Amy -p 123456
agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
agenda logout
agenda m delete -t me1
agenda login -u Bob -p 654321
agenda m delete -t me1
agenda m d -t me2
agenda logout
agenda login -u Amy -p 123456
agenda m d -t me1
agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
agenda m c -t me2 -s 2018-10-27/09:00 -e 2018-10-27/11:00 -p Bob,Cici
agenda m d -a -t me1
agenda m d -a
```

### 测试结果

```sh
agenda register -u Amy -p 123456
agenda r -u Bob -p 654321
agenda r -u Cici -p 123456
agenda r -u Duke -p 654321
agenda r -u Ella -p 123456
agenda login -u Amy -p 123456
agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
agenda logout

agenda m delete -t me1
# 失败，用户未登录

agenda login -u Bob -p 654321

agenda m delete -t me1
# 失败，用户不是会议发起者

agenda m d -t me2
# 失败，找不到该会议

agenda logout
agenda login -u Amy -p 123456

agenda m d -t me1
# 成功

agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
agenda m c -t me2 -s 2018-10-27/09:00 -e 2018-10-27/11:00 -p Bob,Cici

agenda m d -a -t me1
# 失败，不能同时指定-a和-t

agenda m d -a
# 成功，删除两个会议
```

## Meeting/Quit

**controller/meeting**

- 检测是否登录
- 检测议题是否合法

**service/meeting**

- 检测会议是否存在
- 检测当前用户是否为会议发起者，发起者不能退出会议
- 检测当前用户是否为会议参与者
- 检测当前用户退出会议之后会议参与人数是否为0，是否需要删除会议

### 测试命令

```sh
agenda register -u Amy -p 123456
agenda r -u Bob -p 654321
agenda r -u Cici -p 123456
agenda r -u Duke -p 654321
agenda r -u Ella -p 123456
agenda login -u Amy -p 123456
agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
agenda logout
agenda m quit -t me1
agenda login -u Bob -p 654321
agenda m quit -t me1
agenda m q -t me1
agenda m q -t me2
agenda logout
agenda login -u Amy -p 123456
agenda m q -t me1
agenda logout
agenda login -u Cici -p 123456
agenda m q -t me1
```

### 测试结果

```sh
agenda register -u Amy -p 123456
agenda r -u Bob -p 654321
agenda r -u Cici -p 123456
agenda r -u Duke -p 654321
agenda r -u Ella -p 123456
agenda login -u Amy -p 123456
agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
agenda logout

agenda m quit -t me1
# 失败，未登录

agenda login -u Bob -p 654321

agenda m quit -t me1
# 成功

agenda m q -t me1
# 失败，不是会议参与者

agenda m q -t me2
# 失败，不存在该会议

agenda logout
agenda login -u Amy -p 123456

agenda m q -t me1
# 失败，为会议发起者

agenda logout
agenda login -u Cici -p 123456

agenda m q -t me1
# 成功，并删除会议
```

## Meeting/List

**controller/meeting**

- 检测是否登录
- 检测议题是否合法
- 检测时间是否合法

### 测试命令

```sh
agenda register -u Amy -p 123456
agenda r -u Bob -p 654321
agenda r -u Cici -p 123456
agenda r -u Duke -p 654321
agenda r -u Ella -p 123456
agenda login -u Amy -p 123456
agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
agenda logout
agenda login -u Bob -p 654321
agenda m c -t me2 -s 2018-10-24/09:00 -e 2018-10-24/11:00 -p Duke
agenda m l
agenda m l -t me3
agenda m l -t me1
agenda m l -s 2018-10-25/09:00 -e 2019-10-25/09:00
agenda m l -s 2018-10-25/09:00
agenda m l -e 2019-10-25/09:00
agenda m l -t me1 -s 2018-10-25/09:00 -e 2019-10-25/09:00
agenda m l -s 2018-10-25/09:00 -e 2018-10-25/11:00
agenda logout
agenda login -u Amy -p 123456
agenda m l
agenda logout
agenda login -u Ella -p 123456
agenda m l
agenda logout
```

### 测试结果

```sh
agenda register -u Amy -p 123456
agenda r -u Bob -p 654321
agenda r -u Cici -p 123456
agenda r -u Duke -p 654321
agenda r -u Ella -p 123456
agenda login -u Amy -p 123456
agenda m c -t me1 -s 2018-10-26/09:00 -e 2018-10-26/11:00 -p Bob,Cici
agenda logout
agenda login -u Bob -p 654321
agenda m c -t me2 -s 2018-10-24/09:00 -e 2018-10-24/11:00 -p Duke

agenda m l
# me1、me2

agenda m l -t me3
# nil

agenda m l -t me1
# me1

agenda m l -s 2018-10-25/09:00 -e 2019-10-25/09:00
# me1

agenda m l -s 2018-10-25/09:00
# me1

agenda m l -e 2019-10-25/09:00
# me1, me2

agenda m l -t me1 -s 2018-10-25/09:00 -e 2019-10-25/09:00
# me1

agenda m l -s 2018-10-25/09:00 -e 2018-10-25/11:00
# nil

agenda logout
agenda login -u Amy -p 123456

agenda m l
# me1

agenda logout
agenda login -u Ella -p 123456

agenda m l
# nil

agenda logout
```
