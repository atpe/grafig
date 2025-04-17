import { Vec2, Vec, Vec3 } from './core'

export function div(dividend: Vec2, divisor: number | Vec): Vec2
export function div(dividend: Vec3, divisor: number | Vec): Vec3
export function div(dividend: Vec, divisor: number | Vec): Vec
export function div<V extends Vec>(dividend: V, divisor: number | Vec): V {
  return typeof divisor === 'number'
    ? {
        ...dividend,
        x: dividend.x / (divisor || 1),
        y: dividend.y / (divisor || 1),
        z: dividend.z ? dividend.z / (divisor || 1) : undefined,
      }
    : {
        ...dividend,
        x: dividend.x / (divisor.x || 1),
        y: dividend.y / (divisor.y || 1),
        z: dividend.z ? dividend.z / (divisor.z || 1) : undefined,
      }
}
