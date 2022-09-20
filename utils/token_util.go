package utils

import (
	"errors"
	"fmt"
	"math"
	"time"
	"unicode"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/mapsapi/models"
)

func GenerateClientToken(email string, clientId uint) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = clientId
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	fmt.Println(claims["id"])
	fmt.Println(claims["email"])
	fmt.Println(claims["exp"])

	// Generate encoded token and send it as response.
	tokenString, err := token.SignedString([]byte(GetEnvVariable("JWT_CLIENT_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetClientTokenClaims(context *fiber.Ctx) (*models.ClientTokenClaims, error) {
	user := context.Locals("user").(*jwt.Token)
	if claims, ok := user.Claims.(jwt.MapClaims); ok && user.Valid {
		clientTokenClaims := &models.ClientTokenClaims{}

		if claims["id"] != nil {
			clientTokenClaims.ID = uint(math.Round(claims["id"].(float64)))
		}

		if claims["email"] != nil {
			clientTokenClaims.Email = claims["email"].(string)
		}

		if claims["exp"] != nil {
			clientTokenClaims.Exp = uint(math.Round(claims["exp"].(float64)))
		}

		return clientTokenClaims, nil
	} else {
		return nil, errors.New("invalid claims")
	}
}

func IsMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
