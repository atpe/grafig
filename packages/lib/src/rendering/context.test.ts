/**
 * @jest-environment jsdom
 */

import { describe, expect, jest, test } from '@jest/globals'
import { createContext } from './context'

describe('createContext()', () => {
  describe('given a canvas element with non-null context', () => {
    // arrange
    const canvasElement = document.createElement('canvas')

    // act
    const sut = createContext(canvasElement)

    // assert
    test('returns a non-null context object', expect(sut).not.toBeNull)
  })
  describe('given a canvas element with null context', () => {
    // arrange
    const canvasElement = document.createElement('canvas')
    jest.spyOn(canvasElement, 'getContext').mockReturnValueOnce(null)

    // act
    const sut = (): CanvasRenderingContext2D => createContext(canvasElement)

    // assert
    test('throws error', expect(sut).toThrow)
  })
})
