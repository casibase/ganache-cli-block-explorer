# Ganache-CLI-Block-Explorer
Author : Vinay Awasthi

Thanks to Original Author: vivekganesan01@gmail.com
Go the https://github.com/vivekganesan01/ganache-cli-block-explorer and give a like

Ganache-cli-block-explorer is a web based block reader, which connects to your local ganache (powered by truffle) and explore the block details from the local blockchain network.

  - Explore all the transactions within the blocks
  - Gas limit
  - Pending transactions
  - Block mined details and much more ...


> Designed to help block chain learners to
> understand ganache cli in a better way.


### Tech

Ganache-CLI-Block-Explorer involves:

* [HTML CSS Bootstrap] - HTML enhanced for web apps!
* [Go lang ethereum library] - To communicate with ganache client

And of course Dillinger itself is open source with a [public repository][dill]
 on GitHub.

### Installation

`git pull the repository`

The default configuration file is `conf/app.yaml`. You can run the project directly, or modify the configuration as needed:

```yaml
network_host: "http://localhost:8545"  # Ethereum JSON-RPC host
server_addr: "127.0.0.1:5051"           # Web server address
```

If you do not modify the file, the program will use the above default settings.

```go
go run main.go
```

Open your browser and visit your server address to verify deployment:

```sh
127.0.0.1:5051
```

> Note: The application listens on port 5051 by default. Please make sure this port is not occupied.

To change the network or port, simply edit `conf/app.yaml` and restart the service for changes to take effect.

### Development

Want to contribute? Great!
Open your favourite Terminal and run these commands.

First thing:
```sh
 git checkout -b your-branch
 Make changes and create a pull request to `release` branch
```
Note: Checkout from `master`.

### Dependencies
- add go mod, open command prompt and execute the following commands
  * go mod init ganache-cli-block-explorer
  * go mod tidy

- install deps
  * go get github.com/ethereum/go-ethereum/common github.com/ethereum/go-ethereum/core/types github.com/ethereum/go-ethereum/ethclient github.com/gorilla/mux


### Demo

#### Welcome Page

![image](https://user-images.githubusercontent.com/15568499/175276001-023de8a8-fb67-4b26-b058-712119b89a7f.png)

#### Home Page

![image](https://user-images.githubusercontent.com/15568499/175277496-37625532-0c02-4e93-b79c-2677d2fa7d30.png)
![image](https://user-images.githubusercontent.com/15568499/175277614-f19ce8d2-53bc-456c-bf3c-022469b5c137.png)

#### Block Details
![image](https://user-images.githubusercontent.com/15568499/175278638-dd8c5dad-80a9-4653-9c18-94d96de97355.png)


#### Transaction Details
![image](https://user-images.githubusercontent.com/15568499/175278787-5d4e4c6d-da44-480a-98d2-fe0dde236766.png)

#### Account Details
![image](https://user-images.githubusercontent.com/15568499/175279062-a7a2cdc6-2971-4799-99c5-3745a7c8d17a.png)

