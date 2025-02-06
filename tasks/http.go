package tasks

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type HttpTask struct {
	url    string
	data   string
	client *http.Client
}

func NewHttpTask(url string, data string) *HttpTask {
	return &HttpTask{url, data, &http.Client{Timeout: 30 * time.Second}}
}

func (h *HttpTask) Start() error {
	req, err := http.NewRequest(http.MethodGet, h.url, nil)
	if err != nil {
		return err
	}

	res, err := h.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	log.Println(string(body))

	return nil
}

func (h *HttpTask) StartMsg() string {
	return fmt.Sprintf("Http Request Start by URL %s", h.url)
}

func (h *HttpTask) EndMsg() string {
	return fmt.Sprintf("Http Request End by URL %s", h.url)
}
