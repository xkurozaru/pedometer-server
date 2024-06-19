package registry

import (
	"github.com/supabase-community/supabase-go"
	auth_application "github.com/xkurozaru/pedometer-server/application/auth"
	user_application "github.com/xkurozaru/pedometer-server/application/user"
	walking_record_application "github.com/xkurozaru/pedometer-server/application/walking_record"

	"github.com/xkurozaru/pedometer-server/dependency/config"
	"github.com/xkurozaru/pedometer-server/domain/auth"
	"github.com/xkurozaru/pedometer-server/domain/user"
	"github.com/xkurozaru/pedometer-server/domain/walking_record"
	"github.com/xkurozaru/pedometer-server/infrastructure/database"
	user_database "github.com/xkurozaru/pedometer-server/infrastructure/database/user"
	walking_record_database "github.com/xkurozaru/pedometer-server/infrastructure/database/walking_record"
	supabase_client "github.com/xkurozaru/pedometer-server/infrastructure/supabase"
	supabase_auth "github.com/xkurozaru/pedometer-server/infrastructure/supabase/auth"
	"github.com/xkurozaru/pedometer-server/interfaces/auth_interface"
	"github.com/xkurozaru/pedometer-server/interfaces/user_interface"
	"github.com/xkurozaru/pedometer-server/interfaces/walking_record_interface"
	"gorm.io/gorm"
)

type Registry interface {
	NewAuthHandler() auth_interface.AuthHandler
	NewUserHandler() user_interface.UserHandler
	NewWalkingRecordHandler() walking_record_interface.WalkingRecordHandler
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

func (r registry) NewWalkingRecordHandler() walking_record_interface.WalkingRecordHandler {
	return walking_record_interface.NewWalkingRecordHandler(
		r.NewUserApplicationService(),
		r.NewWalkingRecordApplicationService(),
	)
}

func (r registry) NewAuthApplicationService() auth_application.AuthApplicationService {
	return auth_application.NewAuthApplicationService(
		r.NewAuthRepository(),
	)
}

func (r registry) NewUserApplicationService() user_application.UserApplicationService {
	return user_application.NewUserApplicationService(
		r.NewAuthRepository(),
		r.NewUserRepository(),
	)
}

func (r registry) NewWalkingRecordApplicationService() walking_record_application.WalkingRecordApplicationService {
	return walking_record_application.NewWalkingRecordApplicationService(
		r.NewUserRepository(),
		r.NewWalkingRecordRepository(),
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

func (r registry) NewWalkingRecordRepository() walking_record.WalkingRecordRepository {
	return walking_record_database.NewWalkingRecordDatabase(r.NewDB())
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
