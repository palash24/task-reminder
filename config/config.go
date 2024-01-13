package config

type Config struct {
	ServerAddr string
	ServerPort string
	DBAddr     string
	DBPort     string
	DbUser     string
	DBPassword string
	DBName     string
}

// Returning struct simply, as I dont want the calling function to change the values
func NewConfig() Config {
	return Config{
		ServerAddr: "localhost",
		ServerPort: "8080",
		DBAddr:     "localhost",
		DBPort:     "2553",
		DbUser:     "postgres",
		DBPassword: "postgres",
		DBName:     "taskreminder",
	}
}
