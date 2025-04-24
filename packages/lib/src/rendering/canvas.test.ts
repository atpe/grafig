/**
 * @jest-environment jsdom
 */

import { describe, expect, jest, test } from '@jest/globals'
import { createCanvas, useElementById } from './canvas'

describe('createCanvas()', () => {
  const dimensions = { x: 500, y: 250 }
  const createElementSpy = jest.spyOn(document, 'createElement')

  describe('given no root', () => {
    // arrange
    const canvasElement = document.createElement('canvas')

    jest
      .spyOn(document.body, 'clientWidth', 'get')
      .mockReturnValueOnce(dimensions.x)
    jest
      .spyOn(document.body, 'clientHeight', 'get')
      .mockReturnValueOnce(dimensions.y)

    createElementSpy.mockReturnValueOnce(canvasElement)

    // act
    const sut = createCanvas()

    // assert
    test('creates canvas element', expect(createElementSpy).toBeCalled)
    test('returns a non-null canvas object', expect(sut).toBeDefined)
    test('returns a canvas object with expected size', () =>
      expect(sut.dimensions).toEqual(dimensions))
    test('returns a canvas object with expected parent', () =>
      expect(canvasElement.parentElement).toBe(document.body))
  })

  describe('given null root', () => {
    // arrange
    const canvasElement = document.createElement('canvas')
    canvasElement.width = dimensions.x
    canvasElement.height = dimensions.y

    createElementSpy.mockReturnValueOnce(canvasElement)

    // act
    const sut = createCanvas(null)

    // assert
    test('creates canvas element', expect(createElementSpy).toBeCalled)
    test('returns a non-null canvas object', expect(sut).toBeDefined)
    test('returns a canvas object with expected size', () =>
      expect(sut.dimensions).toEqual(dimensions))
    test(
      'returns a canvas object with null parent',
      expect(canvasElement.parentElement).toBeNull
    )
  })

  describe('given a root value', () => {
    // arrange
    const dimensions = { x: 500, y: 250 }

    const canvasElement = document.createElement('canvas')
    const rootElement = document.createElement('div')

    jest
      .spyOn(rootElement, 'clientWidth', 'get')
      .mockReturnValueOnce(dimensions.x)
    jest
      .spyOn(rootElement, 'clientHeight', 'get')
      .mockReturnValueOnce(dimensions.y)

    createElementSpy.mockReturnValueOnce(canvasElement)

    // act
    const sut = createCanvas(rootElement)

    // assert
    test('creates canvas element', expect(createElementSpy).toBeCalled)
    test('returns a non-null canvas object', expect(sut).toBeDefined)
    test('returns a canvas object with expected size', () =>
      expect(sut.dimensions).toEqual(dimensions))
    test('returns a canvas object with expected parent', () =>
      expect(canvasElement.parentElement).toBe(rootElement))
  })

  describe('given a root selector', () => {
    // arrange
    const dimensions = { x: 500, y: 250 }

    const canvasElement = document.createElement('canvas')
    const rootElement = document.createElement('div')

    rootElement.id = 'test-container'
    document.body.appendChild(rootElement)

    jest
      .spyOn(rootElement, 'clientWidth', 'get')
      .mockReturnValueOnce(dimensions.x)
    jest
      .spyOn(rootElement, 'clientHeight', 'get')
      .mockReturnValueOnce(dimensions.y)

    createElementSpy.mockReturnValueOnce(canvasElement)

    // act
    const sut = createCanvas(useElementById('test-container'))

    // assert
    test('creates canvas element', expect(createElementSpy).toBeCalled)
    test('returns a non-null canvas object', expect(sut).toBeDefined)
    test('returns a canvas object with expected size', () =>
      expect(sut.dimensions).toEqual(dimensions))
    test('returns a canvas object with expected parent', () =>
      expect(canvasElement.parentElement).toBe(rootElement))
  })
  describe('given getContext() returns null context', () => {
    // arrange
    const canvasElement = document.createElement('canvas')
    createElementSpy.mockReturnValueOnce(canvasElement)
    jest.spyOn(canvasElement, 'getContext').mockReturnValueOnce(null)

    // act
    const sut = createCanvas

    // assert
    test('throws error', expect(sut).toThrow)
  })
})
