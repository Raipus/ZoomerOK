BLOG_DIR := blog
ACCOUNT_DIR := account

all: test lint

verify:
	@echo "Проверка зависимостей для микросервиса blog..."
	$(MAKE) -C $(BLOG_DIR) verify
	@echo "Проверка зависимостей для микросервиса account..."
	$(MAKE) -C $(ACCOUNT_DIR) verify

test:
	@echo "Запуск тестов для микросервиса blog..."
	$(MAKE) -C $(BLOG_DIR) test
	@echo "Запуск тестов для микросервиса account..."
	$(MAKE) -C $(ACCOUNT_DIR) test

lint:
	@echo "Запуск линтера для микросервиса blog..."
	$(MAKE) -C $(BLOG_DIR) lint
	@echo "Запуск линтера для микросервиса account..."
	$(MAKE) -C $(ACCOUNT_DIR) lint