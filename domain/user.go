package domain

import "time"

// User representa a entidade de um usuário no sistema
type User struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // Não serializar esse campo
	XP           int       `json:"xp"`
	TotalScore   int       `json:"total_score"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "phishing_quest.users"
}

// NewUser é um construtor para criar um novo usuário
func NewUser(username, email, passwordHash string) *User {
	return &User{
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		XP:           0,
		TotalScore:   0,
		//CreatedAt:    time.Now(),
		//UpdatedAt:    time.Now(),
	}
}
