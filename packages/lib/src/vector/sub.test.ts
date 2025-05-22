import { describe, expect, test } from '@jest/globals'
import { Vec2, Vec3 } from './core'
import { sub } from './sub'

const name = 'given %o and %o'

describe('sub()', () => {
  describe.each([
    [{ x: 1, y: 1 }, 1, { x: 0, y: 0 }],
    [{ x: -1, y: 0 }, -1, { x: 0, y: 1 }],
    [{ x: 0.1, y: 1 }, 0.01, { x: 0.09, y: 0.99 }],
  ])(name, (a: Vec2, n: number, res: Vec2) => {
    const sut = sub(a, n)
    test('returns 2D vector', () => {
      expect(sut.z).toBeUndefined()
    })
    test('returns expected component values', () => {
      expect(sut.x).toBeCloseTo(res.x, 5)
      expect(sut.y).toBeCloseTo(res.y, 5)
    })
  })
  describe.each([
    [{ x: 1, y: 1, z: 1 }, 1, { x: 0, y: 0, z: 0 }],
    [{ x: -1, y: 0, z: 1 }, -1, { x: 0, y: 1, z: 2 }],
    [{ x: 0.1, y: 1, z: 10 }, 0.01, { x: 0.09, y: 0.99, z: 9.99 }],
  ])(name, (a: Vec3, n: number, res: Vec3) => {
    const sut = sub(a, n)
    test('returns 3D vector', () => {
      expect(sut.z).toBeDefined()
    })
    test('returns expected component values', () => {
      expect(sut.x).toBeCloseTo(res.x, 5)
      expect(sut.y).toBeCloseTo(res.y, 5)
      expect(sut.z).toBeCloseTo(res.z, 5)
    })
  })
  describe.each([
    [
      { x: 1, y: 1 },
      { x: 1, y: 1 },
      { x: 0, y: 0 },
    ],
    [
      { x: -1, y: 0 },
      { x: -1, y: -2 },
      { x: 0, y: 2 },
    ],
    [
      { x: 0.1, y: 1 },
      { x: 0.01, y: 0.01 },
      { x: 0.09, y: 0.99 },
    ],
  ])(name, (a: Vec2, b: Vec2, res: Vec2) => {
    const sut = sub(a, b)
    test('returns 2D vector', () => {
      expect(sut.z).toBeUndefined()
    })
    test('returns expected component values', () => {
      expect(sut.x).toBeCloseTo(res.x, 5)
      expect(sut.y).toBeCloseTo(res.y, 5)
    })
  })
  describe.each([
    [
      { x: 1, y: 1, z: 1 },
      { x: 1, y: 1 },
      { x: 0, y: 0, z: 1 },
    ],
    [
      { x: -1, y: 0, z: 1 },
      { x: -1, y: -2 },
      { x: 0, y: 2, z: 1 },
    ],
    [
      { x: 0.1, y: 1, z: 10 },
      { x: 0.01, y: 0.01 },
      { x: 0.09, y: 0.99, z: 10 },
    ],
  ])(name, (a: Vec3, b: Vec2, res: Vec3) => {
    const sut = sub(a, b)
    test('returns 3D vector', () => {
      expect(sut.z).toBeDefined()
    })
    test('returns expected component values', () => {
      expect(sut.x).toBeCloseTo(res.x, 5)
      expect(sut.y).toBeCloseTo(res.y, 5)
      expect(sut.z).toBeCloseTo(res.z, 5)
    })
  })
  describe.each([
    [
      { x: 1, y: 1 },
      { x: 1, y: 1, z: 1 },
      { x: 0, y: 0 },
    ],
    [
      { x: -1, y: 0 },
      { x: -1, y: -2, z: -3 },
      { x: 0, y: 2 },
    ],
    [
      { x: 0.1, y: 1 },
      { x: 0.01, y: 0.01, z: 0.01 },
      { x: 0.09, y: 0.99 },
    ],
  ])(name, (a: Vec2, b: Vec3, res: Vec2) => {
    const sut = sub(a, b)
    test('returns 2D vector', () => {
      expect(sut.z).toBeUndefined()
    })
    test('returns expected component values', () => {
      expect(sut.x).toBeCloseTo(res.x, 5)
      expect(sut.y).toBeCloseTo(res.y, 5)
    })
  })
  describe.each([
    [
      { x: 1, y: 1, z: 1 },
      { x: 1, y: 1, z: 1 },
      { x: 0, y: 0, z: 0 },
    ],
    [
      { x: -1, y: 0, z: 1 },
      { x: -1, y: -2, z: -3 },
      { x: 0, y: 2, z: 4 },
    ],
    [
      { x: 0.1, y: 1, z: 10 },
      { x: 0.01, y: 0.01, z: 0.01 },
      { x: 0.09, y: 0.99, z: 9.99 },
    ],
  ])(name, (a: Vec3, b: Vec3, res: Vec3) => {
    const sut = sub(a, b)
    test('returns 3D vector', () => {
      expect(sut.z).toBeDefined()
    })
    test('returns expected component values', () => {
      expect(sut.x).toBeCloseTo(res.x, 5)
      expect(sut.y).toBeCloseTo(res.y, 5)
      expect(sut.z).toBeCloseTo(res.z, 5)
    })
  })
})
