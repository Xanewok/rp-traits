// A script that's used to scrape the traits from the API endpoint; mainly to
// create test cases for the test/simple.js test.
const fetch = require("node-fetch");
const Web3 = require("web3");

const traits = require("./traits.json");
const allTraits = Object.values(traits).flatMap((val) => val);

const namesToIds = {};
for (const value of allTraits) {
  if (!namesToIds[value.feature]) {
    namesToIds[value.feature] = {};
  }
  namesToIds[value.feature][value.name] = value;
}

const FIGHTER_TOKEN = "0x87E738a3d5E5345d6212D8982205A564289e6324";
const SEEDER_V2 = "0x2Ed251752DA7F24F33CFbd38438748BB8eeb44e1";

async function main() {
  // Configuring the connection to an Ethereum node
  const network = process.env.ETHEREUM_NETWORK;
  const web3 = new Web3(
    new Web3.providers.HttpProvider(
      `https://${network}.infura.io/v3/${process.env.INFURA_PROJECT_ID}`
    )
  );

  const contract = new web3.eth.Contract(
    [
      {
        inputs: [
          { internalType: "address", name: "origin", type: "address" },
          { internalType: "uint256", name: "identifier", type: "uint256" },
        ],
        name: "getSeed",
        outputs: [{ internalType: "uint256", name: "", type: "uint256" }],
        stateMutability: "view",
        type: "function",
      },
    ],
    SEEDER_V2
  );

  const testCases = {};
  for (let id = 1; id <= 37202; id++) {
    const seed = await contract.methods.getSeed(FIGHTER_TOKEN, id).call();

    const res = await fetch(`https://api.raid.party/metadata/fighter/${id}`);
    const { attributes } = await res.json();
    const traits = {};
    for (const value of attributes) {
      if (value.trait_type != "Damage" && value.trait_type != "Enhancement") {
        const key =
          value.trait_type.slice(0, 1).toLowerCase() +
          value.trait_type.slice(1);
        const mappedId = namesToIds[key][value.value].id;

        traits[key] = mappedId;
      }
    }
    testCases[id] = { seed, traits };
    // Stream values instead of storing all of it in memory in case of a spurious HTTP failure
    console.log(`\"${id}\": ${JSON.stringify({ seed, traits })} ,`);
  }
}
main();
