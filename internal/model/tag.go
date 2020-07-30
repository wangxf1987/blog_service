package model

import "blog_service/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State string `json:"state"`
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (a Tag) TableName() string {
	return "blog_tag"
}
