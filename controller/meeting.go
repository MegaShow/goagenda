package controller

type MeetingCtrl interface {

}

func GetMeetingCtrl() MeetingCtrl {
	return &ctrl
}
