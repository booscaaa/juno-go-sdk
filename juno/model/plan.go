package model

import (
	"encoding/json"
	"fmt"
	"io"
)

type _Embedded struct {
	Embedded Plans `json:"_embedded"`
}

type Plans struct {
	Plans []Plan `json:"plans"`
}

type Plan struct {
	ID        string  `json:"id"`
	CreatedOn string  `json:"createdOn"`
	Name      string  `json:"name"`
	Frequency string  `json:"frequency"`
	Status    string  `json:"status"`
	Amount    float64 `json:"amount"`
}

func FromJsonJunoPlan(body io.ReadCloser) (*Plan, error) {
	plan := Plan{}
	err := json.NewDecoder(body).Decode(&plan)

	if err != nil {
		return nil, err
	}

	return &plan, nil
}

func FromJsonJunoPlans(body io.ReadCloser) (*[]Plan, error) {
	embedded := _Embedded{}
	err := json.NewDecoder(body).Decode(&embedded)

	fmt.Println(err)

	if err != nil {
		return nil, err
	}

	return &embedded.Embedded.Plans, nil
}
