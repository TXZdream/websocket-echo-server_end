package entities

// ProjectInfo store project information
type ProjectInfo struct {
}

func (u ProjectInfo) TableName() string {
	return "projects"
}
