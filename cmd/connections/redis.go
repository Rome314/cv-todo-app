package connections

import "github.com/go-redis/redis"

func GetRedisConnection(address, password string, db int) (client *redis.Client, err error) {
	options := redis.Options{
		Addr:     address,  // use default Addr
		Password: password, // no password set
		DB:       db,       // use default DB
	}
	client = redis.NewClient(&options)
	_, err = client.Ping().Result()
	return
}
