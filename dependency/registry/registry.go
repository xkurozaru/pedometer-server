package registry

import (
	"github.com/supabase-community/supabase-go"
	"github.com/xkurozaru/pedometer-server/application"
	"github.com/xkurozaru/pedometer-server/dependency/config"
	"github.com/xkurozaru/pedometer-server/domain/auth"
	"github.com/xkurozaru/pedometer-server/domain/user"
	"github.com/xkurozaru/pedometer-server/infrastructure/database"
	user_database "github.com/xkurozaru/pedometer-server/infrastructure/database/user"
	supabase_client "github.com/xkurozaru/pedometer-server/infrastructure/supabase"
	supabase_auth "github.com/xkurozaru/pedometer-server/infrastructure/supabase/auth"
	"github.com/xkurozaru/pedometer-server/interfaces/auth_interface"
	"github.com/xkurozaru/pedometer-server/interfaces/user_interface"
	"gorm.io/gorm"
)

type Registry interface {
	NewAuthHandler() auth_interface.AuthHandler
	NewUserHandler() user_interface.UserHandler
}

type registry struct{}

func NewRegistry() Registry {
	return registry{}
}

func (r registry) NewAuthHandler() auth_interface.AuthHandler {
	return auth_interface.NewAuthHandler(
		r.NewAuthApplicationService(),
	)
}

func (r registry) NewUserHandler() user_interface.UserHandler {
	return user_interface.NewUserHandler(
		r.NewUserApplicationService(),
	)
}

func (r registry) NewAuthApplicationService() application.AuthApplicationService {
	return application.NewAuthApplicationService(
		r.NewAuthRepository(),
	)
}

func (r registry) NewUserApplicationService() application.UserApplicationService {
	return application.NewUserApplicationService(
		r.NewAuthRepository(),
		r.NewUserRepository(),
	)
}

func (r registry) NewAuthRepository() auth.AuthRepository {
	supabaseConfig, err := config.NewSupabaseConfig()
	if err != nil {
		panic(err)
	}
	return supabase_auth.NewAuthAPI(r.NewAuthClient(), supabaseConfig.JWTSecret)
}

func (r registry) NewUserRepository() user.UserRepository {
	return user_database.NewUserDatabase(r.NewDB())
}

func (r registry) NewDB() *gorm.DB {
	dbConfig, err := config.NewDBConfig()
	if err != nil {
		panic(err)
	}
	db, err := database.ConnectDB(dbConfig)
	if err != nil {
		panic(err)
	}
	return db
}

func (r registry) NewAuthClient() *supabase.Client {
	supabaseConfig, err := config.NewSupabaseConfig()
	if err != nil {
		panic(err)
	}
	client, err := supabase_client.ConnectSupabase(supabaseConfig)
	if err != nil {
		panic(err)
	}
	return client
}
