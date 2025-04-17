import { Vec2, Vec, Vec3 } from './core'

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
