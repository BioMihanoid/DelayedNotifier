package api

import (
	"BioMihanoid/DelayedNotifier/internal/models"
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Service interface {
	CreateNotify(ctx context.Context, notify models.Notification) error
	GetNotifyStatus(ctx context.Context, id uuid.UUID) (string, error)
	DeleteNotify(ctx context.Context, id uuid.UUID) error
}

type Handler struct {
	service Service
}

type notifyRequest struct {
	Message string    `json:"name"`
	SendAt  time.Time `json:"send_at"`
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/notify", h.createNotify)
	r.GET("/notify/:id", h.getStatusNotify)
	r.DELETE("/notify/:id", h.cancelNotify)

	return r
}

func (h *Handler) createNotify(c *gin.Context) {
	input := notifyRequest{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("error parsing request"))
		return
	}

	notify := models.Notification{
		ID:      uuid.New(),
		Message: input.Message,
		SendAt:  input.SendAt,
		Status:  models.PendingStatus,
	}

	err := h.service.CreateNotify(c, notify)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("error create notify"))
		return
	}

	c.JSON(http.StatusOK, notify.ID)
}

func (h *Handler) getStatusNotify(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, errors.New("id is requered"))
		return
	}

	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("incorrected id"))
		return
	}

	status, err := h.service.GetNotifyStatus(c, parsedUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("error get notify status"))
		return
	}

	c.JSON(http.StatusOK, status)
}

func (h *Handler) cancelNotify(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, errors.New("id is requered"))
		return
	}

	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("incorrected id"))
		return
	}

	err = h.service.DeleteNotify(c, parsedUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("error delete notify"))
		return
	}

	c.JSON(http.StatusOK, "ok")
}
