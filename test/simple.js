const assert = require("assert");

const GoMathRand = artifacts.require("GoMathRand");
const GoMathRandLimited = artifacts.require("GoMathRandLimited");

function pickTraits(obj) {
  const attrs = ["class", "body", "weapon", "hair", "hairColor", "back", "aura", "top", "bottom"]
  const ret = {};
  for (const key of attrs) {
    ret[key] = obj[key];
  }
  return ret;
}

contract("GoMathRand", (accounts) => {

  it("newRand", async () => {
    const limited = await GoMathRandLimited.deployed();
    const fighters = {
      1918: {
        seed: "2539103265505867606901872458139978623723014567987151303084214062672181312887",
        traits: {
          class: '0',
          body: '2', weapon: '2',
          hair: '5', hairColor: '3',
          back: '0', aura: '0',
          top: '7', bottom: '1'
        }
      },
      20801: {
        seed: "103256593631039286645257230999332642700195251781097825023731579771058446619717",
        traits: {
          class: '4',
          body: '2', weapon: '8',
          hair: '16', hairColor: '3',
          back: '18', aura: '3',
          top: '8', bottom: '3'

        }
      }
    };
    for (const fighter of Object.values(fighters)) {
      const traits = await limited.getTraits(fighter.seed);
      assert.deepEqual(fighter.traits, pickTraits(traits))
    }

    const call1 = await limited.storeTraits(fighters[20801].seed);
    console.log({ gasUsed: call1?.receipt?.gasUsed });

  });
});
