Данное приложение является менеджером задач, в котором можно создавать проекты и "прикреплять" к проекту задачи.

Запуск приложения:
docker-compose up
make test-migration-up
для заполнения таблицы тестовыми данными выполнить скрипт create_test_data.sql, находящийся в internal/db/scripts

Кодогенерация клиента:
make generate
