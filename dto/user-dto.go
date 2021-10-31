package dto

type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"id"`
	Name     string `json:"name" form:"name" binding:"name"`
	Email    string `json:"email" form:"email" binding:"email" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6"`
}

type UserCreateDTO struct {
	Name     string `json:"name" form:"name" binding:"name"`
	Email    string `json:"email" form:"email" binding:"email" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6"`
}
