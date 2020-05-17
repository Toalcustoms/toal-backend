import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type config struct {
	KeycloakRealm string
	DBConfig *dbconfig
}

type dbconfig struct {
	hostname string
	port     string
	username string
	password string
	name     string
}

func (d dbconfig) Info() string {
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		d.hostname, d.port, d.username,
		d.password, d.name)
}

func LoadDBConf() *dbconfig {
	return &dbconfig{
		getEnv("HOST", ""),
		getEnv("DB_PORT", ""),
		getEnv("POSTGRES_USER", ""),
		getEnv("POSTGRES_PASSWORD", ""),
		getEnv("POSTGRES_DB", ""),
	}
}

func loadConfig() *config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("No .env file found!\n")
	}

	dbconf := LoadDBConf()
	keycloakRealm := getEnv("KEYCLOAK_REALM", "")

	return &config{
		KeycloakRealm: keycloakRealm,
		DBConfig: dbconf,
	}
}

func getEnv(key, defaultvalue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultvalue
}