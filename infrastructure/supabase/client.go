package supabase_client

import (
	"github.com/HackU-team04/Pedometer/dependency/config"
	"github.com/supabase-community/supabase-go"
)

func ConnectSupabase(supabaseConfig config.SupabaseConfig) (*supabase.Client, error) {
	client, err := supabase.NewClient(supabaseConfig.APIURL, supabaseConfig.APIKey, nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}
