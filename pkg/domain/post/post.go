package post

import (
	"gorm.io/gorm"
)

type Post struct {
    gorm.Model
    Title string
    Body  string
}

func (Post) TableName() string {
    return "posts"
}
