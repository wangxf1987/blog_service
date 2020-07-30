package service

type CountArticleRequest struct {
	State uint8 `from:"state,default=1" binding:"oneof=0 1"`
}

type TagArticleRequest struct {
	Name  string `from:"name" binding:"max=100"`
	State uint8  `from:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Name          string `from:"name" binding:"required,min=3,max=100"`
	State         uint8  `from:"state,default=1" binding:"oneof=0 1"`
	Desc          string `from:"desc" binding:"max=512"`
	Content       string `from:"content"`
	CoverImageUrl string `from:"cover_image_url" binding:"required,max=512"`
	CreatedBy     string `from:"created_by" binding:"required,min=3,max=100"`
}

type UpdateArticleRequest struct {
	ID            uint32 `from:"id" binding:"required,gte=1"`
	Name          string `from:"name" binding:"required,min=3,max=100"`
	State         uint8  `from:"state,default=1" binding:"oneof=0 1"`
	Desc          string `from:"desc" binding:"max=512"`
	Content       string `from:"content"`
	CoverImageUrl string `from:"cover_image_url" binding:"required,max=512"`
	ModifiedBy    string `from:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `from:"id" binding:"required,gte=1"`
}
