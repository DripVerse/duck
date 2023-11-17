# [DUCK 🦆](https://duck.dripverse.org)

### _Simplifying Decentralised Communication_

If you care about data ownership and need a private free chat, Duck Chat has got your back.

_Have a private conversation in a crowded room without worrying about eavesdroppers._

## Install 🦆

<!-- ### Mac

```sh
brew install dripverse/tools/duck
```

### Ubuntu/Debian based machines


```sh
sudo apt install duck
```

_Currently having some issues with installing from respective package manager repos._ -->

### [Download Binary from latest Release](https://github.com/DripVerse/duck/releases) ⬇️

## Stable Releases:
- [v1.1.0](https://github.com/DripVerse/duck/releases/tag/v1.1.0) - P2P Chat works. Base Version.

## Usage:

```sh
duck
```

![Duck Chat](https://duck.dripverse.org/images/about/cli.png "Duck Chat")

### Features ✨
- Complete Data Ownership
- Enhanced Privacy
- No single point of control
- No third party involvement
- Free Usage
- Open Source
- Lightweight

### Config

DUCK uses config from `.drip` env file located either at `$HOME/.drip` or local directory. The first time you run, this environment file will be set. Sensitive information like keys and secrets can be stored here.
[TODO: Accept ENV variables from command](https://github.com/DripVerse/duck/issues/3)

If you have used any other DripVerse CLI tools, this config should already exist and no further config will be needed.

#### Important Config

| Key | Type | Description | Required/Optional |
|---|---|---|---|
| `ACCOUNT_KEY_EVM` | String | Your Private Key to EVM EOA Account | Required |
| `DP_KEY` | String | DripVerse Key if you want to use DripVerse Features like Account, Wallet, etc. | Optional |

---

### Development 💻

`go install` && `go run main.go`

`make build` && `./duck`

`make install` && `duck`

## Overview

Inspired by Peerchat on Kademlia DHT which is a terminal-based P2P chat application using libp2p and the IPFS network for peer discovery and routing.
It uses a Kademlia DHT from libp2p for peer discovery and routing and supports a more fully featured host. The other components of the libp2p such as TLS encryption, peer active discovery, YAMUX stream multiplexing are integrated as well.

### Dependencies:

- [tview](https://github.com/rivo/tview) and [tcell](https://github.com/gdamore/tcell) for TUI.
- [Logrus](https://github.com/sirupsen/logrus) for Logging.
- [libp2p](https://libp2p.io/) for p2p networking.
- [Viper](https://github.com/spf13/viper) for config management.

## Roadmap 🚀

- [x] P2P CLI Chat.
- [ ] Web UI.
- [ ] Config Management
- [ ] Wallet/Account Support.
- [ ] Add as Utility to DripVerse.
- [ ] ...

_[Have a request for features or add-ons?](https://github.com/DripVerse/duck/issues)_

### Devices Tested:

| Device Type | OS | Status |
| --- | --- | --- |
| Desktop | Mac Os 14 Sonoma | ✅ |
| Desktop | Linux PopOS | ✅ |

_Desktop here is used as a synonymous term for laptops, desktops, Macs and PCs._
