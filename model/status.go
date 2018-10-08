package model

type StatusDB struct {
	Data Status
	Database
}

type Status struct {
	Name string `json:"user"`
}

var StatusModel = StatusDB{ Database: Database{ schema: "Status" } }

func (m *StatusDB) GetStatus() Status {
	initStatusModel()
	return m.Data
}

func (m *StatusDB) SetStatus(status Status) {
	initStatusModel()
	m.isDirty = true
	m.Data = status
}

func initStatusModel() {
	StatusModel.initModel(&StatusModel.Data)
}

func ReleaseStatusModel() {
	StatusModel.releaseModel(&StatusModel.Data)
}

