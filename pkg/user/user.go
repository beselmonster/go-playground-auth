package user

import (
	database "auth/pkg/db"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	Id int `json:"id"`

	Name string `json:"name"`

	Token string `json:"-"`

	CreatedAt string `json:"created_at"`
}

type Repository struct {
	users []User
}

/**
Returns all user from db
*/
func (userRepository *Repository) All() []User {
	rows, err := database.Con.Query("select id, name, created_at from users")

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		var user User

		err := rows.Scan(&user.Id, &user.Name, &user.CreatedAt)

		if err != nil {
			log.Fatal(err)
		}

		userRepository.users = append(userRepository.users, user)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return userRepository.users
}

/**
Repository Constructor
*/
func NewRepository(users []User) *Repository {
	return &Repository{
		users: users,
	}
}
