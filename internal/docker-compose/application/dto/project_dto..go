package dto

type ProjectDTO struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Status      string       `json:"status"`
	CreatedAt   string       `json:"createdAt"`
	UpdatedAt   string       `json:"updatedAt"`
	Services    []ServiceDTO `json:"services"`
}

type ServiceDTO struct {
	Name        string            `json:"name"`
	Image       string            `json:"image"`
	Ports       []PortMapping     `json:"ports"`
	Environment map[string]string `json:"environment"`
}

type PortMapping struct {
	Host      int    `json:"host"`
	Container int    `json:"container"`
	Protocol  string `json:"protocol"`
}
