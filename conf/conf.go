package conf

import (
	"github.com/joho/godotenv"
	"makespace-remaster/module"
	"os"
)

func Init() {
	godotenv.Load()
	module.Database(os.Getenv("MySQL_DSN"))
	module.CheckAdmin()
}
