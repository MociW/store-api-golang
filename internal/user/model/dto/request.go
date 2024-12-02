package dto

/* ------------------------------ User Request ------------------------------ */

type UserRegisterRequest struct {
	FirstName string `json:"first_name" validate:"required,max=100,alpha"`
	LastName  string `json:"last_name" validate:"required,max=100,alpha"`
	Username  string `json:"user_name" validate:"required,max=100,alpha"`
	Email     string `json:"email" validate:"required,max=100,email"`
	Password  string `json:"password" validate:"required,min=8,max=100"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,max=100,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type UserUpdateRequest struct {
	FirstName   string `json:"first_name" validate:"omitempty,max=100,alpha"`
	LastName    string `json:"last_name" validate:"omitempty,max=100,alpha"`
	Email       string `json:"email" validate:"omitempty,max=100,email"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,max=13,numeric"`
}

/* ----------------------------- Address Request ---------------------------- */

type CreateAddressRequest struct {
	Title      string `json:"title" validate:"omitempty,max=100,alpha"`
	Street     string `json:"street" validate:"omitempty,max=100,alpha"`
	Country    string `json:"country" validate:"omitempty,max=100,alpha"`
	City       string `json:"city" validate:"omitempty,max=100,alpha"`
	PostalCode string `json:"postal_code" validate:"omitempty,max=100,alpha"`
}

type UpdateAddressRequest struct {
	ID         uint   `json:"id"`
	Title      string `json:"title" validate:"omitempty,max=100,alpha"`
	Street     string `json:"street" validate:"omitempty,max=100,alpha"`
	Country    string `json:"country" validate:"omitempty,max=100,alpha"`
	City       string `json:"city" validate:"omitempty,max=100,alpha"`
	PostalCode string `json:"postal_code" validate:"omitempty,max=100,alpha"`
}

type DeleteAddressRequest struct {
	ID uint `json:"id"`
}

type FindAddressRequest struct {
	ID uint `json:"id"`
}
