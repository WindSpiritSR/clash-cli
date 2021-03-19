# clash-cli

[![Release](https://img.shields.io/github/v/release/WindSpiritSR/clash-cli.svg)](https://github.com/WindSpiritSR/clash-cli/releases/latest)
[![License](https://img.shields.io/github/license/WindSpiritSR/clash-cli)](https://mit-license.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/WindSpiritSR/clash-cli)](https://goreportcard.com/report/github.com/WindSpiritSR/clash-cli)

A CLI tool for clash

# What can it do

- Change proxy mode
- Switch proxy node
- Realtime Traffic
- Proxy log

# How does it do this

This tool uses bbolt to store the connection data with Clash, then it send GET/PUT requests to RESTful API of Clash to control it

![Function Select](screenshot/menu.png)

![Latency Test](screenshot/latency-test.png)

![Proxy Log](screenshot/proxy-log.png)
