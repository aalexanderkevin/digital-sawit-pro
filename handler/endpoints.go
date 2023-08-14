package handler

import (
	"errors"
	"net/http"

	"digital-sawit-pro/generated"
	"digital-sawit-pro/helper"
	"digital-sawit-pro/model"
	"digital-sawit-pro/repository"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
)

// (POST /register)
func (s *Server) Register(ctx echo.Context) error {
	var req generated.Register
	var res generated.RegisterSuccessful

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	user := &model.User{
		PhoneNumber: &req.PhoneNumber,
		FullName:    &req.FullName,
	}
	if err := user.Validate(); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := model.IsPasswordValid(req.Password); !err {
		return ctx.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid password format",
		})
	}

	passwordSalt := ksuid.New().String()
	user.Password = helper.Pointer(helper.Hash(passwordSalt, req.Password))
	user.PasswordSalt = helper.Pointer(passwordSalt)
	user, err := s.Repository.Add(ctx.Request().Context(), user)
	if err != nil {
		var e model.Error
		if errors.As(err, &e) {
			return ctx.JSON(e.Code, e)
		} else {
			return ctx.JSON(http.StatusInternalServerError, model.Error{
				Message: err.Error(),
				Code:    http.StatusInternalServerError,
			})
		}
	}

	res.UserId = user.Id
	return ctx.JSON(http.StatusOK, res)
}

// (POST /login)
func (s *Server) Login(ctx echo.Context) error {
	var req generated.Login
	var resp generated.LoginSuccessful

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	user := &model.User{
		PhoneNumber: &req.PhoneNumber,
		Password:    &req.Password,
	}
	if err := user.ValidateLogin(); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	user, err := s.Repository.Get(ctx.Request().Context(), repository.UserGetFilter{
		PhoneNumber: &req.PhoneNumber,
	})
	if err != nil {
		return WriteFailResponse(ctx, err)
	}

	token, err := helper.GenerateJwt(*user)
	if err != nil {
		return WriteFailResponse(ctx, err)
	}

	_, err = s.Repository.Update(ctx.Request().Context(), *user.Id, &model.User{
		SuccessfulLogin: helper.Pointer(*user.SuccessfulLogin + 1),
	})
	if err != nil {
		return WriteFailResponse(ctx, err)
	}

	resp.UserId = user.Id
	resp.JwtToken = token
	return ctx.JSON(http.StatusOK, resp)
}

// (GET /profile)
func (s *Server) GetProfile(ctx echo.Context) error {
	var resp generated.GetMyProfileSuccessful

	// Get token
	token := helper.GetAuthToken(ctx)
	if token == "" {
		return WriteFailResponse(ctx, model.Error{Code: http.StatusForbidden, Message: "jwt not found"})
	}

	jwtData, err := helper.DecodeClaims(token)
	if err != nil {
		return WriteFailResponse(ctx, model.Error{Code: http.StatusForbidden, Message: "jwt is not valid"})
	}

	profile, err := s.Repository.Get(ctx.Request().Context(), repository.UserGetFilter{
		PhoneNumber: &jwtData.PhoneNumber,
	})
	if err != nil {
		return WriteFailResponse(ctx, err)
	}

	resp.FullName = profile.FullName
	resp.PhoneNumber = profile.PhoneNumber
	return ctx.JSON(http.StatusOK, resp)
}

// (PUT /profile)
func (s *Server) UpdateProfile(ctx echo.Context) error {
	var req generated.UpdateMyProfile

	// Get token
	token := helper.GetAuthToken(ctx)
	if token == "" {
		return WriteFailResponse(ctx, model.Error{Code: http.StatusForbidden, Message: "jwt not found"})
	}

	jwtData, err := helper.DecodeClaims(token)
	if err != nil {
		return WriteFailResponse(ctx, model.Error{Code: http.StatusForbidden, Message: "jwt is not valid"})
	}

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	user := &model.User{
		PhoneNumber: &req.PhoneNumber,
		FullName:    &req.FullName,
	}
	if err := user.Validate(); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	_, err = s.Repository.Update(ctx.Request().Context(), jwtData.Id, user)
	if err != nil {
		return WriteFailResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, "OK")
}

func WriteFailResponse(ctx echo.Context, err error) error {
	var e model.Error
	if errors.As(err, &e) {
		return ctx.JSON(e.Code, e)
	} else {
		return ctx.JSON(http.StatusInternalServerError, model.Error{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}
}
