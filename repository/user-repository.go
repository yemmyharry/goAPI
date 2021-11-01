package repository

import (
	"goAPI/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type UserRepository interface {
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.User
	ProfileUser(userID string) entity.User
}

type userConnection struct {
	connection *gorm.DB
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		panic("failed to hash password")
	}
	return string(hash)
}

func (u *userConnection) InsertUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	u.connection.Save(&user)
	return user
}

func (u *userConnection) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser entity.User
		u.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}
	u.connection.Save(&user)
	return user
}

func (u *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := u.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (u *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return u.connection.Where("email = ?", email).Take(&user)
}

func (u *userConnection) FindByEmail(email string) entity.User {
	var user entity.User
	u.connection.Where("email = ?", email).Take(&user)
	return user
}

func (u *userConnection) ProfileUser(userID string) entity.User {
	var user entity.User
	u.connection.Find(&user, userID)
	return user
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{connection: db}
}
