package main

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ardanlabs/smartcontract/app/bank/proxy/contract/go/bank"
	"github.com/ethereum/go-ethereum/common"
)

const (
	ownerStoreFile    = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	account1StoreFile = "zarf/ethereum/keystore/UTC--2022-05-13T16-57-20.203544000Z--8e113078adf6888b7ba84967f299f29aece24c55"
	account2StoreFile = "zarf/ethereum/keystore/UTC--2022-05-13T16-59-42.277071000Z--0070742ff6003c3e809e78d524f0fe5dcc5ba7f7"
	account3StoreFile = "zarf/ethereum/keystore/UTC--2022-09-16T16-13-42.375710134Z--7fdfc99999f1760e8dbd75a480b93c7b8386b79a"
	account4StoreFile = "zarf/ethereum/keystore/UTC--2022-09-16T16-13-55.707637523Z--000cf95cb5eb168f57d0befcdf6a201e3e1acea9"

	passPhrase       = "123" // All three accounts use the same passphrase
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

	// Deposit will fail if withdrawing account as insufficient balance
	// So, firsr run make geth-deposit
	/*
	   $ make geth-deposit
	   makefile:162: warning: overriding recipe for target 'basic-test'
	   makefile:125: warning: ignoring old recipe for target 'basic-test'
	   curl -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"eth_sendTransaction", "params": [{"from":"0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd", "to":"0x8E113078ADF6888B7ba84967F299F29AeCe24c55", "value":"0x1000000000000000000"}], "id":1}' localhost:8545
	   {"jsonrpc":"2.0","id":1,"result":"0x721042bdc46b0e054b983a8cc76047c06e6366ed474c0ba4e8269242f112fe87"}
	   curl -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"eth_sendTransaction", "params": [{"from":"0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd", "to":"0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7", "value":"0x1000000000000000000"}], "id":1}' localhost:8545
	   {"jsonrpc":"2.0","id":1,"result":"0xa1593cbda94d7389040626525c17a2e4a60cd39713ec2ffcabd8d32200514d31"}
	   curl -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"eth_sendTransaction", "params": [{"from":"0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd", "to":"0x7FDFc99999f1760e8dBd75a480B93c7B8386B79a", "value":"0x1000000000000000000"}], "id":1}' localhost:8545
	   {"jsonrpc":"2.0","id":1,"result":"0xa35f85b0d42fd6d072bfab4590ab16ab14f87cf72e17694f1b560aed96f522c9"}
	   curl -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"eth_sendTransaction", "params": [{"from":"0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd", "to":"0x000cF95cB5Eb168F57D0bEFcdf6A201e3E1acea9", "value":"0x1000000000000000000"}], "id":1}' localhost:8545
	   {"jsonrpc":"2.0","id":1,"result":"0xcaab05dc27231a1c394642c6c255756a99b5175c886d401934299eeea7f91cd5"}
	*/

	depositAmount := os.Getenv("DEPOSIT_AMOUNT") // env var in Makefile commands
	depositTarget := os.Getenv("DEPOSIT_TARGET") // env var in Makefile commands
	var ethAccount string

	// Validate the deposit target is valid.
	switch depositTarget {
	case "owner":
		ethAccount = ownerStoreFile
	case "account1":
		ethAccount = account1StoreFile // used for demo, see const above
	case "account2":
		ethAccount = account2StoreFile
	case "account3":
		ethAccount = account3StoreFile
	case "account4":
		ethAccount = account4StoreFile
	default:
		ethAccount = account1StoreFile
	}

	backend, err := ethereum.CreateDialedBackend(ctx, ethereum.NetworkLocalhost)
	if err != nil {
		return err
	}
	defer backend.Close()

	privateKey, err := ethereum.PrivateKeyByKeyFile(ethAccount, passPhrase)
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

	valueGwei, err := strconv.ParseFloat(depositAmount, 64)
	if err != nil {
		return fmt.Errorf("converting deposit amount to float: %v", err)
	}

	const gasLimit = 1600000
	const gasPriceGwei = 39.576
	tranOpts, err := clt.NewTransactOpts(ctx, gasLimit, currency.GWei2Wei(big.NewFloat(gasPriceGwei)), big.NewFloat(valueGwei))
	if err != nil {
		return err
	}

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

	proxyContract, err := bank.NewBank(common.HexToAddress(contractID), clt.Backend)
	if err != nil {
		return fmt.Errorf("new proxy connection: %w", err)
	}

	tx, err := proxyContract.Deposit(tranOpts)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransaction(tx))

	// =========================================================================

	receipt, err := clt.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}

/* Deposit 120000 to account1
$ make bank-proxy-deposit
makefile:162: warning: overriding recipe for target 'basic-test'
makefile:125: warning: ignoring old recipe for target 'basic-test'
DEPOSIT_TARGET="account1" DEPOSIT_AMOUNT="120000" CGO_ENABLED=0 go run app/bank/proxy/cmd/deposit/main.go

Input Values
----------------------------------------------------
fromAddress: 0x8E113078ADF6888B7ba84967F299F29AeCe24c55
oneETHToUSD: 2521.4702120453194
oneUSDToETH: 0.0003965940169441219
contractID: 0xA5B76e49bD18E952502f1eB4c4B281B91C727CBD

Transaction Details
----------------------------------------------------
hash            : 0xbcf671076a56ad6090777640ca9e28588cea2fc56eda8d67b8d7bee630ee4ce5
nonce           : 0
gas limit       : 1600000
gas offer price : 39.576 GWei
value           : 120000 GWei
max gas price   : 63441600 GWei
max gas price   : 159.97 USD

Receipt Details
----------------------------------------------------
status          : 1
gas used        : 145022
gas price       : 39.576 GWei
gas price       : 0.00 USD
final gas cost  : 5739390.672 GWei
final gas cost  : 14.47 USD

Logs
----------------------------------------------------
EventLog
map[value:deposit[8e113078adf6888b7ba84967f299f29aece24c55] balance[120000000000000]]
EventLog
map[value:success[true]]

Balance
----------------------------------------------------
balance before  : 4.722366483e+12 GWei
balance after   : 4.722360623e+12 GWei
balance diff    : 5859390.672 GWei
balance diff    : 14.77 USD
*/

/* Check account1 balance 120000000000000
$ make bank-proxy-balance
makefile:162: warning: overriding recipe for target 'basic-test'
makefile:125: warning: ignoring old recipe for target 'basic-test'
BALANCE_TARGET="account1" CGO_ENABLED=0 go run app/bank/proxy/cmd/balance/main.go

Input Values
----------------------------------------------------
fromAddress: 0x8E113078ADF6888B7ba84967F299F29AeCe24c55
oneETHToUSD: 2521.247380988553
oneUSDToETH: 0.00039662906842878354
contractID: 0xA5B76e49bD18E952502f1eB4c4B281B91C727CBD
account balance: 120000000000000
Balance
----------------------------------------------------
balance before  : 4.722360623e+12 GWei
balance after   : 4.722360623e+12 GWei
balance diff    : 0 GWei
balance diff    : 0.00 USD
*/
