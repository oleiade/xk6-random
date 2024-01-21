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

### Random Generator 

The `Random` class provides the following methods for generating random numbers.

It supports a default constructor which will use the current time as a seed, or a constructor which takes a seed as an argument.

```javascript
import { Random } from 'k6/x/random';

export default function() {
    // Create a new Random generator with a seed
    let rng = new Random(1234);
    let int = rng.int();
    console.log(int);
}
```

Armed with your random generator, you can use the following methods to generate random numbers.

#### `Random.int()`

Returns a random integer.

```javascript
import { Random } from 'k6/x/random';

export default function() {
    let rng = new Random();
    let int = rng.int();
    console.log(int);
}
```

#### `Random.float()`

Returns a random real number.

```javascript
import { Random } from 'k6/x/random';

export default function() {
    let rng = new Random();
    let float = rng.float();
    console.log(float);
}
```

#### `Random.boolean()`

Returns a random boolean value.

```javascript
import { Random } from 'k6/x/random';

export default function() {
    let rng = new Random();
    let bool = rng.boolean();
    console.log(bool);
}
```

#### `Random.intBetween(min, max)`

Returns a random integer between min and max.

```javascript
import { Random } from 'k6/x/random';

export default function() {
    let rng = new Random();
    let int = rng.intBetween(1, 10);
    console.log(int);
}
```

#### `Random.floatBetween(min, max)`

Returns a random real number between min and max.

```javascript
import { Random } from 'k6/x/random';

export default function() {
    let rng = new Random();
    let float = rng.floatBetween(1, 10);
    console.log(float);
}
```

#### `Random.pick(array)`

Returns a random element from the given array.

```javascript
import { Random } from 'k6/x/random';

export default function() {
    let rng = new Random();
    let array = [1, 2, 3, 4, 5];
    let pick = rng.pick(array);
    console.log(pick);
}
```

#### `Random.weightedPick(array, weights)`

Returns a random element from the given array, based on the given weights.

```javascript
import { Random } from 'k6/x/random';

export default function() {
    let rng = new Random();
    let array = [1, 2, 3, 4, 5];
    let weights = [0.1, 0.2, 0.3, 0.2, 0.2];
    let pick = rng.weightedPick(array, weights);
    console.log(pick);
}
```

#### `Random.normal(mean, stdev)`

Returns a random real number from a normal distribution.

```javascript
import { Random } from 'k6/x/random';

export default function() {
    let rng = new Random();
    let normal = rng.normal(0, 1);
    console.log(normal);
}
```

#### `Random.logNormal(mean, stdev)`

Returns a random real number from a log-normal distribution.

```javascript
import { Random } from 'k6/x/random';

export default function() {
    let rng = new Random();
    let logNormal = rng.logNormal(0, 1);
    console.log(logNormal);
}
```

#### `Random.bernoulli(probability)`

Returns a random boolean value with the given probability.

```javascript
import { Random } from 'k6/x/random';

export default function() {
    let rng = new Random();
    let bernoulli = rng.bernoulli(0.5);
    console.log(bernoulli);
}
```

#### `Random.binomial(trials, probability)`

Returns a random integer from a binomial distribution.

```javascript
import { Random } from 'k6/x/random';

export default function() {
    let rng = new Random();
    let binomial = rng.binomial(10, 0.5);
    console.log(binomial);
}
```

#### `Random.geometric(probability)`

Returns a random integer from a geometric distribution.

```javascript
import { Random } from 'k6/x/random';

export default function() {
    let rng = new Random();
    let geometric = rng.geometric(0.5);
    console.log(geometric);
}
```

#### `Random.exponential(rate)`

Returns a random real number from an exponential distribution.

```javascript
import { Random } from 'k6/x/random';

export default function() {
    let rng = new Random();
    let exponential = rng.exponential(0.5);
    console.log(exponential);
}
```

### Top-level Functions

Top-level functions are also provided for convenience:

#### `permutation(n)`

Returns a random permutation of the integers [0, n).

```javascript
import { permutation } from 'k6/x/random';

export default function() {
    let perm = permutation(10);
    console.log(perm);
}
```

#### `shuffle(array)`

Shuffles the given array in-place.

```javascript
import { shuffle } from 'k6/x/random';

export default function() {
    let array = [1, 2, 3, 4, 5];
    shuffle(array);
    console.log(array);
}
```

#### `shuffled(array)`

Returns a shuffled copy of the given array.

```javascript
import { shuffled } from 'k6/x/random';

export default function() {
    let array = [1, 2, 3, 4, 5];
    let shuffledArray = shuffled(array);
    console.log(shuffledArray);
}
```
