package user

func NewUserDelivery() UserUsecase {
	userRepository := NewUserRepository(
		"localhost:228",
	)

	return NewUserUsecase(userRepository)
}
