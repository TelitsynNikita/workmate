package handler

import (
	"fmt"
	"sync/atomic"
	"workmate/internal/service"

	"github.com/gofiber/fiber/v3"
)

var IsShutDown atomic.Bool

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) InitRoutes() *fiber.App {
	router := fiber.New()

	// Мидлвара, которая нужна для предупреждения пользователей, что сервер находится в аварийном состоянии
	router.Use(func(c fiber.Ctx) error {
		if IsShutDown.Load() {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"message": "service is shutting down",
			})
		}

		ctx := c.RequestCtx()

		fmt.Println(ctx)
		return c.Next()
	})

	message := router.Group("/link")
	message.Post("/check_by_urls", h.CheckLinksStatusByUrl)
	message.Post("/check_by_id", h.CheckLinksStatusByID)

	return router
}
