// internal/domain/service.go
package domain

// UserService
type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(name, email, password string) (*User, error) {
	createdUser := &User{Username: name, Email: email, Password: password}

	if err := s.userRepository.Save(createdUser); err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (s *UserService) GetUserByID(id int) (*User, error) {
	return s.userRepository.FindByID(id)
}

// AccountService
type AccountService struct {
	accountRepository AccountRepository
}

func NewAccountService(accountRepository AccountRepository) *AccountService {
	return &AccountService{accountRepository: accountRepository}
}

func (s *AccountService) CreateAccount(name string) (*Account, error) {
	user := &Account{Name: name}
	if err := s.accountRepository.Save(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AccountService) GetAccountByID(id int) (*Account, error) {
	return s.accountRepository.FindByID(id)
}
