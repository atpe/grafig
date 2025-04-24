import { Selector, Option, select } from '../option'
import { from, Vec2 } from '../vector'

/**
 * Returns a {@link Selector} that returns the document body
 *
 * @returns a document body {@link Selector} function
 */
export function useBody(): Selector<HTMLElement> {
  return () => document.body
}

/**
 * Returns a {@link Selector} that returns the element with the given id, if it exists
 *
 * @returns an element {@link Selector} function
 */
export function useElementById(
  id = 'canvas-container'
): Selector<HTMLElement | null> {
  return () => document.getElementById(id)
}

/** Represents a specific DOM canvas element */
export interface Canvas {
  /** The underlying {@link CanvasRenderingContext2D} */
  readonly context: CanvasRenderingContext2D
  /** The canvas dimesions */
  dimensions: Vec2
  /** Fills the entire canvas */
  background(style: string): void
  /** Renders text onto the canvas */
  text(text: string, style: string, position: Vec2, dimensions: Vec2): void
}

/**
 * Creates a canvas within the DOM from the given options
 *
 * If the root element is supplied, the canvas will be appended to the element
 * and given the same dimensions. If the root element is undefined, the document
 * body is used by default. If the root element is null, the canvas is created
 * without a parent and remains the same size.
 *
 * @param root the root element option
 * @returns the created {@link Canvas}
 */
export function createCanvas(
  root: Option<HTMLElement | null> = useBody()
): Canvas {
  const element = document.createElement('canvas')
  const context = element.getContext('2d')

  if (context === null) throw new Error('canvas context cannot be null')

  const canvas: Canvas = {
    context,
    get dimensions() {
      return from(element, 'width', 'height')
    },
    set dimensions(value: Vec2) {
      element.width = value.x
      element.height = value.y
    },
    background: (style) => {
      context.fillStyle = style
      context.fillRect(0, 0, element.width, element.height)
    },
    text: (text, style, position, dimensions) => {
      context.font = `${dimensions.y}px sans serif`
      context.fillStyle = style
      context.fillText(text, position.x, position.y, dimensions.x)
    },
  }

  const parent = select(root)
  if (!parent) return canvas

  canvas.dimensions = from(parent, 'clientWidth', 'clientHeight')
  parent.appendChild(element)

  return canvas
}
