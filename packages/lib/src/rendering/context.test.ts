/**
 * @jest-environment jsdom
 */

import { describe, expect, jest, test } from '@jest/globals'
import { Context, ContextRenderer, createContext } from './context'

describe('createContext()', () => {
  const canvas = document.createElement('canvas')
  const getContextSpy = jest.spyOn(canvas, 'getContext')

  describe('given valid canvas element', () => {
    // arrange
    const context = canvas.getContext('2d')
    getContextSpy.mockReturnValueOnce(context)

    // act
    const sut = createContext(canvas)

    // assert
    test('gets 2D canvas context', () =>
      expect(getContextSpy).toBeCalledWith('2d'))
    test('returns a non-null context object', expect(sut).toBeDefined)
  })

  describe('given invalid canvas element', () => {
    // arrange
    getContextSpy.mockReturnValueOnce(null)

    // act
    const sut = (): Context => createContext(canvas)

    // assert
    test('gets 2D canvas context', () =>
      expect(getContextSpy).toBeCalledWith('2d'))
    test('throws error', expect(sut).toThrow)
  })
})

describe('Context.dimensions', () => {
  describe('get', () => {
    // arrange
    const canvas = document.createElement('canvas')
    const context = createContext(canvas)

    // act
    const sut = context.dimensions

    //assert
    test('returns canvas element dimensions', () => {
      expect(sut.x).toBe(canvas.width)
      expect(sut.y).toBe(canvas.height)
    })
  })
  describe('set', () => {
    // arrange
    const canvas = document.createElement('canvas')
    const context = createContext(canvas)

    // act
    context.dimensions = { x: 100, y: 100 }

    //assert
    test('updates canvas element dimensions', () => {
      expect(canvas.width).toBe(100)
      expect(canvas.height).toBe(100)
    })
  })
})

describe('Context.apply()', () => {
  describe('when given context render function', () => {
    // arrange
    const canvas = document.createElement('canvas')
    const context = createContext(canvas)

    const fn = jest.fn<ContextRenderer>()

    // act
    context.apply(fn)

    //assert
    test('applies function to context', () =>
      expect(fn).toBeCalledWith(context))
  })
})
