package helper

import (
	"errors"

	"gorm.io/gorm"
)

func CheckErrNotFound(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return nil
}
