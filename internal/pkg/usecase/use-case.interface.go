package usecase

type UseCaseInterface[Input any, Output any] interface {
	Execute(input Input) (output Output, err error)
}

type UseCaseNoOutputInterface[Input any] interface {
	Execute(input Input) (err error)
}
