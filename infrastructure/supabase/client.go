package supabase_client

import (
	"sync"

	"github.com/supabase-community/supabase-go"
	"github.com/xkurozaru/pedometer-server/dependency/config"
)

var client_instance *supabase.Client
var err_instance error
var once sync.Once

func ConnectSupabase(supabaseConfig config.SupabaseConfig) (*supabase.Client, error) {
	once.Do(func() {
		client, err := supabase.NewClient(supabaseConfig.APIURL, supabaseConfig.APIKey, nil)
		if err != nil {
			err_instance = err
			return
		}
		client_instance = client
	})

	return client_instance, err_instance
}
