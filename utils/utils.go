package utils

import "log"

func DieIfErr(err error, message string) {
	if err != nil {
		log.Fatalf(message + "\n")
	}
}

func WarnIfErr(err error, message string) {
	if err != nil {
		log.Printf(message + "\n")
	}
}
