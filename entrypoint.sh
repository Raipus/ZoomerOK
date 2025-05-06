#!/usr/bin/bash

set -e # Останавливаем скрипт при первой же ошибке

echo "Запускаем Backend..."

# Предполагается, что backend.sh находится в той же директории, что и entrypoint.sh
# Если нет, укажите правильный путь. Например: ../backend.sh
cd Backend
./backend.sh
if [ $? -ne 0 ]; then
  echo "Ошибка при запуске Backend"
  exit 1
fi
echo "Backend запущен."

cd ..
echo "Запускаем Frontend..."
docker compose -f Frontend/docker-compose.yml up -d --build
if [ $? -ne 0 ]; then
  echo "Ошибка при запуске Frontend"
  exit 1
fi
echo "Frontend запущен."

exec "$@"
