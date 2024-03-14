# Homework 19

## Task

### Task 19.2.1

Попробуйте изменить приведённый выше код программы так, чтобы завершение работы потребителей, работающих с каналами,
полученными из функции демультиплексирования, происходило уже не в цикле вот так:

```go
  // После завершения посылки сообщений в основной 
        // канал-источник 
        // данных
        // закрываем все каналы-потребители
        for _, c := range output {
            close(c)
        }
```
А осуществлялось одной командой.

## Solution

### Solution 19.2.1

[Решение находится здесь!](https://github.com/MoJIoToK/learning_go/blob/master/module19/demultiplex_19.2.1.go)

Изначальное решение находится [здесь](https://github.com/MoJIoToK/learning_go/blob/master/module19/demultiplex.go)


## Improvements
