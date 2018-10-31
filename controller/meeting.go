package controller

import (
	"fmt"

	"time"

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
	participator := c.Args
	user := c.Ctx.User.Get()
	if user == "" {
		fmt.Println("you should first login")
		return
	}
	err := ctrl.Srv.Meeting().AddMeeting(title, participator, user)
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Info("add successful")
	}

}

func (c *Controller) MeetingList() {
	title, setT := c.Ctx.GetString("title")
	startTime, _ := c.Ctx.GetTime("startTime")
	endTime, _ := c.Ctx.GetTime("endTime")

	if setT {
		verifyNonNilTitle(title)
	}
	verifyStartTime(startTime)
	verifyEndTime(endTime)

	log.Verbose("check status")
	currentUser := c.Ctx.User.Get()
	if currentUser == "" {
		fmt.Println("you should login")
		return
	}

	log.Verbose("check time")
	if !endTime.Equal(time.Unix(0, 0)) && !startTime.Before(endTime) {
		fmt.Println("start time should be before end time")
		return
	}

	detail, err := c.Srv.Meeting().ListMeetings(currentUser, title, startTime, endTime)
	if err != nil {
		log.Error(err.Error())
	}
	fmt.Println(detail)
}

func (c *Controller) MeetingRemove() {
	title, _ := c.Ctx.GetString("title")

	verifyNonNilTitle(title)

	currentUser := c.Ctx.User.Get()
	if currentUser == "" {
		fmt.Println("you should login")
		return
	}

	err := c.Srv.Meeting().RemoveParticipators(currentUser, title, c.Args)
	if err != nil {
		if err.Error() == "delete_meeting" {
			log.Info("remove participators successfully and delete meeting because of no participator")
		} else {
			log.Error(err.Error())
		}
	} else {
		log.Info("remove participators successfully")
	}
}

func GetMeetingCtrl() MeetingCtrl {
	return &ctrl
}
