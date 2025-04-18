package main_test

import (
	"fmt"
	"testing"

	"github.com/carlosetorresm/go-with-tests/21-acceptance-test/adapters"
	"github.com/carlosetorresm/go-with-tests/21-acceptance-test/adapters/grpcserver"
	"github.com/carlosetorresm/go-with-tests/21-acceptance-test/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "50051"
		dockerFilePath = "./cmd/grpcserver/Dockerfile"
		driver         = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	adapters.StartDockerServer(t, port, dockerFilePath)
	specifications.GreetSpecification(t, &driver)

}
