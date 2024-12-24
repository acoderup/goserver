package etcd

import (
	"context"
	"go.etcd.io/etcd/client/v3"
	"testing"
)

func TestClient_AddFunc(t *testing.T) {
	cli := new(Client)

	cli.AddFunc("abc", func(ctx context.Context, res *clientv3.GetResponse) {
		t.Log("initFunc")
	}, func(ctx context.Context, res *clientv3.WatchResponse) {
		t.Log("watchFunc")
	})

	err := cli.Open([]string{"localhost:2379"}, "", "", 0)
	if err != nil {
		t.Log(err)
		return
	}
	cli.Start()
	cli.Close()
	t.Log("close...")
}
