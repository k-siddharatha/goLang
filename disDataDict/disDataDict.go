package disDataDict

import (
	"go.etcd.io/etcd/client"
)

func New() client.KeysAPI {
  cfg := client.Config{
	   Endpoints: []string{"http://127.0.0.1:2379"},
	   Transport: client.DefaultTransport,
  }

  c, _ := client.New(cfg)


  kAPI := client.NewKeysAPI(c)

  return kAPI

}
