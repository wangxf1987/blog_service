package v1

import "github.com/gin-gonic/gin"

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {

}

// @Summary get tags
// @Produce json
// @Param name query string false "tag name" maxlength(100)
// @Param state query int false "state" Enums(0, 1) default(1)
// @Param page query int false "page number"
// @Param page_size query int false "number of every page"
// @Success 200 {object} model.TagSwagger "request success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal error"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {

}

// @Summary create tag
// @Produce json
// @Param name query string false "tag name" maxlength(100)
// @Param state query int false "state" Enums(0, 1) default(1)
// @Param created_by body string false "creator" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "request success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal error"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {

}

// @Summary update tag
// @Produce json
// @Param id path int true "tag ID"
// @Param name query string false "tag name" maxlength(100)
// @Param state query int false "state" Enums(0, 1) default(1)
// @Param modified_by body string true "modifier" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "request success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal error"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {

}

// @Summary delete tag
// @Produce json
// @Param id path int true "tag ID"
// @Success 200 {object} model.Tag "request success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal error"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {

}
