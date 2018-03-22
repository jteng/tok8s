package main

import (
	"testing"

	"github.com/jteng/tok8s"
	"github.com/jteng/tok8s/mock"
)

var tt *testing.T

func TestMainProcess(t *testing.T) {
	oldGetSub := GetSubscriber
	oldGetDeploy := GetDeploymentConfigs
	oldExec := ExecuteK8sCmd

	defer func() {
		GetSubscriber = oldGetSub
		GetDeploymentConfigs = oldGetDeploy
		ExecuteK8sCmd = oldExec
	}()
	GetSubscriber = getMockSubscriber
	GetDeploymentConfigs = getDepoymentConfigs
	ExecuteK8sCmd = runCommand
	tt = t
	main()

}

func getMockSubscriber(s string) tok8s.Subscriber {
	return mock.Subscriber{
		Messages: []tok8s.BuildMessage{
			tok8s.BuildMessage{Digest: "us.gcr.io/repo/image:12345"},
		},
	}
}

func getDepoymentConfigs() map[string]tok8s.DeploymentConfig {
	return map[string]tok8s.DeploymentConfig{
		"us.gcr.io/repo/image": {
			DeploymentName: "d-name", ContainerName: "c-name",
		},
	}
}

func runCommand(c string) error {
	if c == "" {
		tt.Errorf("expect a k8s command")
	}
	return nil
}
