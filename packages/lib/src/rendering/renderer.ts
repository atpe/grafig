import { Canvas } from './canvas'

export type RenderFunction = (canvas: Canvas) => void

export interface Renderer {
  apply(...fns: RenderFunction[]): void
}

export function createRenderer(canvas: Canvas): Renderer {
  return {
    apply: (...fns: RenderFunction[]): void => fns.forEach((fn) => fn(canvas)),
  }
}
