package repositories

import (
	"github.com/wailman24/Go-chi-starter.git/internal/models"
	"github.com/wailman24/Go-chi-starter.git/pkg/db"
	"gorm.io/gorm"
)

type UserRepositorie struct {
	db *gorm.DB
}

func NewUserRepositorie() *UserRepositorie {
	return &UserRepositorie{
		db: db.Db,
	}
}

func (ur *UserRepositorie) CreateUser(user *models.User) error {

	err := ur.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepositorie) GetUserByEmail(user *models.UserLogin) (*models.UserLogin, error) {

	err := ur.db.Raw(`SELECT u.id,u.email, u.password FROM users u where u.email = ?`, user.Email).Scan(user).Error
	if err != nil {
		return nil, err
	}

	return user, err
}
