package controller

import (
	"demo/cmd/repository"
	"encoding/json"
	"fmt"
)

type Task struct {
	FileName string
}

func  (t Task) AddList(n string , age int)  {
	repo := repository.JSONProvider{FileName: t.FileName}
	v,err := repo.Read()
	if err != nil {
		fmt.Println(err)
	}
	v = append(v,repository.UserModel{Id: len(v)+1,Name: n,Age: age})
	result, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
	} else {
		err := repo.Write(result)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(v)
		}
	}
}

func (t Task) ShowList() {
	repo := repository.JSONProvider{FileName: t.FileName}
	v,err := repo.Read()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
}
