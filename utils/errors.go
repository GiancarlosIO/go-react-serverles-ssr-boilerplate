package utils

import "log"

type ErrorType string

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}