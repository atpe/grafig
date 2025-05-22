import { ContextRenderer } from './context'

/**
 * Returns the given context renderer, unchanged.
 *
 * Provides a simple means of using the trailing lambda syntax.
 *
 * @example
 * // const fn = Renderer {
 * //   <do context stuff here>
 * // }
 *
 * @param render the context renderer to return
 * @returns the given context renderer
 */
export function Renderer(render: ContextRenderer): ContextRenderer {
  return render
}
