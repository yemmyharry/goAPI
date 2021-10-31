package dto

type BookUpdateDTO struct {
	ID          uint64 `json:"id,omitempty" form:"id" binding:"required"`
	Title       string `json:"title,omitempty" form:"title" binding:"required"`
	Description string `json:"description,omitempty" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

type BookCreateDTO struct {
	Title       string `json:"title,omitempty" form:"title" binding:"required"`
	Description string `json:"description,omitempty" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}
