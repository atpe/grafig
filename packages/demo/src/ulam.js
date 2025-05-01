// Code generated from src/ulam.fs by @grafig/syntax. DO NOT EDIT.

import FigScript, { Renderer } from '@grafig/lib'

function prime(n) {
  for(let i = 2, s = Math.sqrt(n); i <= s; i++) {
    if(n % i === 0) return false;
  }
  return n > 1;
}

function next(
  width, // Grid width
  index, // Current index
  direction, // Move direction [0 - 3]
  length, // Spiral step length
  steps, // Number of steps made
  turns // Number of turns made
) {
  const turn = steps % length === 0
  const incr = turns % 2 === 0

  const s = FigScript.add(steps, 1)
  const t = turn ? FigScript.add(turns, 1) : turns
  const d = turn ? (FigScript.add(direction, 1)) % 4 : direction
  const l = turn && incr ? FigScript.add(length, 1) : length
  const i = (() => {
    switch (direction) {
      case 0:
        return FigScript.add(index, 1)
      case 1:
        return FigScript.sub(index, width)
      case 2:
        return FigScript.sub(index, 1)
      case 3:
        return FigScript.add(index, width)
      default:
        return index
    }
  })();

  return [i, d, l, s, t];
}

export function Spiral(dimensions, radius) {
  const max = Math.max(dimensions.x, dimensions.y)
  const width = Math.floor(FigScript.div(max, radius))
  const length = width ** 2
  const offset = width % 2 === 0 ? Math.floor(FigScript.div(width, 2)) : 0
  const center = FigScript.add(Math.floor(FigScript.div(length, 2)), offset)
  const nodes = Array(length)
    .fill(0)
    .reduce(
      (acc, _, n) => n === 0 
        ? [[center, 0, 1, 1, 0]]
        : [...acc, next(width, ...acc[FigScript.sub(acc.length,1)])],
      []
    )
    .map(([i, d, l, s, t], n) => {
      const position = FigScript.mult(({ x: i % width, y: Math.floor(FigScript.div(i, width)) }), (FigScript.mult(radius, 2)))
      if (
        position.x < 0 ||
        position.x > dimensions.x ||
        position.y < 0 ||
        position.y > dimensions.y ||
        !prime(FigScript.add(n, 1))
      )
        return () => {}

      return Renderer((it) => {
        it.circle({ position, radius, fill: 'white' })
      })
    })
  return (context) => {
    for (const node of nodes) {
      context.apply(node)
    }
  }
}