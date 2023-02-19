package main

import (
	"flag"
	"fmt"

	"github.com/hashicorp/vault/api"
)

var (
	token string
)

type DBManager struct {
	client *api.Client
}

func NewDBManager() (*DBManager, error) {
	if len(token) == 0 {
		return nil, fmt.Errorf("token is empty")
	}

	config := api.DefaultConfig()
	vc, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	vc.SetToken(token)

	return &DBManager{
		client: vc,
	}, nil
}

func (mgr *DBManager) GetDBCreds() (*api.Secret, error) {
	if mgr == nil {
		return nil, fmt.Errorf("db manager is nil")
	}

	path := fmt.Sprint("/v1/database/creds/db-reader-role")
	req := mgr.client.NewRequest("GET", path)

	resp, err := mgr.client.RawRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	sr, err := api.ParseSecret(resp.Body)
	if err != nil {
		return nil, err
	}

	return sr, nil
}

func main() {

}

func init() {
	flag.StringVar(&token, "token", "x", "vault root token")
	flag.Parse()
}
