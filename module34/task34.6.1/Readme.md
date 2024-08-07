# Writer, Reader, Regexp

## Задание

Практическое задание поможет вам использовать полученные знания о работе с файлами и регулярными выражениями.

Суть задания — написать программу, которая считывает из файла список математических выражений, считает результат и записывает в другой файл.

Пример входного файла:
```
5+4=?
9+3=?
Сегодня прекрасная погода
13+7=?
4-2=?
```

Пример файла с выводом:
```
5+4=9
9+3=12
13+7=20
4-2=2
```

Пожелания к программе:

1. Использовать методы и структуры пакетов ioutils и regexp. 
2. Программа должна принимать на вход 2 аргумента: имя входного файла и имя файла для вывода результатов. 
3. Если не найден вывод, создать. 
4. Если файл вывода существует, очистить перед записью новых результатов. 
5. Использовать буферизированную запись результатов.

## Решение
Первое решение с открытием на чтение входного файла с использованием метода `os.OpenFile` находится по этой ссылке:

[Решение 1 находится здесь](https://github.com/MoJIoToK/learning_go/tree/master/module34/task34.6.1)

Второе решение с открытием на чтение входного файла с использованием метода `ioutil.ReadFile` находится по этой ссылке:

[Решение 2 находится здесь](https://github.com/MoJIoToK/learning_go/blob/master/module34/task34.6.1/version2.go)

Насколько мне известно пакет `ioutil` - устарел. Поэтому изначально мне показалось нецелесообразно делать решение с ним.
Но поскольку всё-таки в задании сказано использовать этот пакет, я сделал вторую версию кода с данным пакетом. 
Однако, данный пакет используется, во второй версии решения, исключительно для чтения. Т.к. в пп.5 требований указано, 
что запись должна быть буферизирована. Поэтому для записи я ничего не менял.

По сути в этих двух вариантах меняется две три строчки.

## Запуск приложения

Для запуска приложения необходимо создать файл с входными данными и в консоли ввести команду:
`go run main.go`

Далее следовать инструкциям в консоли.