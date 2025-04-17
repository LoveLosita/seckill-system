package dao

import (
	"errors"
	"gorm.io/gorm"
	"kitex-server/init"
	"kitex-server/users/kitex_gen/user"
	"kitex-server/users/model"
	"kitex-server/users/response"
)

func GetUserHashedPassword(userName string) (string, user.Status) {
	var pwdUser model.User
	result := init.Db.Table("users").Where("username = ?", userName).Find(&pwdUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", response.WrongUsrName
		} else {
			return "", response.InternalErr(result.Error)
		}
	}
	return pwdUser.Password, user.Status{}
}

func InsertUserInfo(newUser user.UserRegisterRequest) user.Status {
	result := init.Db.Table("users").Create(&newUser)
	if result.Error != nil {
		return response.InternalErr(result.Error)
	}
	return user.Status{}
}

func IfUsernameExists(name string) (bool, user.Status) {
	var nameUser model.User
	result := init.Db.Table("users").First(&nameUser, "username = ?", name)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, user.Status{}
		}
		return false, response.InternalErr(result.Error)
	}
	return true, user.Status{}
}

func IfEmailExists(email string) (bool, user.Status) {
	var emailUser model.User
	result := init.Db.Table("users").First(&emailUser, "email = ?", email)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, user.Status{}
		}
		return false, response.InternalErr(result.Error)
	}
	return true, user.Status{}
}

func IfPhoneNumberExists(phoneNumber string) (bool, user.Status) {
	var phoneUser model.User
	result := init.Db.Table("users").First(&phoneUser, "phone_number = ?", phoneNumber)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, user.Status{}
		}
		return false, response.InternalErr(result.Error)
	}
	return true, user.Status{}
}

func GetUserIDByName(name string) (int64, user.Status) {
	var userID int64
	result := init.Db.Table("users").Select("id").Where("username = ?", name).Scan(&userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, response.WrongUsrName
		}
		return 0, response.InternalErr(result.Error)
	}
	return userID, user.Status{}
}
