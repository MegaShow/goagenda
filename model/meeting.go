package model

import (
	"os"
	"time"
)

type MeetingModel interface {
	GetMeetingByTitle(title string) Meeting
	GetOccupiedParticipators(startTime, endTime time.Time) map[string]bool
	CreateMeeting(meeting Meeting)
	DeleteMeetingByTitle(title string) bool
	DeleteMeetingsByInitiator(name string) int
}

type MeetingDB struct {
	Data []Meeting
	Database
}

type Meeting struct {
	Title         string    `json:"title"`
	StartTime     time.Time `json:"startTime"`
	EndTime       time.Time `json:"endTime"`
	Initiator     string    `json:"initiator"`
	Participators []string  `json:"participators"`
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

func (m *MeetingDB) GetOccupiedParticipators(startTime, endTime time.Time) map[string]bool {
	occupiedParticipators := make(map[string]bool)
	for _, item := range m.Data {
		if item.EndTime.After(startTime) && item.StartTime.Before(endTime) {
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

func (m *MeetingDB) DeleteMeetingByTitle(title string) bool {
	m.isDirty = true
	for i := 0; i < len(m.Data); i++ {
		if m.Data[i].Title == title {
			m.Data = append(m.Data[:i], m.Data[i+1:]...)
			return true
		}
	}
	return false
}

func (m *MeetingDB) DeleteMeetingsByInitiator(name string) (count int) {
	m.isDirty = true
	for i := 0; i < len(m.Data); i++ {
		if m.Data[i].Initiator == name {
			m.Data = append(m.Data[:i], m.Data[i+1:]...)
			i--
			count++
		}
	}
	return
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
