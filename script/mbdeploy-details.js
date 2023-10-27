const ethDeploy = require("../broadcast/mbdeploy.s.sol/1281/run-latest.json");

const main = async () => {
    //console.log(ethDeploy.transactions);
    fiddyAddress =
        ethDeploy.transactions.filter(t => t.contractName == "FiddyCent")
            .map(t => t.contractAddress)[0];
    console.log(`export FIDDY_MB_ADDRESS=${fiddyAddress}`)

    transporterAddress =
        ethDeploy.transactions.filter(t => t.contractName == "Transporter")
            .map(t => t.contractAddress)[0];
    console.log(`export MB_TRANSPORTER=${transporterAddress}`)
}

main();
