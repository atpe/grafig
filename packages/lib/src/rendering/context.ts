import { from, Vec2 } from '../vector'
import {
  TextProps,
  RectProps,
  SquareProps,
  EllipseProps,
  CircleProps,
  text,
  rect,
  square,
  ellipse,
  circle,
} from '../shape'
import { background } from '../style'

/** The functions available for a context */
export interface ContextFunctions {
  /**
   * Fills the entire canvas with a given color.
   *
   * @param fill the target color
   */
  background(fill: CanvasRenderingContext2D['fillStyle']): void
  /**
   * Renders text onto the canvas.
   *
   * @param props the text props object
   */
  text(props: TextProps): void
  /**
   * Renders a rectangle onto the canvas.
   *
   * @param props the rect props object
   */
  rect(props: RectProps): void
  /**
   * Renders a square onto the canvas.
   *
   * @param props the square props object
   */
  square(props: SquareProps): void
  /**
   * Renders an ellipse onto the canvas.
   *
   * @param props the ellipse props object
   */
  ellipse(props: EllipseProps): void
  /**
   * Renders a circle onto the canvas.
   *
   * @param props the circle props object
   */
  circle(props: CircleProps): void
}

/** A function that applies some side effect to the canvas context */
export type ContextRenderer = (context: Context) => void

/** The canvas rendering context */
export interface Context extends ContextFunctions {
  /** The canvas dimensions */
  dimensions: Vec2
  /**
   * Applies one or more render functions to the canvas.
   *
   * @param fns the functions to apply
   */
  apply(...fns: ContextRenderer[]): void
}

/**
 * Creates a new rendering context for the given canvas element
 *
 * @param canvas the target canvas element
 * @returns the rendering context
 */
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
    apply(...fns: ContextRenderer[]): void {
      fns.forEach((fn) => fn(this))
    },
    background: background(canvas),
    text: text(canvas),
    rect: rect(canvas),
    square: square(canvas),
    ellipse: ellipse(canvas),
    circle: circle(canvas),
  }
}
