package service

import (
	"github.com/pdrum/todo/model"
	"github.com/pdrum/todo/request"
	"github.com/pdrum/todo/response"
)

type ItemService struct {
	Repo model.ItemRepo
}

func (i ItemService) Create(request request.CreateItemRequest) (*response.ItemResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	item := model.Item{
		Title: request.Title,
		Done:  request.Done,
	}
	if err := i.Repo.Create(&item); err != nil {
		return nil, err
	}
	return &response.ItemResponse{
		ID:    item.ID,
		Title: item.Title,
		Done:  item.Done,
	}, nil
}

func (i ItemService) ListAll() ([]response.ItemResponse, error) {
	items, err := i.Repo.ListAll()
	if err != nil {
		return nil, err
	}
	responses := []response.ItemResponse{}
	for _, item := range items {
		responses = append(responses, response.ItemResponse{
			ID:    item.ID,
			Title: item.Title,
			Done:  item.Done,
		})
	}
	return responses, nil
}
