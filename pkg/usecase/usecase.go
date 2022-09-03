package usecase

import (
	"context"

	"github.com/yrcs/yuneto/pkg/repo"
)

type BaseUsecase[E, T any] struct {
	Repo *repo.BaseRepo[E, T]
}

func (u *BaseUsecase[E, T]) Add(ctx context.Context, do E) (E, error) {
	return u.Repo.Create(ctx, do)
}

func (u *BaseUsecase[E, T]) Edit(ctx context.Context, m map[string]any) (E, error) {
	return u.Repo.Update(ctx, m)
}

func (u *BaseUsecase[E, T]) Find(ctx context.Context, id uint64, lockMode ...repo.LockMode) (E, error) {
	return u.Repo.Find(ctx, id, lockMode...)
}

func (u *BaseUsecase[E, T]) FindByField(ctx context.Context, fieldName string, fieldValue any, args ...any) (E, error) {
	return u.Repo.FindByField(ctx, fieldName, fieldValue, args...)
}

func (u *BaseUsecase[E, T]) Delete(ctx context.Context, id uint64) error {
	return u.Repo.Delete(ctx, id)
}

func (u *BaseUsecase[E, T]) DeleteByIDs(ctx context.Context, ids []uint64) error {
	return u.Repo.DeleteByIDs(ctx, ids)
}
