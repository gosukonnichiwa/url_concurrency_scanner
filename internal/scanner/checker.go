package scanner

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type URLChecker struct {
	client http.Client
}

func NewURLChecker(timeout time.Duration) *URLChecker {
	return &URLChecker{
		client: http.Client{Timeout: timeout},
	}
}

func (c *URLChecker) Check(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("request creation failed: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	return resp.Status, nil
}
