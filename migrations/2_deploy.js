const GoMathRand = artifacts.require('GoMathRand');
const GoMathRandLimited = artifacts.require('GoMathRandLimited');

module.exports = async function (deployer) {
    // await deployer.deploy(GoMathRand);
    await deployer.deploy(GoMathRandLimited);
};
