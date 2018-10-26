package model

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"go/types"
	"io"
	"os"
)

type Manager types.Basic

type Database struct {
	isInit  bool
	isDirty bool
	path    string
	file    string
	schema  string
}

func (m *Database) initModel(data interface{}) {
	if m.isInit == false {
		m.path = viper.GetString("Database.Path")
		m.file = viper.GetString("Database." + m.schema + "File")
		if _, err := os.Stat(m.path); err != nil {
			err := os.MkdirAll(m.path, 0777)
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
		}
		f, err := os.OpenFile(m.path+string(os.PathSeparator)+m.file, os.O_CREATE|os.O_RDONLY, 0777)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		defer f.Close()
		decoder := json.NewDecoder(f)
		err = decoder.Decode(data)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			os.Exit(2)
		}
		m.isInit = true
	}
}

func (m *Database) releaseModel(data interface{}) {
	if m.isDirty == true {
		f, err := os.OpenFile(m.path+string(os.PathSeparator)+m.file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		defer f.Close()
		m.releaseModelWithFile(data, f)
		m.isDirty = false
	}
}

func (m *Database) releaseModelWithFile(data interface{}, f *os.File) {
	if m.isDirty == true {
		encoder := json.NewEncoder(f)
		err := encoder.Encode(data)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		m.isDirty = false
	}
}
