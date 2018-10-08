package controller

type MeetingCtrl Controller

var meetingCtrl MeetingCtrl

func GetMeetingCtrl() *MeetingCtrl {
	return (*MeetingCtrl)(initController((*Controller)(&meetingCtrl)))
}
