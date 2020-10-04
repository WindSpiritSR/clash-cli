package step

import (
	"fmt"

	"clash-cli/api"
	"clash-cli/model"
	C "github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/tunnel"
	"github.com/manifoldco/promptui"
)

type SwitchProxy struct {
	*api.Client
	LastStep Step
}

func (sp SwitchProxy) Run() error {
	configs, err := sp.GetConfigs()
	if err != nil {
		return err
	}
	proxies, err := sp.GetProxies()
	if err != nil {
		return err
	}
	switch *configs.Mode {
	case tunnel.Global:
		if err := sp.UpdateProxy(proxies.Proxies, proxies.Proxies[model.PROMPT_PROXY_GLOBAL_GROUPNAME],
			model.PROMPT_PROXY_GLOBAL_GROUPNAME, model.PROMPT_PROXY_GLOBAL_LABEL); err != nil {
			return err
		}
		return sp.LastStep.Run()
	default: // Rule
		var selectors []struct {
			model.Proxy
			Name string
		}

		var proxyGroup []string
		for name, group := range proxies.Proxies {
			if group.Type.Is(C.Selector) && name != model.PROMPT_PROXY_ITEM_ALL {
				proxyGroup = append(proxyGroup, name)
			}
		}

		prompt := promptui.Select{
			Label: model.PROMPT_PROXY_LABEL,
			Items: proxyGroup,
		}
		_, proxyMode, err := prompt.Run()
		if err != nil {
			return err
		}

		for name, group := range proxies.Proxies {
			if group.Type.Is(C.Selector) && name == proxyMode {
				selectors = append(selectors, struct {
					model.Proxy
					Name string
				}{Name: name, Proxy: group})
			}
		}

		switch len(selectors) {
		case 0:
			fmt.Println(promptui.IconWarn, model.WARNING_NOT_PROXY)
			return sp.LastStep.Run()
		case 1:
			if err := sp.UpdateProxy(proxies.Proxies, selectors[0].Proxy,
				selectors[0].Name, selectors[0].Name); err != nil {
				return err
			}
			return sp.LastStep.Run()
		default:
		}
	}
	return nil
}

func (sp SwitchProxy) UpdateProxy(proxies map[string]model.Proxy,
	selector model.Proxy, groupName, label string) error {
	items := []model.ProxyName{{Name: model.PROMPT_PROXY_ITEM_LATENCY_TEST, ItemType: model.ItemTypeLatencyTest}}
	for _, v := range selector.All {
		v.ExtraInfo = proxies[v.Name].Now
		if v.Name == selector.Now {
			v.Now = true
		}
		v.ItemType = model.ItemTypeProxy
		items = append(items, v)
	}
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}
	result, _, err := prompt.Run()
	if err != nil {
		return err
	}
	if items[result].ItemType == model.ItemTypeLatencyTest {
		return LatencyTest{
			Proxies:  items,
			Client:   sp.Client,
			LastStep: sp,
		}.Run()
	}
	if err := sp.Client.UpdateProxy(groupName, items[result].Name); err != nil {
		return err
	}
	return nil
}
