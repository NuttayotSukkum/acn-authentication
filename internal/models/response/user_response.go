package response

import "github.com/NuttayotSukkum/acn/acn-authentication/internal/models"

type ResponseUser struct {
	HTTPStatus string
	Time       string
	Data       *models.User
}
