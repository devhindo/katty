package katty

import (

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		return//  log.Fatal("Error loading .env file")
	}
  
	// env := os.Getenv("KEY")
}