# Занятие 8: Указатели и структуры (Structs)

`"Работа с памятью и создание собственных типов данных"`

---

## 📝 План на сегодня

1.  **Что такое указатели?** Концепция адреса памяти
2.  **Операторы & и *:** Получение адреса и разыменование
3.  **Зачем нужны указатели?** Изменение значений в функциях
4.  **Структуры (Structs):** Создание собственных типов данных
5.  **Методы структур:** Функции, привязанные к структурам
6.  **Практика:** Создание реальных структур данных

---

## 1. Что такое указатели?

**Указатель** - это переменная, которая хранит адрес памяти другой переменной.

### Аналогия:
- **Переменная** = Дом (данные)
- **Указатель** = Адрес дома (где найти данные)

```go
var x int = 10    // Переменная x со значением 10
var p *int = &x   // Указатель p хранит адрес переменной x

fmt.Println("Значение x:", x)     // 10
fmt.Println("Адрес x:", &x)       // 0xc0000180a0 (пример адреса)
fmt.Println("Значение p:", p)     // 0xc0000180a0
fmt.Println("Значение по адресу p:", *p) // 10
```

---

## 2. Операторы & и *

### & (амперсанд) - "адрес от":
```go
name := "Анна"
pointer := &name  // pointer теперь содержит адрес переменной name

fmt.Println("Значение name:", name)      // Анна
fmt.Println("Адрес name:", &name)        // 0xc000010200
fmt.Println("Значение pointer:", pointer) // 0xc000010200
```

### * (звездочка) - "значение по адресу":
```go
age := 25
agePointer := &age

fmt.Println("Возраст:", age)                   // 25
fmt.Println("Адрес возраста:", agePointer)     // 0xc0000180b0
fmt.Println("Значение по адресу:", *agePointer) // 25

// Изменение значения через указатель
*agePointer = 30
fmt.Println("Новый возраст:", age) // 30
```

---

## 3. Зачем нужны указатели?

### Проблема: Функции работают с копиями
```go
func increment(x int) {
    x = x + 1
    fmt.Println("В функции:", x) // 6
}

func main() {
    number := 5
    increment(number)
    fmt.Println("В main:", number) // 5 (не изменилось!)
}
```

### Решение: Использование указателей
```go
func increment(x *int) { // Принимает указатель
    *x = *x + 1         // Изменяем значение по адресу
    fmt.Println("В функции:", *x) // 6
}

func main() {
    number := 5
    increment(&number)   // Передаем адрес переменной
    fmt.Println("В main:", number) // 6 (изменилось!)
}
```

---

## 4. Структуры (Structs)

**Структура** - это пользовательский тип данных, который объединяет несколько полей разных типов.

### Объявление структуры:
```go
// Определяем новый тип Person
type Person struct {
    Name    string
    Age     int
    Email   string
    IsStudent bool
}
```

### Создание экземпляров структуры:
```go
// Способ 1: С указанием имен полей
person1 := Person{
    Name: "Анна",
    Age: 25,
    Email: "anna@example.com",
    IsStudent: false,
}

// Способ 2: Без имен полей (в порядке объявления)
person2 := Person{"Петр", 30, "petr@example.com", false}

// Способ 3: Создание пустой структуры с последующим заполнением
var person3 Person
person3.Name = "Мария"
person3.Age = 22
person3.IsStudent = true
```

---

## 5. Методы структур

**Метод** - это функция, привязанная к конкретной структуре.

### Синтаксис методов:
```go
// (p Person) - "получатель" (receiver)
func (p Person) Introduce() string {
    return fmt.Sprintf("Привет, меня зовут %s, мне %d лет", p.Name, p.Age)
}

// Метод с указателем-получателем (может изменять структуру)
func (p *Person) HaveBirthday() {
    p.Age++
    fmt.Printf("%s празднует день рождения! Теперь ему/ей %d лет\n", p.Name, p.Age)
}

func main() {
    person := Person{"Анна", 25, "anna@example.com", false}
    
    fmt.Println(person.Introduce()) // Привет, меня зовут Анна, мне 25 лет
    person.HaveBirthday()          // Анна празднует день рождения! Теперь ему/ей 26 лет
    fmt.Println(person.Introduce()) // Привет, меня зовут Анна, мне 26 лет
}
```

---

## 🎯 Практика 1: Работа с указателями

**Задача:** Поменять значения двух переменных местами

```go
package main
import "fmt"

// Не работает (передаются копии)
func swapBad(a, b int) {
    temp := a
    a = b
    b = temp
}

// Работает (передаются указатели)
func swapGood(a, b *int) {
    temp := *a
    *a = *b
    *b = temp
}

func main() {
    x, y := 10, 20
    
    fmt.Printf("До swap: x=%d, y=%d\n", x, y) // x=10, y=20
    
    swapBad(x, y)
    fmt.Printf("После swapBad: x=%d, y=%d\n", x, y) // x=10, y=20
    
    swapGood(&x, &y)
    fmt.Printf("После swapGood: x=%d, y=%d\n", x, y) // x=20, y=10
}
```

---

## 🎯 Практика 2: Структура "Прямоугольник"

**Задача:** Создать структуру Rectangle с методами

```go
package main
import "fmt"

type Rectangle struct {
    Width  float64
    Height float64
}

// Метод для вычисления площади
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Метод для вычисления периметра
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Метод для масштабирования (использует указатель)
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// Метод для красивого вывода
func (r Rectangle) Display() {
    fmt.Printf("Прямоугольник: %.1f x %.1f\n", r.Width, r.Height)
    fmt.Printf("Площадь: %.1f\n", r.Area())
    fmt.Printf("Периметр: %.1f\n", r.Perimeter())
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    
    fmt.Println("Исходный прямоугольник:")
    rect.Display()
    
    fmt.Println("\nПосле масштабирования в 2 раза:")
    rect.Scale(2)
    rect.Display()
}
```

---

## 🎯 Практика 3: Система управления банковским счетом

**Задача:** Создать структуру BankAccount с методами

```go
package main
import "fmt"

type BankAccount struct {
    Owner   string
    Balance float64
    Number  string
}

// Метод для пополнения счета
func (acc *BankAccount) Deposit(amount float64) {
    if amount > 0 {
        acc.Balance += amount
        fmt.Printf("Счет пополнен на %.2f. Новый баланс: %.2f\n", amount, acc.Balance)
    } else {
        fmt.Println("Ошибка: сумма должна быть положительной")
    }
}

// Метод для снятия денег
func (acc *BankAccount) Withdraw(amount float64) bool {
    if amount > 0 && amount <= acc.Balance {
        acc.Balance -= amount
        fmt.Printf("Со счета снято %.2f. Новый баланс: %.2f\n", amount, acc.Balance)
        return true
    } else {
        fmt.Println("Ошибка: недостаточно средств или неверная сумма")
        return false
    }
}

// Метод для перевода денег на другой счет
func (acc *BankAccount) Transfer(amount float64, recipient *BankAccount) bool {
    if acc.Withdraw(amount) {
        recipient.Deposit(amount)
        fmt.Printf("Перевод %.2f на счет %s выполнен успешно\n", amount, recipient.Owner)
        return true
    }
    return false
}

// Метод для вывода информации о счете
func (acc BankAccount) Display() {
    fmt.Printf("Владелец: %s\n", acc.Owner)
    fmt.Printf("Номер счета: %s\n", acc.Number)
    fmt.Printf("Баланс: %.2f\n", acc.Balance)
}

func main() {
    // Создаем счета
    account1 := BankAccount{
        Owner:   "Анна",
        Balance: 1000,
        Number:  "1234567890",
    }
    
    account2 := BankAccount{
        Owner:   "Петр", 
        Balance: 500,
        Number:  "0987654321",
    }
    
    fmt.Println("Начальное состояние:")
    account1.Display()
    fmt.Println()
    account2.Display()
    
    fmt.Println("\n--- Операции ---")
    account1.Deposit(200)
    account1.Withdraw(150)
    account1.Transfer(300, &account2)
    
    fmt.Println("\nФинальное состояние:")
    account1.Display()
    fmt.Println()
    account2.Display()
}
```

---

## ❓ Важные моменты

### Когда использовать указатели-получатели:
```go
// Value receiver (работает с копией) - для методов, которые НЕ изменяют структуру
func (p Person) GetName() string {
    return p.Name
}

// Pointer receiver (работает с оригиналом) - для методов, которые изменяют структуру
func (p *Person) SetName(name string) {
    p.Name = name
}
```

### Указатели и nil:
```go
var p *Person // nil pointer
// p.Name = "Анна" // ПАНИКА! Разыменование nil указателя

// Всегда проверяйте на nil
if p != nil {
    p.Name = "Анна"
}
```

---

## 🏠 Домашнее задание

**Задача 1: Структура "Студент"**
Создайте структуру Student с полями:
- Name (string)
- Grades ([]int)
- Age (int)

Добавьте методы:
- `AddGrade(grade int)` - добавляет оценку
- `GetAverage() float64` - вычисляет средний балл
- `IsExcellent() bool` - возвращает true, если средний балл >= 4.5

**Задача 2: Калькулятор с историей**
Создайте структуру Calculator с полем History ([]string).
Добавьте методы для операций, которые записывают каждое действие в историю.

**Задача 3: Менеджер задач**
Создайте структуру Task с полями:
- Title
- Description
- IsCompleted
- Priority

Создайте структуру TaskManager с методами:
- AddTask, CompleteTask, GetPendingTasks, GetCompletedTasks

---

## 🚀 Что ждет на следующем занятии?

*   **Интерфейсы:** Контракты для структур
*   **Полиморфизм в Go**
*   **Пустой интерфейс interface{}**

**Удачи в решении задач! 🎉**