export type Dimension = 2 | 3

export type Vec2Component = 'x' | 'y'
export type Vec3Component = 'x' | 'y' | 'z'
export type VecComponent = Vec2Component | Vec3Component

export interface Vec2 {
  readonly x: number
  readonly y: number
  readonly z?: never
}
export interface Vec3 {
  readonly x: number
  readonly y: number
  readonly z: number
}
export type Vec = Vec2 | Vec3
