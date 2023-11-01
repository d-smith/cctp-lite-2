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
$ go run main.go ethBalances $ACCT1
ETH Balance: 10000000000000000000000
Fiddy Balance: 0
$ go run main.go mbBalances $MBACCT1
ETH Balance: 1208925819614629174706176
Fiddy Balance: 0
```

```console
$ # Use the faucet to obtain some Fiddy on the Ethereum side
$ go run main.go ethDrip $ACCT1 100
nonce 4
Dripped 100 to 0x9949f7e672a568bB3EBEB777D5e8D1c1107e96E5: txn id 0xf6a04f0c2b2613347ebd96588028fc6d03984cc36066ccffeaf0a3e6e8a81fe4
$ go run main.go ethBalances $ACCT1
ETH Balance: 10000000000000000000000
Fiddy Balance: 100
```

```console
$ # Authorize the transport contract to spend Fiddy on behalf of ACCT1
$ go run main.go ethAllowance $ACCT1
Allowance: 0
$ go run main.go ethApprove $ACCT1KEY 25
nonce 0
Approved 25: txn id 0x8fb8840e347838277d373a76c63a496a840252cfa088af9ce76761e3a5f7ff71
$ go run main.go ethAllowance $ACCT1
Allowance: 25
```

```console
$ go run main.go ethDeposit4Burn $ACCT1KEY $MBACCT1 3
Deposited 3: txn hash 0xbbc040b7eb5fcdabe1460215b2bfa41a6365467afe0e6cac7ef11beb20ed14b2
```

```console
$ # Grab the env var export from the listener console and set the vars
$ export ATTESTOR_SIG=0x91a8b4c77e19a30fe80936eae954a4a67d0edf6cd9959f79aa2b5cd24a3f4f4b61cd1bb7828667ecd9b55cd8464131c4664b0d9410df0a86ccc8e6179d3f11471c
$ export MSG=0x00000001000000010000000200000000000000000000000000000000000000009949f7e672a568bb3ebeb777d5e8d1c1107e96e50000000000000000000000003cd0a705a2dc65e5b1e1205896baa2be8a07c6e000000001000000000000000000000000c0a4b9e04fb55b1b498c634faeeb7c8dd5895b530000000000000000000000003cd0a705a2dc65e5b1e1205896baa2be8a07c6e000000000000000000000000000000000000000000000000000000000000000030000000000000000000000009949f7e672a568bb3ebeb777d5e8d1c1107e96e5
```

```console
$ # Claim that Fiddy for the recipient address on Moonbeam!
$ go run main.go mbMintFromBurned $MBACCT1KEY $MSG $ATTESTOR_SIG
Minted: txn hash 0x7eaa8623ec02317efa0bf61dceb479058dbba039bb78e6a7dff418f1421bf9ce
```

```console
$ go run main.go ethBalances $ACCT1
ETH Balance: 9999999795971045943794
Fiddy Balance: 97
$ go run main.go mbBalances $MBACCT1
ETH Balance: 1208925818405840093529545
Fiddy Balance: 3
```





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

