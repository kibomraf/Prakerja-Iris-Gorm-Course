package students

import "time"

// Entity Student for database
type Student struct {
	Id            int `gorm:"primaryKey"`
	Email         string
	Password_hash string
	Name          string
	Address       string
	NoHandphone   string
	Created_at    time.Time
	Updated_at    time.Time
}
