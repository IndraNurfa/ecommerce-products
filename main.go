package main

import (
	"ecommerce-products/cmd"
	"ecommerce-products/helpers"
)

func main() {

	// load config
	helpers.SetupConfig()

	// load log
	helpers.SetupLogger()

	// load db
	helpers.SetupPostgreSQL()

	// load redis
	helpers.SetupRedis()

	// load kafka
	// cmd.ServeKafka()

	// run http
	cmd.ServeHTTP()
}
