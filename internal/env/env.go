package env

import "github.com/omaressameldin/water-bot/internal/utils"

func GetToken() (string, error) {
	return utils.GetEnv(TOKEN_KEY)
}

func GetActionPort() (string, error) {
	return utils.GetEnv(ACTIONS_PORT_KEY)
}
