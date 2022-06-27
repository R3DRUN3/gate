# Gate 🔎🚪
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


*Gate* is a simple and fast port scanner powered by *Go* 

<p align="center"><img width="180" height="180" src="https://github.com/yurijserrano/Github-Profile-Readme-Logos/blob/master/programming%20languages/go.svg"></p>

## Usage

Clone this repo:

```sh
git clone https://github.com/R3DRUN3/gate.git
```

Start the program:

```sh
cd gate
go run main.go
```

Output Sample: 

```console
   __   ____      _      _____   _____  __
  / /  / ___|    / \    |_   _| | ____| \ \
 | |  | |  _    / _ \     | |   |  _|    | |
 | |  | |_| |  / ___ \    | |   | |___   | |
 | |   \____| /_/   \_\   |_|   |_____|  | |
  \_\                                   /_/
  ____    _                       _            ____                   _       ____
 / ___|  (_)  _ __ ___    _ __   | |   ___    |  _ \    ___    _ __  | |_    / ___|    ___    __ _   _ __    _ __     ___   _ __
 \___ \  | | | '_ ` _ \  | '_ \  | |  / _ \   | |_) |  / _ \  | '__| | __|   \___ \   / __|  / _` | | '_ \  | '_ \   / _ \ | '__|
  ___) | | | | | | | | | | |_) | | | |  __/   |  __/  | (_) | | |    | |_     ___) | | (__  | (_| | | | | | | | | | |  __/ | |
 |____/  |_| |_| |_| |_| | .__/  |_|  \___|   |_|      \___/  |_|     \__|   |____/   \___|  \__,_| |_| |_| |_| |_|  \___| |_|
                         |_|

Enter target ip: 192.168.236.195
Enter protocolo (tcp/udp): tcp

Scanning 65535 tcp ports on target 192.168.236.195 ===>

##########################################
#         Port 22: tcp Open
#         Port 443: tcp Open
##########################################

Runtime in seconds:  6.238812893

```


## Usage with Docker
Follow the Docs [here](https://hub.docker.com/repository/docker/r3drun3/gate) on DockerHub


