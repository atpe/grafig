import { Vec } from './core'

export function is<V extends Vec>(v: unknown): v is V {
  return (
    typeof (v as V).x === 'number' &&
    typeof (v as V).y === 'number' &&
    (typeof (v as V).z === 'undefined' || typeof (v as V).z === 'number')
  )
}
