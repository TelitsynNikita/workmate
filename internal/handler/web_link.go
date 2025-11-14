package handler

import (
	"workmate/internal/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	jsoniter "github.com/json-iterator/go"
)

func (h *Handler) CheckLinksStatusByUrl(c fiber.Ctx) error {
	var body model.CheckLinksStatusByUrlRequest

	err := jsoniter.Unmarshal(c.Body(), &body)
	if err != nil {
		return fiber.ErrBadRequest
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err = validate.Struct(&body); err != nil {
		return fiber.ErrBadRequest
	}

	id, err := h.Service.URLService.CheckLinksStatusByUrl(body.Links)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"url": id,
	})
}

func (h *Handler) CheckLinksStatusByID(c fiber.Ctx) error {
	var body model.CheckLinksStatusByIDRequest

	err := jsoniter.Unmarshal(c.Body(), &body)
	if err != nil {
		return fiber.ErrBadRequest
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err = validate.Struct(&body); err != nil {
		return fiber.ErrBadRequest
	}

	id, err := h.Service.URLService.GetUrlByID(body.LinksList)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"url": id,
	})
}
