package api

import (
	"errors"
	"github.com/deepaksinghvi/perfrun/pkg/service"
	"github.com/deepaksinghvi/perfrun/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var URL = "https://be66f1a1-3b11-4cfc-a293-0195be93f244.mock.pstmn.io"

// GetPerfRun godoc
// @Summary Get Perfrun
// @ID get-perfrun-by-id
// @Tags Perfrun
// @Accept  json
// @Produce  json
// @Param id path int64 true "Perfrun ID"
// @Success 200 {object} service.PerfRunResult
// @Failure 404 {object} utils.HTTPError
// @Router /perfrunws/{id} [get]
func GetPerfRun(ctx *gin.Context) {
	perfRunId, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	perfRun, err := service.GetPerfRun(URL, int(perfRunId))
	if nil != err {
		utils.NewError(ctx, http.StatusBadRequest, errors.New("Perfrun not found!"))
		return
	}
	ctx.JSON(http.StatusOK, perfRun)
}

// CreatePerfRun godoc
// @Summary Create New Perfrun
// @ID create-new-perfrun
// @Tags Perfrun
// @Accept json
// @Produce json
// @Param perfRunRequest body service.PerfRunRequest true "Create PerfRun"
// @Success 200 {object} service.PerfRunResult
// @Failure 500 {object} utils.HTTPError
// @Router /perfrunws [post]
func CreatePerfRun(ctx *gin.Context) {
	var perfRunRequest service.PerfRunRequest
	err := ctx.BindJSON(&perfRunRequest)
	if err != nil {
		ctx.JSON(401, err)
		return
	}
	result := service.CreatePerfRun(URL,perfRunRequest)
	ctx.JSON(http.StatusCreated, result)
	return
}