### Установка
go mod init my-cli-app
go get github.com/spf13/cobra@latest

### Сборка
go build -o my-cli .

### Примеры использования:
./my-cli
# Welcome to my-cli! Use --help to see available commands.

./my-cli --help
# Покажет справку по всем командам

./my-cli greet --name Alice --times 3

 Hello, Alice!

 Hello, Alice!

 Hello, Alice!

./my-cli version
# my-cli v1.0.0