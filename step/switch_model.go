package step

import (
	"clash-cli/api"
	"clash-cli/model"
	T "github.com/Dreamacro/clash/tunnel"
	"github.com/manifoldco/promptui"
)

type SwitchMode struct {
	*api.Client
	LastStep Step
}

func (sm SwitchMode) Run() error {
	configs, err := sm.GetConfigs()
	if err != nil {
		return err
	}

	items := make([]string, 3)
	for i, v := range []T.TunnelMode{T.Global, T.Rule, T.Direct} {
		if v != *configs.Mode {
			items[i] = v.String()
			continue
		}
		items[i] = v.String() + promptui.IconGood
	}
	prompt := promptui.Select{
		Label: model.PROMPT_MODEL_LABEL,
		Items: items,
	}
	result, _, err := prompt.Run()

	if err != nil {
		return err
	}
	if err := sm.UpdateMode(T.TunnelMode(result)); err != nil {
		return err
	}
	return sm.LastStep.Run()
}
