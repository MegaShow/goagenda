package controller

type MeetingCtrl interface {
	MeetingCreate()
	MeetingSet()
	MeetingQuit()
}

func (c *Controller) MeetingCreate() {}

func (c *Controller) MeetingSet() {}

func (c *Controller) MeetingQuit() {}

func GetMeetingCtrl() MeetingCtrl {
	return &ctrl
}
