package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ardanlabs/smartcontract/app/bank/proxy/contract/go/bank"
	"github.com/ardanlabs/smartcontract/app/bank/proxy/contract/go/bankapi"
	"github.com/ethereum/go-ethereum/common"
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

	const gasLimit = 1600000
	const gasPriceGwei = 39.576
	const valueGwei = 0.0
	tranOpts, err := clt.NewTransactOpts(ctx, gasLimit, currency.GWei2Wei(big.NewFloat(gasPriceGwei)), big.NewFloat(valueGwei))
	if err != nil {
		return err
	}

	// =========================================================================

	// The API impl address
	address, tx, _, err := bankapi.DeployBankapi(tranOpts, clt.Backend)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransaction(tx))

	fmt.Println("\nContract Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("contract id     :", address.Hex())

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

	contractIDBytes, err := os.ReadFile("zarf/ethereum/bank.cid")
	if err != nil {
		return fmt.Errorf("importing bank.cid file: %w", err)
	}

	contractID := string(contractIDBytes)
	if contractID == "" {
		return errors.New("need to export the bank.cid file")
	}
	fmt.Println("contractID:", contractID)

	fmt.Println("\nSet This Contract To Bank")
	fmt.Println("----------------------------------------------------")
	fmt.Println("bank id         :", contractID)
	fmt.Println("contract id     :", address.Hex())

	bankContract, err := bank.NewBank(common.HexToAddress(contractID), clt.Backend)
	if err != nil {
		return fmt.Errorf("new proxy connection: %w", err)
	}

	tranOpts.Nonce = big.NewInt(0).Add(tranOpts.Nonce, big.NewInt(1))

	// Provide API impl address to Proxy
	tx, err = bankContract.SetContract(tranOpts, common.HexToAddress(address.Hex()))
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransaction(tx))

	// =========================================================================

	fmt.Println("\nWaiting Logs")
	fmt.Println("----------------------------------------------------")
	log.SetDefault(stdOut)

	receipt, err = clt.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))
	log.SetDefault(discard)

	// =========================================================================

	callOpts, err := clt.NewCallOpts(ctx)
	if err != nil {
		return err
	}

	version, err := bankContract.Version(callOpts)
	if err != nil {
		return err
	}

	api, err := bankContract.API(callOpts)
	if err != nil {
		return err
	}

	fmt.Println("\nValidate Version and API")
	fmt.Println("----------------------------------------------------")
	fmt.Println("version         :", version)
	fmt.Println("api             :", api)

	return nil
}

/* After successfully deployed the proxy Back contract, check API Version must be 0.1.0 below:

$ make bank-api-deploy
makefile:162: warning: overriding recipe for target 'basic-test'
makefile:125: warning: ignoring old recipe for target 'basic-test'
CGO_ENABLED=0 go run app/bank/proxy/cmd/deploy/api/main.go

Input Values
----------------------------------------------------
fromAddress: 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd
oneETHToUSD: 2656.4669153656514
oneUSDToETH: 0.0003764398473083766

Transaction Details
----------------------------------------------------
hash            : 0xdf913a97ecea96cdc28cc3d5ece2ee6eea9a2b4c7460ddbeee04ab808f1c208d
nonce           : 7
gas limit       : 1600000
gas offer price : 39.576 GWei
value           : 0 GWei
max gas price   : 63321600 GWei
max gas price   : 168.21 USD

Contract Details
----------------------------------------------------
contract id     : 0x840bB1050D4d79c998667C35EEE4223Ef97127B2

Waiting Logs
----------------------------------------------------

Receipt Details
----------------------------------------------------
status          : 1
gas used        : 1535136
gas price       : 39.576 GWei
gas price       : 0.00 USD
final gas cost  : 60754542.34 GWei
final gas cost  : 161.39 USD

Logs
----------------------------------------------------
contractID: 0xA5B76e49bD18E952502f1eB4c4B281B91C727CBD

Set This Contract To Bank
----------------------------------------------------
bank id         : 0xA5B76e49bD18E952502f1eB4c4B281B91C727CBD
contract id     : 0x840bB1050D4d79c998667C35EEE4223Ef97127B2

Transaction Details
----------------------------------------------------
hash            : 0x67f9fd4771229e21f971047a44eaeef7416da354a4e3b0bcb86bd172c5f27c0a
nonce           : 8
gas limit       : 1600000
gas offer price : 39.576 GWei
value           : 0 GWei
max gas price   : 63321600 GWei
max gas price   : 168.21 USD

Waiting Logs
----------------------------------------------------

Receipt Details
----------------------------------------------------
status          : 1
gas used        : 144945
gas price       : 39.576 GWei
gas price       : 0.00 USD
final gas cost  : 5736343.32 GWei
final gas cost  : 15.24 USD

Logs
----------------------------------------------------
EventLog
map[value:contract[840bb1050d4d79c998667c35eee4223ef97127b2] success[true] version[0.1.0]]

Validate Version and API
----------------------------------------------------
version         : 0.1.0
api             : 0x840bB1050D4d79c998667C35EEE4223Ef97127B2

Balance
----------------------------------------------------
balance before  : 1.157920892e+68 GWei
balance after   : 1.157920892e+68 GWei
balance diff    : 66490885.66 GWei
balance diff    : 176.63 USD

*/
