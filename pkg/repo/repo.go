package repo

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"github.com/yrcs/yuneto/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Repo base repository.
type Repo[E, T any] interface {
	// Exists determines whether po exists (whether to ignore case).
	Exists(ctx context.Context, fieldName string, fieldValue any, ignoreCase ...bool) (bool, error)
	// Unique determines whether po is unique (whether to ignore case).
	Unique(ctx context.Context, id uint64, fieldName string, fieldValue any, ignoreCase ...bool) (bool, error)
	// Find find po by id (with lockMode).
	Find(ctx context.Context, id uint64, lockMode ...LockMode) (E, error)
	// FindByField find po by fieldName and fieldValue (with ignoreCase or lockMode).
	FindByField(ctx context.Context, fieldName string, fieldValue any, args ...any) (E, error)
	// Create create po.
	Create(ctx context.Context, do E) (E, error)
	// Update update po.
	Update(ctx context.Context, m map[string]any) (E, error)
	// Delete delete po.
	Delete(ctx context.Context, id uint64) error
	// Delete delete pos.
	DeleteByIDs(ctx context.Context, ids []uint64) error
}

type BaseRepo[E, T any] struct {
	DO  E
	PO  T
	DB  *gorm.DB
	Log *log.Helper
}

var _ Repo[*any, *any] = (*BaseRepo[*any, *any])(nil)

func (r *BaseRepo[E, T]) Exists(ctx context.Context, fieldName string, fieldValue any, ignoreCase ...bool) (bool, error) {
	db := r.DB.WithContext(ctx)

	var ignored bool
	for _, v := range ignoreCase {
		ignored = v
		break
	}
	var result *gorm.DB
	if ignored {
		result = db.Where("LOWER("+fieldName+") = ?", strings.ToLower(fmt.Sprintf("%v", fieldValue))).First(&r.PO)
	} else {
		result = db.Where(map[string]any{fieldName: fieldValue}).First(&r.PO)
	}
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return true, err
	}
	return result.RowsAffected > 0, nil
}

func (r *BaseRepo[E, T]) Unique(ctx context.Context, id uint64, fieldName string, fieldValue any, ignoreCase ...bool) (bool, error) {
	db := r.DB.WithContext(ctx)

	var ignored bool
	for _, v := range ignoreCase {
		ignored = v
		break
	}
	var result *gorm.DB
	if ignored {
		result = db.Where("id <> ?", id).Where("LOWER("+fieldName+") = ?", strings.ToLower(fmt.Sprintf("%v", fieldValue))).First(&r.PO)
	} else {
		result = db.Where("id <> ?", id).Where(map[string]any{fieldName: fieldValue}).First(&r.PO)
	}
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, err
	}
	return result.RowsAffected <= 0, nil
}

func (r *BaseRepo[E, T]) Find(ctx context.Context, id uint64, lockMode ...LockMode) (E, error) {
	db := r.DB.WithContext(ctx)

	var mode *LockMode
	for _, v := range lockMode {
		mode = &v
		break
	}
	var result *gorm.DB
	if mode != nil {
		result = db.Clauses(clause.Locking{Strength: mode.String()}).First(&r.PO, id)
	} else {
		result = db.First(&r.PO, id)
	}
	if err := result.Error; err != nil {
		var zero E
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return zero, nil
		}
		return zero, err
	}
	r.DO = util.InitStruct(r.DO)
	copier.CopyWithOption(r.DO, r.PO, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	return r.DO, nil
}

func (r *BaseRepo[E, T]) FindByField(ctx context.Context, fieldName string, fieldValue any, args ...any) (E, error) {
	db := r.DB.WithContext(ctx)

	var ignoreCase bool
	var lockMode *LockMode
	for i, v := range args {
		if i > 1 {
			break
		}
		switch v := v.(type) {
		case bool:
			ignoreCase = v
		case LockMode:
			lockMode = &v
		}
	}
	if ignoreCase {
		db = db.Where("LOWER("+fieldName+") = ?", strings.ToLower(fmt.Sprintf("%v", fieldValue)))
	} else {
		db = db.Where(map[string]any{fieldName: fieldValue})
	}
	if lockMode != nil {
		db = db.Clauses(clause.Locking{Strength: lockMode.String()})
	}
	result := db.First(&r.PO)
	if err := result.Error; err != nil {
		var zero E
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return zero, nil
		}
		return zero, err
	}
	r.DO = util.InitStruct(r.DO)
	copier.CopyWithOption(r.DO, r.PO, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	return r.DO, nil
}

func (r *BaseRepo[E, T]) Create(ctx context.Context, do E) (E, error) {
	return do, r.DB.WithContext(ctx).Create(&do).Error
}

func (r *BaseRepo[E, T]) Update(ctx context.Context, m map[string]any) (E, error) {
	db := r.DB.WithContext(ctx)

	result := db.Where("id = ?", m["Id"]).First(&r.PO)
	if err := result.Error; err != nil {
		var zero E
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return zero, nil
		}
		return zero, err
	}

	result = r.DB.WithContext(ctx).Model(r.PO).Omit("Id").Updates(m)
	r.DO = util.InitStruct(r.DO)
	copier.CopyWithOption(r.DO, r.PO, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	return r.DO, result.Error
}

func (r *BaseRepo[E, T]) Delete(ctx context.Context, id uint64) error {
	r.PO = util.InitStruct(r.PO)
	return r.DB.WithContext(ctx).Delete(r.PO, id).Error
}

func (r *BaseRepo[E, T]) DeleteByIDs(ctx context.Context, ids []uint64) error {
	r.PO = util.InitStruct(r.PO)
	return r.DB.WithContext(ctx).Delete(r.PO, ids).Error
}
