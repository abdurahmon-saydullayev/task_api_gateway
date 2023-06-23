package handlers

import (
	"context"
	"fmt"
	"task_go_api_gateway/api/http"
	"task_go_api_gateway/genproto/blogpost_service"
	"task_go_api_gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateBook godoc
// @ID create_book
// @Router /book [POST]
// @Summary Create Book
// @Description  Create Book
// @Tags Book
// @Accept json
// @Produce json
// @Param profile body blogpost_service.CreateBook true "CreateBook"
// @Success 200 {object} http.Response{data=blogpost_service.Book} "GetBookBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateBook(c *gin.Context) {
	var book blogpost_service.CreateBook

	err := c.ShouldBindJSON(&book)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.BookService().Create(
		c.Request.Context(),
		&book,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	fmt.Println(resp)
	fmt.Println("lksdkmldsnmlkdsncklds")
	h.handleResponse(c, http.Created, resp)
}

// GetBookByID godoc
// @ID get_book_by_id
// @Router /book/{id} [GET]
// @Summary Get Book By ID
// @Description Get Book By ID
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=blogpost_service.Book} "Book"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetBookByID(c *gin.Context) {
	BookId := c.Param("id")

	if !util.IsValidUUID(BookId) {
		h.handleResponse(c, http.InvalidArgument, "Book id is an invalid uuid")
		return
	}

	resp, err := h.services.BookService().GetByID(
		context.Background(),
		&blogpost_service.BookPK{
			Id: BookId,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	fmt.Println(resp)
	fmt.Println(resp.Name)
	fmt.Println("lksdkmldsnmlkdsncklds")
	h.handleResponse(c, http.OK, resp)
}

// GetBookList godoc
// @ID get_book_list
// @Router /book [GET]
// @Summary Get Book List
// @Description Get Book List
// @Tags Book
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=blogpost_service.GetBookListResponse} "GetAllBookResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetBookList(c *gin.Context) {

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

	resp, err := h.services.BookService().GetList(
		context.Background(),
		&blogpost_service.GetBookListRequest{
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

// @ID update_book
// @Router /book/{id} [PUT]
// @Summary Update Book
// @Description Update Book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body blogpost_service.UpdateBookResponse true "UpdateBookRequestBody"
// @Success 200 {object} http.Response{data=blogpost_service.Book} "Book data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateBook(c *gin.Context) {

	var Book blogpost_service.UpdateBookResponse

	Book.Id = c.Param("id")

	if !util.IsValidUUID(Book.Id) {
		h.handleResponse(c, http.InvalidArgument, "Book id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&Book)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.BookService().Update(
		c.Request.Context(),
		&Book,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteBook godoc
// @ID delete_Book
// @Router /Book/{id} [DELETE]
// @Summary Delete Book
// @Description Delete Book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Book data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteBook(c *gin.Context) {

	BookId := c.Param("id")

	if !util.IsValidUUID(BookId) {
		h.handleResponse(c, http.InvalidArgument, "Book id is an invalid uuid")
		return
	}

	resp, err := h.services.BookService().Delete(
		c.Request.Context(),
		&blogpost_service.BookPK{Id: BookId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
