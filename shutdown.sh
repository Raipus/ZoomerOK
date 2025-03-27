   #!/bin/bash

   # Перебираем все каталоги первого уровня
   for dir in */; do
       # Убираем завершающий слеш из имени каталога
       dir="${dir%/}"

       # Проверяем, существует ли docker-compose.yml/docker-compose.yaml в каталоге
       if [ -f "$dir/docker-compose.yml" ] || [ -f "$dir/docker-compose.yaml" ]; then
           echo "Stopping docker-compose project in: $dir"
           cd "$dir" || exit 1
           docker-compose down
           cd ..
       fi
   done

   echo "All docker-compose projects stopped."
