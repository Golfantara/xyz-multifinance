package usecase

import "test/features/auth"

type service struct {
}

func New() auth.Usecase {
	return &service{
		
	}
}