package main

import (
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type if_error struct {
	s string
}

func (e *if_error) Error() string {
	return e.s
}
var keys = make([]string, 0, 0)

func UniqueId(id string) int {

	if len(keys) == 0 {
		keys = append(keys, id)
	} else {
		for _, val := range keys {
			if val == id {
				return 0
			} else {
				keys = append(keys, id)
			}

		}

	}

	return 1
}

func CheckForName() error {
	return &if_error{
		"Name Field cannot be null",
	}

}
func CheckForEmail() error {
	return &if_error{
		"Email Field cannot be null",
	}

}
func CheckForPhone() error {
	return &if_error{
		" Mobile number should be of 10 digits",
	}

}
func CheckForActivity(activity string) {
	activity = "0"

}
func Validate(arr [][]string) []string {

	var valid_users []string

	for entry := 1; entry < 8; entry++ {
		if arr[entry][0] == "" {
			id := uuid.New()
			arr[entry][0] = id.String()
		}

		if arr[entry][0] != "" {
			if UniqueId(arr[entry][0]) == 0 {
				log.Warn("User already exist with user id ", arr[entry][0])
				fmt.Println()
			} else {
				if (len(arr[entry][3]) == 10 && arr[entry][1] != "") && (arr[entry][2] != "") {
					valid_users = append(valid_users, arr[entry][0])

				}

			}
			if arr[entry][1] == "" {
				err := CheckForName()
				log.Warn(err, " for user with user id ", arr[entry][0])
				fmt.Println()
			}
			if arr[entry][2] == "" {
				err := CheckForEmail()
				log.Warn(err, " for user with user id ", arr[entry][0])
				fmt.Println()
			}
			if len(arr[entry][3]) != 10 {
				err := CheckForPhone()
				log.Warn(err, " for user with user id", arr[entry][0])
				fmt.Println()
			}
			if arr[entry][4] == "" {
				CheckForActivity(arr[entry][4])
			}

		}

	}
	return valid_users

}
