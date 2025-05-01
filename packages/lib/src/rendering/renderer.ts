import { ContextRenderFunction } from './context'

export function Renderer(render: ContextRenderFunction): ContextRenderFunction {
  return render
}
