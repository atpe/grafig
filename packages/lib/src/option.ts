export type Selector<V, A extends unknown[] = []> = (...args: A) => V

export type Option<V, A extends unknown[] = []> = V | Selector<V, A>

/**
 * Gets the value from a given option
 *
 * @param value the option or value
 * @param args the option arguments
 * @returns the selected value
 */
export function select<V, A extends unknown[] = []>(
  value: Option<V, A>,
  ...args: A
): V {
  return value instanceof Function ? value(...args) : value
}
