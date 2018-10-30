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

func (c *Controller) MeetingQuit() {
	title, _ := c.Ctx.GetString("title")

	verifyNonNilTitle(title)

	currentUser := c.Ctx.User.Get()
	if currentUser == "" {
		fmt.Println("you should login")
		return
	}

	err := c.Srv.Meeting().QuitMeeting(currentUser, title)
	if err != nil {
		if err.Error() == "delete_meeting" {
			log.Info("quit meeting successfully and delete meeting because of no participator")
		} else {
			log.Error(err.Error())
		}
	} else {
		log.Info("quit meeting successfully")
	}
}

func (c *Controller) MeetingDelete() {
	isAll, _ := c.Ctx.GetBool("all")
	title, setT := c.Ctx.GetString("title")

	if isAll && setT {
		fmt.Println("flags -a and -t cannot be set together")
		return
	} else if setT {
		verifyNonNilTitle(title)
	} else if !isAll {
		c.Cmd.Usage()
		return
	}

	currentUser := c.Ctx.User.Get()
	if currentUser == "" {
		fmt.Println("you should login")
		return
	}

	if isAll == true {
		err := c.Srv.Meeting().DeleteMeeting(currentUser, "")
		if err != nil {
			log.Error(err.Error())
		}
		log.Info("delete all meetings successfully")
	} else {
		err := c.Srv.Meeting().DeleteMeeting(currentUser, title)
		if err != nil {
			log.Error(err.Error())
		}
		log.Info("delete meeting \"" + title + "\" successfully")
	}
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
