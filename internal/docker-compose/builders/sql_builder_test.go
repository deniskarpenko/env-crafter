package builders

import (
	"fmt"
	"testing"
)

var mysqlService = ServiceInfo{
	ImageName: "mysql",
	Tag:       "latest",
	Ports:     []RelatedPorts{{LocalPort: 3306, ContainerPort: 3320}},
	Volumes:   []Volume{{"mysql_data", "/var/lib/mysql"}},
	Envs: []string{
		"MYSQL_ROOT_PASSWORD=secret",
		"MYSQL_DATABASE=test_db",
		"MYSQL_USER=admin_db",
		"MYSQL_PASSWORD=password",
	},
	ServiceType: SQL,
	Networks:    []string{"app_network"},
}

func TestSQLBuilder_Create(t *testing.T) {
	builder := SQLBuilder{}

	service, err := builder.Create(&mysqlService)

	if err != nil {
		t.Error(err)
	}

	imageWithTag := fmt.Sprintf("%s:%s", mysqlService.ImageName, mysqlService.Tag)

	if service.Image != imageWithTag {
		t.Error("Image not match")
	}

	ports := ConvertRelatedPortsToString(mysqlService.Ports)

	for _, port := range ports {
		equalPorts := fmt.Sprintf("%d:%d", mysqlService.Ports[0].LocalPort, mysqlService.Ports[0].ContainerPort)

		if port != equalPorts {
			t.Error("Port not match")
		}
	}

	volumes := ConvertVolumeToString(mysqlService.Volumes)

	for _, volume := range volumes {
		equalVolumes := fmt.Sprintf("%s:%s", mysqlService.Volumes[0].LocalPath, mysqlService.Volumes[0].ContainerPath)

		if volume != equalVolumes {
			t.Error("Volume not match")
		}
	}

}
