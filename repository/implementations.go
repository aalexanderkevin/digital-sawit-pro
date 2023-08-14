package repository

import (
	"digital-sawit-pro/model"

	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

func (u *Repository) Add(ctx context.Context, user *model.User) (*model.User, error) {
	gormModel := User{}.FromModel(*user)

	if err := u.Db.WithContext(ctx).Create(&gormModel).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return nil, model.NewDuplicateError()
		}
		return nil, err
	}

	return gormModel.ToModel(), nil
}

func (u *Repository) Get(ctx context.Context, filter UserGetFilter) (*model.User, error) {
	user := User{
		Id:          filter.Id,
		PhoneNumber: filter.PhoneNumber,
	}

	q := u.Db.WithContext(ctx)
	if filter.Id != nil {
		q = q.Where("id = ?", filter.Id)
	}
	if filter.PhoneNumber != nil {
		q = q.Where("phone_number = ?", filter.PhoneNumber)
	}

	err := q.First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.NewNotFoundError()
		}
		return nil, err
	}

	return user.ToModel(), nil
}

func (u *Repository) Update(ctx context.Context, id string, user *model.User) (*model.User, error) {
	_, err := u.Get(ctx, UserGetFilter{Id: &id})
	if err != nil {
		return nil, err
	}

	gormModel := User{}.FromModel(*user)

	tx := u.Db.WithContext(ctx)
	err = tx.Model(&User{Id: &id}).Updates(&gormModel).Error
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return nil, model.NewDuplicateError()
		}
		return nil, err
	}

	return u.Get(ctx, UserGetFilter{Id: &id})
}
