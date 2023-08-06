package handler

import (
	"digital-sawit-pro/helper"
	"digital-sawit-pro/model"
	"digital-sawit-pro/repository"

	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
)

func TestLogin(t *testing.T) {
	type fields struct {
		Repository func() repository.RepositoryInterface
	}
	type args struct {
		ctx func() echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ShouldSuccess",
			fields: fields{
				Repository: func() repository.RepositoryInterface {
					mockCtrl := gomock.NewController(t)
					mockRepo := repository.NewMockRepositoryInterface(mockCtrl)
					mockRepo.EXPECT().Get(gomock.Any(), repository.UserGetFilter{
						PhoneNumber: helper.Pointer("+6281234567890"),
					}).Return(&model.User{
						Id:              helper.Pointer("id"),
						PhoneNumber:     helper.Pointer("+6281234567890"),
						FullName:        helper.Pointer("test aja"),
						Password:        helper.Pointer("password"),
						PasswordSalt:    helper.Pointer("password_salt"),
						SuccessfulLogin: helper.Pointer(0),
					}, nil)

					mockRepo.EXPECT().Update(gomock.Any(), "231", &model.User{
						SuccessfulLogin: helper.Pointer(1),
					}).Return(nil, nil)

					return mockRepo
				},
			},
			args: args{
				ctx: func() echo.Context {
					// Create a mock request with GET method and an empty request body
					req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte(`
					{"password": "Password!23","phone_number": "+6281234567890"}
					`)))

					// Create a response recorder to capture the response
					rec := httptest.NewRecorder()

					c := echo.New().NewContext(req, rec)
					c.Request().Header.Set("Content-Type", "application/json")
					return c
				},
			},
			wantErr: false,
		},
		// {
		// 	name: "failed during get phone number will result error 500",
		// 	fields: fields{
		// 		Repository: func() repository.RepositoryInterface {
		// 			mockCtrl := gomock.NewController(t)
		// 			mockRepo := repository.NewMockRepositoryInterface(mockCtrl)
		// 			mockRepo.EXPECT().GetUserByPhoneNumber(gomock.Any(), "+628128813798").Return(repository.GetUserOutput{}, errors.New("some error"))

		// 			return mockRepo
		// 		},
		// 	},
		// 	args: args{
		// 		ctx: func() echo.Context {
		// 			// Create a mock request with GET method and an empty request body
		// 			req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte(`{"password": "Testaja1!","phone_number": "+628128813798"}`)))

		// 			// Create a response recorder to capture the response
		// 			rec := httptest.NewRecorder()

		// 			c := echo.New().NewContext(req, rec)
		// 			c.Request().Header.Set("Content-Type", "application/json")
		// 			return c
		// 		},
		// 	},
		// 	wantErr: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Repository: tt.fields.Repository(),
			}
			err := s.Login(tt.args.ctx())
			if (err != nil) != tt.wantErr {
				t.Errorf("PostLogin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
