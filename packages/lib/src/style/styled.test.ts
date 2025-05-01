/**
 * @jest-environment jsdom
 */

import { describe, test, expect, jest } from '@jest/globals'
import { styled } from './styled'

describe('styled()', () => {
  const canvas = document.createElement('canvas')
  const context = canvas.getContext('2d')

  if (context == null) throw new ReferenceError()

  describe('given a context and function', () => {
    // arrange
    // act
    const sut = styled(context, () => undefined)

    // assert
    test('returns shape render function', () => {
      expect(sut).toBeDefined()
      expect(sut).toBeInstanceOf(Function)
    })
  })
  describe('given a context and function and called', () => {
    // arrange
    const fn = jest.fn()
    const props = {}

    // act
    styled(context, fn)(props)

    // assert
    test('calls given function with given props', () =>
      expect(fn).toBeCalledWith(props))
  })
})
