package main

import (
	"encoding/json"
	"os"
	"os/exec"

	"github.com/jteng/tok8s"
	log "github.com/sirupsen/logrus"
)

//GetSubscriber, GetDeploymentConfigs and ExecuteK8sCmd are functions
//which can be injected by test
var (
	GetSubscriber        = initSubscriber
	GetDeploymentConfigs = initDeploymentConfigs
	ExecuteK8sCmd        = executeCmd
)

func main() {
	sub := GetSubscriber(os.Getenv("SUBSCRIPTION"))
	msgChan := sub.Subscribe()
	configs := GetDeploymentConfigs()

	for {
		m, ok := <-msgChan
		if !ok {
			log.Infof("channel closed, quit")
			break
		}
		//TODO: should we really ack the message here?
		m.Ack()

		bm := tok8s.BuildMessage{}
		err := json.Unmarshal(m.Data(), &bm)
		if err != nil {
			log.Errorf("failed to unmarshal message %s", err.Error())
			continue
		}
		k8sCmd, err := tok8s.BuildCmd(configs, bm)
		if err != nil {
			log.Errorf("failed to build k8s command %s", err.Error())
			continue
		}
		if k8sCmd == "" {
			log.Infof("skip message, we don't manage the image %s", bm.GetImage())
		}
		if ExecuteK8sCmd(k8sCmd) != nil {
			log.Errorf("failed to execute command [%s]", k8sCmd)
		}
	}
}

func initSubscriber(subscription string) tok8s.Subscriber {
	return nil
}

func initDeploymentConfigs() map[string]tok8s.DeploymentConfig {
	//TODO: the configs should be in a config file or external db, we should read from there

	return map[string]tok8s.DeploymentConfig{
		"us.gcr.io/my-test/hello": {
			DeploymentName: "d-name", ContainerName: "c-name",
		},
	}
}

func executeCmd(k8scmd string) error {
	out, err := exec.Command("sh", "-c", k8scmd).Output()
	if err != nil {
		return err
	}
	log.Infof("command result: %s", out)
	return nil
}
