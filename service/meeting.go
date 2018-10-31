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
	SetMeeting(title string, startTime time.Time, setStart bool, endTime time.Time, setEnd bool,
		initiator string, participators []string, setPars bool) error
}

func CheckFreeParticipators(participators []string, initiator string, occupiedParticipators map[string]bool) (bool, string){
	newParticipators := participators
	newParticipators = append(newParticipators, initiator)
	for _, participator := range newParticipators {
		_, occupied := occupiedParticipators[participator]
		if occupied  {
				return false, participator
		}
	}
	return true, ""
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
		return errors.New("meeting '" + title + "' already exists")
	}

	log.Verbose("check if some participator doesn't exist")
	for _, participator := range participators {
		if s.DB.User().GetUserByName(participator).Name == "" {
			return errors.New("user '" + participator + "' doesn't exist")
		}
	}

	log.Verbose("check if some participator is occupied")
	occupiedParticipators := s.DB.Meeting().GetOccupiedParticipators(title, startTime, endTime)
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
		Title: 			title,
		StartTime:		startTime,
		EndTime:		endTime,
		Initiator:		initiator,
		Participators:	finalParticipators,
	})
	return nil
}

func (s *Service) SetMeeting(title string, startTime time.Time, setStart bool, endTime time.Time, setEnd bool,
	initiator string, participators []string, setPars bool) error{
	log.Verbose("check if meeting exists")
	titleMeeting := s.DB.Meeting().GetMeetingByTitle(title)
	if titleMeeting.Title == "" {
		return errors.New("meeting '" + title + "' doesn't exist")
	}

	log.Verbose("check if current user initiates the meeting")
	if titleMeeting.Initiator != initiator {
		return errors.New("you aren't the initiator of meeting '" + title + "'")
	}

	log.Verbose("check if new time interval is valid")
	if !setStart {
		startTime = titleMeeting.StartTime
	}
	if !setEnd {
		endTime = titleMeeting.EndTime
	}
	if !startTime.Before(endTime) {
		if setStart && !setEnd {
			return errors.New("new start time should be before old end time")
		} else {
			return errors.New("new end time should be after old start time")
		}
	}

	log.Verbose("check if some new participator doesn't exist")
	if setPars {
		for _, participator := range participators {
			if s.DB.User().GetUserByName(participator).Name == "" {
				return errors.New("user '" + participator + "' doesn't exist")
			}
		}
	}

	log.Verbose("check if some participator is occupied")
	if !setPars {
		participators = titleMeeting.Participators
	}
	occupiedParticipators := s.DB.Meeting().GetOccupiedParticipators(title, startTime, endTime)
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
	finalParticipators := participators
	if setPars {
		finalParticipators = RemoveDuplicatedParticipators(participators, initiator)
		sort.Strings(finalParticipators)
	}

	s.DB.Meeting().SetMeeting(title, startTime, setStart, endTime, setEnd, finalParticipators, setPars)
	return nil
}

func (s *Manager) Meeting() MeetingService {
	return s.GetService()
}
