package model

import (
	"encoding/json"
	C "github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/log"
	T "github.com/Dreamacro/clash/tunnel"
	"github.com/dustin/go-humanize"
	"github.com/manifoldco/promptui"
	"strings"
	"sync"
	"time"
)

type Config struct {
	AllowLan  *bool         `json:"allow-lan"`
	LogLevel  *string       `json:"log-level"`
	Mode      *T.TunnelMode `json:"mode"`
	Port      *int          `json:"port"`
	RedirPort *int          `json:"redir-port"`
	SocksPort *int          `json:"socks-port"`
}

type Proxies struct {
	Proxies map[string]Proxy `json:"proxies"`
}

type Proxy struct {
	All  []ProxyName `json:"all"`
	Type AdapterType `json:"type,omitempty"`
	Now  string      `json:"now,omitempty"`
}

const (
	ItemTypeProxy = iota
	ItemTypeLatencyTest
)

var Latencys sync.Map

type ProxyName struct {
	Now       bool
	ItemType  int
	Name      string
	ExtraInfo string
}

func (p ProxyName) String() string {
	s := p.Name
	if p.ExtraInfo != "" {
		s = strings.Join([]string{"[" + p.ExtraInfo + "]", p.Name}, " ")
	}
	if v, ok := Latencys.Load(p.Name); ok {
		s = "[" + v.(string) + "] " + p.Name
	}
	if p.Now {
		s += " " + promptui.IconGood
	}
	return s
}

func (p *ProxyName) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	p.Name = s
	return nil
}

type AdapterType C.AdapterType

func (a *AdapterType) UnmarshalJSON(b []byte) error {
	switch strings.Trim(string(b), "\"") {
	case C.Direct.String():
		*a = AdapterType(C.Direct)
	case C.Fallback.String():
		*a = AdapterType(C.Fallback)
	case C.Reject.String():
		*a = AdapterType(C.Reject)
	case C.Selector.String():
		*a = AdapterType(C.Selector)
	case C.Shadowsocks.String():
		*a = AdapterType(C.Shadowsocks)
	case C.Socks5.String():
		*a = AdapterType(C.Socks5)
	case C.Http.String():
		*a = AdapterType(C.Http)
	case C.URLTest.String():
		*a = AdapterType(C.URLTest)
	case C.Vmess.String():
		*a = AdapterType(C.Vmess)
	case C.LoadBalance.String():
		*a = AdapterType(C.LoadBalance)
	}
	return nil
}

func (a AdapterType) Is(v C.AdapterType) bool {
	return C.AdapterType(a) == v
}

type HumanBytes string

func (h *HumanBytes) UnmarshalJSON(d []byte) error {
	var i uint64
	if err := json.Unmarshal(d, &i); err != nil {
		return err
	}
	*h = HumanBytes(humanize.Bytes(i))
	return nil
}

type Traffic struct {
	Up   HumanBytes `json:"up"`
	Down HumanBytes `json:"down"`
}

type Log struct {
	Type    log.LogLevel `json:"type"`
	Payload string       `json:"payload"`
}

type HumanLatency string

func (h *HumanLatency) UnmarshalJSON(d []byte) error {
	var i int
	if err := json.Unmarshal(d, &i); err != nil {
		return err
	}
	*h = HumanLatency((time.Duration(i) * time.Millisecond).String())
	return nil
}

type Latency struct {
	Delay HumanLatency
}
