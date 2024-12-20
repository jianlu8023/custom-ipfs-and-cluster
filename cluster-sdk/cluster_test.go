package ipfs_cluster

import (
	"context"
	"fmt"
	"testing"

	rclient "github.com/ipfs-cluster/ipfs-cluster/api/rest/client"
	"ipfs-cluster/internal/config"
)

func TestInitSdk(t *testing.T) {
	ok := InitSdk("127.0.0.1", "", "", 9094, "TCP")

	fmt.Println(ok)

}

func TestLbSdk(t *testing.T) {
	client, err := rclient.NewLBClient(
		&rclient.Failover{},
		[]*rclient.Config{
			{
				Username: "",
				Password: "",
				Host:     "127.0.0.1",
				Port:     "9094",
				LogLevel: "info",
			},
		},
		config.Retries,
	)
	if err != nil {
		t.Fatal(err)
	}
	version, err := client.Version(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(version)
	err = client.Health(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	graph, err := client.GetConnectGraph(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(graph)

}
