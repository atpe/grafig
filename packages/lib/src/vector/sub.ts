import { Vec2, Vec, Vec3 } from './core'

/**
 * Subtracts a number or vector from a given vector
 *
 * @param augend the vector being subtracted from
 * @param addend the number or vector being subtracted
 * @returns the vector difference
 */
export function sub(minuend: Vec2, subtrahend: number | Vec): Vec2
export function sub(minuend: Vec3, subtrahend: number | Vec): Vec3
export function sub(minuend: Vec, subtrahend: number | Vec): Vec
export function sub<V extends Vec>(minuend: V, subtrahend: number | Vec): V {
  return typeof subtrahend === 'number'
    ? {
        ...minuend,
        x: minuend.x - subtrahend,
        y: minuend.y - subtrahend,
        z: minuend.z ? minuend.z - subtrahend : undefined,
      }
    : {
        ...minuend,
        x: minuend.x - subtrahend.x,
        y: minuend.y - subtrahend.y,
        z: minuend.z ? minuend.z - (subtrahend.z || 0) : undefined,
      }
}
