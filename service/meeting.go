package service

import (
	"errors"
	"github.com/MegaShow/goagenda/lib/log"
	"github.com/MegaShow/goagenda/model"
	"time"
)

type MeetingService interface {
	CreateMeeting(title string, startTime time.Time, endTime time.Time, initiator string, participators []string) error
}

func CheckFreeParticipators(participators []string, initiator string, occupiedParticipators map[string]bool) bool{
	newParticipators := participators
	newParticipators = append(newParticipators, initiator)
	for _, participator := range newParticipators {
		_, occupied := occupiedParticipators[participator]
		if occupied  {
				return false
		}
	}
	return true
}

func RemoveDuplicatedParticipators(participators []string, initiator string) []string{
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

func (s *Service) CreateMeeting(title string, startTime time.Time, endTime time.Time, initiator string, participators []string) error{
	log.Verbose("check if title exists")
	titleMeeting := s.DB.Meeting().GetMeetingByTitle(title)
	if titleMeeting.Title == title {
		return errors.New("title already exists")
	}

	log.Verbose("check if some participator doesn't exist")
	emptyUser := model.User{}
	for _, participator := range participators {
		if s.DB.User().GetUserByName(participator) == emptyUser {
			return errors.New("some participator doesn't exist")
		}
	}

	log.Verbose("check if some participator is occupied")
	occupiedParticipators := s.DB.Meeting().GetOccupiedParticipators(startTime, endTime)
	if !CheckFreeParticipators(participators, initiator, occupiedParticipators) {
		return errors.New("some participators or you are occupied during the time")
	}

	log.Verbose("remove duplicated participators")
	finalParticipators := RemoveDuplicatedParticipators(participators, initiator)

	s.DB.Meeting().CreateMeeting(model.Meeting{
		Title: 			title,
		StartTime:		startTime,
		EndTime:		endTime,
		Initiator:		initiator,
		Participators:	finalParticipators,
	})
	return nil
}

func (s *Manager) Meeting() MeetingService {
	return s.GetService()
}