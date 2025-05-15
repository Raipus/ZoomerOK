#!/usr/bin/bash

set -e # Останавливаем скрипт при первой же ошибке

echo "Запускаем Backend..."

# Предполагается, что backend.sh находится в той же директории, что и entrypoint.sh
echo "Запускаем Backend..."
docker compose -f Backend/docker-compose.yml up -d --build
if [ $? -ne 0 ]; then
  echo "Ошибка при запуске Backend"
  exit 1
fi
echo "Backend запущен."

sleep 1

docker exec backend_kafka /opt/kafka/bin/kafka-topics.sh --create --bootstrap-server kafka:9092 --partitions 1 --replication-factor 1 --topic account_blog_topic && \
docker exec backend_kafka /opt/kafka/bin/kafka-topics.sh --create --bootstrap-server kafka:9092 --partitions 1 --replication-factor 1 --topic blog_account_topic

echo "Запускаем Frontend..."
docker compose -f Frontend/docker-compose.yml up -d --build
if [ $? -ne 0 ]; then
  echo "Ошибка при запуске Frontend"
  exit 1
fi
echo "Frontend запущен."

exec "$@"
