# FighterTraits

An on-chain implementation of the random trait generation logic used by the Raid Party.

## Additional `/resource`s
- /reference.go - The reference Go implementation used by the API endpoint to calculate traits
- /traits.json - The trait database used by the generation algorithm
- /genFightersData.js - A script used to scrape the endpoint and convert traits to their numerical IDs
- /traits_go_sorted.json - The same contents as traits.json but sorted by the `weight` key using Go's `sort.Slice` (as used by the https://github.com/mroth/weightedrand)
- /inline-trait-selection - A Rust crate that pre-generates Solidity binary lookup functions for the sorted partial weight sums

## Context
The way trait calculation works righ now is as follows:
- When minting a fighter, we request a seed from the `Seeder` contract
- After seed batch arrives, a given fighter token is assigned a `uint256` that serves as the "seed"
- Then, the API endpoint initializes a Go's `math/rand` PRNG with that seed and uses the weighted random algorithm to pick traits in a pre-determined order

To start migrating that data on-chain and to prove fairness of this random trait generation, a simple re-implementation of Go's `math/rand` is used, along with pre-computed weighted random generation for every given trait.

Firstly, Go's `math/rand` uses a simple Lehmer RNG with parameters m = 2^31 - 1 and a = 48271 (the same as C+11's `minstd_rand`) combined with a xorshift-like algo and backed by a pre-cooked, 607-element long feedback register.

We only ever need 9 random numbers, so to save on gas, the original implementation was modified to only be able to generate up to 12 random numbers (to keep a buffer in case the PRNG needs to reroll whenever a generated number ends up in a modulo overflow when generating a number in a bounded range). This limitation allows us to skip initializing the entire register and only do that 2*12 times, while relying on Lehmer's parameters to do a pre-computed modular multiplication to skip over entire elements in constant time.

Additionally, we take the pre-sorted array in the resulting order of Go's `sort.Slice` by `weight`, calculate a partial sum and then bake the pre-generated binary lookup as different functions.
There's currently no way to efficiently use constant lookup tables, see https://github.com/ethereum/solidity/issues/12821 for more work.
