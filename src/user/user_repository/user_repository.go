package userRepository

import (
	"finpro-fenlie/helper"
	"finpro-fenlie/model/entity"
	"finpro-fenlie/src/user"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// RetrieveAllUser implements user.UserRepository.
func (repo *userRepository) RetrieveAllUser(page, size int, email, name, companyId string) ([]entity.User, int, error) {
	var users []entity.User
	var total int64

	err := repo.db.Model(&entity.User{}).Scopes(helper.FindBasedOnCompany(companyId), helper.Paginate(page, size)).Where("users.email LIKE $1 AND users.name LIKE $2", "%"+email+"%", "%"+name+"%").Count(&total).Joins("Company", repo.db.Select("Company.name")).Find(&users).Error

	return users, int(total), err
}

func (repo *userRepository) InsertUser(payload *entity.User) error {
	if err := repo.db.Create(&payload).Error; err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) RetrieveUserByID(id, companyId string) (entity.User, error) {
	var user entity.User
	if err := repo.db.Scopes(helper.FindBasedOnCompany(companyId)).Joins("Company", repo.db.Select("Company.name")).Select("users.id", "users.name", "users.email", "users.role", "users.password").First(&user, "users.id = $1", id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (repo *userRepository) EditUser(payload *entity.User) error {
	err := repo.db.Model(payload).Scopes(helper.FindBasedOnCompany(payload.CompanyID)).Omit("email", "id", "company_id").Updates(&payload).Error
	return err
}

func (repo *userRepository) RemoveUser(id, companyId string) error {
	fmt.Println("IDDDDD", id)
	result := repo.db.Scopes(helper.FindBasedOnCompany(companyId)).Delete(&entity.User{ID: id})
	if result.RowsAffected < 1 {
		return errors.New("cannot find the requested data")
	}

	return nil
}

func (repo *userRepository) RetrieveUserByEmail(email, companyId string) (entity.User, error) {
	var user entity.User
	if err := repo.db.Scopes(helper.FindBasedOnCompany(companyId)).Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// func (repo *userRepository) CountUsers(email, name string, userInfo *middlewareDto.UserInfo) (int64, error) {
// 	var count int64
// 	query := repo.db.Select("id", "name", "email").Where("company_id = ? AND deleted_at IS NULL", userInfo.CompanyID).Model(&userDto.User{})

// 	if email != "" {
// 		query = query.Where("email = ?", email)
// 	}

// 	if name != "" {
// 		query = query.Where("LOWER(name) ILIKE ?", "%"+strings.ToLower(name)+"%")
// 	}

// 	if err := query.Count(&count).Error; err != nil {
// 		return 0, err
// 	}

// 	return count, nil
// }

func NewUserRepository(db *gorm.DB) user.UserRepository {
	return &userRepository{db}
}
