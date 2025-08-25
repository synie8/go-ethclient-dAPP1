package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// 配置结构体
type Config struct {
	SepoliaURL      string
	MainnetURL      string
	PrivateKey      string
	AccountAddress1 string
	AccountAddress2 string
	ContractAddress string
	GasLimit        uint64
	GasPrice        int64
	ChainID         int64
	AlchemyAPIKey   string
	InfuraAPIKey    string
}

// 全局配置实例
var AppConfig *Config

// InitConfig 初始化配置
func InitConfig() {
	// 加载 .env 文件
	err := godotenv.Load("./.env")
	if err != nil {
		log.Printf("Warning: .env file not found, using system environment variables")
	}

	AppConfig = &Config{
		SepoliaURL:      getEnv("SEPOLIA_URL", "https://eth-sepolia.g.alchemy.com/v2/default"),
		MainnetURL:      getEnv("MAINNET_URL", ""),
		PrivateKey:      getEnv("PRIVATE_KEY", ""),
		AccountAddress1: getEnv("ACCOUNT_ADDRESS1", ""),
		AccountAddress2: getEnv("ACCOUNT_ADDRESS2", ""),
		ContractAddress: getEnv("CONTRACT_ADDRESS", ""),
		GasLimit:        getEnvAsUint64("GAS_LIMIT", 21000),
		GasPrice:        getEnvAsInt64("GAS_PRICE", 20),
		ChainID:         getEnvAsInt64("CHAIN_ID", 11155111),
		AlchemyAPIKey:   getEnv("ALCHEMY_API_KEY", ""),
		InfuraAPIKey:    getEnv("INFURA_API_KEY", ""),
	}
}

// getEnv 获取环境变量，如果不存在则使用默认值
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt64 获取环境变量并转换为 int64
func getEnvAsInt64(key string, defaultValue int64) int64 {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	value := int64(0)
	_, err := fmt.Sscanf(valueStr, "%d", &value)
	if err != nil {
		log.Printf("Error parsing %s as int64: %v", key, err)
		return defaultValue
	}
	return value
}

// getEnvAsUint64 获取环境变量并转换为 uint64
func getEnvAsUint64(key string, defaultValue uint64) uint64 {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	value := uint64(0)
	_, err := fmt.Sscanf(valueStr, "%d", &value)
	if err != nil {
		log.Printf("Error parsing %s as uint64: %v", key, err)
		return defaultValue
	}
	return value
}

// GetConfig 获取配置实例
func GetConfig() *Config {
	if AppConfig == nil {
		InitConfig()
	}
	return AppConfig
}
