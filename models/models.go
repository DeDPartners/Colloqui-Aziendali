package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string `json:"username" gorm:"column:username;type:varchar(100);not null;unique"`
	Password string `json:"password" gorm:"column:password;size:256;not null"`
	Token    string `json:"token"    gorm:"column:token;size:256;not null;unique"`
}

type ProjectModel struct {
	gorm.Model
	Title    string      `json:"title"    gorm:"column:title;type:varchar(100);not null;unique"`
	Archived bool        `json:"archived" gorm:"column:archived;type:boolean;default:false"`
	Tasks    []TaskModel `json:"tasks"    gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE;"`
}

type TaskModel struct {
	gorm.Model
	Name      string     `json:"name"       gorm:"column:name;type:varchar(100);not null;unique"`
	Priority  int        `json:"priority"   gorm:"column:priority;type:int;not null;check:priority >= 1 AND priority <= 5"`
	Deadline  *time.Time `json:"deadline"   gorm:"column:deadline;type:timestamp"`
	Done      bool       `json:"done"       gorm:"column:done;type:boolean;default:false"`
	ProjectID uint       `json:"project_id" gorm:"column:project_id;type:bigint;not null"`
}

// INSERT INTO users (created_at, updated_at, deleted_at, username, password, token)
// VALUES (NOW(), NOW(), NULL, 'admin30092024', '4Jx9@mK2', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9');

