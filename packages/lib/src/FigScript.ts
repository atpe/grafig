/* eslint-disable @typescript-eslint/no-explicit-any */
import * as vector from './vector'

function add<T>(a: T, b: any): T
function add<V extends vector.Vec>(a: V, b: number | vector.Vec): V
function add(a: any, b: any): any {
  if (a == null || b == null)
    throw new TypeError(`add() cannot be applied to nullish value`)
  return vector.is(a) && (vector.is(b) || typeof b === 'number')
    ? vector.add(a, b)
    : a + b
}
function sub<T>(a: T, b: any): T
function sub<V extends vector.Vec>(a: V, b: number | vector.Vec): V
function sub(a: any, b: any): any {
  if (a == null || b == null)
    throw new TypeError(`sub() cannot be applied to nullish value`)
  return vector.is(a) && (vector.is(b) || typeof b === 'number')
    ? vector.sub(a, b)
    : a - b
}
function mult<T>(a: T, b: any): T
function mult<V extends vector.Vec>(a: V, b: number | vector.Vec): V
function mult(a: any, b: any): any {
  if (a == null || b == null)
    throw new TypeError(`mult() cannot be applied to nullish value`)
  return vector.is(a) && (vector.is(b) || typeof b === 'number')
    ? vector.mult(a, b)
    : a * b
}
function div<T>(a: T, b: any): T
function div<V extends vector.Vec>(a: V, b: number): V
function div(a: any, b: any): any {
  if (a == null || b == null)
    throw new TypeError(`div() cannot be applied to nullish value`)
  return vector.is(a) && (vector.is(b) || typeof b === 'number')
    ? vector.div(a, b)
    : a / b
}

export default {
  ...vector,
  add,
  sub,
  mult,
  div,
}
