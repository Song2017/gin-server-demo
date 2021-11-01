package apiserver

import (
	model "apiserver/v1/nomad-model"
	"os"
	"strconv"
	"sync"
)

var (
	mu_app sync.Mutex
)
var app_config *model.AppSettings

func AppConfig() *model.AppSettings {
	if app_config == nil {
		mu_app.Lock()
		defer mu_app.Unlock()
		if app_config == nil {
			redis_time_out, err := strconv.Atoi(os.Getenv("APP_TIMEOUT"))
			if err != nil {
				redis_time_out = 5
			}
			app_config = &model.AppSettings{
				AppTimeout:     redis_time_out,
				SecurityCaKey:  os.Getenv("SECURITY_CA_KEY"),
				GinMode:        os.Getenv("GIN_MODE"),
				RedisHost:      os.Getenv("REDIS_SERVER_HOST"),
				RedisPass:      os.Getenv("REDIS_SERVER_PASSWORD"),
				RedisHostPilot: os.Getenv("REDIS_SERVER_HOST_PILOT"),
				RedisPassPilot: os.Getenv("REDIS_SERVER_PASSWORD_PILOT"),
			}
		}
	}
	return app_config
}
