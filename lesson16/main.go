package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var version = "development"

func main() {
	log.Println("version:", version)
	log.Println("password:", os.Getenv("PG_PASSWORD"))
	mux := http.NewServeMux()

	mux.HandleFunc("/login", LoginHandler)
	mux.HandleFunc("/protected", ProtectedHandler)

	rootWithMiddlewares := Chain(ErrorHandler, Logging, CORS, AuthMiddleware2)(mux)

	log.Println(" at :8080")

	http.ListenAndServe(":8080", rootWithMiddlewares)
}

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("check 1")
		//defer func() {
		//	if err := recover(); err != nil {
		//		log.Printf("Panic: %v", err)
		//		writeError(w, http.StatusInternalServerError, "Internal server error")
		//	}
		//}()
		panic("test 2")
		next.ServeHTTP(w, r)
		log.Println("check 2")
	})
}

func writeError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(AppError{
		Code:    code,
		Message: message,
	})
}
func AuthMiddleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer valid-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "habr.ru")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)
		//time.Sleep(3 * time.Second)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

type Middleware func(http.Handler) http.Handler

func Chain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("I am protected route")
	fmt.Fprintf(w, "Welcome home")
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	panic("test panic")
	fmt.Fprintf(w, "Your token:\n%s", "token")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем токен из заголовка Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Формат: "Bearer <token>"
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}

		tokenString := bearerToken[1]

		// Валидируем токен
		claims, err := ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Добавляем claims в контекст запроса
		ctx := context.WithValue(r.Context(), "userClaims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Вспомогательная функция для получения claims из контекста
func GetUserFromContext(r *http.Request) *Claims {
	if claims, ok := r.Context().Value("userClaims").(*Claims); ok {
		return claims
	}
	return nil
}

// Проверка пароля
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Секретный ключ (должен храниться в безопасном месте)
var jwtSecret = []byte("your-secret")

// Claims (данные в токене)
type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// Создание JWT токена
func GenerateToken(userID int, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "your-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Валидация JWT токена
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
