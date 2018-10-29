package service

import (
	"errors"
	"github.com/MegaShow/goagenda/lib/log"
	"github.com/MegaShow/goagenda/model"
	"sort"
	"time"
)

type MeetingService interface {
	CreateMeeting(title string, startTime time.Time, endTime time.Time, initiator string, participators []string) error
	DeleteMeeting(user string, title string) error
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

func (s *Service) DeleteMeeting(user string, title string) error {
	if title == "" {
		if s.DB.Meeting().DeleteMeetingsByInitiator(user) == 0 {
			return errors.New("no meeting matched and deleted")
		}
	} else {
		meeting := s.DB.Meeting().GetMeetingByTitle(title)
		if meeting.Title == "" {
			return errors.New("no such meeting with the title")
		} else if meeting.Initiator == user {
			s.DB.Meeting().DeleteMeetingByTitle(title)
		} else {
			return errors.New("you are not the initiator of this meeting")
		}
	}
	return nil
}

func (s *Manager) Meeting() MeetingService {
	return s.GetService()
}
