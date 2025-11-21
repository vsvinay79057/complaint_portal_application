package controller

import (
	"complaint_portal/models"
	"complaint_portal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ComplaintController struct {
	compUC usecase.ComplaintUsecase
	userUC usecase.UserUsecase
}

func NewComplaintController(c usecase.ComplaintUsecase, u usecase.UserUsecase) *ComplaintController {
	return &ComplaintController{compUC: c, userUC: u}
}

func (c *ComplaintController) SubmitComplaint(ctx echo.Context) error {
	userSecret := ctx.Request().Header.Get("X-SECRET-CODE")
	var req models.ComplaintRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid payload"})
	}
	if req.Title == "" || req.Summary == "" || req.Rating < 1 || req.Rating > 5 {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid fields"})
	}
	cp, err := c.compUC.Submit(userSecret, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusCreated, cp)
}

func (c *ComplaintController) GetAllComplaintsForUser(ctx echo.Context) error {
	userSecret := ctx.Request().Header.Get("X-SECRET-CODE")
	list, err := c.compUC.GetAllForUser(userSecret)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	resp := make([]map[string]interface{}, 0, len(list))
	for _, it := range list {
		resp = append(resp, map[string]interface{}{
			"id":         it.ID,
			"title":      it.Title,
			"summary":    it.Summary,
			"rating":     it.Rating,
			"resolved":   it.Resolved,
			"created_at": it.CreatedAt,
		})
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (c *ComplaintController) GetAllComplaintsForAdmin(ctx echo.Context) error {
	list, err := c.compUC.GetAllForAdmin()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	resp := make([]map[string]interface{}, 0, len(list))
	for _, it := range list {
		resp = append(resp, map[string]interface{}{
			"id":         it.ID,
			"title":      it.Title,
			"summary":    it.Summary,
			"rating":     it.Rating,
			"user_name":  it.User.Name,
			"user_email": it.User.Email,
			"resolved":   it.Resolved,
		})
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (c *ComplaintController) ViewComplaint(ctx echo.Context) error {
	userSecret := ctx.Request().Header.Get("X-SECRET-CODE")
	id := ctx.Param("id")
	cp, err := c.compUC.ViewComplaint(userSecret, id)
	if err != nil {
		return ctx.JSON(http.StatusForbidden, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, cp)
}

func (c *ComplaintController) ResolveComplaint(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "id required"})
	}
	if err := c.compUC.ResolveComplaint(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "complaint resolved"})
}
