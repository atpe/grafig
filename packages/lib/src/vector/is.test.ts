import { describe, expect, test } from '@jest/globals'
import { is } from './is'

describe.each([
  [{ x: 1, y: 1 }, true],
  [{ x: 1, y: 1, z: 1 }, true],
  [{ x: 1, y: 1, a: 1 }, true],
  [{}, false],
  [[], false],
  [1, false],
])('is(%o) => %s', (obj: unknown, expected: boolean) => {
  const sut = is(obj)

  test('returns expected component values', () => {
    expect(sut).toBe(expected)
  })
})
