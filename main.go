package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
)

const (
	testProject  = "test"
	testInstance = "test"
	testDatabase = "test"
)

func main() {
	ctx := context.Background()

	cli, _ := connectTestCloudSpanner()

	_, err := cli.Single().ReadRow(ctx, "test", spanner.Key{0}, []string{"test"})

	println(err.Error())
}

func connectTestCloudSpanner() (*spanner.Client, error) {
	ctx := context.Background()
	return spanner.NewClient(ctx, fmt.Sprintf("projects/%s/instances/%s/databases/%s", testProject, testInstance, testDatabase))
}
