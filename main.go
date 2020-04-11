package main

import (
	"balabanovds/go-social/internal/data"
	"balabanovds/go-social/internal/web"
)

func main() {
	data.Init()
	web.Start()
}
