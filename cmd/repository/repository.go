package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type UserModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type JSONProvider struct {
	FileName string
}

func (j JSONProvider) Read() ([]UserModel,error) {
	var v []UserModel
	file, err := os.OpenFile(j.FileName,os.O_CREATE,0755)
	if err != nil {
		return v, err
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&v)
	if err != nil {
		return v, err
	}
	return v, nil
}

func (j JSONProvider) Write(b []byte) error {
	err := ioutil.WriteFile(j.FileName, b, 0644)
	return err
}
