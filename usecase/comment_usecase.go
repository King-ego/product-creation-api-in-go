package usecase

import (
	"api/model"
	"api/repository"
)

type CommentUseCase struct {
	repository repository.CommentRepository
}

func NewCommentUseCase(repository repository.CommentRepository) CommentUseCase {
	return CommentUseCase{
		repository: repository,
	}
}

func (pc *CommentUseCase) GetComments() ([]model.Comment, error) {
	return pc.repository.GetComments()
}
