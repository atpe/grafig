import { from, Vec2 } from '../vector'
import { ContextRenderer, createContext } from './context'

/** The canvas props object */
export interface CanvasProps {
  /** The parent element for the canvas */
  parent?: HTMLElement | null
  /** The canvas dimensions */
  dimensions?: Vec2
}

/**
 * Binds a given render function to a canvas.
 *
 * @param props the canvas props object
 * @param render the render function to apply
 * @returns the resultant render function
 *
 * @example
 * Canvas({ parent: document.body }, (context) => {
 *   context.background('white')
 *   context.apply(myFunc)
 * })
 * @example
 * Canvas({ parent: document.body }) {
 *   background('white')
 *   apply(myFunc)
 * }
 */
export function Canvas(
  { parent, dimensions }: CanvasProps = {},
  render?: ContextRenderer
): () => void {
  const element = document.createElement('canvas')
  const context = createContext(element)

  if (parent) {
    context.dimensions = from(parent, 'clientWidth', 'clientHeight')
    parent.appendChild(element)
  }

  if (dimensions) {
    context.dimensions = dimensions
  }

  return () => {
    if (render) render(context)
  }
}
