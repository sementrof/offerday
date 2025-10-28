#!/bin/bash

echo "🛑 Остановка Event Manager..."

# Останавливаем все сервисы
docker compose down

echo "✅ Event Manager остановлен!"
echo ""
echo "📝 Для полной очистки выполните:"
echo "  docker compose down -v    # Удаление томов (данные будут потеряны)"
echo "  docker compose down --rmi all    # Удаление образов"