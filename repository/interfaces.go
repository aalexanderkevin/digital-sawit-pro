// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"
	"digital-sawit-pro/model"
)

type RepositoryInterface interface {
	Add(ctx context.Context, user *model.User) (*model.User, error)
	Get(ctx context.Context, filter UserGetFilter) (*model.User, error)
	Update(ctx context.Context, id string, user *model.User) (*model.User, error)
}
