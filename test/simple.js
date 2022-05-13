const assert = require("assert");

const GoMathRand = artifacts.require("GoMathRand");

contract("GoMathRand", (accounts) => {

  it("newRand", async () => {
    const rng = await GoMathRand.deployed();

    const myRng = await rng.newRand(911);
    const value1 = await rng.Intn(myRng, 35);
    assert.equal(value1, 33);
  });
});
