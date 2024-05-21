package userRepository

import (
	"errors"
	"finpro-fenlie/model/dto/middlewareDto"
	"finpro-fenlie/model/dto/userDto"
	"finpro-fenlie/pkg/middleware"
	"finpro-fenlie/pkg/validation"
	"finpro-fenlie/src/user"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepository {
	return &userRepository{db}
}

func (repo *userRepository) RetrieveLoginUser(ctx *gin.Context, req middlewareDto.LoginRequest, user userDto.User) (string, error) {
	// Validate password
	if !comparePassword(req.Password, user.Password) {
		return "", errors.New("invalid password")
	}

	companyID := user.CompanyID.String()

	// Generate JWT token
	token, err := middleware.GenerateTokenJwt(req.Email, user.Role, companyID, 60)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (repo *userRepository) InsertUser(c *gin.Context, user userDto.User, checkEmail, checkPass bool) error {
	userLogged, err := middleware.GetUserInfo(c)
	if err != nil {
		return err
	}

	if !checkEmail || !checkPass {
		return errors.New("email or password validation failed")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	companyID, err := uuid.Parse(userLogged.CompanyID)
	if err != nil {
		return err
	}

	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}

	user.Password = string(hash)
	user.CompanyID = companyID
	if err := repo.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) RetrieveAllUser(c *gin.Context, page, size int, totalData int64, email, name string) (userDto.GetResponse, error) {
	var users []userDto.User
	var response userDto.GetResponse
	userLogged, err := middleware.GetUserInfo(c)
	if err != nil {
		return userDto.GetResponse{}, err
	}

	query := repo.db.Model(&userDto.User{})

	if email != "" {
		query = query.Where("email = ?", email)
	}

	if name != "" {
		query = query.Where("LOWER(name) ILIKE ?", "%"+strings.ToLower(name)+"%")
	}

	if page != 0 && size != 0 {
		query = query.Limit(size).Offset((page - 1) * size)
	}

	if err := query.Select("id", "name", "email", "company_id").Where("company_id = ? AND deleted_at IS NULL", userLogged.CompanyID).Find(&users).Error; err != nil {
		return response, err
	}

	response.Data = users
	response.TotalData = totalData
	if page != 0 && size != 0 {
		response.Pagination = userDto.Paging{Page: page, Size: size}
	} else {
		response.Pagination = userDto.Paging{Page: 1, Size: 10}

	}

	return response, nil
}

func (repo *userRepository) RetrieveUserByID(c *gin.Context, id string) (userDto.User, error) {
	userLogged, err := middleware.GetUserInfo(c)
	if err != nil {
		return userDto.User{}, err
	}

	var user userDto.User
	if err := repo.db.Select("id", "name", "email", "company_id").Where("company_id = ? AND id = ? AND deleted_at IS NULL", userLogged.CompanyID, id).First(&user).Error; err != nil {
		return userDto.User{}, err
	}
	return user, nil
}

func (repo *userRepository) EditUser(c *gin.Context, id string, userUpdates map[string]interface{}) error {
	userLogged, err := middleware.GetUserInfo(c)
	if err != nil {
		return err
	}

	var existingUser userDto.User
	if err := repo.db.Where("id = ? AND company_id = ? AND deleted_at IS NULL", id, userLogged.CompanyID).First(&existingUser).Error; err != nil {
		return err
	}

	if password, ok := userUpdates["password"]; ok {
		if err := validation.ValidatePassword(password.(string)); err != nil {
			return err
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(password.(string)), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		userUpdates["password"] = string(hash)
	}

	if err := repo.db.Model(&existingUser).Updates(userUpdates).Error; err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) RemoveUser(c *gin.Context, id string) error {
	userLogged, err := middleware.GetUserInfo(c)
	if err != nil {
		return errors.New("unauthorized")
	}

	var dbUser userDto.User
	if err := repo.db.Select("email").Where("company_id = ? AND id = ? AND deleted_at IS NULL", userLogged.CompanyID, id).First(&dbUser).Error; err != nil {
		return err
	}

	if userLogged.Email == dbUser.Email {
		return errors.New("you can't delete your own account")
	}

	if err := repo.db.Model(&userDto.User{}).Where("company_id = ? AND id = ?", userLogged.CompanyID, id).Update("deleted_at", gorm.Expr("CURRENT_TIMESTAMP")).Error; err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) RetrieveUserByEmail(email string) (userDto.User, error) {
	var user userDto.User
	if err := repo.db.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (repo *userRepository) CountUsers(c *gin.Context, email, name string) (int64, error) {
	userLogged, err := middleware.GetUserInfo(c)
	if err != nil {
		return 0, err
	}

	var count int64
	query := repo.db.Select("id", "name", "email").Where("company_id = ? AND deleted_at IS NULL", userLogged.CompanyID).Model(&userDto.User{})

	if email != "" {
		query = query.Where("email = ?", email)
	}

	if name != "" {
		query = query.Where("LOWER(name) ILIKE ?", "%"+strings.ToLower(name)+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (repo *userRepository) CheckUserEmailPassword(user userDto.User) (bool, bool, error) {
	var dbUser userDto.User
	err := repo.db.Where("email = ?", user.Email).First(&dbUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			dbUser = userDto.User{}
		} else {
			return false, false, err
		}
	}

	if err := validation.ValidateEmail(user.Email, dbUser.Email); err != nil {
		return false, false, err
	}

	if err := validation.ValidatePassword(user.Password); err != nil {
		return false, false, err
	}

	return true, true, nil
}

func comparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
