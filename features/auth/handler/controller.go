package handler

import "test/features/auth"

type controller struct {
}

func New() auth.Handler {
	return &controller{
		
	}
}