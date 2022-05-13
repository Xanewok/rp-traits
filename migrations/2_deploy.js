const GoMathRand = artifacts.require('GoMathRand');

module.exports = async function (deployer) {
    await deployer.deploy(GoMathRand);
};
