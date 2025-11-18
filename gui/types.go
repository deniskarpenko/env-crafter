package main

type ImageWithTag struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type ContainerConfig struct {
	Ports    []string `json:"ports"`
	Volumes  []string `json:"volumes"`
	EnvFiles [][]byte `json:"envFiles"` // В Go файлы передаются как пути (строки)
	Envs     []string `json:"envs"`
}

type ImageWithTagConfig struct {
	Image  *ImageWithTag    `json:"image"`  // pointer для null значений
	Config *ContainerConfig `json:"config"` // pointer для null значений
}

type ProjectConfig struct {
	Backend *ImageWithTagConfig `json:"backend"`
	SQL     *ImageWithTagConfig `json:"sql"`
	Nosql   *ImageWithTagConfig `json:"nosql"`
	Web     *ImageWithTagConfig `json:"web"`
}

type ApplicationConfig struct {
	Projects []ProjectConfig `json:"projects"`
}
