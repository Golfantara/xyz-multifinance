package repository

import "test/features/auth"

type model struct {
}

func New() auth.Repository {
	return &model{
	}
}