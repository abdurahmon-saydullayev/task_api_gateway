package handlers

import (
	"context"
	"task_go_api_gateway/api/http"
	"task_go_api_gateway/genproto/blogpost_service"
	"task_go_api_gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateAuthor godoc
// @ID create_author
// @Router /author [POST]
// @Summary Create Author
// @Description  Create Author
// @Tags Author
// @Accept json
// @Produce json
// @Param profile body blogpost_service.CreateAuthor true "CreateAuthor"
// @Success 200 {object} http.Response{data=blogpost_service.Author} "GetAuthorBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateAuthor(c *gin.Context) {
	var Author blogpost_service.CreateAuthor

	err := c.ShouldBindJSON(&Author)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.AuthorService().Create(
		c.Request.Context(),
		&Author,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetAuthorByID godoc
// @ID get_Author_by_id
// @Router /author/{id} [GET]
// @Summary Get Author By ID
// @Description Get Author By ID
// @Tags Author
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=blogpost_service.Author} "Author"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAuthorByID(c *gin.Context) {
	AuthorId := c.Param("id")

	if !util.IsValidUUID(AuthorId) {
		h.handleResponse(c, http.InvalidArgument, "Author id is an invalid uuid")
		return
	}

	resp, err := h.services.AuthorService().GetByID(
		context.Background(),
		&blogpost_service.AuthorPK{
			Id: AuthorId,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetAuthorList godoc
// @ID get_Author_list
// @Router /author [GET]
// @Summary Get Author List
// @Description Get Author List
// @Tags Author
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=blogpost_service.GetAuthorListResponse} "GetAllAuthorResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAuthorList(c *gin.Context) {

	offset, err := h.getOffsetParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.AuthorService().GetList(
		context.Background(),
		&blogpost_service.GetAuthorListRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @ID update_Author
// @Router /author/{id} [PUT]
// @Summary Update Author
// @Description Update Author
// @Tags Author
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body blogpost_service.UpdateAuthorResponse true "UpdateAuthorRequestBody"
// @Success 200 {object} http.Response{data=blogpost_service.Author} "Author data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateAuthor(c *gin.Context) {

	var Author blogpost_service.UpdateAuthorResponse

	Author.Id = c.Param("id")

	if !util.IsValidUUID(Author.Id) {
		h.handleResponse(c, http.InvalidArgument, "Author id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&Author)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.AuthorService().Update(
		c.Request.Context(),
		&Author,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteAuthor godoc
// @ID delete_Author
// @Router /author/{id} [DELETE]
// @Summary Delete Author
// @Description Delete Author
// @Tags Author
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Author data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteAuthor(c *gin.Context) {

	AuthorId := c.Param("id")

	if !util.IsValidUUID(AuthorId) {
		h.handleResponse(c, http.InvalidArgument, "Author id is an invalid uuid")
		return
	}

	resp, err := h.services.AuthorService().Delete(
		c.Request.Context(),
		&blogpost_service.AuthorPK{Id: AuthorId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
