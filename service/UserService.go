package service

type IUserService interface {
	GetName(userId int) string
	DelUser(userId int) string
}

type UserService struct {
}

func (s UserService) GetName(userId int) string {
	if userId == 101 {
		return "calvin"
	}
	return "guest"
}

func (s UserService) DelUser(userId int) string {
	if userId == 101 { // 模拟不可以删除超级管理员
		return "无权限"
	}
	return "删除成功"
}
