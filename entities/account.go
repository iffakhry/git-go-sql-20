package entities

import "time"

type Account struct {
	ID        uint
	Username  string
	Name      string
	Bio       string
	JoinDate  time.Time
	Location  string
	Email     string
	Password  string
	CreatedAt time.Time
}
