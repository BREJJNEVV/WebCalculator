WEBCALCULATOR APPLICATION

Для запуска нужно:

Зпустить докер контейнер
CMD:
  docker pull postgres
  docker run --name postgres-container -e POSTGRES_PASSWORD=yourpassword -d -p 5432:5432 postgres

Запустить фронденд
CMD:
  npm install
  npm start
