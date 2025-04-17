import { Vec2 } from '../vector'
import { RenderFunction } from './renderer'

export function createContext(
  canvas: HTMLCanvasElement
): CanvasRenderingContext2D {
  const ctx = canvas.getContext('2d')
  if (ctx === null) throw new Error('canvas context cannot be null')
  return ctx
}

export function background(style: string): RenderFunction {
  return ({ context, dimensions }): void => {
    context.fillStyle = style
    context.fillRect(0, 0, dimensions.x, dimensions.y)
  }
}

export function text(
  text: string,
  style: string,
  position: Vec2,
  dimensions: Vec2
): RenderFunction {
  return ({ context }): void => {
    context.font = `${dimensions.y}px sans serif`
    context.fillStyle = style
    context.fillText(text, position.x, position.y, dimensions.x)
  }
}
