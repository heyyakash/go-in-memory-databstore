package helper

import (
	"net/http"
	"strconv"
)

func CheckExpiry(w http.ResponseWriter, command string, time *int) error {
	i, err := strconv.Atoi(command)
	if err != nil {
		return err
	}
	*time = i
	return nil
}
