package userRepository

import (
	"finpro-fenlie/model/entity"
	"finpro-fenlie/src/user"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepository {
	return &userRepository{db}
}

func (repo *userRepository) InsertUser(payload *entity.User) error {
	if err := repo.db.Create(&payload).Error; err != nil {
		return err
	}

	return nil
}

// func (repo *userRepository) RetrieveAllUser(page, size int, totalData int64, email, name string, userInfo *middlewareDto.UserInfo) (userDto.GetResponse, error) {
// 	var users []userDto.User
// 	var response userDto.GetResponse

// 	query := repo.db.Model(&userDto.User{})

// 	if email != "" {
// 		query = query.Where("email = ?", email)
// 	}

// 	if name != "" {
// 		query = query.Where("LOWER(name) ILIKE ?", "%"+strings.ToLower(name)+"%")
// 	}

// 	if page != 0 && size != 0 {
// 		query = query.Limit(size).Offset((page - 1) * size)
// 	}

// 	if err := query.Select("id", "name", "email", "company_id").Where("company_id = ? AND deleted_at IS NULL", userInfo.CompanyID).Find(&users).Error; err != nil {
// 		return response, err
// 	}

// 	response.Data = users
// 	response.TotalData = totalData
// 	if page != 0 && size != 0 {
// 		response.Pagination = userDto.Paging{Page: page, Size: size}
// 	} else {
// 		response.Pagination = userDto.Paging{Page: 1, Size: 10}

// 	}

// 	return response, nil
// }

// // func (repo *userRepository) RetrieveUserByID(id string) (entity.User, error) {
// 	var user entity.User
// 	if err := repo.db.Select("id", "name", "email", "company_id").Where("id = ? AND deleted_at IS NULL", id).First(&user).Error; err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }

// func (repo *userRepository) EditUser(payload *entity.User) error {
// 	var existingUser userDto.User
// 	if err := repo.db.Where("id = ? AND company_id = ? AND deleted_at IS NULL", id, userInfo.CompanyID).First(&existingUser).Error; err != nil {
// 		return err
// 	}

// 	if password, ok := userUpdates["password"]; ok {
// 		if err := validation.ValidatePassword(password.(string)); err != nil {
// 			return err
// 		}
// 		hash, err := bcrypt.GenerateFromPassword([]byte(password.(string)), bcrypt.DefaultCost)
// 		if err != nil {
// 			return err
// 		}
// 		userUpdates["password"] = string(hash)
// 	}

// 	if err := repo.db.Model(&existingUser).Updates(userUpdates).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (repo *userRepository) RemoveUser(id string, userInfo *middlewareDto.UserInfo) error {
// 	var dbUser userDto.User
// 	if err := repo.db.Select("email").Where("company_id = ? AND id = ? AND deleted_at IS NULL", userInfo.CompanyID, id).First(&dbUser).Error; err != nil {
// 		return err
// 	}

// 	if userInfo.Email == dbUser.Email {
// 		return errors.New("you can't delete your own account")
// 	}

// 	if err := repo.db.Model(&userDto.User{}).Where("company_id = ? AND id = ?", userInfo.CompanyID, id).Update("deleted_at", gorm.Expr("CURRENT_TIMESTAMP")).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

func (repo *userRepository) RetrieveUserByEmail(email string) (entity.User, error) {
	var user entity.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
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
