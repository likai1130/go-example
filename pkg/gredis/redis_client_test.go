package gredis

import "testing"

func ClientTest(t *testing.T) {
	client, err := NewRedisClient()
	if err != nil {
		panic(err)
	}

	err = client.Set("app", "go-example", 0).Err()
	if err != nil {
		panic(err)
	}

	err = client.Get("app").Err()
	if err != nil {
		panic(err)
	}
}
