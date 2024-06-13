package usecase

import "context"

type UseCaseInterface[Input any, Output any] interface {
	Execute(ctx context.Context, input Input) (output Output, err error)
}

type UseCaseNoOutputInterface[Input any] interface {
	Execute(ctx context.Context, input Input) (err error)
}
