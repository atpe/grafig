/**
 * @jest-environment jsdom
 */

import * as functions from './index'

import { describe, expect, jest, test } from '@jest/globals'

const fns = ['circle', 'ellipse', 'rect', 'square', 'text'] as const

describe.each(fns)('%s()', (key) => {
  const canvas = document.createElement('canvas')
  const getContextSpy = jest.spyOn(canvas, 'getContext')

  describe('given valid canvas element', () => {
    // arrange

    // act
    const sut = functions[key](canvas)

    // assert
    test('gets 2D canvas context', () =>
      expect(getContextSpy).toBeCalledWith('2d'))
    test('returns a non-null render function', () => {
      expect(sut).toBeDefined()
      expect(sut).toBeInstanceOf(Function)
    })
  })
  describe('given invalid canvas element', () => {
    // arrange
    getContextSpy.mockReturnValueOnce(null)

    // act
    const sut = (): unknown => functions[key](canvas)

    // assert
    test('gets 2D canvas context', () =>
      expect(getContextSpy).toBeCalledWith('2d'))
    test('throws error', () => expect(sut).toThrow())
  })
})
