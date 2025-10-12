package builders

type Service struct {
	Image         string   `yaml:"image,omitempty"`
	Ports         []string `yaml:"ports,omitempty"`
	Volumes       []string `yaml:"volumes,omitempty"`
	Build         Build    `yaml:"build,omitempty"`
	ContainerName string   `yaml:"container_name"`
	EnvFile       string   `yaml:"env_file,omitempty"`
	DependsOn     []string `yaml:"depends_on,omitempty"`
	Networks      []string `yaml:"networks,omitempty"`
}

type Build struct {
	Context    string `yaml:"context"`
	Dockerfile string `yaml:"dockerfile"`
}

type ServiceBuilder struct {
}

func (b *ServiceBuilder) CreateService(info ServiceInfo) (Service, error) {

	ib, err := GetImageBuilderByImageName(info)

	if err != nil {
		return Service{}, err
	}

	service, err := ib.Create(&info)

	if err != nil {
		return Service{}, err
	}

	return service, nil
}

func NewServiceBuilder() *ServiceBuilder {
	return &ServiceBuilder{}
}
