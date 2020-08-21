package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (monster *Monster) Store(filePath string) bool {
	str, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("Marshal失败 err=", err)
		return false
	}
	err = ioutil.WriteFile(filePath, str, 0666)
	if err != nil {
		panic(err)
		return false
	}
	return true
}

func (monster *Monster) ReStore(filePath string, name string) bool {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
		return false
	}
	err = json.Unmarshal(data, monster)
	if err != nil {
		fmt.Println("UnMarshal失败 err=", err)
		return false
	}
	if monster.Name != name {
		return false
	}
	return true

}

// func Store(filePath string, monster Monster) bool {
// 	str := json.Marshal(monster)
// 	err := ioutil.WriteFile(filePath, []byte(str), 0666)
// 	res := true
// 	if err != nil {
// 		panic(err)
// 		res := false
// 	}
// 	return res
// }
