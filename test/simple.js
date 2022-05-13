const assert = require("assert");

const GoMathRand = artifacts.require("GoMathRand");
const GoMathRandLimited = artifacts.require("GoMathRandLimited");

contract("GoMathRand", (accounts) => {

  it("newRand", async () => {
    const rng = await GoMathRand.deployed();
    const limited = await GoMathRandLimited.deployed();

    const myRng = await rng.newRand(911);
    const value1 = await rng.Intn(myRng, 35);
    assert.equal(value1, 33);

    const myLimited = await limited.newRand(911);
    const value1Limited = await limited.Int31n(myLimited, 35);
    assert.equal(value1Limited, 33);
    await limited.generateMany(911, 35);
    let expected;
    expected = [
      33,
      22,
      22,
      16,
      17,
      29,
      13,
      7,
      18,
      1,
    ];
    for (let i = 0; i < 10; i++) {
      const value = `${await limited.values(i)}`;
      assert.equal(value, expected[i])
    }
    const { receipt } = await limited.generateMany("105779926529366228504990970003713286107530024193944566341142813727459338091771", 50000000);
    console.log({ gasUsed: receipt.gasUsed });

    expected = [
      13875352,
      33488040,
      3075775,
      34509352,
      1990175,
      14640839,
      21896810,
      28369212,
      46008401,
      36697042,
    ];
    for (let i = 0; i < 10; i++) {
      const value = `${await limited.values(i)}`;
      assert.equal(value, expected[i])
    }

  });
});
