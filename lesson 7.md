# Занятие 7: Функции в Go

`"Создаем переиспользуемый код"`

---

## 📝 План на сегодня

1.  **Что такое функции и зачем они нужны?**
2.  **Объявление и вызов** функций
3.  **Параметры и аргументы**
4.  **Возвращаемые значения** (один и несколько)
5.  **Именованные возвращаемые значения**
6.  **Анонимные функции и замыкания**
7.  **Практика:** Рефакторинг кода с функциями

---

## 1. Что такое функции?

**Функция** - это блок кода, который выполняет определенную задачу и может быть вызван многократно.

### Преимущества функций:
- **Повторное использование** кода
- **Упрощение** сложных программ
- **Улучшение читаемости**
- **Облегчение тестирования**

```go
// Вместо этого (дублирование кода):
fmt.Println("Привет, Анна!")
fmt.Println("Привет, Петр!")
fmt.Println("Привет, Мария!")

// Лучше так (с функцией):
greet("Анна")
greet("Петр") 
greet("Мария")
```

---

## 2. Объявление и вызов функций

### Базовый синтаксис:
```go
func имяФункции(параметры) возвращаемоеЗначение {
    // тело функции
}
```

### Простая функция:
```go
// Объявление функции
func sayHello() {
    fmt.Println("Привет, мир!")
}

// Вызов функции
func main() {
    sayHello() // Выводит: Привет, мир!
    sayHello() // Можно вызывать много раз
}
```

---

## 3. Параметры и аргументы

### Функция с параметрами:
```go
// Параметр - то, что функция принимает
func greet(name string) {
    fmt.Printf("Привет, %s!\n", name)
}

func main() {
    // Аргумент - то, что мы передаем при вызове
    greet("Анна")   // Привет, Анна!
    greet("Петр")   // Привет, Петр!
}
```

### Несколько параметров:
```go
func introduce(name string, age int) {
    fmt.Printf("Меня зовут %s, мне %d лет\n", name, age)
}

func main() {
    introduce("Анна", 25)
    introduce("Петр", 30)
}
```

---

## 4. Возвращаемые значения

### Одно возвращаемое значение:
```go
func add(a int, b int) int {
    result := a + b
    return result
}

func main() {
    sum := add(5, 3)
    fmt.Println("Сумма:", sum) // Сумма: 8
}
```

### Несколько возвращаемых значений:
```go
// Функция возвращает результат и остаток от деления
func divide(dividend, divisor int) (int, int) {
    quotient := dividend / divisor
    remainder := dividend % divisor
    return quotient, remainder
}

func main() {
    q, r := divide(10, 3)
    fmt.Printf("10 / 3 = %d (остаток %d)\n", q, r) // 10 / 3 = 3 (остаток 1)
}
```

---

## 5. Именованные возвращаемые значения

### Явное именование результатов:
```go
// Результаты имеют имена (как переменные)
func calculate(a, b int) (sum int, product int) {
    sum = a + b      // Не нужно :=, так как sum уже объявлен
    product = a * b
    return           // Возвращает sum и product автоматически
}

func main() {
    s, p := calculate(4, 5)
    fmt.Printf("Сумма: %d, Произведение: %d\n", s, p)
}
```

### Преимущества именованных возвращаемых значений:
- Улучшают читаемость
- Упрощают документацию
- Автоматически возвращаются при пустом `return`

---

## 6. Анонимные функции и замыкания

### Анонимная функция (лямбда):
```go
func main() {
    // Функция без имени
    double := func(x int) int {
        return x * 2
    }
    
    fmt.Println(double(5)) // 10
    
    // Немедленный вызов
    result := func(a, b int) int {
        return a + b
    }(3, 4)
    
    fmt.Println(result) // 7
}
```

### Замыкания (функции, запоминающие окружение):
```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    myCounter := counter()
    fmt.Println(myCounter()) // 1
    fmt.Println(myCounter()) // 2
    fmt.Println(myCounter()) // 3
    
    anotherCounter := counter()
    fmt.Println(anotherCounter()) // 1 (новая переменная count)
}
```

---

## 🎯 Практика 1: Рефакторинг калькулятора

**Задача:** Разбить калькулятор на функции

```go
package main
import "fmt"

// Функции для операций
func add(a, b float64) float64 {
    return a + b
}

func subtract(a, b float64) float64 {
    return a - b
}

func multiply(a, b float64) float64 {
    return a * b
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("деление на ноль")
    }
    return a / b, nil
}

func main() {
    var a, b float64
    var operation string
    
    fmt.Print("Введите первое число: ")
    fmt.Scan(&a)
    
    fmt.Print("Введите второе число: ")
    fmt.Scan(&b)
    
    fmt.Print("Выберите операцию (+, -, *, /): ")
    fmt.Scan(&operation)
    
    switch operation {
    case "+":
        fmt.Printf("Результат: %.2f\n", add(a, b))
    case "-":
        fmt.Printf("Результат: %.2f\n", subtract(a, b))
    case "*":
        fmt.Printf("Результат: %.2f\n", multiply(a, b))
    case "/":
        result, err := divide(a, b)
        if err != nil {
            fmt.Println("Ошибка:", err)
        } else {
            fmt.Printf("Результат: %.2f\n", result)
        }
    default:
        fmt.Println("Неизвестная операция")
    }
}
```

---

## 🎯 Практика 2: Функции для работы со срезами

**Задача:** Создать функции для частых операций со срезами

```go
package main
import "fmt"

// Находит максимальное число в срезе
func findMax(numbers []int) int {
    if len(numbers) == 0 {
        return 0
    }
    
    max := numbers[0]
    for _, num := range numbers {
        if num > max {
            max = num
        }
    }
    return max
}

// Фильтрует четные числа
func filterEven(numbers []int) []int {
    var result []int
    for _, num := range numbers {
        if num % 2 == 0 {
            result = append(result, num)
        }
    }
    return result
}

// Считает среднее значение
func calculateAverage(numbers []int) float64 {
    if len(numbers) == 0 {
        return 0
    }
    
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return float64(sum) / float64(len(numbers))
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    fmt.Println("Исходный срез:", numbers)
    fmt.Println("Максимум:", findMax(numbers))
    fmt.Println("Четные числа:", filterEven(numbers))
    fmt.Printf("Среднее: %.2f\n", calculateAverage(numbers))
}
```

---

## 🎯 Практика 3: Функция с различными возвращаемыми значениями

**Задача:** Создать функцию для обработки пользовательских данных

```go
package main
import "fmt"

// Обрабатывает пользовательские данные
func processUserData(name string, age int) (string, bool, error) {
    // Проверяем валидность данных
    if name == "" {
        return "", false, fmt.Errorf("имя не может быть пустым")
    }
    
    if age < 0 || age > 150 {
        return "", false, fmt.Errorf("некорректный возраст: %d", age)
    }
    
    // Формируем приветствие
    greeting := fmt.Sprintf("Привет, %s! Тебе %d лет.", name, age)
    
    // Проверяем совершеннолетие
    isAdult := age >= 18
    
    return greeting, isAdult, nil
}

func main() {
    // Тестируем функцию с разными данными
    testCases := []struct{
        name string
        age  int
    }{
        {"Анна", 25},
        {"Петр", 17},
        {"", 30},
        {"Мария", -5},
    }
    
    for _, tc := range testCases {
        fmt.Printf("\nТест: %s, %d лет\n", tc.name, tc.age)
        
        greeting, isAdult, err := processUserData(tc.name, tc.age)
        if err != nil {
            fmt.Println("Ошибка:", err)
            continue
        }
        
        fmt.Println(greeting)
        if isAdult {
            fmt.Println("Совершеннолетний")
        } else {
            fmt.Println("Несовершеннолетний")
        }
    }
}
```

---

## ❓ Важные моменты

### Переменное количество аргументов:
```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3))       // 6
    fmt.Println(sum(1, 2, 3, 4, 5)) // 15
}
```

### Функции как значения:
```go
func applyOperation(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

func main() {
    result := applyOperation(10, 5, func(x, y int) int {
        return x * y
    })
    fmt.Println(result) // 50
}
```

---

## 🏠 Домашнее задание

**Задача 1: Улучшенный калькулятор**
Перепишите калькулятор из практики, добавив:
1. Функцию `calculate` которая принимает операцию и возвращает результат
2. Функцию `displayResult` для красивого вывода результатов
3. Функцию `getInput` для получения данных от пользователя

**Задача 2: Библиотека функций для работы со строками**
Создайте набор функций:
- `reverseString(s string) string` - переворачивает строку
- `countVowels(s string) int` - считает количество гласных
- `isPalindrome(s string) bool` - проверяет, является ли строка палиндромом

**Задача 3: Генератор отчетов**
Напишите функцию, которая принимает срез чисел и возвращает отчет:
```go
func generateReport(numbers []int) (min, max, average float64, evenCount int) {
    // ваша реализация
}
```

---

## 🚀 Что ждет на следующем занятии?

*   **Указатели:** Работа с памятью
*   **Структуры (Structs):** Создание собственных типов данных
*   **Объединение данных в логические группы**

**Удачи в решении задач! 🎉**