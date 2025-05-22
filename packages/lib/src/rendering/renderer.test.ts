import { describe, expect, jest, test } from '@jest/globals'
import { Renderer } from './renderer'
import { ContextRenderer } from './context'

describe('Renderer()', () => {
  describe('when given context render function', () => {
    // arrange
    const fn = jest.fn<ContextRenderer>()

    // act
    const sut = Renderer(fn)

    //assert
    test('returns context render function', () => expect(sut).toBe(fn))
  })
})
