package config

type Config struct {
	Log      Log
	Database Database
}

type Log struct {
	IsOpen bool
	Path   string
	File   string
}

type Database struct {
	Path        string
	UserFile    string
	MeetingFile string
	StatusFile  string
}

func Default() Config {
	return Config{
		Log: Log{
			IsOpen: true,
			Path:   ".",
			File:   "agenda.log",
		},
		Database: Database{
			Path:        "data",
			UserFile:    "user.json",
			MeetingFile: "meeting.json",
			StatusFile:  "status.json",
		},
	}
}
