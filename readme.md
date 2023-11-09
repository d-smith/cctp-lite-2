# cctp-lite-2

Sample project implementing a lightweight fascimile of CCTP.

Note: [this project](https://github.com/d-smith/cctpcli) provides a CLI for interacting with the contracts deployed by this project.

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


### Metamask Integration

Rough notes

* Add local network and moonbeam local network in network settings
  * Settings -> Networks -> Add a Network -> Add a network manually

  For local eth:

```Network name
Ethereum Dev Node
New RPC URL
http://localhost:8545
Chain ID
1337
Currency symbol
ETH
```

For local moonbeam:

```
Network name
Moonbeam Dev Node
New RPC URL
http://localhost:9933
Chain ID
1281
Currency symbol
ETH
```

To add dev accounts to metamask, use the private keys from the env - just import them and add a name.


