const ethDeploy = require("../broadcast/deploy.s.sol/31337/run-latest.json");

const main = async () => {
    //console.log(ethDeploy.transactions);
    fiddyAddress =
        ethDeploy.transactions.filter(t => t.contractName == "FiddyCent")
            .map(t => t.contractAddress)[0];
    console.log(`FiddyCent deployed at ${fiddyAddress}`);
    console.log(`export FIDDY_CENT=${fiddyAddress}`)
}

main();