package main

import (
	"fmt"

	"github.com/shandler/go-expert-stress-test/internal/infra"
	"github.com/shandler/go-expert-stress-test/internal/usecase"
)

func main() {
	sr := usecase.NewStressRequest(
		infra.NewMakeHttpRequest(),
		usecase.NewJsonPresenter(),
	)
	execResult, err := sr.Execute(setInputFlags())
	if err != nil {
		fmt.Printf(fmt.Sprintf(`{"error": "%s"}`, err.Error()))
		return
	}

	fmt.Println(string(execResult))
}
