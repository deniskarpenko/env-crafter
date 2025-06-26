package models

import "time"

type Image struct {
	Id           int
	Image        string
	IsDockerFile bool
	CreatedAt    time.Time
}
