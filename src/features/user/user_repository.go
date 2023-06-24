package user

type userRepository struct {
	databaseURL string
}

type UserRepository interface {
	CreateUser(name string) (*User, error)
}

func NewUserRepository(serverURL string) UserRepository {
	return &userRepository{
		databaseURL: serverURL,
	}
}

func (ur userRepository) CreateUser(name string) (*User, error) {
	var user = &User{
		Name: name,
	}

	return user, nil
}
