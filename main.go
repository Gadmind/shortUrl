package main

import (
	"ShortUrl/route"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	route.API()
}
