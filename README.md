# Fibonacci service

Возвращает срез последовательности чисел Фибоначчи 

## Для запуска приложения


```bash
make run
```

## HTTP REST

-POST /sequence  
Входные параметры в формате json: start number, end number- начальный и конечный порядковые номера возвращаемой последовательности(от 0 до 93)  
Выходные данные в формате json: '{"index":number, "index":number,...}'

## gRPC
Call GetSequence  
Входные параметры: start int, end int- начальный и конечный порядковые номера возвращаемой последовательности(от 0 до 93)  
Возвращает: {"result":{"index":number,...}}
