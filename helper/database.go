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

func Paginate(page, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}

func FindBasedOnCompany(companyId string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("company_id = ?", companyId)
	}
}
