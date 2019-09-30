package env

import "github.com/omaressameldin/water-bot/internal/utils"

func GetToken() (string, error) {
	return utils.GetEnv(TOKEN_KEY)
}

func GetActionPort() (string, error) {
	return utils.GetEnv(ACTIONS_PORT_KEY)
}

func GetPostCode() (string, error) {
	return utils.GetEnv(POSTCODE)
}

func GetStillWaterLink() (string, error) {
	return utils.GetEnv(STILL_WATER_LINK)
}

func GetSparklingWaterLink() (string, error) {
	return utils.GetEnv(SPARKLING_WATER_LINK)
}
