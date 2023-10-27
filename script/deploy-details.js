const ethDeploy = require("../broadcast/deploy.s.sol/31337/run-latest.json");

const main = async () => {
    //console.log(ethDeploy.transactions);
    fiddyAddress =
        ethDeploy.transactions.filter(t => t.contractName == "FiddyCent")
            .map(t => t.contractAddress)[0];
    console.log(`export FIDDY_ETH_ADDRESS=${fiddyAddress}`)

    transporterAddress =
        ethDeploy.transactions.filter(t => t.contractName == "Transporter")
            .map(t => t.contractAddress)[0];
    console.log(`export TRANSPORTER=${transporterAddress}`)
}

main();
