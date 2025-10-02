-**Флаги командной строки** (через `cobra`)
-**Переменные окружения**
-**Файл конфигурации** (YAML, но легко переключить на JSON/TOML)
-**Приоритет**: флаги > переменные окружения > файл конфигурации > значения по умолчанию

Используется **Viper + Cobra** — стандарт де-факто для сложных Go-приложений.

---

## `.env` (опционально)

```env
APP_PORT=7777
APP_HOST=from-env.local
DB_NAME=env_db
```

---

## `config.yaml`

```yaml
app:
  host: from-config.local
  port: 6666
db:
  name: config_db
```

---

## Приоритет источников конфигурации

Viper применяет **следующий порядок приоритета** (от самого высокого к низкому):

1. **Флаги командной строки** (`--app.port 9999`)
2. **Переменные окружения** (`APP_APP_PORT=8888`)
3. **Файл конфигурации** (`config.yaml`)
4. **Значения по умолчанию** (`SetDefault`)

> Обратите внимание: переменные окружения должны использовать **префикс `APP_`**, а вложенные ключи — **нижнее подчёркивание**:
>
> - `app.host` → `APP_APP_HOST`
> - `db.name` → `APP_DB_NAME`

---

## Как запустить и проверить

### 1. Базовый запуск (только значения по умолчанию)
```bash
go run main.go
```
**Вывод:**
```
✅ Final configuration:
  App Host: localhost
  App Port: 8080
  DB Name:  default_db
```

---

### 2. С файлом конфигурации
```bash
go run main.go
```
(при наличии `config.yaml`)
**Вывод:**
```
App Host: from-config.local
App Port: 6666
DB Name:  config_db
```

---

### 3. С `.env`
```bash
go run main.go
```
(при наличии `.env`, даже без `config.yaml`)
**Вывод:**
```
App Host: from-env.local
App Port: 7777
DB Name:  env_db
```

> `.env` переопределяет `config.yaml`, потому что `godotenv.Load()` вызывается до `viper.ReadInConfig()`, и переменные окружения имеют более высокий приоритет.

---

### 4. С флагами (самый высокий приоритет)
```bash
go run main.go --app.port 9999 --db.name flag_db
```
**Вывод:**
```
✅ Final configuration:
App Host: localhost        # из .env или config.yaml или default
App Port: 9999             # из флага!
DB Name:  flag_db          # из флага!
```

---

### 5. Комбинированный пример
```bash
APP_APP_HOST=cli-env.local go run main.go --app.port 1234
```
**Вывод:**
```
✅ Final configuration:
App Host: cli-env.local    # из env
App Port: 1234             # из флага (выше env!)
DB Name:  env_db           # из .env
```

---

##Преимущества этого подхода

- Единая точка управления конфигурацией.
- Гибкий приоритет источников.
- Поддержка `.env`, что удобно для разработки.
- CLI-флаги для быстрого переопределения.
- Лёгкое тестирование разных окружений.

---

## Очистка (если нужно)

Удалите файлы, чтобы проверить поведение по умолчанию:
```bash
rm -f config.yaml .env
```
