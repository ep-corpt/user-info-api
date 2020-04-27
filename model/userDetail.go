package model

type UserDetailWrapper struct {
	UserDetail UserDetail `json:"userDetail"`
}

type UserDetail struct{
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
}
