# Contributing

更新于2018.10.18。

本文适用于《服务计算》课程同一小组的同学。

## Coding and Pull Request

1. 前往[Teambition](https://www.teambition.com/project/5bc6ffbaf10ae90018184bd0/)领取任务。
2. 克隆本项目。
3. 编写任务逻辑，且自己Review。
4. Pull request至本项目。

## Code Architecture

### 如何编写一个命令集？

> **Go Agenda的命令集架构已编写好，但阅读本节有助于了解Go Agenda的架构。**

Go Agenda使用**Command、Controller、Services、Models**的架构形式实现命令集。

什么是命令集？在Go Agenda中，我们将一系列类似、相同对象的命令的集合称为命令集。命令集方便我们对数量庞大的命令加以管理，我们将同一个命令集的命令定义在同一个`.go`文件中。比如，下列命令是一个命令集，并且下列命令均有相同的父命令。

```sh
$ agenda user set
$ agenda user list
$ agenda user delete
```

当然，不相似的命令一样可以组成命令集，并不要求均有相同的父命令。在Go Agenda中，以下的命令均在同一个`.go`文件中定义。

```sh
$ agenda register
$ agenda login
$ agenda logout
$ agenda status
```

为什么使用命令集？Go Agenda将同一个命令集的命令定义在同一个`.go`文件中，并且同一个命令集的命令共享一个Controller，可以大大减少程序运行时Controller的数量。

下面我们将以编写命令集`user`为例，介绍Go Agenda的Command、Controller前面两层模型是如何工作的。(为了简化代码，部分细节或属性省去)

首先，需要编写命令集的根命令(父命令)，如果命令集内的命令不具有相同父命令，那再按实际情况修改代码。

```go
var userRootCmd = &cobra.Command{
	Use:     "user",
	Aliases: []string{"u"},
}
```

然后编写命令集三个命令。

```go
var userDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	PreRun:  userPreRun,
	Run:     controller.GetUserCtrl().Delete,
}

var userListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	PreRun:  userPreRun,
	Run:     controller.GetUserCtrl().List,
}

var userSetCmd = &cobra.Command{
	Use:     "set",
	Aliases: []string{"s"},
	PreRun:  userPreRun,
	Run:     controller.GetUserCtrl().Set,
}
```

现在你可能疑惑`PreRun`和`Run`的两个函数有什么用，我们可以先不管它，先来编写导入命令和处理命令Argument、Flag的逻辑。

在`cmd/*.go`文件的`init`函数中，将父命令添加到`rootCmd`中，并将命令集的命令均添加到该父命令中。对于Argument和Flag的处理，均采用`XxxxP`的形式处理。

```go
func init() {
	rootCmd.AddCommand(userRootCmd)
	userRootCmd.AddCommand(userDeleteCmd)
	userRootCmd.AddCommand(userListCmd)
	userRootCmd.AddCommand(userSetCmd)

	userListCmd.Flags().StringP("user", "u", "", "the username searched")

	userSetCmd.Flags().StringP("password", "p", "", "new password")
	userSetCmd.Flags().StringP("email", "e", "", "new email")
	userSetCmd.Flags().StringP("telephone", "t", "", "new telephone")
}
```

如果Argument为必须项，记得使用`MarkFlagRequired`函数标记。

接下来我们实现Controller。



