package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ardanlabs/smartcontract/app/bank/proxy/contract/go/bank"
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

	const gasLimit = 1800000 // increased from 1600000 due to contract creation out of gas
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

	if err := os.WriteFile("zarf/ethereum/bank.cid", []byte(address.Hex()), 0644); err != nil {
		return fmt.Errorf("exporting bank.cid file: %w", err)
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

	return nil
}

/* Out of gas err:

$ make bank-proxy-deploy
makefile:161: warning: overriding recipe for target 'basic-test'
makefile:124: warning: ignoring old recipe for target 'basic-test'
CGO_ENABLED=0 go run app/bank/proxy/cmd/deploy/bank/main.go

Input Values
----------------------------------------------------
fromAddress: 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd
oneETHToUSD: 2587.606793106456
oneUSDToETH: 0.00038645747980877994

Transaction Details
----------------------------------------------------
hash            : 0x679cd69d29cdd977b03aa23b5205dd60088d9f2261f15381db6e828d3e35c21a
nonce           : 2
gas limit       : 1600000
gas offer price : 39.576 GWei
value           : 0 GWei
max gas price   : 63321600 GWei
max gas price   : 163.85 USD

Contract Details
----------------------------------------------------
contract id     : 0x87A061ED19dcA76EC5B01643b054f5eae2730a85

Waiting Logs
----------------------------------------------------

Balance
----------------------------------------------------
balance before  : 1.157920892e+68 GWei
balance after   : 1.157920892e+68 GWei
balance diff    : 63321600 GWei
balance diff    : 163.85 USD
extracting tx error: contract creation code storage out of gas
exit status 1
make: *** [makefile:188: bank-proxy-deploy] Error 1
*/

/* Increase const gasLimit = 1800000 from 1600000
To do more research on gas price for contract creation,
https://support.metamask.io/transactions-and-gas/gas-fees/why-did-my-transaction-fail-with-an-out-of-gas-error-how-can-i-fix-it/
https://ethereum.stackexchange.com/questions/155241/despite-calculation-i-get-contract-creation-code-storage-out-of-gas
https://gist.github.com/miguelmota/117caf685b84cba8317f07e1ac6cd0da
https://ethereum.stackexchange.com/questions/19725/contract-creation-code-storage-out-of-gas
https://www.rareskills.io/post/smart-contract-creation-cost

$ make bank-proxy-deploy
makefile:162: warning: overriding recipe for target 'basic-test'
makefile:125: warning: ignoring old recipe for target 'basic-test'
CGO_ENABLED=0 go run app/bank/proxy/cmd/deploy/bank/main.go

Input Values
----------------------------------------------------
fromAddress: 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd
oneETHToUSD: 2690.309687582103
oneUSDToETH: 0.0003717044192405756

Transaction Details
----------------------------------------------------
hash            : 0x2de78ece2b82d4035e45fed036d8145eb19d85e205cbcb05ddce6fe39e3cbce3
nonce           : 6
gas limit       : 1800000
gas offer price : 39.576 GWei
value           : 0 GWei
max gas price   : 71236800 GWei
max gas price   : 191.65 USD

Contract Details
----------------------------------------------------
contract id     : 0xA5B76e49bD18E952502f1eB4c4B281B91C727CBD

Waiting Logs
----------------------------------------------------

Receipt Details
----------------------------------------------------
status          : 1
gas used        : 1610565
gas price       : 39.576 GWei
gas price       : 0.00 USD
final gas cost  : 63739720.44 GWei
final gas cost  : 171.48 USD

Logs
----------------------------------------------------

Balance
----------------------------------------------------
balance before  : 1.157920892e+68 GWei
balance after   : 1.157920892e+68 GWei
balance diff    : 63739720.44 GWei
balance diff    : 171.48 USD
*/
