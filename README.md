# xk6-random

xk6-random is an extension for [k6](https://k6.io), providing advanced random number generation capabilities.
It supports various distributions, array shuffling, and picking random elements from arrays.

```javascript
import { shuffle, Random } from 'k6/x/random';

export default function () {
    // Create a new Random generator
    let rng = new Random();

    // Shuffle an array
    let colors = ['red', 'green', 'blue', 'yellow'];
    shuffle(colors);
    console.log(colors);

    // Pick a random element from an array
    let pick = rng.pick(colors);
    console.log(`picked color: ${pick}`);
    
    // Generate a random integer between 1 and 10
    let int = rng.intBetween(1, 10);
    console.log(`random integer: ${int}`);
}
```

## Features

* Generation of random integers and floating-point numbers.
* Support for different distributions: uniform, normal (Gaussian), log-normal, Bernoulli, binomial, geometric, and exponential.
* Functions to shuffle arrays, generate random permutations, and pick random elements (including weighted picking).

## Installation

To build a k6 binary with the xk6-random extension, first ensure you have the prerequisites:

* Go installed (version 1.17 or later).
* Git installed.

Then, install xk6:

```bash
go install go.k6.io/xk6/cmd/xk6@latest
```

Build the binary with the xk6-random extension:

```bash
xk6 build --with github.com/oleiade/xk6-random@latest
```

## Usage

After building the k6 binary with the xk6-random extension, you can use it in your k6 scripts:

```javascript
import {shuffle, Random} from 'k6/x/random';

// Example usage
export default function () {
    // Create a new Random generator
    let rng = new Random();

    // Generate a random integer
    console.log(rng.int());

    // Generate a random float
    console.log(rng.float());

    // Generate a random boolean
    console.log(rng.boolean());

    // Generate a random number from a normal distribution
    console.log(rng.normal(0, 1));

    // Shuffle an array
    let array = [1, 2, 3, 4, 5];
    shuffle(array);
    console.log(array);
}
```

## API Reference

### Random Generator Methods

The `Random` class provides the following methods for generating random numbers:
* `int()`: Returns a random integer.
* `float()`: Returns a random real number.
* `boolean()`: Returns a random boolean value.
* `intBetween(min, max)`: Returns a random integer between min and max.
* `floatBetween(min, max)`: Returns a random real number between min and max.
* `pick(array)`: Returns a random element from the given array.
* `weightedPick(array, weights)`: Returns a random element from the given array, based on the given weights.
* `normal(mean, stdev)`: Returns a random real number from a normal distribution.
* `logNormal(mean, stdev)`: Returns a random real number from a log-normal distribution.
* `bernoulli(probability)`: Returns a random boolean value with the given probability.
* `binomial(trials, probability)`: Returns a random integer from a binomial distribution.
* `geometric(probability)`: Returns a random integer from a geometric distribution.
* `exponential(rate)`: Returns a random real number from an exponential distribution.
 
### Other Functions

Top-level functions are also provided for convenience:
* `permutation(n)`: Returns a random permutation of the integers [0, n).
* `shuffle(array)`: Shuffles the given array in-place.
* `shuffled(array)`: Returns a shuffled copy of the given array.