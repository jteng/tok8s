package tok8s

//Deployment is a deployment in kubernetes
type Deployment struct {
	Spec Spec `json:"spec" yaml:"spec"`
}

//Spec is a spec in deployment
type Spec struct {
	Template Template `json:"template" yaml:"template"`
}

//Container is a container inside template spec
type Container struct {
	Name  string `json:"name" yaml:"name"`
	Image string `json:"image" yaml:"image"`
}

//TemplateSpec is spec in template
type TemplateSpec struct {
	Containers []Container `json:"containers" yaml:"containers"`
}

//Template is a template in spec in deployment
type Template struct {
	Spec TemplateSpec `json:"spec" yaml:"spec"`
}

//DeploymentConfig is to define the managed docker image name to deployment/container
type DeploymentConfig struct {
	DeploymentName string
	ContainerName  string
}
