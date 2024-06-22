package supabase_auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
	"github.com/xkurozaru/pedometer-server/dependency/config"
	"github.com/xkurozaru/pedometer-server/domain/auth"
	model_errors "github.com/xkurozaru/pedometer-server/domain/errors"
	"github.com/xkurozaru/pedometer-server/domain/user"
)

type authAPI struct {
	supabaseClient *supabase.Client
	supabaseConfig config.SupabaseConfig
}

func NewAuthAPI(
	supabaseClient *supabase.Client,
	supabaseConfig config.SupabaseConfig,
) auth.AuthRepository {
	return authAPI{
		supabaseClient: supabaseClient,
		supabaseConfig: supabaseConfig,
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
		return []byte(a.supabaseConfig.JWTSecret), nil
	})
	if err != nil {
		return "", model_errors.NewInfrastructureError(err.Error())
	}
	if !token.Valid {
		return "", model_errors.NewInfrastructureError("")
	}

	authID := claims["sub"].(string)
	return authID, nil
}

func (a authAPI) Delete(u user.User) error {
	req := types.AdminDeleteUserRequest{
		UserID: uuid.MustParse(u.AuthID()),
	}

	err := a.supabaseClient.Auth.WithToken(a.supabaseConfig.APIKey).AdminDeleteUser(req)
	if err != nil {
		return model_errors.NewInfrastructureError(err.Error())
	}
	return nil
}
