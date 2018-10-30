package model

import (
	"fmt"
	"time"
)

type MeetingModel interface {
	GetMeetingByTitle(title string) Meeting
	GetOccupiedParticipators(startTime, endTime time.Time) map[string]bool
	CreateMeeting(meeting Meeting)
	AddMeeting(title string, participators []string)
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

func (m *MeetingDB) AddMeeting(title string, participators []string) {
	m.isDirty = true
	for index, item := range m.Data {
		if item.Title == title {
			participatorMap := make(map[string]bool)
			for _, participator := range item.Participators {
				participatorMap[participator] = true
			}
			for _, participator := range participators {
				_, hasAdd := participatorMap[participator]
				if !hasAdd {
					participatorMap[participator] = true
					m.Data[index].Participators = append(item.Participators, participator)
					fmt.Println(participator)
				}
			}
			break
		}
	}
}

func ReleaseMeetingModel() {
	meetingDB.releaseModel(&meetingDB.Data)
}

func (m *Manager) Meeting() MeetingModel {
	if meetingDB.isInit == false {
		meetingDB.initModel(&meetingDB.Data)
	}
	return &meetingDB
}
