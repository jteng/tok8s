package tok8s

import (
	"encoding/json"
)

//BuildCmd returns a kuberctl patch command if the image is managed, otherwise empty string.
//It will return error if there is error during json.Marshal(Deployment) process
func BuildCmd(config map[string]DeploymentConfig, bm BuildMessage) (string, error) {
	c, ok := config[bm.GetImage()]
	if !ok {
		return "", nil
	}
	d := Deployment{
		Spec: Spec{
			Template: Template{
				Spec: TemplateSpec{
					Containers: []Container{
						{Name: c.ContainerName, Image: bm.Digest},
					},
				},
			},
		},
	}
	ds, err := json.Marshal(d)
	if err != nil {
		return "", err
	}

	return "kubectl patch deployment " + c.DeploymentName + " -p '" + string(ds) + "'", nil
}
