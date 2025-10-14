package builders

type ServiceInfo struct {
	ImageName   string
	Tag         string
	ServiceType ServiceType
	Ports       []RelatedPorts
	Volumes     []Volume
	Networks    []string
	Envs        []string
	DependsOn   []string
}

type RelatedPorts struct {
	LocalPort     int
	ContainerPort int
}

type Volume struct {
	LocalPath     string
	ContainerPath string
}

type ServiceType string

const (
	Backend   ServiceType = "backend"
	Frontend  ServiceType = "frontend"
	SQL       ServiceType = "sql"
	NOSQL     ServiceType = "no-sql"
	WebServer ServiceType = "web"
	Utility   ServiceType = "utility"
)
