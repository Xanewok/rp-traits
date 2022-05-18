const assert = require("assert");
const fs = require('fs');

const FighterTraits = artifacts.require("FighterTraits");

function pickTraits(obj) {
  const attrs = ["class", "body", "weapon", "hair", "hairColor", "back", "aura", "top", "bottom"]
  const ret = {};
  for (const key of attrs) {
    ret[key] = Number(obj[key]);
  }
  return ret;
}

contract("FighterTraits", (accounts) => {
  it("calculated traits matches those served by the API", async () => {
    // NOTE: Make sure you run `genFightersData` script first to generate this file
    const fighters = JSON.parse(fs.readFileSync('./resource/fighters.json'));
    const limited = await FighterTraits.deployed();

    for (const [id, fighter] of Object.entries(fighters)) {
      if (id % 1000 == 0) {
        console.log({ id, fighter: { seed: fighter.seed } });
      }
      const traits = await limited.getStartingTraitsBySeed(fighter.seed);
      try {
        assert.deepEqual(fighter.traits, pickTraits(traits))
      } catch (e) {
        // NOTE: There's a bug where Overgrown hair has both ID 0 and 9 - simply
        // check if everything matches except for that
        if ((fighter.traits.hair == 0 && traits.hair == 9) || (fighter.traits.hair == 9 && traits.hair == 0)) {
          fighter.traits.hair = traits.hair;
          assert.deepEqual(fighter.traits, pickTraits(traits))
        } else {
          throw e;
        }
      }
    }
  });
});
