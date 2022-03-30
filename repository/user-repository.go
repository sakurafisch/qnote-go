package repository

import (
	"errors"

	"github.com/astaxie/beego/logs"
	"github.com/sakurafisch/qnote-go/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var UserRepository userRepository

type userRepository struct{}

func (*userRepository) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	db := MainDB.First(&user, "email = ?", email)
	if err := db.Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, nil
		}
		logs.Error(err)
		return nil, err
	}
	return &user, nil
}

func (this *userRepository) Register(email string, password string) error {
	user, err := this.GetByEmail(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if user != nil {
		return errors.New("Username or Email already exist")
	}
	passwdHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logs.Error(err)
		return errors.New("user-repository: Failed to generate passwdHash using bcrypt")
	}
	db := MainDB.Create(&entity.User{
		Email:      email,
		PasswdHash: string(passwdHash),
	})
	if err := db.Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
