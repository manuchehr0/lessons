# Занятие 3: Условные операторы и логика
"Научим компьютер принимать решения"

### 📝 План на сегодня
* Операторы сравнения: Как сравнивать данные
* 
* Условный оператор if: Основа принятия решений
* 
* Else и else if: Множественные условия
* 
* Логические операторы: И, ИЛИ, НЕ
* 
* Оператор switch: Альтернатива множественным if
* 
* Практика: Реальные примеры использования

### 1. Операторы сравнения
   Операторы сравнения возвращают true (истина) или false (ложь).


`a := 10
b := 5

fmt.Println(a == b)  // false - равно ли?
fmt.Println(a != b)  // true  - не равно?
fmt.Println(a > b)   // true  - больше?
fmt.Println(a < b)   // false - меньше?
fmt.Println(a >= b)  // true  - больше или равно?
fmt.Println(a <= b)  // false - меньше или равно?`
### 2. Условный оператор if
   if выполняет код, если условие истинно.

Базовый синтаксис:
go
if условие {
// код выполнится, если условие true
}
Простой пример:
go
age := 18

if age >= 18 {
fmt.Println("Вы совершеннолетний")
}
### 3. Else и else if
   Else - "в противном случае":
   go
   temperature := 25

if temperature > 30 {
fmt.Println("Жарко")
} else {
fmt.Println("Не жарко")
}
Else if - множественные условия:
go
score := 85

if score >= 90 {
fmt.Println("Оценка: A")
} else if score >= 80 {
fmt.Println("Оценка: B")  // ← выполнится это
} else if score >= 70 {
fmt.Println("Оценка: C")
} else {
fmt.Println("Оценка: F")
}
### 4. Логические операторы
   И (AND) - &&
   go
   age := 20
   hasLicense := true

if age >= 18 && hasLicense {
fmt.Println("Можно водить машину")  // Оба условия true
}
ИЛИ (OR) - ||
go
day := "суббота"

if day == "суббота" || day == "воскресенье" {
fmt.Println("Выходной день!")
}
НЕ (NOT) - !
go
isRaining := false

if !isRaining {
fmt.Println("Можно идти гулять")  // isRaining = false, !false = true
}
### 5. Оператор switch
   Switch - удобная альтернатива множественным if-else.

Простой switch:
go
day := "понедельник"

switch day {
case "понедельник":
fmt.Println("Начало недели")
case "пятница":
fmt.Println("Сколько выходные!")
default:
fmt.Println("Обычный день")
}
Switch без выражения:
go
score := 85

switch {
case score >= 90:
fmt.Println("Отлично!")
case score >= 70:
fmt.Println("Хорошо")  // ← выполнится это
default:
fmt.Println("Плохо")
}
### 🎯 Практика 1: Проверка четности
Задача: Определить, четное ли число

`go
package main
import "fmt"

func main() {
number := 7

    if number % 2 == 0 {
        fmt.Println(number, "- четное число")
    } else {
        fmt.Println(number, "- нечетное число")
    }
}`
% - оператор остатка от деления. Если число делится на 2 без остатка - оно четное.

### 🎯 Практика 2: Калькулятор с операциями
Задача: Улучшим калькулятор - пользователь выбирает операцию

`go
package main
import "fmt"

func main() {
var a, b float64
var operation string

    fmt.Print("Введите первое число: ")
    fmt.Scan(&a)
    
    fmt.Print("Введите второе число: ")
    fmt.Scan(&b)
    
    fmt.Print("Выберите операцию (+, -, *, /): ")
    fmt.Scan(&operation)
    
    if operation == "+" {
        fmt.Println("Результат:", a + b)
    } else if operation == "-" {
        fmt.Println("Результат:", a - b)
    } else if operation == "*" {
        fmt.Println("Результат:", a * b)
    } else if operation == "/" {
        if b != 0 {
            fmt.Println("Результат:", a / b)
        } else {
            fmt.Println("Ошибка: деление на ноль!")
        }
    } else {
        fmt.Println("Неизвестная операция")
    }
}`
### 🎯 Практика 3: Проверка возраста
Задача: Определить категорию по возрасту

`go
package main
import "fmt"

func main() {
var age int
fmt.Print("Введите ваш возраст: ")
fmt.Scan(&age)

    switch {
    case age < 0:
        fmt.Println("Ошибка: возраст не может быть отрицательным")
    case age < 7:
        fmt.Println("Дошкольник")
    case age < 18:
        fmt.Println("Школьник")
    case age < 65:
        fmt.Println("Взрослый")
    default:
        fmt.Println("Пенсионер")
    }
}`
### ❓ Важные моменты
Фигурные скобки ОБЯЗАТЕЛЬНЫ:
`go
// ПРАВИЛЬНО:
if condition {
// код
}

// ОШИБКА:
if condition
fmt.Println("test")
Инициализация в условии:
go
// Можно объявить переменную прямо в условии:
if age := 20; age >= 18 {
fmt.Println("Совершеннолетний")
}`
// age доступна только внутри блока if
### 🏠 Домашнее задание
#### Задача 1: FizzBuzz
* Напишите программу, которая для чисел от 1 до 100:

* Выводит "Fizz", если число делится на 3

* Выводит "Buzz", если число делится на 5

* Выводит "FizzBuzz", если число делится и на 3, и на 5

Иначе выводит само число

Пример вывода для числа 15: FizzBuzz

#### Задача 2: Калькулятор ИМТ
Напишите программу, которая:

Запрашивает рост (в метрах) и вес (в кг)

Вычисляет ИМТ по формуле: вес / (рост * рост)

Выводит категорию:

Меньше 18.5: "Недостаточный вес"

18.5-24.9: "Нормальный вес"

25-29.9: "Избыточный вес"

30 и больше: "Ожирение"

### 🚀 Что ждет на следующем занятии?
Циклы: Как повторять действия много раз

For loop: Единственный, но мощный цикл в Go

Break и continue: Управление выполнением цикла

Удачи в решении задач! 🎉

