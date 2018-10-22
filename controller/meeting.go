package controller

import (
	"fmt"
)

type MeetingCtrl interface {
	DeleteMeeting()
	Add()
	ListMeeting()
	Remove()
}

func (c *Controller) DeleteMeeting() {
	if c.Ctx.GetBool("all") == true {
		fmt.Println("delete all")
	} else {
		fmt.Println("title: ", c.Ctx.GetString("title"))
	}
	//TODO

}

func (c *Controller) Add() {
	fmt.Println("title: ", c.Ctx.GetString("title"))
	participater := c.Args
	if len(participater) == 0 {
		fmt.Println(0)
	}
	for i := 0; i < len(participater); i++ {
		fmt.Println(participater[i])
	}
}

func (c *Controller) ListMeeting() {
	fmt.Println("title: ", c.Ctx.GetString("title"))
	fmt.Println("startTime: ", c.Ctx.GetString("startTime"))
	fmt.Println("endTime: ", c.Ctx.GetString("endTime"))
	fmt.Println("userName: ", c.Ctx.GetString("userName"))
}

func (c *Controller) Remove() {
	fmt.Println("title: ", c.Ctx.GetString("title"))
	participater := c.Args
	if len(participater) == 0 {
		fmt.Println(0)
	}
	for i := 0; i < len(participater); i++ {
		fmt.Println(participater[i])
	}
}

func GetMeetingCtrl() MeetingCtrl {
	return &ctrl
}
