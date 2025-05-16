package scanner

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Проверяет доступность URL с таймаутом
type URLChecker struct {
	client http.Client
}

// Создает проверщик с заданным таймаутом запроса
func NewURLChecker(timeout time.Duration) *URLChecker {
	return &URLChecker{
		client: http.Client{Timeout: timeout},
	}
}

// Выполняет GET-запрос с учетом контекста
// Возвращает статус или ошибку (при коде >=400 или проблемах запроса)
func (c *URLChecker) Check(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("не удалось создать запрос: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("ошибка запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	return resp.Status, nil
}
