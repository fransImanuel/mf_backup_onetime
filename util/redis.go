package util

import (
	"errors"
	"mf_backup_onetime/dto"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

type Redis struct {
	Config *dto.RedisConfig
	Client *redis.Client
}

func InitRedis(config *dto.RedisConfig) *Redis {
	return &Redis{
		Config: config,
	}
}

func (r *Redis) New() (err error) {
	log.Info("Redis - New() - starting...")

	if r.Client != nil {
		_ = r.Client.Close()
	}

	client := redis.NewClient(&redis.Options{
		Addr:     r.Config.Host,
		Password: r.Config.Password,
		DB:       r.Config.DB,
	})

	result, err := client.Ping().Result()
	if err != nil {
		return err
	}
	log.Info("Redis - New() - result Ping: %v", result)

	r.Client = client
	log.Info("Redis - New() - finished.")
	return nil
}

func (r *Redis) Get(key string) (res string, err error) {
	if key == "" {
		return res, errors.New("key required")
	}

	res, err = r.Client.Get(key).Result()
	if err != nil {
		return res, err
	}

	return res, nil
}
