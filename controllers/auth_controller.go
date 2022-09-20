package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/mapsapi/models"
	"github.com/nikola43/mapsapi/services"
	"github.com/nikola43/mapsapi/utils"
)

func Login(context *fiber.Ctx) error {
	loginRequest := new(models.LoginClientRequest)

	err := context.BodyParser(loginRequest)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	clientLoginResponse, err := services.LoginClient(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusNotFound, err, context)
	}

	return context.JSON(clientLoginResponse)
}
