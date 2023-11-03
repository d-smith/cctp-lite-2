# cctp-lite-2

Sample project implementing a lightweight fascimile of CCTP.

## Set Up

### Nodes

Start a local ethereum node:

```
cd anvil
./start-anvil.sh
```

Local moonbeam:

```
docker pull purestake/moonbeam:v0.32.2
docker run --rm --name moonbeam_development -p 9944:9944 -p 9933:9933 purestake/moonbeam:v0.32.2 --dev --ws-external --rpc-external 
```

### Contract Deployment

Deploy the contracts to eth:

```
forge script script/deploy.s.sol:DeployScript --broadcast --rpc-url http://127.0.0.1:8545 --extra-output-files abi --extra-output-files bin
```

Deploy the contracts to moonbeam:

```
forge script script/mbdeploy.s.sol:DeployScript --broadcast --rpc-url http://127.0.0.1:9933 --legacy --extra-output-files abi --extra-output-files bin
```

### Script Environment

Set the deployment-related environment variables by running the deploy details scripts to get the exports needed to run the scripts.

```
node script/deploy-details.js
node script/mbdeploy-details.js
```

The environment variables for the demo can be set sourcing the .env file in cctpcli. Note the environment variables from the deploy details scripts are also needed.

## Basic Demo

Note: This will be simplified later when the listener stores burn receipts for accounts, and the mint
from burn command is modified to lookup burn receipts for the account.

### Running the event listener

In a dedicated shell, set up the environment and run the ethereum network smart contract event listener:

```console
$ export FIDDY_ETH_ADDRESS=0xC0a4b9e04fB55B1b498c634FAEeb7C8dD5895b53
$ export TRANSPORTER=0xa7F08a6F40a00f4ba0eE5700F730421C5810f848
$ export FIDDY_MB_ADDRESS=0x970951a12F975E6762482ACA81E57D5A2A4e73F4
$ export MB_TRANSPORTER=0x3ed62137c5DB927cb137c26455969116BF0c23Cb
$ . .env
$ go run main.go runEthEventListener
```

The listener subscribes to the MessageSent event emitted by the smart contract when the Transporter burns
a deposit for mint on a remote chain, in this instance Fiddy tokens deposited and burned on the ethereum
side for transport to moonbeam.


### Transport FiddyCent from Ethereum to Moonbeam


```console
$ export FIDDY_ETH_ADDRESS=0xC0a4b9e04fB55B1b498c634FAEeb7C8dD5895b53
$ export TRANSPORTER=0xa7F08a6F40a00f4ba0eE5700F730421C5810f848
$ export FIDDY_MB_ADDRESS=0x970951a12F975E6762482ACA81E57D5A2A4e73F4
$ export MB_TRANSPORTER=0x3ed62137c5DB927cb137c26455969116BF0c23Cb
$ . .env

```

```console
$ # Show starting balances
$ cctpcli ethBalances $ACCT3
ETH Balance: 10000000000000000000000
Fiddy Balance: 0
No claims found for address

$ cctpcli mbBalances $MBACCT3
ETH Balance: 1208925819614629174706176
Fiddy Balance: 0
No claims found for address
```

```console
$ # Use the faucet to obtain some Fiddy on the Ethereum side
$ cctpcli ethDrip $ACCT3 50
Dripped 50 to 0x7bA7d161F9E8B707694f434d65c218a1F0853B1C: txn id 0x6ceaa362fac94e32e834745d06e50ed9569c5621b9b387603d93a4d8cd85eeed

$ cctpcli ethBalances $ACCT3
ETH Balance: 10000000000000000000000
Fiddy Balance: 50
```

```console
$ # Authorize the transport contract to spend Fiddy on behalf of ACCT1
$ cctpcli ethAllowance $ACCT3
Allowance: 0

$ go run main.go ethApprove $ACCT1KEY 25
Approved 25: txn id 0x8fb8840e347838277d373a76c63a496a840252cfa088af9ce76761e3a5f7ff71

$ cctpcli ethAllowance $ACCT3
Allowance: 25
```

```console
$ # Deposit 5 Fiddy on the Ethereum side for transport to Moonbeam
$ cctpcli ethDeposit4Burn $ACCT3KEY $MBACCT3 5
Deposited 5: txn hash 0x957bfe8a77ef16cd3f67077bc645f24966bffdfdb233d90bb1ebb132372e9325

$ cctpcli ethDeposit4Burn $ACCT3KEY $MBACCT3 10
Deposited 10: txn hash 0x5efcc307b5df7b3312086c538ed14b4d65b02c20df93d933ca46d9b47234ae50
```

```console
$ # View available claims on the Moonbeam side
$ cctpcli mbBalances $MBACCT3
ETH Balance: 1208925819614629174706176
Fiddy Balance: 0
Claims:
  Claim id 1 :: Source domain 1 -> Destination domain 2, Claimable Amount 5
  Claim id 2 :: Source domain 1 -> Destination domain 2, Claimable Amount 10
```

```console
$ # Claim 10 Fiddy on the Moonbeam side
$ cctpcli mbMintFromBurned $MBACCT3 $MBACCT3KEY 2
Minted: txn hash 0x036f2d3cd8500cf4d48c37a89c38cc8556cc466475c9577f6cb50afb655c0f44

$ cctpcli mbBalances $MBACCT3
ETH Balance: 1208925818413956864309442
Fiddy Balance: 10
Claims:
  Claim id 1 :: Source domain 1 -> Destination domain 2, Claimable Amount 5
```

```console
$ # Show final balances
$ cctpcli ethBalances $ACCT3
ETH Balance: 9999999829348848426872
Fiddy Balance: 35
No claims found for address


$ cctpcli mbBalances $MBACCT3
ETH Balance: 1208925818413956864309442
Fiddy Balance: 10
Claims:
  Claim id 1 :: Source domain 1 -> Destination domain 2, Claimable Amount 5
```


## Refining the Implementation

1. Initial implementation - receive smart contract MessageSent event, sign an attestation and emit it to stdout. Use the output with the command line to claim the tokens on the remote chain.
2. (Current) Modify the event listener to store the attestation in a database by invoking an API. Add a command to claim the tokens on the remote chain using the attestation id.
3. Modify the API to store the message details in the data with producing an attestation signature. Modify the retrieval API to check the number of confirmations of the transaction, and return the attestation only after a threshold of confirmations is reached.


## Misc

Create the db in the cmd directory `sqlite3 attestor.db < att.sql`

To do some FiddyCent token operations, run `scripts\deploy-details.sh` to get the contract address and an export command. Run the export command, then play with the contract.

```
cast call $FIDDY_CENT "totalSupply()(uint256)"  --rpc-url  http://127.0.0.1:8545

cast send $FIDDY_CENT "transfer(address,uint256)" --private-key $DEPLOYER_KEY $ACCT1 50

cast call $FIDDY_CENT "balanceOf(address)" $ACCT1

cast call $MB_FIDDY_CENT "totalSupply()(uint256)" --rpc-url  http://127.0.0.1:9933

cast call $MB_FIDDY_CENT "balanceOf(address)" $MB_ACCT1 --rpc-url  http://127.0.0.1:9933

cast send $MB_FIDDY_CENT "transfer(address,uint256)" --private-key $MB_DEPLOYER_KEY $MB_ACCT1 50 --rpc-url  http://127.0.0.1:9933



```


Bootstrapping the project: `forge init cctp-lite-2`

OpenZeppelin dependencies: `forge install OpenZeppelin/openzeppelin-contracts`

Updating foundry tool - `foundryup`

Adding cli commands: `cobra-cli add command`


go run main.go ethBalances $ACCT1
go run main.go ethDrip $ACCT1 100
go run main.go ethBalances $ACCT1
go run main.go ethAllowance $ACCT1
go run main.go ethApprove $ACCT1KEY 25
go run main.go ethAllowance $ACCT1
go run main.go ethDeposit4Burn $ACCT1KEY $MBACCT1 5

