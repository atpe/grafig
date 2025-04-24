import { Vec2, Vec3, Vec } from './core'

/**
 * Creates a vector from a given object and property names
 *
 * @param obj the target object
 * @param x the property key for the x-component
 * @param y the property key for the y-component
 * @returns the resultant vector
 */
export function from<
  C extends string,
  T extends Record<C, number> = Record<C, number>
>(obj: T, x: C, y: C): Vec2
/**
 * Creates a vector from a given object and property names
 *
 * @param obj the target object
 * @param x the property key for the x-component
 * @param y the property key for the y-component
 * @param z the property key for the z-component
 * @returns the resultant vector
 */
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
