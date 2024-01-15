package user

type UseCase interface {
}

type Router struct {
	useCase UseCase
}

func New(useCase UseCase) Router {
	return Router{
		useCase: useCase,
	}
}
