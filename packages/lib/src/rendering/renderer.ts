import { Vec2 } from '../vector'
import { Canvas } from './canvas'

/**
 * A function that manipulate a {@link Canvas}
 *
 * @param canvas the target {@link Canvas}
 */
export type RenderFunction = (canvas: Canvas) => void

/** An object able to apply {@link RenderFunction}s */
export interface Renderer {
  /**
   * Applies a list of {@link RenderFunction}s to the bound {@link Canvas}
   *
   * @param functions list of {@link RenderFunction}s to apply
   */
  apply(...functions: RenderFunction[]): void
  /**
   * Creates a copy of the {@link Renderer} appending the given functions to the internal list of {@link RenderFunction}s
   *
   * @param functions list of {@link RenderFunction}s to append
   */
  with(...functions: RenderFunction[]): Renderer
  /**
   * Applies the internal list of {@link RenderFunction}s to the bound {@link Canvas}
   */
  render(): void
}

/**
 * Creates a {@link Renderer} able to manipulate the given {@link Canvas}
 *
 * The supplied list of {@link RenderFunction}s will only be applied to the canvas upon calling the render() function, allowing
 *
 * @param canvas the target {@link Canvas}
 * @param fns the list of {@link RenderFunction}s to append
 * @returns the created {@link Renderer}
 */
export function createRenderer(
  canvas: Canvas,
  ...functions: RenderFunction[]
): Renderer {
  const apply = (fn: RenderFunction): void => fn(canvas)
  return {
    apply: (...fns) => fns.forEach(apply),
    with: (...fns) => createRenderer(canvas, ...functions, ...fns),
    render: () => functions.forEach(apply),
  }
}

/**
 * Creates a {@link Renderer} and immediately applies the {@link RenderFunction}s to the given {@link Canvas}
 *
 * @param canvas the target {@link Canvas}
 * @param fns the list of {@link RenderFunction}s to append
 */
export function useRenderer(
  canvas: Canvas,
  ...functions: RenderFunction[]
): void {
  createRenderer(canvas, ...functions).render()
}

/**
 * Creates a {@link RenderFunction} that fills the entire canvas
 *
 * @param style the fill style
 * @returns the created {@link RenderFunction}
 */
export function background(style: string): RenderFunction {
  return (canvas): void => canvas.background(style)
}

/**
 * Creates a {@link RenderFunction} that renders text to the canvas
 *
 * @param text the text to render
 * @param style the fill style
 * @param position the  position
 * @param dimensions the text dimensions
 * @returns the created {@link RenderFunction}
 */
export function text(
  text: string,
  style: string,
  position: Vec2,
  dimensions: Vec2
): RenderFunction {
  return (canvas): void => canvas.text(text, style, position, dimensions)
}
