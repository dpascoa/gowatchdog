package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/dpascoa/gowatchdog/checker"
	"github.com/dpascoa/gowatchdog/logger"
)

type Config struct {
	IntervalSeconds int      `json:"interval_seconds"`
	URLs            []string `json:"urls"`
}

func loadConfig(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var config Config
	err = decoder.Decode(&config)
	return config, err
}

func main() {
	fmt.Println("🐶 GoWatchdog starting...")

	config, err := loadConfig("config.json")
	if err != nil {
		panic(err)
	}

	fmt.Printf("🔁 Checking %d URLs every %d seconds\n", len(config.URLs), config.IntervalSeconds)

	err = logger.InitLog("gowatchdog.log")
	if err != nil {
		panic(err)
	}
	defer logger.CloseLog()

	go func() {
		tmpl := template.Must(template.New("index.html").
			Funcs(template.FuncMap{
				"statusClass": func(status string) string {
					if strings.Contains(status, "UP") {
						return "up"
					}
					return "down"
				},
			}).
			ParseFiles("web/templates/index.html"))

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			data := struct {
				Now      string
				Statuses map[string]string
			}{
				Now:      time.Now().Format("2006-01-02 15:04:05"),
				Statuses: logger.Latest,
			}

			err := tmpl.Execute(w, data)
			if err != nil {
				http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
				fmt.Println("Template error:", err)
			}
		})

		fs := http.FileServer(http.Dir("web/static"))
		http.Handle("/static/", http.StripPrefix("/static/", fs))

		fmt.Println("📡 Dashboard at http://localhost:8080")
		http.ListenAndServe(":8080", nil)
	}()

	for {
		var wg sync.WaitGroup

		for _, url := range config.URLs {
			wg.Add(1)
			go func(u string) {
				defer wg.Done()
				result := checker.CheckURL(u)
				logger.LogStatus(result.URL, result.StatusCode, result.Alive)
			}(url)
		}

		wg.Wait()

		time.Sleep(time.Duration(config.IntervalSeconds) * time.Second)
		fmt.Println("🔁 Rechecking...")
	}
}
