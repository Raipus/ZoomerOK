#!/usr/bin/bash

set -e # Останавливаем скрипт при первой же ошибке

# Проверяем, существует ли сеть gateway_network
if ! docker network inspect gateway_network > /dev/null 2>&1; then
  echo "Создаем сеть gateway_network..."
  docker network create gateway_network
  if [ $? -ne 0 ]; then
    echo "Ошибка при создании сети gateway_network"
    exit 1
  fi
  echo "Сеть gateway_network создана."
else
  echo "Сеть gateway_network уже существует."
fi

# Проверяем, существует ли сеть kafka_network
if ! docker network inspect kafka_network > /dev/null 2>&1; then
  echo "Создаем сеть kafka_network..."
  docker network create kafka_network
  if [ $? -ne 0 ]; then
    echo "Ошибка при создании сети kafka_network"
    exit 1
  fi
  echo "Сеть kafka_network создана."
else
  echo "Сеть kafka_network уже существует."
fi

# Проверяем, существует ли volume frontend_static
if ! docker volume inspect frontend_static > /dev/null 2>&1; then
  echo "Создаем volume frontend_static..."
  docker volume create frontend_static
  if [ $? -ne 0 ]; then
    echo "Ошибка при создании сети frontend_static"
    exit 1
  fi
  echo "volume frontend_static создана."
else
  echo "volume frontend_static уже существует."
fi

echo "Запускаем Backend/account..."
docker compose -f Backend/account/docker-compose.yml up -d --build
if [ $? -ne 0 ]; then
  echo "Ошибка при запуске Backend/account"
  exit 1
fi
echo "Backend/account запущен."

echo "Запускаем Backend/blog..."
docker compose -f Backend/blog/docker-compose.yml up -d --build
if [ $? -ne 0 ]; then
  echo "Ошибка при запуске Backend/blog"
  exit 1
fi
echo "Backend/blog запущен."

echo "Запускаем Backend/gateway..."
docker compose -f Backend/gateway/docker-compose.yml up -d --build
if [ $? -ne 0 ]; then
  echo "Ошибка при запуске Backend/gateway"
  exit 1
fi
echo "Backend/gateway запущен."

echo "Запускаем Backend/kafka..."
docker compose -f Backend/docker-compose.kafka.yml up -d --build
if [ $? -ne 0 ]; then
  echo "Ошибка при запуске Backend/kafka"
  exit 1
fi
echo "Backend/kafka запущен."

echo "Запускаем Frontend..."
docker compose -f Frontend/docker-compose.yml up -d --build
if [ $? -ne 0 ]; then
  echo "Ошибка при запуске Frontend"
  exit 1
fi
echo "Frontend запущен."

exec "$@"
