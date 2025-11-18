package dto

type BuildRequest struct {
	ProjectID string       `json:"projectId"`
	Options   BuildOptions `json:"options"`
}

type BuildOptions struct {
	NoCache  bool     `json:"noCache"`
	Parallel bool     `json:"parallel"`
	Services []string `json:"services"` // пустой = все сервисы
}

type BuildResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Logs    string `json:"logs,omitempty"`
}
