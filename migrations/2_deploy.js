const FighterTraits = artifacts.require("FighterTraits");

module.exports = async function (deployer) {
  await deployer.deploy(
    FighterTraits,
    "0x87E738a3d5E5345d6212D8982205A564289e6324",
    "0x2Ed251752DA7F24F33CFbd38438748BB8eeb44e1"
  );
};
