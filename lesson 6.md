# Занятие 5: Карты (Maps) и структуры данных

`"Хранение данных по ключам - как в словаре"`

---

## 📝 План на сегодня

1.  **Что такое карта (map)?** Концепция ключ-значение
2.  **Создание и инициализация** карт
3.  **Основные операции:** добавление, получение, удаление
4.  **Проверка существования** ключей
5.  **Обход карт** с помощью range
6.  **Практика:** Реальные примеры использования

---

## 1. Что такое карта (Map)?

**Карта** - это коллекция пар "ключ-значение", где каждый ключ уникален.

| Ключ | Значение |
|------|----------|
| "user1" | "Анна" |
| "user2" | "Петр" |
| "user3" | "Мария" |

> **Аналогия:** Как словарь - слово (ключ) и его определение (значение)

---

## 2. Создание и инициализация карт

### Пустая карта:
```go
// Ключи string, значения int
var scores map[string]int

// С помощью make (рекомендуемый способ)
ages := make(map[string]int)
```

### Инициализация с данными:
```go
// Карта с начальными значениями
users := map[string]string{
    "user1": "Анна",
    "user2": "Петр", 
    "user3": "Мария",
}

// Разные типы ключей и значений
config := map[string]interface{}{
    "port":     8080,
    "host":     "localhost",
    "debug":    true,
    "database": "postgres",
}
```

---

## 3. Основные операции с картами

### Добавление и обновление:
```go
colors := make(map[string]string)

// Добавление элементов
colors["red"] = "#FF0000"
colors["green"] = "#00FF00"
colors["blue"] = "#0000FF"

// Обновление
colors["red"] = "#FF1111"
```

### Получение значений:
```go
fmt.Println(colors["red"])   // #FF1111
fmt.Println(colors["black"]) // пустое значение (для string - "")
```

### Удаление элементов:
```go
delete(colors, "green")
fmt.Println(colors["green"]) // пустая строка
```

---

## 4. Проверка существования ключей

### Правильный способ проверки:
```go
grades := map[string]int{
    "Анна": 85,
    "Петр": 92,
}

// Неправильно (если оценка 0, мы не поймем - нет ключа или оценка 0)
grade := grades["Мария"] // 0

// Правильно - проверяем существование
grade, exists := grades["Мария"]
if exists {
    fmt.Println("Оценка:", grade)
} else {
    fmt.Println("Студент не найден")
}

// Или короче:
if grade, exists := grades["Анна"]; exists {
    fmt.Println("Оценка Анны:", grade)
}
```

---

## 5. Обход карт с помощью range

### Базовый обход:
```go
users := map[string]string{
    "alice": "Алиса",
    "bob":   "Боб",
    "carol": "Каролина",
}

// Обход ключей и значений
for key, value := range users {
    fmt.Printf("Ключ: %s, Значение: %s\n", key, value)
}

// Только ключи
for key := range users {
    fmt.Println("Ключ:", key)
}

// Только значения
for _, value := range users {
    fmt.Println("Значение:", value)
}
```

---

## 🎯 Практика 1: Телефонная книга

**Задача:** Создать простую телефонную книгу

```go
package main
import "fmt"

func main() {
    phonebook := make(map[string]string)
    
    // Добавляем контакты
    phonebook["Анна"] = "+7-900-123-45-67"
    phonebook["Петр"] = "+7-900-987-65-43" 
    phonebook["Мария"] = "+7-900-555-35-35"
    
    // Поиск номера
    name := "Анна"
    if phone, exists := phonebook[name]; exists {
        fmt.Printf("Номер %s: %s\n", name, phone)
    } else {
        fmt.Printf("Контакт %s не найден\n", name)
    }
    
    // Вывод всех контактов
    fmt.Println("\nВсе контакты:")
    for name, phone := range phonebook {
        fmt.Printf("%s: %s\n", name, phone)
    }
    
    // Удаление контакта
    delete(phonebook, "Петр")
    fmt.Println("\nПосле удаления Петра:")
    for name, phone := range phonebook {
        fmt.Printf("%s: %s\n", name, phone)
    }
}
```

---

## 🎯 Практика 2: Подсчет слов в тексте

**Задача:** Посчитать частоту встречаемости слов

```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    text := "привет мир привет гофер привет программа"
    words := strings.Split(text, " ") // Разбиваем текст на слова
    
    wordCount := make(map[string]int)
    
    // Подсчитываем слова
    for _, word := range words {
        wordCount[word]++ // Увеличиваем счетчик для каждого слова
    }
    
    // Выводим результаты
    fmt.Println("Частота слов:")
    for word, count := range wordCount {
        fmt.Printf("%s: %d раз(а)\n", word, count)
    }
}
```

---

## 🎯 Практика 3: Система голосования

**Задача:** Подсчитать голоса за разных кандидатов

```go
package main
import "fmt"

func main() {
    votes := []string{"Анна", "Петр", "Анна", "Мария", "Петр", "Анна"}
    results := make(map[string]int)
    
    // Подсчет голосов
    for _, candidate := range votes {
        results[candidate]++
    }
    
    // Находим победителя
    winner := ""
    maxVotes := 0
    
    for candidate, count := range results {
        fmt.Printf("%s: %d голосов\n", candidate, count)
        
        if count > maxVotes {
            maxVotes = count
            winner = candidate
        }
    }
    
    fmt.Printf("\nПобедитель: %s с %d голосами!\n", winner, maxVotes)
}
```

---

## 🎯 Практика 4: Вложенные карты

**Задача:** Хранение информации о студентах

```go
package main
import "fmt"

func main() {
    students := make(map[string]map[string]interface{})
    
    // Добавляем студентов
    students["student1"] = map[string]interface{}{
        "name": "Анна",
        "age":  20,
        "grades": []int{85, 92, 78},
    }
    
    students["student2"] = map[string]interface{}{
        "name": "Петр", 
        "age":  22,
        "grades": []int{90, 88, 95},
    }
    
    // Выводим информацию
    for id, info := range students {
        fmt.Printf("Студент %s:\n", id)
        fmt.Printf("  Имя: %s\n", info["name"])
        fmt.Printf("  Возраст: %d\n", info["age"])
        fmt.Printf("  Оценки: %v\n", info["grades"])
        fmt.Println()
    }
    
    // Получаем конкретного студента
    if student, exists := students["student1"]; exists {
        fmt.Printf("Информация о student1: %s, %d лет\n", 
            student["name"], student["age"])
    }
}
```

---

## ❓ Важные моменты

### Карта vs Срез:
```go
// Срез - доступ по индексу (число)
users := []string{"Анна", "Петр"}
fmt.Println(users[0]) // "Анна"

// Карта - доступ по ключу (любой сравниваемый тип)
usersMap := map[string]string{"user1": "Анna", "user2": "Петр"}
fmt.Println(usersMap["user1"]) // "Анна"
```

### Нулевое значение карты:
```go
var nilMap map[string]int
// nilMap["key"] = 1 // ПАНИКА! Карта не инициализирована

// Всегда используйте make для создания карт
safeMap := make(map[string]int)
safeMap["key"] = 1 // OK
```

---

## 🏠 Домашнее задание

**Задача 1: Улучшенная телефонная книга**
Создайте телефонную книгу, которая позволяет:
1. Добавлять контакты (имя и телефон)
2. Удалять контакты
3. Искать контакты по имени
4. Показывать все контакты
5. Проверять, существует ли контакт перед добавлением

**Задача 2: Анализатор текста**
Напишите программу, которая:
1. Принимает текст от пользователя
2. Подсчитывает частоту каждого слова
3. Находит самое частое слово
4. Выводит топ-3 самых частых слов

**Задача 3: Система учета товаров**
Создайте систему учета товаров:
```go
products := map[string]map[string]interface{}{
    "item1": {"name": "Ноутбук", "price": 50000, "quantity": 5},
    "item2": {"name": "Мышь", "price": 1000, "quantity": 20},
}
```
Добавьте функции для:
- Просмотра всех товаров
- Поиска товара по имени
- Расчет общей стоимости всех товаров на складе

---

## 🚀 Что ждет на следующем занятии?

*   **Функции:** Создание переиспользуемого кода
*   **Параметры и возвращаемые значения**
*   **Замыкания и анонимные функции**

**Удачи в решении задач! 🎉**