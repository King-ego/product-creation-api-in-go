package repository

import (
	"api/model"
	"database/sql"
	"fmt"
)

type CommentRepository struct {
	connection *sql.DB
}

func NewCommentRepository(connection *sql.DB) CommentRepository {
	return CommentRepository{
		connection: connection,
	}
}

func (pr *CommentRepository) GetComments() ([]model.Comment, error) {
	query := "SELECT * FROM comments"

	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Comment{}, err
	}

	var CommentObj model.Comment
	var CommentList []model.Comment

	for rows.Next() {
		err = rows.Scan(
			&CommentObj.ID,
			&CommentObj.Message,
			&CommentObj.ProductId,
			&CommentObj.UserId,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Comment{}, err
		}

		CommentList = append(CommentList, CommentObj)
	}

	rows.Close()

	return CommentList, err
}
