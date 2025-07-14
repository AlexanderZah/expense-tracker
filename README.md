# Expense Tracker

приложение для учета расходов, написанное на Go.

## Описание

Это приложение позволяет:
- Добавлять новые расходы
- Просматривать историю расходов
- Хранить данные в JSON-файле

## Установка

1. Убедитесь, что у вас установлен Go (версии 1.16 или выше)
2. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/AleksaAlekszZakharov/expense-tracker.git
   ```
3. Перейдите в директорию проекта
    ```
    cd expense-tracker
    ```
4. Сборка
    ```
    go build
    ```
5. Запуск
    ```
    ./expense-tracker
    ```
## Доступные команды
```
добавить новый расход
.\expense-tracker add --amount ? --description ? 
показать все расходы
.\expense-tracker list
Удалить по id
.\expense-tracker delete --id ? 
Показать сумму расходов
.\expense-tracker summary ? 
```