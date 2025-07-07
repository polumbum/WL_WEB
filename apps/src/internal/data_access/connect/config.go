package connect

import (
	"encoding/json"
	"fmt"
	"os"
)

type DatabaseConfig struct {
	Host     string            `json:"host"`
	Port     int               `json:"port"`
	Users    map[string]string `json:"users"` // Словарь для хранения паролей по именам пользователей
	DBName   string            `json:"dbname"`
	Timezone string            `json:"timezone"`
}

type Config struct {
	Database DatabaseConfig `json:"database"`
}

func LoadConfig(filename string) (Config, error) {
	var config Config
	data, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func (dbConfig *DatabaseConfig) GetPostgresConnectionStr(user string) string {
	password, exists := dbConfig.Users[user] // Получаем пароль по имени пользователя
	if !exists {
		return fmt.Sprintf("Error: user '%s' not found", user)
	}

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable timezone=%s",
		dbConfig.Host, dbConfig.Port, user, password,
		dbConfig.DBName, dbConfig.Timezone)
}
