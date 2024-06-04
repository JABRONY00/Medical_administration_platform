package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

/*func GetEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Panicf("Error loading .env file: %v", err)
	}
	return os.Getenv(key)
}*/

/*func CheckRequiredEnvs() {
	requiredEnvVars := []string{"SERVER_PORT", "LOG_LEVEL"}
	for _, envVar := range requiredEnvVars {
		if value, exists := os.LookupEnv(envVar); !exists || value == "" {
			log.Panic(fmt.Sprintf("Error: Environment variable %v is not set.", envVar))
		}
	}
}*/

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func CheckPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
