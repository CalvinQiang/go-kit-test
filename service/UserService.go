package service

type IUserService interface {
	GetName(userId int) string
}

type UserService struct {
}

func (this UserService) GetName(userId int) string {
	if userId == 101 {
		return "calvin"
	}
	return "guest"
}
