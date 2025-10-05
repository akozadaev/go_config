## Unified Configuration Demo

Демонстрация унифицированной загрузки конфигурации из трёх источников с автоматическим приоритетом:

1. Флаги командной строки — самый высокий приоритет
2. Переменные окружения (включая файл `.env`)
3. Файл конфигурации (`config.yaml`)
4. Значения по умолчанию — самый низкий приоритет

Реализовано с использованием:
- github.com/spf13/viper — управление конфигурацией
- github.com/joho/godotenv — загрузка переменных из `.env`
- github.com/spf13/cobra — создание CLI-интерфейса

---

## Формат конфигурации

### Структура
```yaml
app:
  host: string
  port: int
db:
  name: string
```

### Пример config.yaml
```yaml
app:
  host: from-config.local
  port: 6666
db:
  name: config_db
```

---

## Переменные окружения

Viper настроен следующим образом:
- Префикс: `APP_`
- Точки в ключах заменяются на подчёркивания: `app.host` → `APP_APP_HOST`

### Требуемый формат .env
```env
APP_APP_HOST=from-env.local
APP_APP_PORT=7777
APP_DB_NAME=env_db
```

Примечание: двойное `APP_` возникает потому, что префикс `APP_` применяется ко всему ключу, а `app.host` преобразуется в `APP_HOST`, что в совокупности даёт `APP_APP_HOST`.

---

## Запуск

### Сборка
```bash
go build -o unified-config-demo .
```

### Примеры запуска

Без дополнительных файлов (используются значения по умолчанию):
```bash
./unified-config-demo
```

С файлом `.env` в рабочей директории:
```bash
./unified-config-demo
```

С указанием конфигурационного файла:
```bash
./unified-config-demo --config config.yaml
```

С переопределением через флаги:
```bash
./unified-config-demo --app.host=my-cli-host --app.port=5555 --db.name=cli_db
```

Комбинированный запуск:
```bash
./unified-config-demo --config config.yaml --db.name=override
```

---

## Приоритет источников (от высшего к низшему)

1. Флаги командной строки
2. Переменные окружения
3. Файл config.yaml
4. Значения по умолчанию

Viper автоматически применяет этот порядок — дополнительная логика не требуется.

---

## Отладка

При запуске программа выводит диагностическую информацию:
- Удалось ли загрузить `.env`
- Используется ли файл конфигурации и его путь
- Текущие значения ключей из Viper и соответствующие переменные окружения

Пример вывода:
```
Loaded .env file
Using config file: /path/to/config.yaml
Viper configuration debug:
  app.host = from-env.local (env: APP_APP_HOST = "from-env.local")
  app.port = 7777 (env: APP_APP_PORT = "7777")
  db.name = env_db (env: APP_DB_NAME = "env_db")

Final configuration:
  App Host: from-env.local
  App Port: 7777
  DB Name:  env_db
```

---

## Требования

- Go 1.16 или новее
- Опционально: файлы `.env` и `config.yaml` в рабочей директории

---

## Пример полного сценария

```bash
# Создание .env
cat > .env <<EOF
APP_APP_HOST=my-env-host
APP_APP_PORT=9000
APP_DB_NAME=env_database
EOF

# Создание config.yaml
cat > config.yaml <<EOF
app:
  host: fallback-from-file
  port: 8080
db:
  name: file_database
EOF

# Запуск
./unified-config-demo
```

Результат: значения из `.env` будут использованы, так как переменные окружения имеют более высокий приоритет, чем файл конфигурации.

--- 

Теперь README полностью текстовый, без эмодзи, символов-декораций и изображений.