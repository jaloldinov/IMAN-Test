package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

// InsertPosts godoc
// @Description Gets post from open api and saves it to db
// @Summary saves post to db
// @Tags First Service
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseModel{data=first_service.InsertPostsResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
// @Router /v1/api [GET]
func (h *handlerV1) InsertPosts(c *gin.Context) {

	resp, err := h.services.FirstService().InsertPosts(
		c.Request.Context(),
		&emptypb.Empty{},
	)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while inserting posts", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// ListPosts godoc
// @Description Checks latest operation whether save post is successful or not
// @Summary check if post is saved to db
// @Tags First Service
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseModel{data=first_service.CheckPostsResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
// @Router /v1/api/check [GET]
func (h *handlerV1) CheckPosts(c *gin.Context) {

	resp, err := h.services.FirstService().CheckPosts(
		c.Request.Context(),
		&emptypb.Empty{},
	)

	if err != nil {
		h.handleErrorResponse(c, 500, "connection refused", err)
		return
	}

	h.handleSuccessResponse(c, 200, "ok", resp)
}
