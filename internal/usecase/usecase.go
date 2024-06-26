package usecase

import "context"

type UseCase[I, O any] interface {
	Invoke(ctx context.Context, input I) (O, error)
}
