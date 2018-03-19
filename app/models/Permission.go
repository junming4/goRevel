package models

type Permissions struct {
	PermissionId int64 `gorm:"default:'permission_id'"`
	Name         string
	Slug         string
	Description  string
	Model        string
}

func (permissions Permissions) CreatePermission() bool {

	return true
}