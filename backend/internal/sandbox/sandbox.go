package sandbox

import (
	"context"
	"fmt"
	"github.com/docker/distribution/uuid"
	"github.com/docker/docker/api/types"
	apiContainer "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type SandboxInterface interface {
	Start() error
	//SetConfiguration(configuration io.Reader) error
	//SetConfigurationStr(configuration string) error
	Stop() error
	Destroy() error
	StopAndDestroy() error
}

type impl struct {
	containerID string
	client      *client.Client
}

func NewSandbox(client *client.Client) (SandboxInterface, error) {
	uuid := uuid.Generate()

	containerResponse, err := client.ContainerCreate(
		context.Background(),
		&apiContainer.Config{
			Image: "nginx",
			Tty:   false,
		},
		nil,
		nil,
		nil,
		fmt.Sprintf("nginx-debugger-nginx-%s", uuid.String()),
	)

	if err != nil {
		return nil, fmt.Errorf("error creating container while sandbox creation: %s", err.Error())
	}

	return &impl{
		containerID: containerResponse.ID,
		client:      client,
	}, nil
}

func (s *impl) Start() error {
	if err := s.client.ContainerStart(context.Background(), s.containerID, types.ContainerStartOptions{}); err != nil {
		return fmt.Errorf("error while starting container: %s", err.Error())
	}
	return nil
}

//func (s *impl) SetConfigurationStr(configuration string) error {
//	reader := strings.NewReader(configuration)
//	if err := s.SetConfiguration(reader); err != nil {
//		return fmt.Errorf("error while setting configuration: %s", err.Error())
//	}
//	return nil
//}

//func (s *impl) SetConfiguration(configuration io.Reader) error {
//	if err := s.client.CopyToContainer(
//		context.Background(),
//		s.containerID,
//		"/etc/nginx/conf.d/default.conf",
//		configuration,
//		types.CopyToContainerOptions{
//			AllowOverwriteDirWithFile: true,
//		},
//	); err != nil {
//		return fmt.Errorf("error while copying configuration to container: %s", err.Error())
//	}
//
//	if err := s.client.ContainerRestart(
//		context.Background(),
//		s.containerID,
//		apiContainer.StopOptions{},
//	); err != nil {
//		return fmt.Errorf("error while restarting container: %s", err.Error())
//	}
//	return nil
//}

func (s *impl) Stop() error {
	if err := s.client.ContainerStop(context.Background(), s.containerID, apiContainer.StopOptions{}); err != nil {
		return fmt.Errorf("error while stopping container: %s", err.Error())
	}
	return nil
}

func (s *impl) Destroy() error {
	if err := s.client.ContainerRemove(context.Background(), s.containerID, types.ContainerRemoveOptions{}); err != nil {
		return fmt.Errorf("error while destroying container: %s", err.Error())
	}
	return nil
}

func (s *impl) StopAndDestroy() error {
	if err := s.Stop(); err != nil {
		return fmt.Errorf("error while stopping container: %s", err.Error())
	}

	if err := s.Destroy(); err != nil {
		return fmt.Errorf("error while destroying container: %s", err.Error())
	}

	return nil
}
