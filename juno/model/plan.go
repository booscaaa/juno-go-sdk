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

func (plan Plan) isValid() (bool, error) {
	if plan.ID == "" {
		return false, fmt.Errorf("ID not be empty")
	}

	if plan.CreatedOn == "" {
		return false, fmt.Errorf("CreatedOn not be empty")
	}

	if plan.Name == "" {
		return false, fmt.Errorf("Name not be empty")
	}

	if plan.Frequency == "" {
		return false, fmt.Errorf("Frequency not be empty")
	}

	if plan.Status == "" {
		return false, fmt.Errorf("Status not be empty")
	}

	if plan.Amount == 0.0 {
		return false, fmt.Errorf("Amount not be empty")
	}

	return true, nil
}

func FromJsonJunoPlan(body io.ReadCloser) (*Plan, error) {
	plan := Plan{}
	json.NewDecoder(body).Decode(&plan)

	if isValid, err := plan.isValid(); !isValid {
		return nil, err
	}

	return &plan, nil
}

func FromJsonJunoPlans(body io.ReadCloser) (*[]Plan, error) {
	embedded := _Embedded{}
	json.NewDecoder(body).Decode(&embedded)

	return &embedded.Embedded.Plans, nil
}
