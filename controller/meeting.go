package controller

import (
	"fmt"

	"github.com/MegaShow/goagenda/lib/log"
)

type MeetingCtrl interface {
	MeetingCreate()
	MeetingSet()
	MeetingQuit()
	MeetingDelete()
	MeetingAdd()
	MeetingList()
	MeetingRemove()
}

func (c *Controller) MeetingCreate() {
	title, _ := c.Ctx.GetString("title")
	startTime, _ := c.Ctx.GetTime("startTime")
	endTime, _ := c.Ctx.GetTime("endTime")
	participators, _ := c.Ctx.GetStringSlice("participator")

	verifyNonNilTitle(title)
	verifyNonNilStartTime(startTime)
	verifyNonNilEndTime(endTime)
	verifyNonNilParticipator(participators)

	log.Verbose("check status")
	currentUser := c.Ctx.User.Get()
	if currentUser == "" {
		fmt.Println("you should login")
		return
	}

	log.Verbose("check time")
	if !startTime.Before(endTime) {
		fmt.Println("start time should be before end time")
		return
	}

	err := ctrl.Srv.Meeting().CreateMeeting(title, startTime, endTime, currentUser, participators)
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("create meeting successfully")
}

func (c *Controller) MeetingSet() {}

func (c *Controller) MeetingQuit() {}

func (c *Controller) MeetingDelete() {
	isAll, _ := c.Ctx.GetBool("all")
	title, _ := c.Ctx.GetString("title")
	if isAll == true {
		fmt.Println("delete all")
	} else {
		fmt.Println("title: ", title)
	}
	//TODO
}

func (c *Controller) MeetingAdd() {
	title, _ := c.Ctx.GetString("title")
	fmt.Println("title: ", title)
	participator := c.Args
	if len(participator) == 0 {
		fmt.Println(0)
	}
	for i := 0; i < len(participator); i++ {
		fmt.Println(participator[i])
	}
	err := ctrl.Srv.Meeting().AddMeeting(title, participator, c.Ctx.User.Get())
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("add successful")
}

func (c *Controller) MeetingList() {
	title, _ := c.Ctx.GetString("title")
	startTime, _ := c.Ctx.GetString("startTime")
	endTime, _ := c.Ctx.GetString("endTime")
	userName, _ := c.Ctx.GetString("userName")
	fmt.Println("title: ", title)
	fmt.Println("startTime: ", startTime)
	fmt.Println("endTime: ", endTime)
	fmt.Println("userName: ", userName)
}

func (c *Controller) MeetingRemove() {
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
