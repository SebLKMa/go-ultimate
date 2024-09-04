package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ardanlabs/smartcontract/app/basic/contract/go/basic"
	"github.com/ethereum/go-ethereum/common"
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

	converter, err := currency.NewConverter(basic.BasicMetaData.ABI, coinMarketCapKey)
	if err != nil {
		converter = currency.NewDefaultConverter(basic.BasicMetaData.ABI)
	}
	oneETHToUSD, oneUSDToETH := converter.Values()

	fmt.Println("oneETHToUSD:", oneETHToUSD)
	fmt.Println("oneUSDToETH:", oneUSDToETH)

	// =========================================================================

	contractIDBytes, err := os.ReadFile("zarf/ethereum/basic.cid")
	if err != nil {
		return fmt.Errorf("importing basic.cid file: %w", err)
	}

	contractID := string(contractIDBytes)
	if contractID == "" {
		return errors.New("need to export the basic.cid file")
	}
	fmt.Println("contractID:", contractID)

	storeCon, err := basic.NewBasic(common.HexToAddress(contractID), clt.Backend)
	if err != nil {
		return fmt.Errorf("new contract: %w", err)
	}

	version, err := storeCon.Version(nil)
	if err != nil {
		return err
	}
	fmt.Println("version:", version)

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

	key := "bill"
	value := big.NewInt(1_000_000)

	tx, err := storeCon.SetItem(tranOpts, key, value)
	if err != nil {
		log.Fatal("SetItem ERROR:", err)
	}
	fmt.Print(converter.FmtTransaction(tx))

	receipt, err := clt.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}

/* Sample output:
$ make basic-write
CGO_ENABLED=0 go run app/basic/cmd/write/main.go

Input Values
----------------------------------------------------
fromAddress: 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd
oneETHToUSD: 2546.094722667976
oneUSDToETH: 0.00039275836483888946
contractID: 0x531130464929826c57BBBF989e44085a02eeB120
version: 1.1

Transaction Details
----------------------------------------------------
hash            : 0xa93e1e747a9dd1f8988f428fa8378741b1fbe7e87ca0326fdd95123014737341
nonce           : 1
gas limit       : 1600000
gas offer price : 39.576 GWei
value           : 0 GWei
max gas price   : 63321600 GWei
max gas price   : 161.22 USD

Receipt Details
----------------------------------------------------
status          : 1
gas used        : 47579
gas price       : 39.576 GWei
gas price       : 0.00 USD
final gas cost  : 1882986.504 GWei
final gas cost  : 4.79 USD

Logs
----------------------------------------------------
ItemSet
map[key:bill value:1000000]

Balance
----------------------------------------------------
balance before  : 1.157920892e+68 GWei
balance after   : 1.157920892e+68 GWei
balance diff    : 1882986.504 GWei
balance diff    : 4.79 USD
*/
