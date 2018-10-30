package service

import (
	"errors"
	"sort"
	"strings"
	"time"

	"github.com/MegaShow/goagenda/lib/log"
	"github.com/MegaShow/goagenda/model"
)

type MeetingService interface {
	CreateMeeting(title string, startTime, endTime time.Time, initiator string, participators []string) error
	DeleteMeeting(user, title string) error
	QuitMeeting(user, title string) error
	RemoveParticipators(user, title string, participators []string) error
	ListMeetings(user, title string, startTime, endTime time.Time) (string, error)
	AddMeeting(title string, participators []string, name string) error
}

func CheckFreeParticipators(participators []string, initiator string, occupiedParticipators map[string]bool) (bool, string) {
	newParticipators := participators
	newParticipators = append(newParticipators, initiator)
	for _, participator := range newParticipators {
		_, occupied := occupiedParticipators[participator]
		if occupied {
			return false, participator
		}
	}
	return true, ""
}

func RemoveDuplicatedParticipators(participators []string, initiator string) []string {
	clearMap := make(map[string]bool)
	for _, item := range participators {
		clearMap[item] = true
	}
	delete(clearMap, initiator)
	var finalParticipators []string
	for participator := range clearMap {
		finalParticipators = append(finalParticipators, participator)
	}
	return finalParticipators
}

func (s *Service) CreateMeeting(title string, startTime time.Time, endTime time.Time, initiator string, participators []string) error {
	log.Verbose("check if title exists")
	titleMeeting := s.DB.Meeting().GetMeetingByTitle(title)
	if titleMeeting.Title == title {
		return errors.New("title already exists")
	}

	log.Verbose("check if some participator doesn't exist")
	for _, participator := range participators {
		if s.DB.User().GetUserByName(participator).Name == "" {
			return errors.New("user '" + participator + "' doesn't exist")
		}
	}

	log.Verbose("check if some participator is occupied")
	occupiedParticipators := s.DB.Meeting().GetOccupiedParticipators(startTime, endTime)
	free, occupiedOne := CheckFreeParticipators(participators, initiator, occupiedParticipators)
	if !free {
		var begin string
		if initiator == occupiedOne {
			begin = "you are"
		} else {
			begin = "user '" + occupiedOne + "' is"
		}
		return errors.New(begin + " occupied during the time")
	}

	log.Verbose("remove duplicated participators")
	finalParticipators := RemoveDuplicatedParticipators(participators, initiator)
	sort.Strings(finalParticipators)

	s.DB.Meeting().CreateMeeting(model.Meeting{
		Title:         title,
		StartTime:     startTime,
		EndTime:       endTime,
		Initiator:     initiator,
		Participators: finalParticipators,
	})
	return nil
}

func (s *Service) AddMeeting(title string, participators []string, name string) error {
	log.Verbose("check if meeting is exit")
	meeting := s.DB.Meeting().GetMeetingByTitle(title)
	if meeting.Title == "" {
		return errors.New("meeting with title" + title + "doesn't exit")
	}

	log.Verbose("check if the user is the sponsor of the meeting")
	if meeting.Initiator != name {
		return errors.New("you can only add participators to your meeting")
	}

	log.Verbose("check if some participator doesn't exist")
	for _, participator := range participators {
		if s.DB.User().GetUserByName(participator).Name == "" {
			return errors.New("user '" + participator + "' doesn't exist")
		}
	}

	log.Verbose("check if some participator is occupied")
	occupiedParticipators := s.DB.Meeting().GetOccupiedParticipators(meeting.StartTime, meeting.EndTime)
	free, occupiedOne := CheckFreeParticipators(participators, "", occupiedParticipators)
	if !free {
		var begin string
		begin = "user '" + occupiedOne + "' is"
		return errors.New(begin + " occupied during the time")
	}

	s.DB.Meeting().AddMeeting(title, participators)
	return nil
}

func (s *Service) DeleteMeeting(user, title string) error {
	if title == "" {
		if s.DB.Meeting().DeleteMeetingsByInitiator(user) == 0 {
			return errors.New("no meeting matched and deleted")
		}
	} else {
		log.Verbose("check if title exists")
		meeting := s.DB.Meeting().GetMeetingByTitle(title)
		if meeting.Title == "" {
			return errors.New("no such meeting with the title \"" + title + "\"")
		}
		log.Verbose("check if you are the initiator of meeting")
		if meeting.Initiator == user {
			s.DB.Meeting().DeleteMeetingByTitle(title)
		} else {
			return errors.New("you are not the initiator of this meeting")
		}
	}
	return nil
}

func (s *Service) QuitMeeting(user, title string) error {
	meeting := s.DB.Meeting().GetMeetingByTitle(title)
	log.Verbose("check if title exists")
	if meeting.Title == "" {
		return errors.New("no such meeting with the title \"" + title + "\"")
	}
	log.Verbose("check if you are the initiator of meeting")
	if meeting.Initiator == user {
		return errors.New("you are the initiator of this meeting, please use delete command")
	} else if s.DB.Meeting().QuitMeeting(title, user) == false {
		return errors.New("you are not the participator of this meeting")
	}
	log.Verbose("check if it's necessary to delete the meeting")
	if len(meeting.Participators)-1 == 0 {
		s.DB.Meeting().DeleteMeetingByTitle(title)
		return errors.New("delete_meeting")
	}
	return nil
}

func (s *Service) RemoveParticipators(user, title string, participators []string) error {
	meeting := s.DB.Meeting().GetMeetingByTitle(title)
	log.Verbose("check if title exists")
	if meeting.Title == "" {
		return errors.New("no such meeting with the title \"" + title + "\"")
	}
	log.Verbose("check if you are the initiator of meeting")
	if meeting.Initiator != user {
		return errors.New("you are not the initiator of this meeting")
	}
	for i := 0; i < len(participators); i++ {
		var flag bool
		for j := 0; j < len(meeting.Participators); j++ {
			if participators[i] == meeting.Participators[j] {
				flag = true
			}
		}
		if !flag {
			return errors.New("user \"" + participators[i] + "\" is not participator of this meeting")
		}
	}
	s.DB.Meeting().RemoveParticipators(title, participators)
	if len(s.DB.Meeting().GetMeetingByTitle(title).Participators) == 0 {
		s.DB.Meeting().DeleteMeetingByTitle(title)
		return errors.New("delete_meeting")
	}
	return nil
}

func (s *Service) ListMeetings(user, title string, startTime, endTime time.Time) (string, error) {
	var meetings []model.Meeting
	log.Verbose("filter data by user and title")
	if title == "" {
		meetings = s.DB.Meeting().GetMeetingsByUser(user)
	} else {
		meetings = []model.Meeting{s.DB.Meeting().GetMeetingByTitle(title)}
		if len(meetings) == 1 && meetings[0].Initiator != user && sort.SearchStrings(meetings[0].Participators, user) == len(meetings[0].Participators) {
			meetings = []model.Meeting{}
		}
	}
	log.Verbose("filter data by time")
	for i := 0; i < len(meetings); i++ {
		if !meetings[i].StartTime.Before(startTime) && (endTime.Equal(time.Unix(0, 0)) || !meetings[i].EndTime.After(endTime)) {
			continue
		}
		meetings = append(meetings[:i], meetings[i+1:]...)
		i--
	}
	if len(meetings) == 0 {
		return "", errors.New("no meeting list")
	}
	sort.Slice(meetings, func(i, j int) bool { return meetings[i].StartTime.Before(meetings[j].StartTime) })
	output := "\n  ---------------------------\n"
	for _, item := range meetings {
		str := "  Title: " + item.Title + "\n  Initiator: " + item.Initiator + "\n  Participators: " + strings.Join(item.Participators, ",")
		str += "\n  Start Time: " + item.StartTime.Format("2006-01-02/15:04") + "\n  End Time: " + item.EndTime.Format("2006-01-02/15:04")
		str += "\n  ---------------------------\n"
		output += str
	}
	return output, nil
}

func (s *Manager) Meeting() MeetingService {
	return s.GetService()
}
