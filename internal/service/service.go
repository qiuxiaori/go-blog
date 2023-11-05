package service

import "context"

type Service struct {
	ctx context.Context
	// dao *
	// model *
}

func New(ctx context.Context) Service {
	svc := Service{ctx}
	return svc
}
