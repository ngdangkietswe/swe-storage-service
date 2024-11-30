package configs

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	GlobalConfig = &Configuration{}
)

func init() {
	env := os.Getenv("K8S_ENV")
	if env == "prod" {
		viper.AutomaticEnv()
	} else {
		viper.AddConfigPath("./configs")
		viper.SetConfigName("config")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Can't read config file: %v", err)
			return
		}
	}

	err := viper.Unmarshal(&GlobalConfig)
	if err != nil {
		log.Fatalf("Can't unmarshal config: %v", err)
	}
}

type Configuration struct {
	AppName            string `mapstructure:"APP_NAME"`
	GrpcPort           int    `mapstructure:"GRPC_PORT"`
	MinIOHost          string `mapstructure:"MINIO_HOST"`
	MinIOPort          int    `mapstructure:"MINIO_PORT"`
	MinIOAccessKey     string `mapstructure:"MINIO_ACCESS_KEY"`
	MinIOSecretKey     string `mapstructure:"MINIO_SECRET_KEY"`
	MinIOBucketPublic  string `mapstructure:"MINIO_BUCKET_PUBLIC"`
	MinIOBucketPrivate string `mapstructure:"MINIO_BUCKET_PRIVATE"`
	MinIOUseSSL        bool   `mapstructure:"MINIO_USE_SSL"`
	MinIODefaultExpiry int    `mapstructure:"MINIO_DEFAULT_EXPIRY"`
	JwtSecret          string `mapstructure:"JWT_SECRET"`
}
