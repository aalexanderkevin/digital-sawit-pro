//go:build integration
// +build integration

package repository_test

import (
	"digital-sawit-pro/helper"
	"digital-sawit-pro/model"
	"digital-sawit-pro/repository"
	"digital-sawit-pro/storage"

	"context"
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_Add(t *testing.T) {
	t.Run("ShouldInsertUser", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		fakeUser := &model.User{
			PhoneNumber:  helper.Pointer(fake.CharactersN(10)),
			FullName:     helper.Pointer(fake.CharactersN(10)),
			Password:     helper.Pointer(fake.CharactersN(10)),
			PasswordSalt: helper.Pointer(fake.CharactersN(10)),
		}

		//-- code under test
		userRepo := repository.NewUserRepository(db)
		addedUser, err := userRepo.Add(context.TODO(), fakeUser)

		//-- assert
		require.NoError(t, err)
		require.NotNil(t, addedUser)
		require.Equal(t, *fakeUser.PhoneNumber, *addedUser.PhoneNumber)
		require.Equal(t, *fakeUser.FullName, *addedUser.FullName)
		require.Equal(t, *fakeUser.Password, *addedUser.Password)
		require.Equal(t, 0, *addedUser.SuccessfulLogin)
	})

	t.Run("ShouldReturnError_WhenInsertTheSamePhoneNumber", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		userRepo := repository.NewUserRepository(db)

		user1, err := userRepo.Add(context.TODO(), &model.User{
			PhoneNumber:  helper.Pointer(fake.CharactersN(10)),
			FullName:     helper.Pointer(fake.CharactersN(10)),
			Password:     helper.Pointer(fake.CharactersN(10)),
			PasswordSalt: helper.Pointer(fake.CharactersN(10)),
		})

		//-- code under test
		addedUser, err := userRepo.Add(context.TODO(), &model.User{
			PhoneNumber:  user1.PhoneNumber,
			FullName:     helper.Pointer(fake.CharactersN(10)),
			Password:     helper.Pointer(fake.CharactersN(10)),
			PasswordSalt: helper.Pointer(fake.CharactersN(10)),
		})

		//-- assert
		require.Error(t, err)
		require.EqualError(t, err, model.NewDuplicateError().Message)
		require.Nil(t, addedUser)
	})
}

func TestUserRepository_Get(t *testing.T) {
	t.Run("ShouldReturnNotFoundError_WhenThePhoneNumberIsNotExist", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		//-- code under test
		userRepo := repository.NewUserRepository(db)
		user, err := userRepo.Get(context.TODO(), repository.UserGetFilter{
			PhoneNumber: helper.Pointer("phone_number"),
		})
		require.Error(t, err)

		//-- assert
		require.EqualError(t, err, model.NewNotFoundError().Error())
		require.Nil(t, user)
	})

	t.Run("ShouldReturnNotFoundError_WhenTheIdIsNotExist", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		//-- code under test
		userRepo := repository.NewUserRepository(db)
		user, err := userRepo.Get(context.TODO(), repository.UserGetFilter{
			Id: helper.Pointer("id"),
		})
		require.Error(t, err)

		//-- assert
		require.EqualError(t, err, model.NewNotFoundError().Error())
		require.Nil(t, user)
	})

	t.Run("ShouldGetByPhoneNumber_WhenThePhoneNumberExist", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		userRepo := repository.NewUserRepository(db)
		fakeUser, err := userRepo.Add(context.TODO(), &model.User{
			PhoneNumber:  helper.Pointer(fake.CharactersN(10)),
			FullName:     helper.Pointer(fake.CharactersN(10)),
			Password:     helper.Pointer(fake.CharactersN(10)),
			PasswordSalt: helper.Pointer(fake.CharactersN(10)),
		})

		//-- code under test
		user, err := userRepo.Get(context.TODO(), repository.UserGetFilter{PhoneNumber: fakeUser.PhoneNumber})
		require.NoError(t, err)

		//-- assert
		require.NotNil(t, user)
		require.Equal(t, *fakeUser.Id, *user.Id)
		require.Equal(t, *fakeUser.PhoneNumber, *user.PhoneNumber)
		require.Equal(t, *fakeUser.FullName, *user.FullName)
		require.Equal(t, *fakeUser.Password, *user.Password)
		require.Equal(t, *fakeUser.PasswordSalt, *user.PasswordSalt)
		require.Equal(t, 0, *user.SuccessfulLogin)
	})
}

func TestUserRepository_Update(t *testing.T) {
	t.Run("ShouldReturnNotFoundError_WhenTheIdIsNotExist", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		//-- code under test
		userRepo := repository.NewUserRepository(db)
		user, err := userRepo.Update(context.TODO(), "id", &model.User{
			PhoneNumber: helper.Pointer("update_phone_number"),
		})
		require.Error(t, err)

		//-- assert
		require.EqualError(t, err, model.NewNotFoundError().Error())
		require.Nil(t, user)
	})

	t.Run("ShouldReturnUpdatedUser", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		existingUser := &model.User{
			PhoneNumber:  helper.Pointer(fake.CharactersN(10)),
			FullName:     helper.Pointer(fake.CharactersN(10)),
			Password:     helper.Pointer(fake.CharactersN(10)),
			PasswordSalt: helper.Pointer(fake.CharactersN(10)),
		}

		//-- code under test
		userRepo := repository.NewUserRepository(db)
		addedUser, err := userRepo.Add(context.TODO(), existingUser)
		require.NoError(t, err)

		user, err := userRepo.Update(context.TODO(), *addedUser.Id, &model.User{
			PhoneNumber:     helper.Pointer("1234567890"),
			SuccessfulLogin: helper.Pointer(*addedUser.SuccessfulLogin + 1),
		})
		require.NoError(t, err)

		//-- assert
		require.NoError(t, err)
		require.NotNil(t, user)
		require.Equal(t, *addedUser.Id, *user.Id)
		require.Equal(t, "1234567890", *user.PhoneNumber)
		require.Equal(t, *addedUser.FullName, *user.FullName)
		require.Equal(t, *addedUser.Password, *user.Password)
		require.Equal(t, *addedUser.PasswordSalt, *user.PasswordSalt)
		require.Equal(t, 1, *user.SuccessfulLogin)
	})

}
