package env

import "github.com/omaressameldin/water-bot/internal/utils"

func GetToken() (string, error) {
	return utils.GetEnv(TOKEN_KEY)
}
