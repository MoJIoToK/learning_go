package middleware

import (
	"context"
	"github.com/sqids/sqids-go"
	"net/http"
	"time"
)

// ctxKey - тип ключа для ID запроса внутри контекста.
type ctxKey int

// RequestIDKey - ключ ID запроса внутри контекста.
const RequestIDKey ctxKey = 0

// RequestIDHeader - HTTP заголовок ID запроса.
const RequestIDHeader = "X-Request-Id"

// RequestID - проверяет наличие уникального ID запроса в заголовках и записывает значение в контекст.
// Если ID не найдено, то идёт генерация нового ID.
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get(RequestIDHeader)

		if requestID == "" {
			tm := time.Now()
			sec := uint64(tm.Unix())
			nano := uint64(tm.UnixNano())

			s, _ := sqids.New(sqids.Options{MinLength: 10})
			id, err := s.Encode([]uint64{sec, nano})
			if err != nil {
				id = "unknown RequestID"
			}
			requestID = id

			r.Header.Set(RequestIDHeader, requestID)
		}

		ctx := context.WithValue(r.Context(), RequestIDKey, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetRequestID - возвращает ID запроса из контекста в виде строки
func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	if requestID, ok := ctx.Value(RequestIDKey).(string); ok {
		return requestID
	}

	return ""
}
