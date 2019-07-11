package request

import "errors"

type CreateItemRequest struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func (c CreateItemRequest) Validate() error {
	if c.Title == "" {
		return errors.New("title is required")
	}
	return nil
}
