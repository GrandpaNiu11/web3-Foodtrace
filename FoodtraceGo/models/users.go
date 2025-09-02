package models

type Users struct {
	ID       uint
	Username string
	Password string
	Address  string
	Private  string
}

func (f Users) TableName() string {
	return "users"
}
