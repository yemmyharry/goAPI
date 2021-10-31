package dto

//dto-data transfer object

type LoginDTO struct {
	Email    string `json:"email" form:"email" validate:"email" binding:"required"`
	Password string `json:"password" form:"password" validate:"min:6" binding:"required"`
}
