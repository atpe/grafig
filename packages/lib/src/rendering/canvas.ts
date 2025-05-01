import { from, Vec2 } from '../vector'
import { ContextRenderFunction, createContext } from './context'

export interface CanvasProps {
  parent?: HTMLElement | null
  dimensions?: Vec2
}
export function Canvas(
  { parent, dimensions }: CanvasProps = {},
  render?: ContextRenderFunction
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
