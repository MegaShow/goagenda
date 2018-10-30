package model

import (
	"os"
	"time"
)

type MeetingModel interface {
	GetMeetingByTitle(title string) Meeting
	GetOccupiedParticipators(title string, startTime, endTime time.Time) map[string]bool
	CreateMeeting(meeting Meeting)
	SetMeeting(title string, startTime time.Time, setStart bool,
		endTime time.Time, setEnd bool, participators []string, setPars bool)
}

type MeetingDB struct {
	Data []Meeting
	Database
}

type Meeting struct {
	Title			string		`json:"title"`
	StartTime		time.Time	`json:"startTime"`
	EndTime			time.Time	`json:"endTime"`
	Initiator		string		`json:"initiator"`
	Participators	[]string	`json:"participators"`
}

var meetingDB = MeetingDB{Database: Database{schema: "Meeting"}}

func (m *MeetingDB) GetMeetingByTitle(title string) Meeting {
	for _, item := range m.Data {
		if item.Title == title {
			return item
		}
	}
	return Meeting{}
}

func (m *MeetingDB) GetOccupiedParticipators(title string, startTime, endTime time.Time) map[string]bool{
	occupiedParticipators := make(map[string]bool)
	for _, item := range m.Data {
		if item.Title != title && item.EndTime.After(startTime) && item.StartTime.Before(endTime) {
			occupiedParticipators[item.Initiator] = true
			for _, participator := range item.Participators {
				occupiedParticipators[participator] = true
			}
		}
	}
	return occupiedParticipators
}

func (m *MeetingDB) CreateMeeting(meeting Meeting) {
	m.isDirty = true
	m.Data = append(m.Data, meeting)
}

func (m *MeetingDB) SetMeeting(title string, startTime time.Time, setStart bool,
	endTime time.Time, setEnd bool, participators []string, setPars bool) {
	m.isDirty = true
	for index, item := range m.Data {
		if item.Title == title {
			if setStart {
				m.Data[index].StartTime = startTime
			}
			if setEnd {
				m.Data[index].EndTime = endTime
			}
			if setPars {
				m.Data[index].Participators = participators
			}
			return
		}
	}
}

func ReleaseMeetingModel() {
	meetingDB.releaseModel(&meetingDB.Data)
}

func ReleaseMeetingModelWithFile(f *os.File) {
	meetingDB.releaseModelWithFile(&meetingDB.Data, f)
}

func (m *Manager) Meeting() MeetingModel {
	if meetingDB.isInit == false {
		meetingDB.initModel(&meetingDB.Data)
	}
	return &meetingDB
}
