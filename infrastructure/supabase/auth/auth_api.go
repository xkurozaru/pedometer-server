package supabase_auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
	"github.com/xkurozaru/pedometer-server/domain/auth"
	model_errors "github.com/xkurozaru/pedometer-server/domain/errors"
)

type authAPI struct {
	supabaseClient *supabase.Client
	jWTSecret      string
}

func NewAuthAPI(
	supabaseClient *supabase.Client,
	jWTSecret string,
) auth.AuthRepository {
	return authAPI{
		supabaseClient: supabaseClient,
		jWTSecret:      jWTSecret,
	}
}

func (a authAPI) Register(email string, password string) (string, error) {
	req := types.SignupRequest{
		Email:    email,
		Password: password,
	}
	res, err := a.supabaseClient.Auth.Signup(req)
	if err != nil {
		return "", model_errors.NewInfrastructureError(err.Error())
	}
	return res.User.ID.String(), nil
}

func (a authAPI) Login(email string, password string) (string, error) {
	res, err := a.supabaseClient.Auth.SignInWithEmailPassword(email, password)
	if err != nil {
		return "", model_errors.NewInfrastructureError(err.Error())
	}
	return res.AccessToken, nil
}

func (a authAPI) Verify(jWT string) (string, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(jWT, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.jWTSecret), nil
	})
	if err != nil {
		return "", model_errors.NewInfrastructureError(err.Error())
	}
	if !token.Valid {
		return "", model_errors.NewInfrastructureError(err.Error())
	}

	userID := claims["sub"].(string)
	return userID, nil
}
