package dto

type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min:1"`
	Email    string `json:"email" form:"email" validate:"email" binding:"required"`
	Password string `json:"password" form:"password" validate:"min:6" binding:"required"`
}
