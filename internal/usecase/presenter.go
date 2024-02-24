package usecase

import (
	"encoding/json"

	"github.com/shandler/go-expert-stress-test/internal/dto"
)

type Presenter interface {
	Present(output dto.RequestFlagOutput) ([]byte, error)
}

type JsonPresenter struct{}

func NewJsonPresenter() *JsonPresenter {
	return &JsonPresenter{}
}

func (*JsonPresenter) Present(output dto.RequestFlagOutput) ([]byte, error) {
	bytes, err := json.Marshal(output)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
