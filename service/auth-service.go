package service

import (
	"github.com/mashingan/smapping"
	"goAPI/dto"
	"goAPI/entity"
	"goAPI/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.UserCreateDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicate(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func comparePassword(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (a *authService) VerifyCredential(email string, password string) interface{} {
	res := a.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparedPwd := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPwd {
			return res
		}
		return false
	}
	return false
}

func (a authService) CreateUser(user dto.UserCreateDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("failed map %v", err)
	}
	res := a.userRepository.InsertUser(userToCreate)
	return res
}

func (a authService) FindByEmail(email string) entity.User {
	return a.userRepository.FindByEmail(email)
}

func (a authService) IsDuplicate(email string) bool {
	res := a.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepository: userRepo}
}
