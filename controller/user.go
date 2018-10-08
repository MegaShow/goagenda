package controller

var userCtrl UserCtrl

func GetUserCtrl() *UserCtrl {
	return (*UserCtrl)(initController((*Controller)(&userCtrl)))
}
