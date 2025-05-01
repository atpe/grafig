import { from, Vec2 } from '../vector'
import {
  TextProps,
  RectProps,
  SquareProps,
  EllipseProps,
  CircleProps,
  background,
  text,
  rect,
  square,
  ellipse,
  circle,
} from '../shape'

export interface ContextFunctions {
  /** Fills the entire canvas */
  background(fill: CanvasRenderingContext2D['fillStyle']): void
  /** Renders text onto the canvas */
  text(props: TextProps): void
  /** Renders a rectangle onto the canvas */
  rect(props: RectProps): void
  /** Renders a square onto the canvas */
  square(props: SquareProps): void
  /** Renders an ellipse onto the canvas */
  ellipse(props: EllipseProps): void
  /** Renders a circle onto the canvas */
  circle(props: CircleProps): void
}

export type ContextRenderFunction = (context: Context) => void
export interface Context extends ContextFunctions {
  /** The canvas dimensions */
  dimensions: Vec2
  apply(fn: ContextRenderFunction): void
}

export function createContext(canvas: HTMLCanvasElement): Context {
  const ctx = canvas.getContext('2d')

  if (ctx === null) throw new Error('canvas context cannot be null')

  return {
    get dimensions(): Vec2 {
      return from(canvas, 'width', 'height')
    },
    set dimensions(value) {
      canvas.width = value.x
      canvas.height = value.y
    },
    apply(fn: ContextRenderFunction): void {
      fn(this)
    },
    background: background(canvas),
    text: text(canvas),
    rect: rect(canvas),
    square: square(canvas),
    ellipse: ellipse(canvas),
    circle: circle(canvas),
  }
}
