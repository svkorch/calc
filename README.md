# Тестовое задание для Golang-разработчиков

Нужно создать функцию `calculate`. Функция должна принимать арифметические операции двух чисел в виде строки и возвращать строку с результатом их выполнения. Функция принимает данные из аргументов и возвращает с помощью `return`.

Вдобавок нужно написать табличные тесты для функции `calculate`, создать репозиторий на сайте GitHub, расшарить доступы пользователю с логином [iampokrovsky](https://github.com/iampokrovsky), запушить задание в ветку `task` и сделать PR в ветку `main`. Сообщения об ошибках можно просто выводить в консоль.

## Требования

 - **Успешно решённое задание должно пройти все тесты.**
 - Калькулятор умеет выполнять операции сложения, вычитания, умножения и деления с двумя числами: `a + b, a - b, a * b, a / b`. Данные передаются в виде одной строки!
 - Калькулятор умеет работать как с арабскими (1,2,3,4,5…), так и с римскими (I,II,III,IV,V…) числами.
 - Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более. На выходе числа не ограничиваются по величине и могут быть любыми.
 - Калькулятор умеет работать только с целыми числами.
 - Результат на выходе всегда строка с целым числом. В делении учитываем только целую часть.
 - Калькулятор умеет работать только с арабскими или римскими цифрами одновременно, при вводе пользователем строки вроде `3 + II` калькулятор должен выбросить исключение и прекратить свою работу.
 - Т.к. в римской системе нет нуля и отрицательных чисел - возвращаем пустую строку.
 - При вводе пользователем неподходящих чисел функция `calculate` печатает сообщение об ошибке и возвращает непустую строку.
 - При вводе пользователем строки не соответствующей одной из вышеописанных арифметических операций функция `calculate` печатает сообщение об ошибке и возвращает непустую строку.
 - Сообщения об ошибках просто выводить в консоль. Без аварийного завершения программы.

## Пример работы программы

```
calculate('1 + 2'); // '3'
calculate('VI / III'); // 'II'
calculate('VII / III'); // 'II'
calculate('I + II'); // 'III'
calculate('I - II'); // ''
calculate('I + 1'); // throws Error
calculate('I'); // throws Error
calculate('1 + 1 + 1'); // throws Error
```

## Принципы оценки работы

Разбейте ваше решение на функции, отвечающие за разные области. Решения, где будет только одна функция `calculate`, будут низко оценены.


# Отчёт по выполнению задания и выводы

## Результаты тестирования производительности

Конвертация с использованием мапок многократно эффективнее традиционной (библиотечной) реализации преобразований строк в числа только на римских цифрах. На арабских же наблюдается противоположная ситуация.

```
$ go test -bench=. ./internal/calc/ -count 1 -run=^#
goos: linux
goarch: amd64
pkg: github.com/svkorch/calc/internal/calc
cpu: Intel(R) Core(TM) i3-8350K CPU @ 4.00GHz
Benchmark_libRomanStrToInt/input_[I]-4         	22600912	        53.31 ns/op
Benchmark_libRomanStrToInt/input_[II]-4        	14725027	        81.52 ns/op
Benchmark_libRomanStrToInt/input_[III]-4       	13927584	        86.12 ns/op
Benchmark_libRomanStrToInt/input_[IV]-4        	16283077	        73.23 ns/op
Benchmark_libRomanStrToInt/input_[V]-4         	25158380	        47.64 ns/op
Benchmark_libRomanStrToInt/input_[VI]-4        	15374280	        77.78 ns/op
Benchmark_libRomanStrToInt/input_[VII]-4       	14032304	        86.04 ns/op
Benchmark_libRomanStrToInt/input_[VIII]-4      	13293289	        90.19 ns/op
Benchmark_libRomanStrToInt/input_[IX]-4        	19998057	        58.75 ns/op
Benchmark_libRomanStrToInt/input_[X]-4         	27437202	        43.59 ns/op
Benchmark_libArabicStrToInt/input_[1]-4        	397850910	         3.013 ns/op
Benchmark_libArabicStrToInt/input_[2]-4        	373284021	         3.212 ns/op
Benchmark_libArabicStrToInt/input_[3]-4        	397928860	         3.013 ns/op
Benchmark_libArabicStrToInt/input_[4]-4        	371347044	         3.229 ns/op
Benchmark_libArabicStrToInt/input_[5]-4        	397867782	         3.013 ns/op
Benchmark_libArabicStrToInt/input_[6]-4        	397834383	         3.014 ns/op
Benchmark_libArabicStrToInt/input_[7]-4        	397885545	         3.013 ns/op
Benchmark_libArabicStrToInt/input_[8]-4        	397790018	         3.013 ns/op
Benchmark_libArabicStrToInt/input_[9]-4        	398014803	         3.013 ns/op
Benchmark_libArabicStrToInt/input_[10]-4       	293012348	         4.087 ns/op
Benchmark_mapRomanStrToInt/input_[I]-4         	145843437	         8.192 ns/op
Benchmark_mapRomanStrToInt/input_[II]-4        	145440103	         8.168 ns/op
Benchmark_mapRomanStrToInt/input_[III]-4       	135964024	         8.766 ns/op
Benchmark_mapRomanStrToInt/input_[IV]-4        	133185764	         8.920 ns/op
Benchmark_mapRomanStrToInt/input_[V]-4         	131423974	         9.112 ns/op
Benchmark_mapRomanStrToInt/input_[VI]-4        	120879825	         9.833 ns/op
Benchmark_mapRomanStrToInt/input_[VII]-4       	127533234	         9.317 ns/op
Benchmark_mapRomanStrToInt/input_[VIII]-4      	128982254	         9.264 ns/op
Benchmark_mapRomanStrToInt/input_[IX]-4        	124092463	         9.850 ns/op
Benchmark_mapRomanStrToInt/input_[X]-4         	100000000	        10.56 ns/op
Benchmark_mapArabicStrToInt/input_[1]-4        	100000000	        11.17 ns/op
Benchmark_mapArabicStrToInt/input_[2]-4        	100000000	        11.65 ns/op
Benchmark_mapArabicStrToInt/input_[3]-4        	100000000	        11.16 ns/op
Benchmark_mapArabicStrToInt/input_[4]-4        	100000000	        11.68 ns/op
Benchmark_mapArabicStrToInt/input_[5]-4        	98715091	        12.46 ns/op
Benchmark_mapArabicStrToInt/input_[6]-4        	97264310	        12.43 ns/op
Benchmark_mapArabicStrToInt/input_[7]-4        	92982422	        13.01 ns/op
Benchmark_mapArabicStrToInt/input_[8]-4        	91929247	        13.09 ns/op
Benchmark_mapArabicStrToInt/input_[9]-4        	90545331	        13.42 ns/op
Benchmark_mapArabicStrToInt/input_[10]-4       	88454785	        13.42 ns/op
PASS
ok  	github.com/svkorch/calc/internal/calc	59.929s
```

В тех случаях, когда на вход в функцию преобразования в числа строк с представлением чисел арабскими цифрами подаётся строка содержащая **некорректное десятичное число** (а, например, **число, записанное римскими цифрами**), скорость попытки преобразования многократно отстаёт от реализации с поиском в мапе.

```
$ go test -bench=. ./internal/calc/ -count 1 -run=^# -bench=Arabic
goos: linux
goarch: amd64
pkg: github.com/svkorch/calc/internal/calc
cpu: Intel(R) Core(TM) i3-8350K CPU @ 4.00GHz
Benchmark_libArabicStrToInt/input_[1]-4         	397760674	         3.013 ns/op
Benchmark_mapArabicStrToInt/input_[1]-4         	100000000	        11.13 ns/op
PASS
ok  	github.com/svkorch/calc/internal/calc	2.632s

$ go test -bench=. ./internal/calc/ -count 1 -run=^# -bench=Arabic
goos: linux
goarch: amd64
pkg: github.com/svkorch/calc/internal/calc
cpu: Intel(R) Core(TM) i3-8350K CPU @ 4.00GHz
Benchmark_libArabicStrToInt/input_[VIII]-4         	22323192	        55.32 ns/op
Benchmark_mapArabicStrToInt/input_[VIII]-4         	81427831	        12.63 ns/op
PASS
ok  	github.com/svkorch/calc/internal/calc	2.338s

```

## Вывод
Поскольку представление числа заранее неизвестно и, к тому же, ожидаемый операнд может и вовсе не быть числом, то наиболее производительным решением будет **использовать поиск по мапе, но только по одной единственной** (содержащей в ключах все возможные представления всех возможных корректных входных значений операндов. Их будет всего `2 * 10 = 20`.)
