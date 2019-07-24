package disDataDict

import (
	"log"
	"go.etcd.io/etcd/client"
)

func New() client.KeysAPI {
  cfg := client.Config{
	   Endpoints: []string{"http://127.0.0.1:2379"},
	   Transport: client.DefaultTransport,
  }

  c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}


  kAPI := client.NewKeysAPI(c)

  return kAPI

}
