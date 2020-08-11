package step

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"clash-cli/api"
	"clash-cli/model"
	"golang.org/x/sync/errgroup"
)

type LatencyTest struct {
	*api.Client
	LastStep Step
	Proxies  []model.ProxyName
}

func (dt LatencyTest) Run() error {
	fmt.Print(model.EXIT_BY_CTRL_C)
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-interrupt
		signal.Stop(interrupt)
		cancel()
		fmt.Print("\r")
	}()
	e, _ := errgroup.WithContext(ctx)
	for _, v := range dt.Proxies {
		if v.ItemType == model.ItemTypeProxy {
			e.Go(getLatency(ctx, dt.Client, v.Name))
		}
	}
	if err := e.Wait(); err != nil && !IsCanceled(err) {
		return err
	}
	return dt.LastStep.Run()
}

func getLatency(ctx context.Context, client *api.Client, name string) func() error {
	return func() error {
		latency, err := client.GetLatency(ctx, name)
		if err == nil {
			model.Latencys.Store(name, string(latency.Delay))
			return nil
		}

		switch strings.Split(err.Error(), ":")[0] {
		case model.REQUEST_LATENCY_TEST_ERROR_CODE:
			model.Latencys.Store(name, model.REQUEST_LATENCY_TEST_ERROR_MSG)
		case model.REQUEST_LATENCY_TEST_TIMEOUT_CODE:
			model.Latencys.Store(name, model.REQUEST_LATENCY_TEST_TIMEOUT_MSG)
		default:
			return err
		}
		return nil
	}
}
