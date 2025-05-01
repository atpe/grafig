/**
 * @jest-environment jsdom
 */

import { describe, expect, jest, test } from '@jest/globals'
import { Canvas } from './canvas'

describe('Canvas()', () => {
  const dimensions = { x: 500, y: 250 }
  const createElementSpy = jest.spyOn(document, 'createElement')

  describe('given no parent element', () => {
    // arrange
    const canvas = document.createElement('canvas')

    createElementSpy.mockReturnValueOnce(canvas)

    // act
    const sut = Canvas()

    // assert
    test('creates canvas element', expect(createElementSpy).toBeCalled)
    test('returns a non-null render function', () => {
      expect(sut).toBeDefined()
      expect(sut).toBeInstanceOf(Function)
    })
    test(
      'returns an unparented render function',
      expect(canvas.parentElement).toBeNull
    )
  })

  describe('given null parent', () => {
    // arrange
    const canvas = document.createElement('canvas')
    canvas.width = dimensions.x
    canvas.height = dimensions.y

    createElementSpy.mockReturnValueOnce(canvas)

    // act
    const sut = Canvas({ parent: null })

    // assert
    test('creates canvas element', expect(createElementSpy).toBeCalled)
    test('returns a non-null render function', () => {
      expect(sut).toBeDefined()
      expect(sut).toBeInstanceOf(Function)
    })
    test('returns a canvas object with expected size', () => {
      expect(canvas.width).toEqual(dimensions.x)
      expect(canvas.height).toEqual(dimensions.y)
    })
    test(
      'returns an unparented render function',
      expect(canvas.parentElement).toBeNull
    )
  })

  describe('given a parent element', () => {
    // arrange
    const dimensions = { x: 500, y: 250 }

    const canvas = document.createElement('canvas')
    const parent = document.createElement('div')

    jest.spyOn(parent, 'clientWidth', 'get').mockReturnValueOnce(dimensions.x)
    jest.spyOn(parent, 'clientHeight', 'get').mockReturnValueOnce(dimensions.y)

    createElementSpy.mockReturnValueOnce(canvas)

    // act
    const sut = Canvas({ parent })

    // assert
    test('creates canvas element', expect(createElementSpy).toBeCalled)
    test('returns a non-null render function', () => {
      expect(sut).toBeDefined()
      expect(sut).toBeInstanceOf(Function)
    })
    test('returns a canvas object with expected size', () => {
      expect(canvas.width).toEqual(dimensions.x)
      expect(canvas.height).toEqual(dimensions.y)
    })
    test('returns a parented render function', () =>
      expect(canvas.parentElement).toBe(parent))
  })

  describe('given dimensions', () => {
    // arrange
    const dimensions = { x: 500, y: 250 }

    const canvas = document.createElement('canvas')

    createElementSpy.mockReturnValueOnce(canvas)

    // act
    const sut = Canvas({ dimensions })

    // assert
    test('creates canvas element', expect(createElementSpy).toBeCalled)
    test('returns a non-null render function', () => {
      expect(sut).toBeDefined()
      expect(sut).toBeInstanceOf(Function)
    })
    test('returns a canvas object with expected size', () => {
      expect(canvas.width).toEqual(dimensions.x)
      expect(canvas.height).toEqual(dimensions.y)
    })
  })
})
