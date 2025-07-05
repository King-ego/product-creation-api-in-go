package repository

import (
	"api/model"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (pr *UserRepository) GetUsers() ([]model.User, error) {
	query := "SELECT * FROM users"

	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Print(err)
		return []model.User{}, err
	}

	var UserList []model.User
	var UserObj model.User

	for rows.Next() {
		err = rows.Scan(
			&UserObj.ID,
			&UserObj.Name,
			&UserObj.Email,
		)

		if err != nil {
			fmt.Println(err)
			return []model.User{}, err
		}

		UserList = append(UserList, UserObj)
	}

	rows.Close()

	fmt.Println(UserList)

	return UserList, err
}
