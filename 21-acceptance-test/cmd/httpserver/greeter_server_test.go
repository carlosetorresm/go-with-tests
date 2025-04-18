package main_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/carlosetorresm/go-with-tests/21-acceptance-test/adapters"
	"github.com/carlosetorresm/go-with-tests/21-acceptance-test/adapters/httpserver"
	"github.com/carlosetorresm/go-with-tests/21-acceptance-test/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "8080"
		dockerFilePath = "./cmd/httpserver/Dockerfile"
		baseUrl        = fmt.Sprintf("http://localhost:$s", port)
		driver         = httpserver.Driver{BaseURL: baseUrl, Client: &http.Client{
			Timeout: 1 * time.Second,
		}}
	)

	adapters.StartDockerServer(t, port, dockerFilePath)
	specifications.GreetSpecification(t, driver)
}
