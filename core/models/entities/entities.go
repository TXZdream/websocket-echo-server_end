package entities

// ProjectInfo store project information
type ProjectInfo struct {
	ProjectID   int `xorm:"pk autoincr 'project_id'"`
	ProjectName string
	UserID      int `xorm:"'user_id'"`
	Language    int
}

func (u ProjectInfo) TableName() string {
	return "projects"
}
