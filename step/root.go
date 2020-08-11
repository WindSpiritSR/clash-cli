package step

import (
	"clash-cli/api"
	"clash-cli/model"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
)

type Root struct {
	*api.Client
}

func (r Root) Run() error {
	selectItems := []string{model.PROMPT_ROOT_ITEM_TYPE, model.PROMPT_ROOT_ITEM_PROXY,
		model.PROMPT_ROOT_ITEM_TRAFFIC, model.PROMPT_ROOT_ITEM_LOG,
		model.PROMPT_ROOT_ITEM_CLICONF, model.PROMPT_ROOT_ITEM_EXIT}

	checkConn, err := r.checkConn()
	if !checkConn {
		fmt.Printf("%s\n", model.WARNING_CANNOT_CONN_CLASH)
		selectItems = selectItems[4:]
	} else {
		clashConf, err := r.GetConfigs()
		if err != nil {
			return err
		}
		if clashConf.Mode.String() == model.CLASH_CONF_MODE_DIRECT {
			for index, item := range selectItems {
				if item == model.PROMPT_ROOT_ITEM_PROXY {
					selectItems = append(selectItems[:index], selectItems[index+1:]...)
				}
			}
		}
	}

	prompt := promptui.Select{
		Label:     model.PROMPT_ROOT_LABEL,
		Size:      10,
		Items:     selectItems,
		IsVimMode: false,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return err
	}
	var step Step
	switch result {
	case model.PROMPT_ROOT_ITEM_TYPE:
		step = SwitchMode{
			Client:   r.Client,
			LastStep: r,
		}
	case model.PROMPT_ROOT_ITEM_PROXY:
		step = SwitchProxy{
			Client:   r.Client,
			LastStep: r,
		}
	case model.PROMPT_ROOT_ITEM_TRAFFIC:
		step = Traffic{
			Client:   r.Client,
			LastStep: r,
		}
	case model.PROMPT_ROOT_ITEM_LOG:
		step = Log{
			Client:   r.Client,
			LastStep: r,
		}
	case model.PROMPT_ROOT_ITEM_CLICONF:
		step = Config{
			Client:   r.Client,
			LastStep: r,
		}
	default:
		os.Exit(0)
	}
	return step.Run()
}

func (r Root) checkConn() (bool, error) {
	_, err := r.GetProxies()
	if err != nil {
		return false, err
	}
	return true, nil
}
