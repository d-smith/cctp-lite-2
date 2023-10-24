const ethDeploy = require("../broadcast/mbdeploy.s.sol/1281/run-latest.json");

const main = async () => {
    //console.log(ethDeploy.transactions);
    fiddyAddress =
        ethDeploy.transactions.filter(t => t.contractName == "FiddyCent")
            .map(t => t.contractAddress)[0];
    console.log(`FiddyCent deployed at ${fiddyAddress}`);
    console.log(`export MB_FIDDY_CENT=${fiddyAddress}`)
}

main();