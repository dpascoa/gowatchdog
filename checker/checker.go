// checker/checker.go
package checker

import (
	"net/http"
	"time"
)

type Result struct {
	URL        string
	StatusCode int
	Alive      bool
	CheckedAt  time.Time
}

func CheckURL(url string) Result {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	result := Result{
		URL:       url,
		CheckedAt: time.Now(),
	}

	if err != nil {
		result.Alive = false
		result.StatusCode = 0
		return result
	}
	defer resp.Body.Close()

	result.Alive = resp.StatusCode >= 200 && resp.StatusCode < 400
	result.StatusCode = resp.StatusCode

	return result
}
