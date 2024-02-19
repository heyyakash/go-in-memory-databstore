package helper

import (
	"errors"
	"net/http"
	"strconv"
)

func CheckExpiry(w http.ResponseWriter, command string, time interface{}) error {
	i, err := strconv.ParseFloat(command, 64)
	if err != nil {
		return err
	}

	switch v := time.(type) {
	case *float64:
		*v = i
	case *int:
		*v = int(i)
	default:
		return errors.New("unsupported type for time")
	}

	return nil
}
