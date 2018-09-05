package user

// Service is a interface for user service
type Service interface {
	AddUser()
	UpdateUser()
	DeleteUser()
	GetUser()
	GetUserByID()
	RegisterUser()
	RegisterCheck()
}

// BaseUserService is a implements for user service
type BaseUserService struct{}
