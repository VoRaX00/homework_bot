package scheduleParser

import (
	"encoding/json"
	"homework_bot/internal/domain"
	"io/ioutil"
	"net/http"
)

type FefuParser struct {
}

func NewFefuParser() *FefuParser {
	return &FefuParser{}
}

func ByteToSchedule(body []byte) (domain.Schedule, error) {
	var schedule domain.Schedule
	err := json.Unmarshal(body, &schedule)
	return schedule, err
}

func (p *FefuParser) ParseSchedule(link string) (domain.Schedule, error) {
	resp, err := http.Get(link)
	if err != nil {
		return domain.Schedule{}, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return domain.Schedule{}, err
	}

	return ByteToSchedule(body)
}
