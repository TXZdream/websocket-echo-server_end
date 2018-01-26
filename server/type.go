package server

// Command is the JSON format between web server and docker server
type Command struct {
	Command    string   `json:"command"`
	Entrypoint []string `json:"entrypoint"`
	PWD        string   `json:"pwd"`
	ENV        []string `json:"env"`
	UserName   string   `json:"user"`
}
