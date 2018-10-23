package controller

import (
	"github.com/MegaShow/goagenda/lib/log"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type Ctx struct {
	Value *viper.Viper
	Visit map[string]bool
	User  User
}

type user struct {
	get func() string
	set func(string) error
}

type User interface {
	Get() string
	Set(string) error
}

func (u *user) Get() string {
	return u.get()
}

func (u *user) Set(name string) error {
	return u.set(name)
}

func (c *Ctx) Get(key string) (interface{}, bool) {
	value, visit := c.Value.Get(key), c.Visit[key]
	if visit {
		log.AddParams(key, value)
	}
	return value, visit
}

func (c *Ctx) GetSecret(key string) (interface{}, bool) {
	value, visit := c.Value.Get(key), c.Visit[key]
	if visit {
		log.AddParams(key, "*")
	}
	return value, visit
}

func (c *Ctx) GetString(key string) (string, bool) {
	value, visit := c.Get(key)
	return cast.ToString(value), visit
}

func (c *Ctx) GetSecretString(key string) (string, bool) {
	value, visit := c.GetSecret(key)
	return cast.ToString(value), visit
}
