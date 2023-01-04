package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis/v9"
)

var RedisURL = os.Getenv("REDIS_URL")

func RedisPing(w http.ResponseWriter, req *http.Request) {

	var ctx = context.Background()

	fmt.Fprintf(w, "Using Redis addr %v\n", RedisURL)

	rdb := redis.NewClient(&redis.Options{
		Addr:     RedisURL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer rdb.Close()

	res := rdb.Ping(ctx)
	succ, err := res.Result()
	if err != nil {
		fmt.Fprintf(w, "Error pinging redis %v\n", err.Error())
	} else {

		fmt.Fprintf(w, "Pinged Redis successfully %v\n", succ)
	}

	err = rdb.Set(ctx, "key", "value niklas", 0).Err()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintln(w, "key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Fprintln(w, "key2 does not exist")
	} else if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintln(w, "key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Yeay!\n")
	})
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/redis", RedisPing)

	http.ListenAndServe(":8888", nil)
}
