package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config — структура всей конфигурации
type Config struct {
	App struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"app"`
	DB struct {
		Name string `mapstructure:"name"`
	} `mapstructure:"db"`
}

var (
	cfgFile string
	config  Config
)

var rootCmd = &cobra.Command{
	Use:   "unified-config-demo",
	Short: "Demo: unified config from flags, env, and file",
	Run: func(cmd *cobra.Command, args []string) {
		// Применяем приоритет: Viper сам обрабатывает это
		if err := viper.Unmarshal(&config); err != nil {
			fmt.Printf("Unable to decode config: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("✅ Final configuration:")
		fmt.Printf("  App Host: %s\n", config.App.Host)
		fmt.Printf("  App Port: %d\n", config.App.Port)
		fmt.Printf("  DB Name:  %s\n", config.DB.Name)
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	// Флаги командной строки
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config.yaml)")
	rootCmd.Flags().String("app.host", "localhost", "Application host")
	rootCmd.Flags().Int("app.port", 8080, "Application port")
	rootCmd.Flags().String("db.name", "default_db", "Database name")

	// Привязка флагов к Viper
	viper.BindPFlag("app.host", rootCmd.Flags().Lookup("app.host"))
	viper.BindPFlag("app.port", rootCmd.Flags().Lookup("app.port"))
	viper.BindPFlag("db.name", rootCmd.Flags().Lookup("db.name"))

	// Установка префикса для переменных окружения
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv() // читает все переменные окружения с префиксом APP_
}

func initConfig() {
	// Загружаем .env (если есть)
	_ = godotenv.Load() // игнорируем ошибку, если файла нет

	// Настройка Viper
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	// Загружаем конфиг-файл (если есть)
	_ = viper.ReadInConfig() // игнорируем ошибку, если файла нет

	// Устанавливаем значения по умолчанию (низший приоритет)
	viper.SetDefault("app.host", "localhost")
	viper.SetDefault("app.port", 8080)
	viper.SetDefault("db.name", "default_db")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
