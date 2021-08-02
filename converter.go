package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Employee struct {
	Id       string `json:"id"`
	Name     string `json :"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone_number"`
	IsActive bool   `json:active_status`
}

func conToJson(csvData [][]string, valid_users []string) {
	var oneRecord Employee
	var allValidRecords []Employee

	for _, each := range csvData {
		for pos, user := range valid_users {
			if user == each[0] {
				oneRecord.Id = each[0]
				oneRecord.Name = each[1]
				oneRecord.Email = each[2]
				oneRecord.Phone = each[3]
				oneRecord.IsActive, _ = strconv.ParseBool(each[4])
				allValidRecords = append(allValidRecords, oneRecord)
				valid_users = append(valid_users[:pos], valid_users[pos+1:]...)

			}
		}
	}
	jsondata, err := json.MarshalIndent(allValidRecords,"","  ") // convert to JSON

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Information of valid users in the json format ")
	fmt.Println()
	fmt.Println(string(jsondata))
	jsonFile, err := os.Create("./data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(jsondata)
	jsonFile.Close()

}
