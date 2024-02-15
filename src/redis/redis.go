package main

import "github.com/redis/go-redis/v9"


r_lcl := redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

ctx := context.Background()

value, err := r_lcl.Get()