# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What This Repo Is

A personal fork of [Ardan Labs' Ultimate Go training material](https://github.com/ardanlabs/gotraining), extended with personal data-structure work (`ds/`) and an Ethereum smart-contract module (`smartcontract/`). The repo module name is `github.com/ardanlabs/gotraining` (root `go.mod`).

## Module Layout

This is a **Go workspace** (`go.work`). The workspace currently includes only the `smartcontract` sub-module; the root module is used independently.

| Directory | Module / purpose |
|-----------|-----------------|
| `topics/go/` | Core Ultimate Go training code (language, design, concurrency, generics, testing, profiling, packages, algorithms) |
| `ds/` | Personal data-structure implementations (heap variants, LRU cache, binary tree, linked list, stack, priority queues) — several have their own `go.mod` |
| `smartcontract/` | Separate Go module (`github.com/ardanlabs/smartcontract`) for Ethereum smart-contract work |
| `security/` | Password utility |
| `endians/` | Byte-order examples |
| `tools/mpl/` | MPL plotting tool |
| `topics/go/videos/miki/` | Extra video-series examples (build tags, cgo, embed, generics, interfaces) |

## Commands

### Root module — run tests

```sh
# Run all test files (mirrors CI)
find . -name "*_test.go" -not -wholename "*vendor*" -and -not -wholename "*student*" \
  -exec dirname {} \; | uniq | xargs go test

# Run a single package's tests
go test ./ds/heap/heapsort/...
go test ./topics/go/testing/strings/...
```

### smartcontract module

All commands are `make` targets defined in `smartcontract/makefile`. Run them from the `smartcontract/` directory.

```sh
# Run all Go tests + vet + staticcheck + govulncheck
make test

# Dependency hygiene
make tidy
```

#### Ethereum dev-node lifecycle (requires geth installed)

```sh
make geth-up        # start geth in dev mode (mines on demand)
make geth-down      # send SIGINT to geth
make geth-reset     # wipe local chain data (zarf/ethereum/geth/)
make geth-attach    # open JS console
make geth-deposit   # fund the three secondary accounts with 1 ETH each
```

#### Smart-contract build → deploy → interact cycle

Each app follows the same three-step pattern (shown for `basic`):

```sh
# 1. Compile Solidity → ABI + BIN → generate Go bindings
make basic-build

# 2. Deploy to local geth (requires geth-up)
make basic-deploy

# 3. Interact
make basic-read
make basic-write
```

Same pattern applies for `bank-single`, `bank-proxy`, and the proxy's versioned API builds (`bank-api-v1-build`, `bank-api-v2-build`, `bank-api-v3-build`).

## Architecture — smartcontract

Each app under `smartcontract/app/<name>/` follows a layered layout:

```
app/<name>/
  contract/
    src/       Solidity source (.sol)
    abi/       Compiled ABI + BIN (committed)
    go/        Auto-generated Go bindings (abigen output, committed)
  cmd/         Runnable Go programs (deploy, read/write, balance, etc.)
  pkg/         Optional reusable Go library over the bindings
```

`abigen` (from go-ethereum) converts the compiled ABI/BIN into a Go package that exposes the contract as a typed struct. The `cmd/` programs import this package and use `github.com/ardanlabs/ethereum` to connect to geth via JSON-RPC.

The `bank/proxy` app demonstrates the **upgradeable proxy pattern**: `Bank.sol` is the immutable proxy; `BankAPI.sol` holds the logic and can be redeployed at new addresses while the proxy address stays stable.

## Architecture — ds

Data-structure packages in `ds/` are standalone experiments, not a shared library. Several (`heap`, `lrucache`, `heap/heapint`, `heap/heapsort`, `heap/priorityqueueint`) have their own `go.mod` and must be run/tested from inside their own directory. `ds/heap/dijkstra/` contains reference material (PDF + notes) for a Dijkstra implementation using a heap — the implementation itself is not yet written.

## Prerequisites for smartcontract work

- `geth` — `sudo add-apt-repository -y ppa:ethereum/ethereum && sudo apt-get install ethereum`
- `solc` (Solidity compiler) — same PPA or platform-specific binary
- `abigen` — ships with `go-ethereum`; install with `go install github.com/ethereum/go-ethereum/cmd/abigen@latest`
- `staticcheck` and `govulncheck` — required by `make test`

The three pre-funded accounts used throughout the examples:
- Coinbase: `0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd` (passkey: `123`)
- Account 2: `0x8e113078adf6888b7ba84967f299f29aece24c55`
- Account 3: `0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7`
