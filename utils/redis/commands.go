package redis

func Set(key string, value interface{}) (err error) {
	return Cli.Set(ctx, key, value, 0).Err()
}

func Get(key string) (value string, err error) {
	return Cli.Get(ctx, key).Result()
}

func Del(key string) (err error) {
	return Cli.Del(ctx, key).Err()
}
