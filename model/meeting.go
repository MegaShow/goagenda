package model

type MeetingModel interface {
}

type MeetingDB struct {
	Data []Meeting
	Database
}

type Meeting struct {
}

var meetingDB = UserDB{Database: Database{schema: "Meeting"}}

func ReleaseMeetingModel() {
	meetingDB.releaseModel(&meetingDB.Data)
}

func (m *Manager) Meeting() MeetingModel {
	if meetingDB.isInit == false {
		meetingDB.initModel(&meetingDB.Data)
	}
	return &meetingDB
}
