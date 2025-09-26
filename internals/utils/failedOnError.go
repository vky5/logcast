package utils

import "log"

func FailedOnError(err error, packageName string, msg string) error {
	if err != nil {
		log.Printf("🚨 [%s] %s: %s", packageName, msg, err)
		return err
	}

	return nil
}
