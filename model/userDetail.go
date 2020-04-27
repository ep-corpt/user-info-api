package model

type UserDetailWrapper struct {
	UserDetail UserDetail `json:"userDetail"`
	CredentialDetail CredentialDetail `json:"credentialDetail"`
}

type UserDetail struct{
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
}

type CredentialDetail struct {
	Username string `json:"username"`
	Password string `json:"password"`
}