/**
 * @jest-environment jsdom
 */

import { describe, expect, jest, test } from '@jest/globals'
import { createRenderer, RenderFunction } from './renderer'
import { createCanvas } from './canvas'

describe('createRenderer()', () => {
  describe('given a canvas', () => {
    // arrange
    const canvas = createCanvas()

    // act
    const sut = createRenderer(canvas)

    // assert
    test('returns a non-null renderer object', expect(sut).not.toBeNull)
  })
})

describe('Renderer.apply()', () => {
  describe('given a list of functions', () => {
    // arrange
    const canvas = createCanvas()
    const renderer = createRenderer(canvas)

    const rectMock = jest.fn<RenderFunction>()
    const fillMock = jest.fn<RenderFunction>()

    // act
    renderer.apply(rectMock, rectMock, fillMock)

    // assert
    test('calls each function', () => {
      expect(rectMock).nthCalledWith(1, canvas)
      expect(rectMock).nthCalledWith(2, canvas)
      expect(fillMock).nthCalledWith(1, canvas)
    })
  })
})
