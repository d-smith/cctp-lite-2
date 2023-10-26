# cctp-lite-2

Sample project implementing a lightweight fascimile of CCTP.

## Deploy and Run

Local ethereum node:

```
cd anvil
./start-anvil.sh
```

Local moonbeam:

```
docker pull purestake/moonbeam:v0.32.2
docker run --rm --name moonbeam_development -p 9944:9944 -p 9933:9933 purestake/moonbeam:v0.32.2 --dev --ws-external --rpc-external 

```

Deploy the contracts to eth:

```
forge script script/deploy.s.sol:DeployScript --broadcast --rpc-url http://127.0.0.1:8545 --extra-output-files abi --extra-output-files bin
```

Deploy the contracts to moonbeam:

```
forge script script/mbdeploy.s.sol:DeployScript --broadcast --rpc-url http://127.0.0.1:9933 --legacy --extra-output-files abi --extra-output-files bin
```

To do some FiddyCent token operations, run `scripts\deploy-details.sh` to get the contract address and an export command. Run the export command, then play with the contract.

```
cast call $FIDDY_CENT "totalSupply()(uint256)"  --rpc-url  http://127.0.0.1:8545

cast send $FIDDY_CENT "transfer(address,uint256)" --private-key $DEPLOYER_KEY $ACCT1 50

cast call $FIDDY_CENT "balanceOf(address)" $ACCT1

cast call $MB_FIDDY_CENT "totalSupply()(uint256)" --rpc-url  http://127.0.0.1:9933

cast call $MB_FIDDY_CENT "balanceOf(address)" $MB_ACCT1 --rpc-url  http://127.0.0.1:9933

cast send $MB_FIDDY_CENT "transfer(address,uint256)" --private-key $MB_DEPLOYER_KEY $MB_ACCT1 50 --rpc-url  http://127.0.0.1:9933



```

## Misc

Bootstrapping the project: `forge init cctp-lite-2`

OpenZeppelin dependencies: `forge install OpenZeppelin/openzeppelin-contracts`

Updating foundry tool - `foundryup`