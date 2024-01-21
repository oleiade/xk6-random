import { permutation, shuffle, shuffled, Random } from 'k6/x/random';

export const options = {
    scenarios: {
        shuffle: {
            executor: 'shared-iterations',
            vus: 10,
            iterations: 100,
            exec: 'shuffleExample'
        },
        pick: {
            executor: 'shared-iterations',
            vus: 10,
            iterations: 100,
            exec: 'pickExample'
        },
        weightedPick: {
            executor: 'shared-iterations',
            vus: 10,
            iterations: 100,
            exec: 'weightedPickExample'
        },
        permutation: {
            executor: 'shared-iterations',
            vus: 10,
            iterations: 100,
            exec: 'permExample'
        },
        readmeUsageExample: {
            executor: 'shared-iterations',
            vus: 10,
            iterations: 100,
            exec: 'readmeUsageExample'
        },
        intExample: {
            executor: 'shared-iterations',
            vus: 10,
            iterations: 100,
            exec: 'intExample'
        }
    }
}

const random = new Random();


export function shuffleExample() {
    // Create a new Random generator
    let rng = new Random();

    // Shuffle an array
    let colors = ['red', 'green', 'blue', 'yellow'];
    shuffle(colors);
    console.log(colors);

    let pick = rng.pick(colors);
    console.log(`picked color: ${pick}`);
}

// pick exhibits picking a random element from an array.
export function pickExample() {
    const colors = ['red', 'green', 'blue'];

    const picked = [0, 0, 0]  // red, green, blue
    for (let i = 0; i < 1000; i++) {
        const color = random.pick(colors);
        switch (color) {
            case 'red':
                picked[0]++;
                break;
            case 'green':
                picked[1]++;
                break;
            case 'blue':
                picked[2]++;
                break;
        }
    }
}

// weightedPick exhibits picking a random element from an array with
// weights defining each element's probability of being picked.
export function weightedPickExample() {
    const colors = ['red', 'green', 'blue'];

    const picked = [0, 0, 0]  // red, green, blue
    for (let i = 0; i < 1000; i++) {
        const color = random.weightedPick(colors, [0.1, 0.8, 0.1]);
        switch (color) {
            case 'red':
                picked[0]++;
                break;
            case 'green':
                picked[1]++;
                break;
            case 'blue':
                picked[2]++;
                break;
        }
    }
}

export function permExample() {
    const colors = ['red', 'green', 'blue'];

    const firstPermutation = permutation(3);
    for (let i = 0; i < 3; i++) {
        const color = colors[firstPermutation[i]];
        console.log(color);
    }

    const secondPermutation = permutation(3);
    for (let i = 0; i < 3; i++) {
        const color = secondPermutation[i];
        console.log(color);
    }
}

export function readmeUsageExample() {
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

export function intExample() {
    let rng = new Random();
    let int = rng.int();
    console.log(int);
}