package service

import (
	"blog_service/internal/model"
	"blog_service/pkg/app"
)

type CountTagRequest struct {
	Name  string `from:"name" binding:"max=100"`
	State uint8  `from:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `from:"name" binding:"max=100"`
	State uint8  `from:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `from:"name" binding:"required,min=3,max=100"`
	State     uint8  `from:"state,default=1" binding:"oneof=0 1"`
	CreatedBy string `from:"created_by" binding:"required,min=3,max=100"`
}

type UpdateTagRequest struct {
	ID         uint32 `from:"id" binding:"required,gte=1"`
	Name       string `from:"name" binding:"required,min=3,max=100"`
	State      uint8  `from:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy string `from:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `from:"id" binding:"required,gte=1"`
}

func (svc *Service) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}
