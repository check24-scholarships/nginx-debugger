package sandboxFactory

import (
	"context"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"nginx_debugger/internal/sandbox"
)

type SandboxFactoryInterface interface {
	InitializeDocker() error
	CreateSandbox() (sandbox.SandboxInterface, error)
}

type impl struct {
	client *client.Client
}

func NewSandboxFactory() (SandboxFactoryInterface, error) {
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("error while getting docker client: %s", err.Error())
	}
	return &impl{
		client: client,
	}, nil
}

func (i *impl) InitializeDocker() error {
	reader, err := i.client.ImagePull(
		context.Background(),
		"docker.io/library/nginx",
		types.ImagePullOptions{},
	)
	if err != nil {
		return fmt.Errorf("error while pulling image: %s", err.Error())
	}

	defer reader.Close()

	data := make([]byte, 100)
	for {
		_, err = reader.Read(data)
		if err != nil && errors.Is(err, io.EOF) {
			break
		}
	}

	return nil
}

func (i *impl) CreateSandbox() (sandbox.SandboxInterface, error) {
	sandbox, err := sandbox.NewSandbox(i.client)
	if err != nil {
		return nil, fmt.Errorf("error while constructing sandbox: %s", err.Error())
	}
	return sandbox, nil
}
