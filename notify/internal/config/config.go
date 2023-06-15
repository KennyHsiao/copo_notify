package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	service.ServiceConf
	Mysql struct {
		Host       string
		Port       int
		DBName     string
		UserName   string
		Password   string
		DebugLevel string
	}
	RedisCache struct {
		RedisSentinelNode string
		RedisMasterName   string
		RedisDB           int
	}
	Notify struct {
		Telegram struct {
			BotToken     string
			NotifyChatId int
		}
		Slack struct {
			BotToken string
		}
		Line struct {
			ClientID     string
			ClientSecret string
			CallbackURL  string
		}
	}
	ApiKey struct {
		PublicKey string
		PayKey    string
		ProxyKey  string
		LineKey   string
	}
}
