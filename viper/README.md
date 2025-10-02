go get github.com/spf13/viper github.com/spf13/cobra
go run main.go                # использует config.yaml → порт 8081
go run main_viper.go --port 9999    # переопределяет флагом → порт 9999