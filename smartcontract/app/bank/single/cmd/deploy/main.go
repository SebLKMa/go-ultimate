package main

import (
	"context"
	"fmt"
	"io"
	"math/big"
	"os"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ardanlabs/smartcontract/app/bank/single/contract/go/bank"
	"github.com/ethereum/go-ethereum/log"
)

const (
	keyStoreFile     = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	passPhrase       = "123"
	coinMarketCapKey = "a8cd12fb-d056-423f-877b-659046af0aa5"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() (err error) {
	ctx := context.Background()

	stdOut := log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, log.LevelInfo, true))
	discard := log.NewLogger(log.NewTerminalHandlerWithLevel(io.Discard, log.LevelInfo, true))

	backend, err := ethereum.CreateDialedBackend(ctx, ethereum.NetworkLocalhost)
	if err != nil {
		return err
	}
	defer backend.Close()

	privateKey, err := ethereum.PrivateKeyByKeyFile(keyStoreFile, passPhrase)
	if err != nil {
		return err
	}

	clt, err := ethereum.NewClient(backend, privateKey)
	if err != nil {
		return err
	}

	fmt.Println("\nInput Values")
	fmt.Println("----------------------------------------------------")
	fmt.Println("fromAddress:", clt.Address())

	// =========================================================================

	converter, err := currency.NewConverter(bank.BankMetaData.ABI, coinMarketCapKey)
	if err != nil {
		converter = currency.NewDefaultConverter(bank.BankMetaData.ABI)
	}
	oneETHToUSD, oneUSDToETH := converter.Values()

	fmt.Println("oneETHToUSD:", oneETHToUSD)
	fmt.Println("oneUSDToETH:", oneUSDToETH)

	// =========================================================================

	startingBalance, err := clt.Balance(ctx)
	if err != nil {
		return err
	}
	defer func() {
		endingBalance, dErr := clt.Balance(ctx)
		if dErr != nil {
			err = dErr
			return
		}
		fmt.Print(converter.FmtBalanceSheet(startingBalance, endingBalance))
	}()

	// =========================================================================

	const gasLimit = 1700000
	const gasPriceGwei = 39.576
	const valueGwei = 0.0
	tranOpts, err := clt.NewTransactOpts(ctx, gasLimit, currency.GWei2Wei(big.NewFloat(gasPriceGwei)), big.NewFloat(valueGwei))
	if err != nil {
		return err
	}

	// =========================================================================

	address, tx, _, err := bank.DeployBank(tranOpts, clt.Backend)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransaction(tx))

	fmt.Println("\nContract Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("contract id     :", address.Hex())

	if err := os.WriteFile("zarf/ethereum/bank_single.cid", []byte(address.Hex()), 0644); err != nil {
		return fmt.Errorf("exporting bank_single.cid file: %w", err)
	}

	// =========================================================================

	fmt.Println("\nWaiting Logs")
	fmt.Println("----------------------------------------------------")
	log.SetDefault(stdOut)

	receipt, err := clt.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))
	log.SetDefault(discard)

	// =========================================================================

	bankContract, err := bank.NewBank(address, clt.Backend)
	if err != nil {
		return fmt.Errorf("new proxy connection: %w", err)
	}

	callOpts, err := clt.NewCallOpts(ctx)
	if err != nil {
		return err
	}

	version, err := bankContract.Version(callOpts)
	if err != nil {
		return err
	}

	fmt.Println("\nValidate Version")
	fmt.Println("----------------------------------------------------")
	fmt.Println("version         :", version)

	return nil
}

/* Sample output:
$ make bank-single-deploy
CGO_ENABLED=0 go run app/bank/single/cmd/deploy/main.go

Input Values
----------------------------------------------------
fromAddress: 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd
oneETHToUSD: 2569.7544568991243
oneUSDToETH: 0.0003891422378178037

Transaction Details
----------------------------------------------------
hash            : 0xc6f5dcbdb400461de7446e52f2335bbfdbb05bede2245d3901c75590d2b1152a
nonce           : 0
gas limit       : 1700000
gas offer price : 39.576 GWei
value           : 0 GWei
max gas price   : 67279200 GWei
max gas price   : 172.89 USD

Contract Details
----------------------------------------------------
contract id     : 0x531130464929826c57BBBF989e44085a02eeB120

Waiting Logs
----------------------------------------------------

Receipt Details
----------------------------------------------------
status          : 1
gas used        : 1629118
gas price       : 39.576 GWei
gas price       : 0.00 USD
final gas cost  : 64473973.97 GWei
final gas cost  : 165.68 USD

Logs
----------------------------------------------------

Validate Version
----------------------------------------------------
version         : 0.1.0

Balance
----------------------------------------------------
balance before  : 1.157920892e+68 GWei
balance after   : 1.157920892e+68 GWei
balance diff    : 65337878.47 GWei
balance diff    : 167.90 USD
*/
