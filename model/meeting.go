package model

import (
	"os"
	"time"
)

type MeetingModel interface {
	GetMeetingByTitle(title string) Meeting
	GetMeetingsByUser(user string) []Meeting
	GetOccupiedParticipators(title string, startTime, endTime time.Time) map[string]bool
	CreateMeeting(meeting Meeting)
	AddMeeting(title string, participators []string)
  SetMeeting(title string, startTime time.Time, setStart bool, endTime time.Time, setEnd bool, participators []string, setPars bool)
	DeleteMeetingByTitle(title string) bool
	DeleteMeetingsByInitiator(name string) int
	QuitMeeting(title, user string) bool
	RemoveParticipators(title string, participators []string) bool
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

func search(s []string, x string) (index int) {
	for index = 0; index < len(s); index++ {
		if s[index] == x {
			return
		}
	}
	return
}

func (m *MeetingDB) GetMeetingByTitle(title string) Meeting {
	for _, item := range m.Data {
		if item.Title == title {
			return item
		}
	}
	return Meeting{}
}

func (m *MeetingDB) GetMeetingsByUser(user string) (res []Meeting) {
	for _, item := range m.Data {
		if item.Initiator == user || search(item.Participators, user) != len(item.Participators) {
			res = append(res, item)
		}
	}
	return
}

func (m *MeetingDB) GetOccupiedParticipators(title string, startTime, endTime time.Time) map[string]bool {
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
				}
			}
			break
		}
	}
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

func (m *MeetingDB) QuitMeeting(title, user string) bool {
	m.isDirty = true
	for i := 0; i < len(m.Data); i++ {
		if m.Data[i].Title == title {
			for j := 0; j < len(m.Data[i].Participators); j++ {
				if m.Data[i].Participators[j] == user {
					m.Data[i].Participators = append(m.Data[i].Participators[:j], m.Data[i].Participators[j+1:]...)
					return true
				}
			}
			return false
		}
	}
	return false
}

func (m *MeetingDB) RemoveParticipators(title string, participators []string) bool {
	m.isDirty = true
	for i := 0; i < len(m.Data); i++ {
		if m.Data[i].Title == title {
			for j := 0; j < len(m.Data[i].Participators); j++ {
				if search(participators, m.Data[i].Participators[j]) != len(participators) {
					m.Data[i].Participators = append(m.Data[i].Participators[:j], m.Data[i].Participators[j+1:]...)
				}
			}
			return true
		}
	}
	return false
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
