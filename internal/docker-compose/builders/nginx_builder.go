package builders

import (
	"fmt"
	"log"
	"os"
)

const nginxContent = "server {\n    listen 80;\n    server_name localhost;\n    root /var/www/public;\n\n    index index.php index.html;\n\n    location / {\n        try_files $uri /index.php$is_args$args;\n    }\n\n    location ~ \\.php$ {\n        include fastcgi_params;\n        fastcgi_pass php_app:9000; # Match with the PHP service name\n        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;\n        fastcgi_param PATH_INFO $fastcgi_path_info;\n    }\n\n    location ~ /\\.ht {\n        deny all;\n    }\n}"

type NginxBuilder struct{}

func (b *NginxBuilder) Create(info *ServiceInfo) (Service, error) {
	rootDirectory, _ := GetAppRoot()

	var dirPath string = fmt.Sprintf("%s/docker_app/nginx", rootDirectory)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {

		err = os.MkdirAll(dirPath, DirPerm)
		if err != nil {
			return Service{}, err
		}
	}

	file, err := os.Create(fmt.Sprintf("%s/%s", dirPath, "nginx.conf"))

	defer file.Close()

	if err != nil {
		log.Fatalf("error creating nginx.conf - %s", err)
	}

	_, err = file.WriteString(nginxContent)

	if err != nil {
		log.Fatalf("error writing nginx.conf - %s", err)
	}

	return Service{
		Image:         fmt.Sprintf("%s:%s", info.ImageName, info.Tag),
		Ports:         ConvertRelatedPortsToString(info.Ports),
		Volumes:       ConvertVolumeToString(info.Volumes),
		ContainerName: fmt.Sprintf("%s_%s", info.ImageName, info.Tag),
		DependsOn:     info.DependsOn,
		Networks:      info.Networks,
	}, nil
}
