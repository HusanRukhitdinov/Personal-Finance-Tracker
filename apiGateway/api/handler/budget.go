package handler

import (
	_ "api/api/docs"
	pb "api/genproto/budgeting_service"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const layout = "2006-01-02T15:04:05Z07:00"
const defaultTime = "14:31:27.953" // Example time with milliseconds

const timeZoneOffset = "+00:00"

// CreateBudgetHandler   godoc
// @Router       /budget_service/v2/budget [POST]
// @Summary      Create a new budget
// @Description  Create a new budget
// @Tags         budget
// @Accept       json
// @Produce      json
// @Param        budget body models.CreateBudget true "budget"
// @Success      201  {object}  models.Budget
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
// @Security 	 ApiKeyAuth
func (h *Handler) CreateBudgetHandler(ctx *gin.Context) {
	var (
		request  pb.BudgetRequest
		response *pb.Budget
		err      error
	)

	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		handleResponse(ctx, h.log, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	request.StartTime += "T" + defaultTime + timeZoneOffset
	request.EndTime += "T" + defaultTime + timeZoneOffset

	startTime, err := time.Parse(layout, request.GetStartTime())
	if err != nil {
		handleResponse(ctx, h.log, "start_time parse error", http.StatusBadRequest, err.Error())
		return
	}

	endTime, err := time.Parse(layout, request.GetEndTime())
	if err != nil {
		handleResponse(ctx, h.log, "end_time parse error", http.StatusBadRequest, err.Error())
		return
	}

	request.StartTime = startTime.Format(time.RFC3339)
	request.EndTime = endTime.Format(time.RFC3339)

	response, err = h.services.BudgetService().CreateBudget(ctx, &request)
	if err != nil {
		handleResponse(ctx, h.log, "error updating budget in message broker", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(ctx, h.log, "success", http.StatusCreated, response)
}

// UpdateBudgetHandler godoc
// @Router       /budget_service/v2/budget/{id} [PUT]
// @Summary      Update budget
// @Description  Update budget
// @Tags         budget
// @Accept       json
// @Produce      json
// @Param        id path string true "budget_id"
// @Param        budget body models.UpdateBudget false "budget"
// @Success      200  {object}  models.Budget
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
// @Security 	 ApiKeyAuth
func (h *Handler) UpdateBudgetHandler(ctx *gin.Context) {

	var (
		request pb.Budget
		id      string
		err     error
	)
	id = ctx.Param("id")
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		handleResponse(ctx, h.log, "error while reading by body", http.StatusBadRequest, err.Error())
		return
	}
	request.StartTime += "T" + defaultTime + timeZoneOffset
	request.EndTime += "T" + defaultTime + timeZoneOffset

	startTime, err := time.Parse(layout, request.GetStartTime())
	if err != nil {
		handleResponse(ctx, h.log, "start_time parse error", http.StatusBadRequest, err.Error())
		return
	}

	endTime, err := time.Parse(layout, request.GetEndTime())
	if err != nil {
		handleResponse(ctx, h.log, "end_time parse error", http.StatusBadRequest, err.Error())
		return
	}

	request.StartTime = startTime.Format(time.RFC3339)
	request.EndTime = endTime.Format(time.RFC3339)
	request.Id = id
	jsonData, err := json.Marshal(&request)
	if err != nil {
		handleResponse(ctx, h.log, "error marshalling request to JSON", http.StatusInternalServerError, err.Error())
		return
	}
	if err := h.rabbitMqProducer.ProduceMassage("budget_updated", jsonData); err != nil {
		handleResponse(ctx, h.log, "error sending message to RabbitMQ", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(ctx, h.log, "Budget request successfully updated", http.StatusCreated, "Budget request successfully updated")

}

// GetBudgetHandler      godoc
// @Router       /budget_service/v2/budget/{id} [GET]
// @Summary      Get budget by id
// @Description  Get budget by id
// @Tags         budget
// @Accept       json
// @Produce      json
// @Param        id path string true "budget_id"
// @Success      200  {object}  models.Budget
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
// @Security 	 ApiKeyAuth
func (h *Handler) GetBudgetHandler(ctx *gin.Context) {
	var (
		request  pb.PrimaryKey
		response *pb.Budget
		err      error
	)
	request.Id = ctx.Param("id")
	response, err = h.services.BudgetService().GetBudget(ctx, &request)
	if err != nil {
		handleResponse(ctx, h.log, "error is while get budget", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(ctx, h.log, "success", http.StatusCreated, response)

}

// DeleteBudgetHandler godoc
// @Router       /budget_service/v2/budget/{id} [DELETE]
// @Summary      Delete budget
// @Description  Delete budget
// @Tags         budget
// @Accept       json
// @Produce      json
// @Param        id path string true "budget_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
// @Security 	 ApiKeyAuth
func (h *Handler) DeleteBudgetHandler(ctx *gin.Context) {
	var (
		err     error
		request pb.PrimaryKey
	)
	request.Id = ctx.Param("id")
	_, err = h.services.BudgetService().DeleteBudget(ctx, &request)
	if err != nil {
		handleResponse(ctx, h.log, "error is while delete budget", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(ctx, h.log, "", http.StatusOK, "budget successfully deleted!")

}

// GetAllBudgetsHandler godoc
// @Router       /budget_service/v2/budgets [GET]
// @Summary      Get all budgets
// @Description  Get all budgets
// @Tags         budget
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      200  {object}  models.BudgetsResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
// @Security 	 ApiKeyAuth
func (h Handler) GetAllBudgetsHandler(c *gin.Context) {
	var (
		page, limit int
		search      string
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, h.log, "error is while converting page", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "100")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, h.log, "error is while converting limit", http.StatusBadRequest, err.Error())
		return
	}

	search = c.Query("search")

	response, err := h.services.BudgetService().GetListBudgets(context.Background(), &pb.GetListRequest{
		Page:   int32(page),
		Limit:  int32(limit),
		Search: search,
	})

	if err != nil {
		handleResponse(c, h.log, "error is while get all locations", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, response)
}
