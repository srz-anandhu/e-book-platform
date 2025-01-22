package main

import (
	"ebook/cmd"

	_ "github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	cmd.Execute()
}
