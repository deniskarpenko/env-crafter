package builders

import (
	"fmt"
)

type ImageBuilder interface {
	Create(info *ServiceInfo) (Service, error)
}

func GetImageBuilderByImageName(info ServiceInfo) (ImageBuilder, error) {
	switch info.ImageName {
	case "php":
		return &PhpImageBuilder{}, nil
	case "mysql":
		return &SQLBuilder{}, nil
	case "postgres":
		return &SQLBuilder{}, nil
	case "nginx":
		return &NginxBuilder{}, nil
	}

	return nil, fmt.Errorf("unknown image %s", info.ImageName)
}

func ConvertRelatedPortsToString(ports []RelatedPorts) []string {
	var result []string

	for _, port := range ports {
		result = append(result, fmt.Sprintf("%d:%d", port.LocalPort, port.ContainerPort))
	}

	return result
}

func ConvertVolumeToString(volumes []Volume) []string {
	var result []string

	for _, v := range volumes {
		result = append(result, fmt.Sprintf("%s:%s", v.LocalPath, v.ContainerPath))
	}

	return result
}
