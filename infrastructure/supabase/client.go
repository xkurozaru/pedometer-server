package supabase_client

import (
	"github.com/supabase-community/supabase-go"
	"github.com/xkurozaru/pedometer-server/dependency/config"
)

func ConnectSupabase(supabaseConfig config.SupabaseConfig) (*supabase.Client, error) {
	client, err := supabase.NewClient(supabaseConfig.APIURL, supabaseConfig.APIKey, nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}
