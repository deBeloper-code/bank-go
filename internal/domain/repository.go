package domain

// User
type UserRepository interface {
	Save(user *User) error
	FindByID(id int) (*User, error)
}

// Account
type AccountRepository interface {
	Save(account *Account) error
	FindByID(id int) (*Account, error)
}
