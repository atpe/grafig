import { Vec2, Vec, Vec3 } from './core'

/**
 * Multiplies a given vector by a number
 *
 * To multiply the components of a vector with another vector, use cross()
 *
 * @param multiplicand the number being multiplied
 * @param multiplier the number being multiplied by
 * @returns the vector product
 */
export function mult(multiplicand: Vec2, multiplier: number): Vec2
export function mult(multiplicand: Vec3, multiplier: number): Vec3
export function mult(multiplicand: Vec, multiplier: number): Vec
export function mult<V extends Vec>(multiplicand: V, multiplier: number): V {
  return {
    ...multiplicand,
    x: multiplicand.x * multiplier,
    y: multiplicand.y * multiplier,
    z: multiplicand.z ? multiplicand.z * multiplier : undefined,
  }
}
