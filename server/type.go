package server

// Command is the JSON format between web server and docker server
type Command struct {
	Operation   string `json:"operation"`
	Username    string `json:"username"`
	ProjectName string `json:"projectname"`
	Mainfile    string `json:"mainfile"`
}
