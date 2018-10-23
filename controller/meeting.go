package controller

import (
	"fmt"
)

type MeetingCtrl interface {
	MeetingCreate()
	MeetingSet()
	MeetingQuit()
	DeleteMeeting()
	Add()
	ListMeeting()
	Remove()
}

func (c *Controller) MeetingCreate() {}

func (c *Controller) MeetingSet() {}

func (c *Controller) MeetingQuit() {}


func (c *Controller) DeleteMeeting() {
	isAll, _ := c.Ctx.GetBool("all")
	title, _ := c.Ctx.GetString("title")
	if isAll == true {
		fmt.Println("delete all")
	} else {
		fmt.Println("title: ", title)
	}
	//TODO

}

func (c *Controller) Add() {
	title, _ := c.Ctx.GetString("title")
	fmt.Println("title: ", title)
	participator := c.Args
	if len(participator) == 0 {
		fmt.Println(0)
	}
	for i := 0; i < len(participator); i++ {
		fmt.Println(participator[i])
	}
}

func (c *Controller) ListMeeting() {
	title, _ := c.Ctx.GetString("title")
	startTime, _ := c.Ctx.GetString("startTime")
	endTime, _ := c.Ctx.GetString("endTime")
	userName, _ := c.Ctx.GetString("userName")
	fmt.Println("title: ", title)
	fmt.Println("startTime: ", startTime)
	fmt.Println("endTime: ", endTime)
	fmt.Println("userName: ", userName)
}

func (c *Controller) Remove() {
	title, _ := c.Ctx.GetString("title")
	fmt.Println("title: ", title)
	participator := c.Args
	if len(participator) == 0 {
		fmt.Println(0)
	}
	for i := 0; i < len(participator); i++ {
		fmt.Println(participator[i])
	}
}

func GetMeetingCtrl() MeetingCtrl {
	return &ctrl
}
