package handler

import (
	"fmt"
)

func errValidation(m map[string]string) error {
	return fmt.Errorf("missing fields are required: %s", fmt.Sprint(m))
}

// Create Opening
type CreateOpeningRequest struct {
	Role     string  `json:"role"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Remote   *bool   `json:"remote"`
	Link     string  `json:"link"`
	Salary   float64 `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	validationErrorsMap := make(map[string]string)

	if r.Role == "" {
		validationErrorsMap["role"] = "string"
	}
	if r.Company == "" {
		validationErrorsMap["company"] = "string"
	}
	if r.Location == "" {
		validationErrorsMap["location"] = "string"
	}
	if r.Link == "" {
		validationErrorsMap["link"] = "string"
	}
	if r.Remote == nil {
		validationErrorsMap["remote"] = "bool"
	}
	if r.Salary <= 0 {
		validationErrorsMap["salary"] = "float64"
	}

	if len(validationErrorsMap) > 0 {
		return errValidation(validationErrorsMap)
	}

	return nil
}

// UpdateOpening
type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	// If any field is provided, validation is truthy
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}
	// If none of the fields were provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")
}
