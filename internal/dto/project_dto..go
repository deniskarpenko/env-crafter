package dto

type ProjectDTO struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Status      string       `json:"status"`
	CreatedAt   string       `json:"createdAt"`
	Services    []ServiceDTO `json:"services"`
}

type ServiceDTO struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Ports []int  `json:"ports"`
}

type BuildOptions struct {
	NoCache  bool     `json:"noCache"`
	Parallel bool     `json:"parallel"`
	Services []string `json:"services"` // какие сервисы собирать
}

type BuildResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Logs    string `json:"logs,omitempty"`
}
