package service

import (
	"github.com/txzdream/websocket-echo-server_end/core/models/entities"
)

func NewProject(projectName string, language int) entities.ProjectInfo {
	// Get username
	userID := GetUserID()
	project := entities.ProjectInfo{
		ProjectName: projectName,
		UserID: userID,
		Language: language,
	}
	return project
}

// CreateProject insert a new project to the db
func CreateProject(project entities.ProjectInfo) (bool, error) {
	_, err := entities.Engine.InsertOne(project)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetProjects(userID int) []entities.ProjectInfo {
	projects := make([]entities.ProjectInfo, 0)
	entities.Engine.Where("projects.user_id = ?", userID).Find(&projects)
	return projects
}

func GetUserID() int {
	// TODO: return actual user id
	return 1
}
