export type Selector<V, A extends unknown[] = []> = (...args: A) => V

export type Option<V, A extends unknown[] = []> = V | Selector<V, A>

export function select<V, A extends unknown[] = []>(
  value: V | Selector<V, A>,
  ...args: A
): V {
  return value instanceof Function ? value(...args) : value
}
