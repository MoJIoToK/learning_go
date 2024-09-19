# Module 39

# Методы доставки приложений

_Continuous Integration_ (_CI_) - непрерывная интеграция

## Подготовка окружения

### Проект

[Проект для данного модуля](https://github.com/SkillfactoryCoding/go-39-metoddostavkipril)

Для использования проекта необходимо клонировать его на локальную машину и перепривязать в собственномоу репозиторию:
```
git clone https://github.com/SkillfactoryCoding/go-39-metoddostavkipril
cd go-39-metoddostavkipril
rm -rf .git/
git init
git remote add origin git@github.com:<YOUR_ACCOUNT_NAME>/go-39-metoddostavkipril.git
```

### Аккаунт в хранилище образов

Хранилище _docker_ образов называют _registry_ (произносится «реджистри»). Наибольшая публичная библиотека образов это _docker.io_

Для публикации своих собственных образов в реджистри на локальной машине необходимо выполнить вход в свой профиль: `docker login docker.io`

### Последние приготовления

В клонированном проекте в файле _Makefile_ необходимо заменить значение переменной `PUBLIC_REGISTRY_OWNER` на имя пользователя в реджистри.

А затем тоже самое необходимо сделать для `variables` -> `PUBLIC_REGISTRY_OWNER` в файле `.gitlab-ci.yml`


## Литература
[Проект](https://github.com/SkillfactoryCoding/go-39-metoddostavkipril?tab=readme-ov-file)
[Многоступенчатая сборка](https://docs.docker.com/build/building/multi-stage/)