package step

import (
	"encoding/json"
	"fmt"
	"strings"

	"clash-cli/api"
	"clash-cli/model"
)

type Traffic struct {
	*api.Client
	LastStep Step
}

func (t Traffic) Run() error {
	if err := handleReaderWithInterrupt(t.GetTraffics, func(bytes []byte) error {
		traffic := &model.Traffic{}
		if err := json.Unmarshal(bytes, traffic); err != nil {
			return err
		}

		fmt.Printf("\rUpload：%s  Download：%s %s",
			traffic.Up+"/s", traffic.Down+"/s", strings.Repeat(" ", 10))
		return nil
	}); err != nil && !IsCanceled(err) {
		return err
	}
	return t.LastStep.Run()
}
