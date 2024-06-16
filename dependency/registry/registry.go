package registry

import (
	"github.com/HackU-team04/Pedometer/application"
	"github.com/HackU-team04/Pedometer/dependency/config"
	"github.com/HackU-team04/Pedometer/domain/auth"
	"github.com/HackU-team04/Pedometer/domain/user"
	"github.com/HackU-team04/Pedometer/infrastructure/database"
	user_database "github.com/HackU-team04/Pedometer/infrastructure/database/user"
	supabase_client "github.com/HackU-team04/Pedometer/infrastructure/supabase"
	supabase_auth "github.com/HackU-team04/Pedometer/infrastructure/supabase/auth"
	"github.com/HackU-team04/Pedometer/interfaces/auth_interface"
	"github.com/HackU-team04/Pedometer/interfaces/user_interface"
	"github.com/supabase-community/supabase-go"
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
