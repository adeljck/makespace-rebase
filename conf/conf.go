package conf

import (
	"github.com/joho/godotenv"
	"makespace-remaster/module"
	"os"
)

func Init() {
	godotenv.Load()
	RedisInit()
	MongoInit()
	module.Database(os.Getenv("MYSQL_DSN"))
	module.CheckAdmin()
}
