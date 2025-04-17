import { Selector, Option, select } from '../option'
import { from, Vec2 } from '../vector'

import { createContext } from './context'

export function useBody(): Selector<HTMLElement> {
  return () => document.body
}
export function useElementById(
  id = 'canvas-container'
): Selector<HTMLElement | null> {
  return () => document.getElementById(id)
}

export interface Canvas {
  readonly context: CanvasRenderingContext2D
  dimensions: Vec2
}

export function createCanvas(
  root: Option<HTMLElement | null> = useBody()
): Canvas {
  const element = document.createElement('canvas')

  const canvas: Canvas = {
    context: createContext(element),
    get dimensions() {
      return from(element, 'width', 'height')
    },
    set dimensions(value: Vec2) {
      element.width = value.x
      element.height = value.y
    },
  }

  const parent = select(root)
  if (!parent) return canvas

  canvas.dimensions = from(parent, 'clientWidth', 'clientHeight')
  parent.appendChild(element)

  return canvas
}
