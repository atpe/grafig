import { Vec2, Vec, Vec3 } from './core'

/**
 * Multiplies a given vector by a number or vector
 *
 * @param multiplicand the number being multiplied
 * @param multiplier the number or vector being multiplied by
 * @returns the vector product
 */
export function mult(multiplicand: Vec2, multiplier: number | Vec): Vec2
export function mult(multiplicand: Vec3, multiplier: number | Vec): Vec3
export function mult(multiplicand: Vec, multiplier: number | Vec): Vec
export function mult<V extends Vec>(
  multiplicand: V,
  multiplier: number | Vec
): V {
  return typeof multiplier === 'number'
    ? {
        ...multiplicand,
        x: multiplicand.x * multiplier,
        y: multiplicand.y * multiplier,
        z: multiplicand.z ? multiplicand.z * multiplier : undefined,
      }
    : {
        ...multiplicand,
        x: multiplicand.x * multiplier.x,
        y: multiplicand.y * multiplier.y,
        z: multiplicand.z ? multiplicand.z * (multiplier.z ?? 1) : undefined,
      }
}
