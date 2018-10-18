package model

type StatusModel interface {
	GetStatus() Status
	SetStatus(status Status)
}

type StatusDB struct {
	Data Status
	Database
}

type Status struct {
	Name string `json:"user"`
}

var statusDB = StatusDB{Database: Database{schema: "Status"}}

func (m *StatusDB) GetStatus() Status {
	return m.Data
}

func (m *StatusDB) SetStatus(status Status) {
	m.isDirty = true
	m.Data = status
}

func ReleaseStatusModel() {
	statusDB.releaseModel(&statusDB.Data)
}

func (m *Manager) Status() StatusModel {
	if statusDB.isInit == false {
		statusDB.initModel(&statusDB.Data)
	}
	return &statusDB
}
