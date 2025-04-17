import { Vec2, Vec3, Vec } from './core'

export function from<
  C extends string,
  T extends Record<C, number> = Record<C, number>
>(obj: T, x: C, y: C): Vec2
export function from<
  C extends string,
  T extends Record<C, number> = Record<C, number>
>(obj: T, x: C, y: C, z: C): Vec3
export function from<
  C extends string,
  T extends Record<C, number> = Record<C, number>
>(obj: T, x: C, y: C, z?: C): Vec {
  return { x: obj[x], y: obj[y], z: z ? obj[z] : undefined }
}
