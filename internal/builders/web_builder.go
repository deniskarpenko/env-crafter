package docker

import (
	"bitbucket.org/go-scrapper/docker-compose/internal/builders/builders"
	"fmt"
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
)

type WebProjectStructure struct {
	Backend   []builders.ServiceInfo
	SQL       []builders.ServiceInfo
	NOSQL     []builders.ServiceInfo
	Frontend  []builders.ServiceInfo
	WebServer []builders.ServiceInfo
	Utilities []builders.ServiceInfo
}

type WebBuilder struct {
	Logger *slog.Logger
}

func (wb *WebBuilder) Build(ps WebProjectStructure) error {
	wb.Logger.Info("Building docker compose project")

	sb := builders.NewServiceBuilder()

	dbServices, err, dbServiceNames := wb.getServices(ps.SQL, sb, []string{})

	backendServices, err, backendServiceNames := wb.getServices(ps.Backend, sb, dbServiceNames)

	webServices, err, _ := wb.getServices(ps.WebServer, sb, backendServiceNames)

	allServices := append([]builders.Service{}, dbServices...)
	allServices = append(allServices, backendServices...)
	allServices = append(allServices, webServices...)

	if err != nil {
		return err
	}

	serviceMap := make(map[string]builders.Service)

	for _, service := range allServices {
		serviceMap[service.ContainerName] = service
	}

	volumeName := ps.SQL[0].Volumes[0].LocalPath

	output := map[string]interface{}{
		"services": serviceMap,
		"volumes": map[string]interface{}{
			volumeName: map[string]interface{}{},
		},
		"networks": map[string]interface{}{
			"app_network": map[string]interface{}{
				"driver": "bridge",
			},
		},
	}

	data, err := yaml.Marshal(output)

	if err != nil {
		return err
	}

	rootDirectory, _ := builders.GetAppRoot()

	var filePath string = fmt.Sprintf("%s/docker_app/docker-compose.yaml", rootDirectory)

	err = os.WriteFile(filePath, data, 0644)

	if err != nil {
		return err
	}

	return nil
}

func (wb *WebBuilder) getServices(
	serviceInfo []builders.ServiceInfo,
	sb *builders.ServiceBuilder,
	dependsOn []string,
) ([]builders.Service, error, []string) {
	services := []builders.Service{}

	serviceNames := []string{}

	for _, info := range serviceInfo {
		info.DependsOn = dependsOn

		service, err := sb.CreateService(info)
		if err != nil {
			return services, err, serviceNames
		}

		service.DependsOn = dependsOn

		services = append(services, service)
	}

	return services, nil, serviceNames
}
