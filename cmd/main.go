package main

import (
	"bitbucket.org/go-scrapper/docker-compose/internal/docker-compose"
	"bitbucket.org/go-scrapper/docker-compose/internal/docker-compose/builders"
	"fmt"
	"log"
	"log/slog"
	"os"
)

var phpServise = builders.ServiceInfo{
	ImageName: "php",
	Tag:       "8.3-fpm",
	Volumes:   []builders.Volume{{".", "/var/www"}},
	Networks:  []string{"app_network"},
	Envs: []string{
		"APP_ENV=local",
		"DB_HOST=mysql",
		"DB_DATABASE=test_db",
		"DB_USERNAME=admin_db",
		"DB_PASSWORD=password",
	},
	ServiceType: builders.Backend,
}

var mysqlServise = builders.ServiceInfo{
	ImageName: "mysql",
	Tag:       "latest",
	Volumes:   []builders.Volume{{"mysql_data", "/var/lib/mysql"}},
	Envs: []string{
		"MYSQL_ROOT_PASSWORD=secret",
		"MYSQL_DATABASE=test_db",
		"MYSQL_USER=admin_db",
		"MYSQL_PASSWORD=password",
	},
	ServiceType: builders.SQL,
	Networks:    []string{"app_network"},
}

var nginxServise = builders.ServiceInfo{
	ImageName:   "nginx",
	Tag:         "alpine",
	Ports:       []builders.RelatedPorts{{8080, 8080}},
	Volumes:     []builders.Volume{{".", "/var/www"}},
	ServiceType: builders.WebServer,
	Networks:    []string{"app_network"},
}

func main() {
	appDir, err := builders.GetAppRoot()

	if err != nil {
		fmt.Println("Cant find app directory:", err)
	}

	logPath := appDir + "/logs/docker-compose.log"

	fileLogger, err := NewFileLogger(logPath)

	if err != nil {
		log.Fatalln("Can't create file logger:", err)
	}

	builder, err := docker.NewWebBuilder(fileLogger)

	if err != nil {
		fmt.Printf("Can't create builder: %s\n", err)
	}

	ps := docker.WebProjectStructure{
		Backend:   []builders.ServiceInfo{phpServise},
		SQL:       []builders.ServiceInfo{mysqlServise},
		WebServer: []builders.ServiceInfo{nginxServise},
	}

	error := builder.Build(ps)

	if error != nil {
		fmt.Println(error)
	}
}

func NewFileLogger(logPath string) (*slog.Logger, error) {
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return nil, err
	}

	handler := slog.NewJSONHandler(f, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	return slog.New(handler), nil
}
