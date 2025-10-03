go run main.go
 Привет, World!

go run main.go --name Алексей --age 45
 Привет, Алексей! Тебе 45 лет.

go run main.go -n Алексей -a 45 -v
 Режим отладки включён
 Получено имя: Алексей
 Получен возраст: 45
 Привет, Алексей! Тебе 45 лет.

go run main.go --name Charlie extra-arg
 Привет, Charlie!
 Дополнительные аргументы: [extra-arg]