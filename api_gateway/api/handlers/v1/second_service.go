package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jaloldinov/IMAN-Updated/api_gateway/api/models"
	"github.com/jaloldinov/IMAN-Updated/api_gateway/genproto/second_service"
)

// ListPosts godoc
// @Description List posts with limit, offset and based on search query
// @Summary Get list of posts
// @Tags Second Service
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} models.ResponseModel{data=second_service.ListPostsResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
// @Router /v1/post/list/ [GET]
func (h *handlerV1) ListPosts(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.SecondService().ListPosts(
		c.Request.Context(),
		&second_service.ListPostsRequest{
			Limit:  uint32(limit),
			Offset: uint32(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, 500, "error while getting posts", err)
		return
	}

	h.handleSuccessResponse(c, 200, "ok", resp)
}

// GetPost godoc
// @Description Gets post by id
// @Summary retruns unique post by id
// @Tags Second Service
// @Accept json
// @Produce json
// @Param post_id path integer true "post_id"
// @Success 200 {object} models.ResponseModel{data=models.Post} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
// @Router /v1/post/{post_id} [GET]
func (h *handlerV1) GetPost(c *gin.Context) {
	var post models.Post

	post_id := c.Param("post_id")
	coverted_post_id, err := strconv.ParseInt(post_id, 10, 64)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while converting post_id", err)
		return
	}

	resp, err := h.services.SecondService().GetPost(
		c.Request.Context(),
		&second_service.GetPostRequest{
			PostId: int64(coverted_post_id),
		},
	)
	if !handleError(h.log, c, err, "error while getting post") {
		return
	}

	post = models.Post{
		ID:     resp.PostId,
		UserId: resp.UserId,
		Title:  resp.Title,
		Body:   resp.Body,
	}
	h.handleSuccessResponse(c, http.StatusOK, "ok", &post)
}

// UpdatePost godoc
// @Description Updates post by post_id
// @Summary updates post by  post_id
// @Tags Second Service
// @Accept json
// @Produce json
// @Param post body second_service.Post true "post"
// @Success 200 {object} models.ResponseModel{data=second_service.Result} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
// @Router /v1/post/{post_id} [PUT]
func (h *handlerV1) UpdatePost(c *gin.Context) {
	var (
		status models.Status
		post   second_service.Post
	)

	if err := c.BindJSON(&post); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding JSON", err)
		return
	}

	resp, err := h.services.SecondService().UpdatePost(c.Request.Context(), &post)

	if !handleError(h.log, c, err, "error while updating post") {
		return
	}

	status = models.Status{
		Status: resp.Message,
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", &status)
}

// DeletePost godoc
// @Description Deletes post by id
// @Summary deletes post by id
// @Tags Second Service
// @Accept json
// @Produce json
// @Param post_id path integer true "post_id"
// @Success 200 {object} models.ResponseModel{data=second_service.Result}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
// @Router /v1/post/{post_id} [DELETE]
func (h *handlerV1) DeletePost(c *gin.Context) {
	var (
		status models.Status
	)

	post_id := c.Param("post_id")
	coverted_post_id, err := strconv.ParseInt(post_id, 10, 64)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while converting post_id", err)
		return
	}

	resp, err := h.services.SecondService().DeletePost(
		c.Request.Context(),
		&second_service.GetPostRequest{
			PostId: int64(coverted_post_id),
		},
	)

	if !handleError(h.log, c, err, "error while deleting post") {
		return
	}

	status = models.Status{
		Status: resp.Message,
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", &status)
}
