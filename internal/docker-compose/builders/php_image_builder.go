package builders

import (
	"fmt"
	"log"
	"os"
)

const envPhp = "php.env"

const dockerFileName = "php.Dockerfile"

type PhpImageBuilder struct{}

func (b *PhpImageBuilder) Create(info *ServiceInfo) (Service, error) {
	content := b.getDockerfileContent(info)

	dockerFileBuilder := DockerfileBuilder{}

	rootDirectory, _ := GetAppRoot()

	var dirDockerPath string = fmt.Sprintf("%s/docker_app/dockerfiles/", rootDirectory)

	if _, err := os.Stat(dirDockerPath); os.IsNotExist(err) {

		err = os.MkdirAll(dirDockerPath, DirPerm)
		if err != nil {
			return Service{}, err
		}
	}

	filePath := dirDockerPath + dockerFileName

	err := dockerFileBuilder.Build(filePath, content)

	if err != nil {
		log.Println("Can't create php.Dockerfile:", err)
	}

	eb := EnvBuilder{}

	dirEnvPath := fmt.Sprintf("%s/docker_app/env/", rootDirectory)

	if _, err := os.Stat(dirEnvPath); os.IsNotExist(err) {

		err = os.MkdirAll(dirEnvPath, DirPerm)
		if err != nil {
			return Service{}, err
		}
	}

	filePath = dirEnvPath + envPhp

	_, err = eb.Build(info.Envs, filePath)

	if err != nil {
		log.Println("Can't create php.env:", err)
	}

	return Service{
		Image:         "",
		Ports:         ConvertRelatedPortsToString(info.Ports),
		Volumes:       ConvertVolumeToString(info.Volumes),
		ContainerName: fmt.Sprintf("%s_%s", info.ImageName, info.Tag),
		EnvFile:       "env/" + envPhp,
		Build: Build{
			Context:    ".",
			Dockerfile: "dockerfiles/" + dockerFileName,
		},
		DependsOn: info.DependsOn,
		Networks:  info.Networks,
	}, nil
}

func (b *PhpImageBuilder) getDockerfileContent(info *ServiceInfo) string {
	return fmt.Sprintf(`
FROM %s:%s AS builder

WORKDIR /var/www

RUN apt-get update && apt-get install -y \
    libpng-dev \
    libjpeg-dev \
    libfreetype6-dev \
    zip unzip git curl \
    libonig-dev libxml2-dev \
    && docker-php-ext-install pdo mbstring gd pdo_mysql mysqli pcntl

RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

RUN chown -R www-data:www-data /var/www

# ===== Final =====
FROM %s:%s

WORKDIR /var/www

RUN apt-get update && apt-get install -y \
    libpng-dev \
    libjpeg-dev \
    libfreetype6-dev \
    libonig-dev \
    libxml2-dev \
    && docker-php-ext-install pdo mbstring gd pdo_mysql mysqli pcntl

COPY --from=builder /var/www /var/www
RUN chown -R www-data:www-data /var/www

EXPOSE 9000
CMD ["php-fpm"]`, info.ImageName, info.Tag, info.ImageName, info.Tag)
}
