package main

import (
	"os"
	"time"
	"sync"
	"bytes"
	"syscall"
	"context"
	"net/http"
	"os/signal"
	"encoding/json"
)

const (
	URI = ""
	DATA_CONTENT_TYPE = "application/json"
	INTERVAL = time.Duration(10 * time.Second)
)


type Data struct {
	Text string `json:"text"`
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	go func () {
		wg.Add(1)
		ticker := time.NewTicker(INTERVAL)
		defer ticker.Stop()

	Loop:
		for {
			select {
			case <-ctx.Done():
				break Loop
			case <-ticker.C:
				data, err := json.Marshal(&Data{
					Text: "Hello World!",
				})
				if err != nil {
					panic(err)
				}

				resp, err := http.Post(URI, DATA_CONTENT_TYPE, bytes.NewBuffer(data))
				if err != nil {
					panic(err)
				}
				resp.Body.Close()
			default:
				break
			}
		}
		wg.Done()
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func () {
		<-sigCh
		cancel()
	}()

	wg.Wait()
}
