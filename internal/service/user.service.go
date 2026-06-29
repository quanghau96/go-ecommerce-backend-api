package service

import rp "github.com/quanghau96/go-ecommerce-backend-api/internal/repo"

type UserService struct {
	userRepo *rp.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: rp.NewUserRepo(),
	}
}

func (us *UserService) GetInfoUser() string {
	return us.userRepo.GetInfoUser()
}
