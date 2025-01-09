package dto

/* ------------------------------ User Request ------------------------------ */

type UserRegisterRequest struct {
	FirstName string `json:"first_name" validate:"required,max=100,alpha"`
	LastName  string `json:"last_name" validate:"required,max=100"`
	Username  string `json:"username" validate:"required,max=100,alpha"`
	Email     string `json:"email" validate:"required,max=100,email"`
	Password  string `json:"password" validate:"required,min=8,max=100"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,max=100,email"`
	Password string `json:"password"`
}

type UserUpdateRequest struct {
	UserID      string `json:"user_id" validate:"required"`
	FirstName   string `json:"first_name" validate:"omitempty,max=100,alpha"`
	LastName    string `json:"last_name" validate:"omitempty,max=100,alpha"`
	Email       string `json:"email" validate:"omitempty,max=100,email"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,max=13,numeric"`
}

type UserValidate struct {
	ReferenceID string `json:"reference_id"`
	OTP         string `json:"otp"`
}

/* ----------------------------- Address Request ---------------------------- */

type CreateAddressRequest struct {
	UserID     string `json:"user_id"`
	Title      string `json:"title" validate:"omitempty,max=100,alpha"`
	Street     string `json:"street" validate:"omitempty,max=100,alpha"`
	Country    string `json:"country" validate:"omitempty,max=100,alpha"`
	City       string `json:"city" validate:"omitempty,max=100,alpha"`
	PostalCode string `json:"postal_code" validate:"omitempty,max=100,alpha"`
}

type UpdateAddressRequest struct {
	ID         uint   `json:"id"`
	UserID     string `json:"user_id"`
	Title      string `json:"title" validate:"omitempty,max=100,alpha"`
	Street     string `json:"street" validate:"omitempty,max=100,alpha"`
	Country    string `json:"country" validate:"omitempty,max=100,alpha"`
	City       string `json:"city" validate:"omitempty,max=100,alpha"`
	PostalCode string `json:"postal_code" validate:"omitempty,max=100,numeric"`
}

type DeleteAddressRequest struct {
	UserID string `json:"user_id" validate:"required"`
	ID     uint   `json:"id" validate:"required,gt=0"`
}

type FindAddressRequest struct {
	UserID string `json:"user_id"`
	ID     uint   `json:"id"`
}
