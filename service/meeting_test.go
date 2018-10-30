package service

import (
	"encoding/json"
	"fmt"
	"github.com/MegaShow/goagenda/model"
	"github.com/spf13/viper"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func t0(tStr string) time.Time {
	parse, e := time.Parse("2006-1-2/15:4", tStr)
	if e != nil {
		panic(e)
	}
	return parse
}

func compareMeetingsWithFile(filename1, filename2 string) bool {
	var d1, d2 []model.Meeting
	f1, err := os.OpenFile(filename1, os.O_CREATE|os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	decoder := json.NewDecoder(f1)
	err = decoder.Decode(&d1)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		os.Exit(2)
	}
	f2, err := os.OpenFile(filename2, os.O_CREATE|os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	decoder = json.NewDecoder(f2)
	err = decoder.Decode(&d2)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		os.Exit(2)
	}
	return compareMeetings(d1, d2)
}

func compareMeetings(d1, d2 []model.Meeting) bool {
	if len(d1) != len(d2) {
		fmt.Println("Length of meetings diff")
		return false
	}
	for i := 0; i < len(d1); i++ {
		if d1[i].Title != d2[i].Title || d1[i].Initiator != d2[i].Initiator || !d1[i].StartTime.Equal(d2[i].StartTime) || !d1[i].EndTime.Equal(d2[i].EndTime) {
			fmt.Printf("DB[%d] of meetings diff\n", i)
			fmt.Printf("Expected: t %s, u %s, s %s, e %s, p %v\n", d1[i].Title, d1[i].Initiator, d1[i].StartTime.String(), d1[i].EndTime.String(), d1[i].Participators)
			fmt.Printf("Actual:   t %s, u %s, s %s, e %s, p %v\n", d2[i].Title, d2[i].Initiator, d2[i].StartTime.String(), d2[i].EndTime.String(), d2[i].Participators)
			return false
		}
		if len(d1[i].Participators) != len(d2[i].Participators) {
			fmt.Printf("Participators of DB[%d] of meetings diff\n", i)
			fmt.Printf("Expected: %v\n", d1[i].Participators)
			fmt.Printf("Actual:   %v\n", d2[i].Participators)
			return false
		}
		for j := 0; j < len(d1[i].Participators); j++ {
			if d1[i].Participators[j] != d2[i].Participators[j] {
				fmt.Printf("Participators of DB[%d] of meetings diff\n", i)
				fmt.Printf("Expected: %v\n", d1[i].Participators)
				fmt.Printf("Actual:   %v\n", d2[i].Participators)
				return false
			}
		}
	}
	return true
}

func (c *callor) call_m_c(s Manager, title, startTime, endTime, initiator string, participators []string) error {
	c.test.Logf("agenda m c -t \"%s\" -s \"%s\" -e \"%s\" -p \"%s\"", title, startTime, endTime, strings.Join(participators, ","))
	return s.Meeting().CreateMeeting(title, t0(startTime), t0(endTime), initiator, participators)
}

func TestService_CreateMeeting(t *testing.T) {
	c := callor{test: t}
	var s Manager
	viper.Set("Database.Path", "../test/userfile")
	viper.Set("Database.UserFile", "5_users.json")
	s.DB.User()
	viper.Set("Database.Path", "../test/meetingfile")
	viper.Set("Database.MeetingFile", "empty.json")

	err := c.call_m_c(s, "me1", "2018-10-26/09:00", "2018-10-26/10:00", "Amy", []string{"Bob", "Frank"})
	if err.Error() != "user 'Frank' doesn't exist" {
		t.Error("Info: err != \"user 'Frank' doesn't exist\", err = " + err.Error())
		t.FailNow()
	}

	err = c.call_m_c(s, "me1", "2018-10-26/09:00", "2018-10-26/11:00", "Amy", []string{"Bob", "Cici"})
	if err != nil {
		t.Error("Error: err = " + err.Error())
		t.FailNow()
	}

	err = c.call_m_c(s, "me1", "2018-10-26/13:00", "2018-10-26/15:00", "Amy", []string{"Bob", "Cici"})
	if err.Error() != "title already exists" {
		t.Error("Info: err != \"title already exists\", err = " + err.Error())
		t.FailNow()
	}

	err = c.call_m_c(s, "me2", "2018-10-26/08:00", "2018-10-26/9:30", "Amy", []string{"Bob", "Duke"})
	if err.Error() != "user 'Bob' is occupied during the time" {
		t.Error("Info: err != \"user 'Bob' is occupied during the time\", err = " + err.Error())
		t.FailNow()
	}

	err = c.call_m_c(s, "me2", "2018-10-26/08:00", "2018-10-26/9:00", "Amy", []string{"Bob", "Duke", "Duke", "Amy"})
	if err != nil {
		t.Error("Error: err = " + err.Error())
		t.FailNow()
	}

	err = c.call_m_c(s, "me3", "2018-10-26/10:5", "2018-10-26/12:00", "Duke", []string{"Amy", "Ella"})
	if err.Error() != "user 'Amy' is occupied during the time" {
		t.Error("Info: err != \"user 'Amy' is occupied during the time\", err = " + err.Error())
		t.FailNow()
	}

	err = c.call_m_c(s, "me3", "2018-10-26/8:7", "2018-10-26/12:00", "Duke", []string{"Ella"})
	if err.Error() != "you are occupied during the time" {
		t.Error("Info: err != \"you are occupied during the time\", err = " + err.Error())
		t.FailNow()
	}

	err = c.call_m_c(s, "me3", "2018-10-27/16:00", "2018-10-27/17:00", "Duke", []string{"Cici", "Ella"})
	if err != nil {
		t.Error("Error: err = " + err.Error())
		t.FailNow()
	}

	err = c.call_m_c(s, "me4", "2018-10-26/9:30", "2018-10-26/10:30", "Duke", []string{"Ella"})
	if err != nil {
		t.Error("Error: err = " + err.Error())
		t.FailNow()
	}

	f, err := os.OpenFile("../test/result/meeting.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer f.Close()
	model.ReleaseMeetingModelWithFile(f)

	if !compareMeetingsWithFile("../test/meetingfile/m_c_result.json", "../test/result/meeting.json") {
		t.FailNow()
	}
}
