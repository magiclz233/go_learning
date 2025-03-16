package service

import (
	"go_learning/todolist/internal/data"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (u *UserService) CreateUser(user *data.User) error {
	return u.db.Create(user).Error
}

func (u *UserService) UpdateUser(user *data.User) error {
	    // 先获取原有记录
		var oldUser data.User
		if err := u.db.First(&oldUser, user.ID).Error; err != nil {
			return err
		}
	
		// 只更新非零值字段，保留原有的 CreatedAt 等字段
		return u.db.Model(&oldUser).Updates(user).Error
}

func (u *UserService) DeleteUser(id uint) error {
	return u.db.Delete(&data.User{}, id).Error
}

func (u *UserService) GetUserList() (users []*data.User, e error) {
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserService) GetUserByID(id uint) (*data.User, error) {
	var user data.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

