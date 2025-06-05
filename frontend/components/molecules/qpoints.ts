export enum Gender {
    Female,
    Male,
}

export function QPointsCalculator(gender: Gender, bodyweight: number, total: number, age: number): (number) {
    let b0: number = gender ? 416.70 : 266.50
    let b1: number = gender ? -47.87 : -19.44
    let b2: number = gender ? 18.93 : 18.61
    let tMax: number = gender ? 463.26 : 306.54

    const qPoints = total * (tMax / (b0 + b1 * Math.pow(bodyweight / 100, -2) + b2 * Math.pow(bodyweight / 100, 2)));

    const factor = Math.pow(10, 3); // Calculate 10^decimalPlaces
    return Math.ceil(qPoints * factor) / factor;
}