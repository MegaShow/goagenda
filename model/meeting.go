package model

type MeetingDB struct {
	Data []Meeting
	Database
}

type Meeting struct {
}

var MeetingModel = UserDB{Database: Database{schema: "Meeting"}}

func initMeetingModel() {
	MeetingModel.initModel(&MeetingModel.Data)
}

func ReleaseMeetingModel() {
	MeetingModel.releaseModel(&MeetingModel.Data)
}
