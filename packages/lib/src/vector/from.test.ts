import { describe, expect, test } from '@jest/globals'
import { Vec2, Vec3 } from './core'
import { from } from './from'

describe.each([
  [{ x: 1, y: 1 }, 'x', 'y', { x: 1, y: 1 }],
  [{ width: -1, height: 0 }, 'width', 'height', { x: -1, y: 0 }],
  [{ u: 0.1, v: 1, w: 0 }, 'u', 'v', { x: 0.1, y: 1 }],
])(
  'from(%o, %s, %s) => %o',
  <C extends string>(obj: Record<C, number>, x: C, y: C, res: Vec2) => {
    const sut = from(obj, x, y)
    test('returns 2D vector', () => {
      expect(sut.z).toBeUndefined()
    })
    test('returns expected component values', () => {
      expect(sut.x).toBeCloseTo(res.x, 5)
      expect(sut.y).toBeCloseTo(res.y, 5)
    })
  }
)
describe.each([
  [{ x: 1, y: 1, z: 1 }, 'x', 'y', 'z', { x: 1, y: 1, z: 1 }],
  [
    { width: -1, height: 0, depth: 1 },
    'width',
    'height',
    'depth',
    { x: -1, y: 0, z: 1 },
  ],
  [{ u: 0.1, v: 1, w: 10 }, 'u', 'v', 'v', { x: 0.1, y: 1, z: 1 }],
])(
  'from(%o, %s, %s, %s) => %o',
  <C extends string>(obj: Record<C, number>, x: C, y: C, z: C, res: Vec3) => {
    const sut = from(obj, x, y, z)
    test('returns 3D vector', () => {
      expect(sut.z).toBeDefined()
    })
    test('returns expected component values', () => {
      expect(sut.x).toBeCloseTo(res.x, 5)
      expect(sut.y).toBeCloseTo(res.y, 5)
      expect(sut.z).toBeCloseTo(res.z, 5)
    })
  }
)
