export enum Gender {
    Female,
    Male,
}

export interface AgeFactor {
    age: number | null;
    ageFactorWomen: number | null;
    ageFactorMen: number;
}

const ageFactors: AgeFactor[] = [
    { age: 30, ageFactorWomen: 1.000, ageFactorMen: 1.000 },
    { age: 31, ageFactorWomen: 1.010, ageFactorMen: 1.010 },
    { age: 32, ageFactorWomen: 1.021, ageFactorMen: 1.018 },
    { age: 33, ageFactorWomen: 1.031, ageFactorMen: 1.026 },
    { age: 34, ageFactorWomen: 1.042, ageFactorMen: 1.038 },
    { age: 35, ageFactorWomen: 1.052, ageFactorMen: 1.052 },
    { age: 36, ageFactorWomen: 1.063, ageFactorMen: 1.064 },
    { age: 37, ageFactorWomen: 1.073, ageFactorMen: 1.076 },
    { age: 38, ageFactorWomen: 1.084, ageFactorMen: 1.088 },
    { age: 39, ageFactorWomen: 1.096, ageFactorMen: 1.100 },
    { age: 40, ageFactorWomen: 1.108, ageFactorMen: 1.112 },
    { age: 41, ageFactorWomen: 1.122, ageFactorMen: 1.124 },
    { age: 42, ageFactorWomen: 1.138, ageFactorMen: 1.136 },
    { age: 43, ageFactorWomen: 1.155, ageFactorMen: 1.148 },
    { age: 44, ageFactorWomen: 1.173, ageFactorMen: 1.160 },
    { age: 45, ageFactorWomen: 1.194, ageFactorMen: 1.173 },
    { age: 46, ageFactorWomen: 1.216, ageFactorMen: 1.187 },
    { age: 47, ageFactorWomen: 1.240, ageFactorMen: 1.201 },
    { age: 48, ageFactorWomen: 1.265, ageFactorMen: 1.215 },
    { age: 49, ageFactorWomen: 1.292, ageFactorMen: 1.230 },
    { age: 50, ageFactorWomen: 1.321, ageFactorMen: 1.247 },
    { age: 51, ageFactorWomen: 1.352, ageFactorMen: 1.264 },
    { age: 52, ageFactorWomen: 1.384, ageFactorMen: 1.283 },
    { age: 53, ageFactorWomen: 1.419, ageFactorMen: 1.304 },
    { age: 54, ageFactorWomen: 1.456, ageFactorMen: 1.327 },
    { age: 55, ageFactorWomen: 1.494, ageFactorMen: 1.351 },
    { age: 56, ageFactorWomen: 1.534, ageFactorMen: 1.376 },
    { age: 57, ageFactorWomen: 1.575, ageFactorMen: 1.401 },
    { age: 58, ageFactorWomen: 1.617, ageFactorMen: 1.425 },
    { age: 59, ageFactorWomen: 1.660, ageFactorMen: 1.451 },
    { age: 60, ageFactorWomen: 1.704, ageFactorMen: 1.477 },
    { age: 61, ageFactorWomen: 1.748, ageFactorMen: 1.504 },
    { age: 62, ageFactorWomen: 1.794, ageFactorMen: 1.531 },
    { age: 63, ageFactorWomen: 1.841, ageFactorMen: 1.560 },
    { age: 64, ageFactorWomen: 1.890, ageFactorMen: 1.589 },
    { age: 65, ageFactorWomen: 1.942, ageFactorMen: 1.620 },
    { age: 66, ageFactorWomen: 1.996, ageFactorMen: 1.654 },
    { age: 67, ageFactorWomen: 2.052, ageFactorMen: 1.693 },
    { age: 68, ageFactorWomen: 2.109, ageFactorMen: 1.736 },
    { age: 69, ageFactorWomen: 2.168, ageFactorMen: 1.784 },
    { age: 70, ageFactorWomen: 2.226, ageFactorMen: 1.833 },
    { age: 71, ageFactorWomen: 2.285, ageFactorMen: 1.883 },
    { age: 72, ageFactorWomen: 2.343, ageFactorMen: 1.932 },
    { age: 73, ageFactorWomen: 2.402, ageFactorMen: 1.981 },
    { age: 74, ageFactorWomen: 2.464, ageFactorMen: 2.031 },
    { age: 75, ageFactorWomen: 2.528, ageFactorMen: 2.083 },
    { age: 76, ageFactorWomen: 2.597, ageFactorMen: 2.139 },
    { age: 77, ageFactorWomen: 2.670, ageFactorMen: 2.202 },
    { age: 78, ageFactorWomen: 2.749, ageFactorMen: 2.271 },
    { age: 79, ageFactorWomen: 2.831, ageFactorMen: 2.348 },
    { age: 80, ageFactorWomen: 2.918, ageFactorMen: 2.430 },
    { age: 81, ageFactorWomen: 3.009, ageFactorMen: 2.524 },
    { age: 82, ageFactorWomen: 3.104, ageFactorMen: 2.635 },
    { age: 83, ageFactorWomen: 3.201, ageFactorMen: 2.755 },
    { age: 84, ageFactorWomen: 3.301, ageFactorMen: 2.877 },
    { age: 85, ageFactorWomen: 3.403, ageFactorMen: 3.008 },
    { age: 86, ageFactorWomen: 3.507, ageFactorMen: 3.168 },
    { age: 87, ageFactorWomen: 3.613, ageFactorMen: 3.356 },
    { age: 88, ageFactorWomen: 3.720, ageFactorMen: 3.545 },
    { age: 89, ageFactorWomen: 3.827, ageFactorMen: 3.709 },
    { age: 90, ageFactorWomen: 3.935, ageFactorMen: 3.880 },
    { age: null, ageFactorWomen: null, ageFactorMen: 4.059 },
    { age: null, ageFactorWomen: null, ageFactorMen: 4.247 },
    { age: null, ageFactorWomen: null, ageFactorMen: 4.443 },
    { age: null, ageFactorWomen: null, ageFactorMen: 4.648 },
    { age: null, ageFactorWomen: null, ageFactorMen: 4.863 },
];

export function QPointsCalculator(gender: Gender, bodyweight: number, total: number, age: number = 25): (number) {
    let b0: number = gender ? 416.70 : 266.50
    let b1: number = gender ? -47.87 : -19.44
    let b2: number = gender ? 18.93 : 18.61
    let tMax: number = gender ? 463.26 : 306.54

    const qPoints = total * (tMax / (b0 + b1 * Math.pow(bodyweight / 100, -2) + b2 * Math.pow(bodyweight / 100, 2)));

    const factor = Math.pow(10, 3); // Calculate 10^decimalPlaces
    if (age >= 30) {
      const mastersFactor = ageFactors.find((row) => row.age === age )
      if (gender === Gender.Male) {
        if (mastersFactor?.ageFactorMen === undefined) {
          return 999.999
        }
        return Math.ceil(mastersFactor?.ageFactorMen * (qPoints * factor)) / factor
      } else if (gender === Gender.Female) {
        if (mastersFactor?.ageFactorWomen === undefined || mastersFactor?.ageFactorWomen === null) {
          return 999.999
        }
        return Math.ceil(mastersFactor?.ageFactorWomen * (qPoints * factor)) / factor
      } else {
        return 999.999
      }
    }
    return Math.ceil(qPoints * factor) / factor;
}