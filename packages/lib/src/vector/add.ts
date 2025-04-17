import { Vec, Vec2, Vec3 } from './core'

export function add(augend: Vec2, addend: number | Vec): Vec2
export function add(augend: Vec3, addend: number | Vec): Vec3
export function add(augend: Vec, addend: number | Vec): Vec
export function add<V extends Vec>(augend: V, addend: number | Vec): V {
  return typeof addend === 'number'
    ? {
        ...augend,
        x: augend.x + addend,
        y: augend.y + addend,
        z: augend.z ? augend.z + addend : undefined,
      }
    : {
        ...augend,
        x: augend.x + addend.x,
        y: augend.y + addend.y,
        z: augend.z ? augend.z + (addend.z || 0) : undefined,
      }
}
