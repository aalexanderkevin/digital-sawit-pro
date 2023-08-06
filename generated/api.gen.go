// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.3 DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// FullName defines model for FullName.
type FullName = string

// Login defines model for Login.
type Login struct {
	Password    Password    `json:"password"`
	PhoneNumber PhoneNumber `json:"phone_number"`
}

// Password defines model for Password.
type Password = string

// PhoneNumber defines model for PhoneNumber.
type PhoneNumber = string

// Register defines model for Register.
type Register struct {
	FullName    FullName    `json:"full_name"`
	Password    Password    `json:"password"`
	PhoneNumber PhoneNumber `json:"phone_number"`
}

// UpdateMyProfile defines model for UpdateMyProfile.
type UpdateMyProfile struct {
	FullName    FullName    `json:"full_name"`
	PhoneNumber PhoneNumber `json:"phone_number"`
}

// BadRequest defines model for BadRequest.
type BadRequest struct {
	// ErrorMessage The error message
	ErrorMessage *string `json:"error_message,omitempty"`
}

// Conflict defines model for Conflict.
type Conflict struct {
	// ErrorMessage The error message
	ErrorMessage *string `json:"error_message,omitempty"`
}

// Forbidden defines model for Forbidden.
type Forbidden struct {
	// ErrorMessage The error message
	ErrorMessage *string `json:"error_message,omitempty"`
}

// GetMyProfileSuccessful defines model for GetMyProfileSuccessful.
type GetMyProfileSuccessful struct {
	// FullName The name of the user
	FullName *string `json:"full_name,omitempty"`

	// PhoneNumber The phone number
	PhoneNumber *string `json:"phone_number,omitempty"`
}

// InternalServerError defines model for InternalServerError.
type InternalServerError struct {
	// ErrorMessage The error message
	ErrorMessage *string `json:"error_message,omitempty"`
}

// LoginSuccessful defines model for LoginSuccessful.
type LoginSuccessful struct {
	// JwtToken JWT with algorithm RS256
	JwtToken *string `json:"jwt_token,omitempty"`

	// UserId The ID of the user
	UserId *string `json:"user_id,omitempty"`
}

// RegisterSuccessful defines model for RegisterSuccessful.
type RegisterSuccessful struct {
	// UserId The ID of user
	UserId *string `json:"user_id,omitempty"`
}

// LoginJSONRequestBody defines body for Login for application/json ContentType.
type LoginJSONRequestBody = Login

// UpdateProfileJSONRequestBody defines body for UpdateProfile for application/json ContentType.
type UpdateProfileJSONRequestBody = UpdateMyProfile

// RegisterJSONRequestBody defines body for Register for application/json ContentType.
type RegisterJSONRequestBody = Register

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Login existing user
	// (POST /login)
	Login(ctx echo.Context) error
	// Get my profile
	// (GET /profile)
	GetProfile(ctx echo.Context) error
	// Update my profile
	// (PUT /profile)
	UpdateProfile(ctx echo.Context) error
	// Register new user
	// (POST /register)
	Register(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Login converts echo context to params.
func (w *ServerInterfaceWrapper) Login(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Login(ctx)
	return err
}

// GetProfile converts echo context to params.
func (w *ServerInterfaceWrapper) GetProfile(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetProfile(ctx)
	return err
}

// UpdateProfile converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateProfile(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateProfile(ctx)
	return err
}

// Register converts echo context to params.
func (w *ServerInterfaceWrapper) Register(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Register(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/login", wrapper.Login)
	router.GET(baseURL+"/profile", wrapper.GetProfile)
	router.PUT(baseURL+"/profile", wrapper.UpdateProfile)
	router.POST(baseURL+"/register", wrapper.Register)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xXYW/bNhD9K9o1A5qGieQkNRIBw9Zsa5Ct7YIkxYDFrsFIZ4mZRGok1dQL+N8HUpYs",
	"NUrspenmbxR9PL57d3x3voVI5IXgyLWC8BYkqkJwhe7jiMZn+FeJStuvSHCN3C1pUWQsopoJ7l8rwe2e",
	"ilLMqV0VUhQoNaucoJRCTnJUiiZoN2JUkWSFPQwhXKToOROvNiGgZwVCCEpLxhMwptkRV9cYaTB2q+vm",
	"iMaenGM1BH4UfJqxaP1xN0ANgddCXrE4Rr72qBdIDYFj1G9np1JMWYbnZRShUtMy+4IYpmWWTTjN78Fv",
	"f/HE1NMpeqVCeTcEAkUqOE54mV+h7PfiLLy5xaNIOENdSu5ZtO8sJMrjyuu7yqkhcMI1Sk6zc5QfUf5s",
	"iV/73NaYvQq0V6E2BN6IhPEnSfD1jZ5o8WdV6N3bf/n9wrthOvVolgjJdJp7Z+e7L4d9SbbJn7C4n4mT",
	"nx6ukVWocBF7ahGyIXCGCVMa5ZMQsUIEj0dfI+0EYMgcjLv/9bx27Tqnn94gT3QK4TAgkDNef+71cO+o",
	"uRtQQZW6EdJFtCFxCiE88xcdxp/f7Z/Wdj1v9cFz7fdlCFjJZxJjCC+7fsgCyvgOVwROWzjbge93Ah9a",
	"L9o+Bwjhw/Pvv9t5cflq+4/xZrUMtg/r5Tc/PNv49sOoDILd4Yvnm5Ot8aa1pNt/B9uHd3/c2uir53Zw",
	"XVyDvQ6uQdABNhptDXdHo/j2gAwGptd1XQtLxPYh6ptaMWSd0rzAvyTl74uYamw61VMx8bWjGve+dMan",
	"wl6XsQi5cogr5PD25MLi0kzbGOG9QumUnEWWoY8oVaUOg51gJ7CWokBOCwYh7LktV1qpY8TPmlcuqhnQ",
	"8uWk7SSu9RGqQFDpIxHP/pUYPkRW5dt0edKyRLfRGlJ3g+A+X42d/3nvMgT2VznXmoANgZerHOnr+U52",
	"yzyncta0FfzElGY8qRTeWvjFojQT7OH7GHVdvY8h4Z5BzXGxt/x4Z+j7MiowKiXTMwgvb+EIqUT5qrTC",
	"djk24zZTx6i9fObVvNj3VvbwUj3tNjVPX4+fy8fqldntyr/9+jjC94PD5Sfa/yb+mwxVtHSSZCtZtvtN",
	"r3Y0HenrpKtxv1KeBsuZ6pn7/lcRaaY7jje1hLjE2QPK5a2UGYSQal2Evp+JiGapUDo8CA4CMGPzTwAA",
	"AP//ljMH7PsPAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
