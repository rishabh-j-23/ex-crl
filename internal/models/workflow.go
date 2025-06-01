package models

type WorkflowStep struct {
	RequestName string `json:"request-name"`
	Exec        bool   `json:"exec"`
}

type Workflow struct {
	Workflow []WorkflowStep `json:"workflow"`
}
