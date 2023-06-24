package user

type userUsecase struct {
	userRepository UserRepository
}

type UserUsecase interface {
	CreateUser(name string) (*User, error)
}

func NewUserUsecase(userRepository UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (uc userUsecase) CreateUser(name string) (*User, error) {
	user, err := uc.userRepository.CreateUser(name)

	if err != nil {
		return nil, err
	}

	return user, nil
}
