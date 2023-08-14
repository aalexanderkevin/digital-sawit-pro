package handler_test

import (
	"bytes"
	"digital-sawit-pro/generated"
	"digital-sawit-pro/handler"
	"digital-sawit-pro/helper"
	"digital-sawit-pro/helper/test"
	"digital-sawit-pro/model"
	"digital-sawit-pro/repository"
	"encoding/json"
	"errors"

	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestRegister(t *testing.T) {
	t.Parallel()
	t.Run("ShouldReturnBadRequest_WhenMissingPhoneNumber", func(t *testing.T) {
		t.Parallel()
		// INIT
		header := map[string]string{
			"Content-Type": "application/json",
		}
		reqBody := generated.Register{
			FullName: "fullname",
			Password: "Passw0rd!",
		}

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(reqBody)
		require.NoError(t, err)

		router := test.SetupHttpHandler(t, handler.NewServerOptions{
			Repository: nil,
		})

		// CODE UNDER TEST
		w, err := performRequest(router, "POST", "/register", &buf, header, nil)
		require.NoError(t, err)
		defer printOnFailed(t)(w.Body.String())

		// EXPECTATION
		require.Equal(t, http.StatusBadRequest, w.Code)

		resBody := model.Error{}
		err = json.NewDecoder(w.Body).Decode(&resBody)
		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, resBody.Code)
		require.Equal(t, "phone_number: cannot be blank.", resBody.Message)
	})

	t.Run("ShouldReturnBadRequest_WhenMissingFullName", func(t *testing.T) {
		t.Parallel()
		// INIT
		header := map[string]string{
			"Content-Type": "application/json",
		}
		reqBody := generated.Register{
			PhoneNumber: "+62822987602",
			Password:    "Passw0rd!",
		}

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(reqBody)
		require.NoError(t, err)

		router := test.SetupHttpHandler(t, handler.NewServerOptions{
			Repository: nil,
		})

		// CODE UNDER TEST
		w, err := performRequest(router, "POST", "/register", &buf, header, nil)
		require.NoError(t, err)
		defer printOnFailed(t)(w.Body.String())

		// EXPECTATION
		require.Equal(t, http.StatusBadRequest, w.Code)

		resBody := model.Error{}
		err = json.NewDecoder(w.Body).Decode(&resBody)
		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, resBody.Code)
		require.Equal(t, "full_name: cannot be blank.", resBody.Message)
	})

	t.Run("ShouldReturnBadRequest_WhenInvalidPasswordFormat", func(t *testing.T) {
		t.Parallel()
		// INIT
		header := map[string]string{
			"Content-Type": "application/json",
		}
		reqBody := generated.Register{
			FullName:    "full name",
			PhoneNumber: "+62822987602",
			Password:    "Passw0rd",
		}

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(reqBody)
		require.NoError(t, err)

		router := test.SetupHttpHandler(t, handler.NewServerOptions{
			Repository: nil,
		})

		// CODE UNDER TEST
		w, err := performRequest(router, "POST", "/register", &buf, header, nil)
		require.NoError(t, err)
		defer printOnFailed(t)(w.Body.String())

		// EXPECTATION
		require.Equal(t, http.StatusBadRequest, w.Code)

		resBody := model.Error{}
		err = json.NewDecoder(w.Body).Decode(&resBody)
		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, resBody.Code)
		require.Equal(t, "invalid password format", resBody.Message)
	})

	t.Run("ShouldReturnInternalError_WhenFailedInsertUser", func(t *testing.T) {
		t.Parallel()
		// INIT
		header := map[string]string{
			"Content-Type": "application/json",
		}
		reqBody := generated.Register{
			FullName:    "full name",
			PhoneNumber: "+62822987602",
			Password:    "Passw0rd!",
		}

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(reqBody)
		require.NoError(t, err)

		mockCtrl := gomock.NewController(t)
		mockRepo := repository.NewMockRepositoryInterface(mockCtrl)
		mockRepo.EXPECT().Add(gomock.Any(), mock.MatchedBy(func(u *model.User) bool {
			require.Equal(t, reqBody.FullName, *u.FullName)
			require.Equal(t, reqBody.PhoneNumber, *u.PhoneNumber)
			return true
		})).Return(nil, errors.New("error insert"))

		router := test.SetupHttpHandler(t, handler.NewServerOptions{
			Repository: mockRepo,
		})

		// CODE UNDER TEST
		w, err := performRequest(router, "POST", "/register", &buf, header, nil)
		require.NoError(t, err)
		defer printOnFailed(t)(w.Body.String())

		// EXPECTATION
		require.Equal(t, http.StatusInternalServerError, w.Code)

		resBody := model.Error{}
		err = json.NewDecoder(w.Body).Decode(&resBody)
		require.NoError(t, err)
		require.Equal(t, "error insert", resBody.Message)
	})

	t.Run("ShouldReturnUserId_WhenSuccessRegister", func(t *testing.T) {
		t.Parallel()
		// INIT
		header := map[string]string{
			"Content-Type": "application/json",
		}
		reqBody := generated.Register{
			FullName:    "full name",
			PhoneNumber: "+62822987602",
			Password:    "Passw0rd!",
		}

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(reqBody)
		require.NoError(t, err)

		mockCtrl := gomock.NewController(t)
		mockRepo := repository.NewMockRepositoryInterface(mockCtrl)
		mockRepo.EXPECT().Add(gomock.Any(), mock.MatchedBy(func(u *model.User) bool {
			require.Equal(t, reqBody.FullName, *u.FullName)
			require.Equal(t, reqBody.PhoneNumber, *u.PhoneNumber)
			return true
		})).Return(&model.User{
			Id: helper.Pointer("id"),
		}, nil)

		router := test.SetupHttpHandler(t, handler.NewServerOptions{
			Repository: mockRepo,
		})

		// CODE UNDER TEST
		w, err := performRequest(router, "POST", "/register", &buf, header, nil)
		require.NoError(t, err)
		defer printOnFailed(t)(w.Body.String())

		// EXPECTATION
		require.Equal(t, http.StatusOK, w.Code)

		resBody := generated.RegisterSuccessful{}
		err = json.NewDecoder(w.Body).Decode(&resBody)
		require.NoError(t, err)
		require.Equal(t, "id", *resBody.UserId)
	})

}
