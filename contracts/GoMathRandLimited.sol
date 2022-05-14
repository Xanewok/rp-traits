// SPDX-License-Identifier: BSD-3-Clause
pragma solidity ^0.8;

// https://cs.opensource.google/go/go/+/master:src/math/rand/rand.go;l=5;drc=690ac4071fa3e07113bf371c9e74394ab54d6749
contract GoMathRandLimited {
    int64[] public values = new int64[](10);

    function generateSingle(uint256 seed, int32 n) public {
        Source memory rng = newRand(seed);
        values[0] = Int31n(rng, n);
    }

    function generateMany(uint256 seed, int32 n) public {
        Source memory rng = newRand(seed);
        for (uint256 i = 0; i < 10; i++) {
            values[i] = Int31n(rng, n);
        }
    }

    function computeMany(uint256 seed, int32 n) public pure {
        Source memory rng = newRand(seed);
        Int31n(rng, n);
        Int31n(rng, n);
        Int31n(rng, n);
        Int31n(rng, n);
        Int31n(rng, n);
        Int31n(rng, n);
        Int31n(rng, n);
        Int31n(rng, n);
        Int31n(rng, n);
        Int31n(rng, n);
    }

    // rand.go

    /// @dev Equivalent of Go's `rand.New(rand.NewSource(seed.Int64))`
    /// We only support `rand.Intn`, so we use `Source` directly
    function newRand(uint256 seed) public pure returns (Source memory) {
        // https://etherscan.io/address/0x2ed251752da7f24f33cfbd38438748bb8eeb44e1#readContract
        // seed = getSeed(
        //   origin = Summon@0x87e738a3d5e5345d6212d8982205a564289e6324,
        //   identifier
        // )
        // Equivalent of `seed.Int64()`
        // https://cs.opensource.google/go/go/+/master:src/math/big/int.go;l=368;drc=831f1168289e65a7ef49942ad8d16cf14af2ef43
        // Takes the least significant 64 bits of x and then preserves the sign
        // NOTE: Go's `math/big` uses little-endian encoding, so does Solidity
        // for the numbers so just use the truncating conversion directly
        uint64 truncatedSeed = uint64(seed);
        // Equivalent of `rand.NewSource(..)`
        Source memory source = NewSource(int64(truncatedSeed));
        // We only care about `rand.Intn` so we use the underlying `Source` directly
        return source;
    }

    // https://cs.opensource.google/go/go/+/master:src/math/rand/rand.go;l=43;drc=690ac4071fa3e07113bf371c9e74394ab54d6749
    function NewSource(int64 seed) public pure returns (Source memory) {
        Source memory source;
        Seed(source, seed);
        return source;
    }

    // Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n).
    // It panics if n <= 0.
    function Int31n(Source memory rng, int32 n) public pure returns (int32) {
        unchecked {
            require(n > 0, "invalid argument to Int31n");
            if (n & (n - 1) == 0) {
                // n is power of two, can mask
                return Int31(rng) & (n - 1);
            }
            int32 max = int32((1 << 31) - 1 - ((1 << 31) % uint32(n)));
            int32 v = Int31(rng);
            while (v > max) {
                v = Int31(rng);
            }
            return v % n;
        }
    }

    function Int31(Source memory rng) public pure returns (int32) {
        return int32(Int63(rng) >> 32);
    }

    // rng.go

    // NOTE: This is a modified version that allows only for a stream of up to
    // 10 random numbers to save the PRNG initialization gas cost.

    uint16 constant RNG_LEN = 607;
    uint16 constant RNG_TAP = 273;
    uint64 constant RNG_MASK = uint64(type(int64).max);
    int32 constant int32max = type(int32).max;

    uint8 constant RNG_COUNT = 10;

    struct Source {
        uint16 tap;
        uint16 feed;
        // int64[RNG_LEN] vec;
        int64[RNG_COUNT * 2] vec;
    }

    // https://cs.opensource.google/go/go/+/master:src/math/rand/rng.go;l=204;drc=2bea43b0e7f3e636ffc8239f9d3fccdd5d763c8b
    // NOTE: We assume seed is not zero
    function Seed(Source memory rng, int64 seed) internal pure {
        rng.tap = 0;
        rng.feed = RNG_COUNT;
        // rng.feed = RNG_LEN - RNG_TAP;

        unchecked {
            seed = seed % int32max;
            if (seed < 0) {
                seed += int32max;
            }
            if (seed == 0) {
                seed = 89482311;
            }
        }

        // We keep the seed in a full word instead to save on constant widening
        uint256 x = uint64(seed);
        uint256 u;
        unchecked {
            // NOTE: This is split into two loops comparing to the original to save
            // on type conversions
            // We're dealing with Lehmer (multiplicative congruential) generator,
            // so we can amortize some of the computations due to the fact that
            // x_i = a^i * x_0 mod m = (a^i mod m) * x_0 mod m
            // where x_0 is the seed, a is the multiplier and m is the modulus.
            // Originally we were calling it initially 20 times, so pre-compute
            // the amortized multiplier with 48271^20 mod 0x7fffffff = 2075782095.
            assembly {
                x := mulmod(x, 2075782095, 0x7fffffff)
            }
            // Then, we optimize by only computing necessary generator state for
            // up to a given random number count (RNG_COUNT). Because of this and
            // because the original algorithm started with numbers from the middle
            // of the pre-cooked values, simply skip the initial phases, which
            // only internally generated random state.
            // Here we skip 324 iterations of generating 2 values, so
            // 48271^(3*324) mod 0x7fffffff = 750037089. (324 = RNG_LEN - RNG_TAP - RNG_COUNT)
            assembly {
                x := mulmod(x, 750037089, 0x7fffffff)
            }
            // Then, we process the part of the generator array originally read
            // by the `feed` cursor (starting from index = RNG_LEN - RNG_TAP - RNG_COUNT)
            int64[10] memory RNG_COOKED_FEED = [
                -6564663803938238204,
                -8060058171802589521,
                581945337509520675,
                3648778920718647903,
                -4799698790548231394,
                -7602572252857820065,
                220828013409515943,
                -1072987336855386047,
                4287360518296753003,
                -4633371852008891965
            ];
            for (uint256 i = 0; i < RNG_COUNT; i++) {
                x = seedrand(x);

                u = x << 40;
                x = seedrand(x);
                u ^= x << 20;
                x = seedrand(x);
                u ^= x;
                u ^= uint64(RNG_COOKED_FEED[i]);

                rng.vec[i] = int64(uint64(u));
            }
            // Again, we skip again the unnedeed feedback register values...
            // 48271^(3*263) mod 0x7fffffff = 1483819319.
            assembly {
                x := mulmod(x, 1483819319, 0x7fffffff)
            }
            // And finally we read the last values originally read by the `tap`
            // cursor (starting from index = RNG_LEN - RNG_COUNT)
            int64[10] memory RNG_COOKED_TAP = [
                -6344160503358350167,
                5896236396443472108,
                -758328221503023383,
                -1894351639983151068,
                -307900319840287220,
                -6278469401177312761,
                -2171292963361310674,
                8382142935188824023,
                9103922860780351547,
                4152330101494654406
            ];
            for (uint256 i = 0; i < RNG_COUNT; i++) {
                x = seedrand(x);
                u = x << 40;
                x = seedrand(x);
                u ^= x << 20;
                x = seedrand(x);
                u ^= x;
                u ^= uint64(RNG_COOKED_TAP[i]);

                rng.vec[RNG_COUNT + i] = int64(uint64(u));
            }
        }
    }

    // https://en.wikipedia.org/wiki/Lehmer_random_number_generator
    // seed rng x[n+1] = 48271 * x[n] mod (2**31 - 1)
    function seedrand(uint256 x) internal pure returns (uint256 r) {
        assembly {
            r := mulmod(x, 48271, 0x7fffffff)
        }
    }

    function Uint64(Source memory rng) public pure returns (uint64) {
        unchecked {
            if (rng.tap == 0) {
                // rng.tap = RNG_LEN - 1;
                rng.tap = (2 * RNG_COUNT) - 1;
            } else {
                rng.tap--;
            }

            if (rng.feed == 0) {
                // rng.feed = RNG_LEN - 1;
                rng.feed = (2 * RNG_COUNT) - 1;
            } else {
                rng.feed--;
            }

            // NOTE: Go version relies on wrapping arithmetic
            int64 x = rng.vec[rng.feed] + rng.vec[rng.tap];
            rng.vec[rng.feed] = x;
            return uint64(x);
        }
    }

    // Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
    function Int63(Source memory rng) public pure returns (int64) {
        return int64(Uint64(rng) & RNG_MASK);
    }
}
