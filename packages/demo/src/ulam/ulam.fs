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

  const s = steps + 1
  const t = turn ? turns + 1 : turns
  const d = turn ? (direction + 1) % 4 : direction
  const l = turn && incr ? length + 1 : length
  const i = (() => {
    switch (direction) {
      case 0:
        return index + 1
      case 1:
        return index - width
      case 2:
        return index - 1
      case 3:
        return index + width
      default:
        return index
    }
  })();

  return [i, d, l, s, t];
}

export function Spiral(dimensions, radius) {
  const max = Math.max(dimensions.x, dimensions.y)
  const width = Math.floor(max / radius)
  const length = width ** 2
  const offset = width % 2 === 0 ? Math.floor(width / 2) : 0
  const center = Math.floor(length / 2) + offset
  const getSpiralIndices = (acc, _, n) => n === 0 
    ? [[center, 0, 1, 1, 0]]
    : [...acc, next(width, ...acc[acc.length-1])]
  const toRenderFunction = (acc, [i], n) => {
    const position = (i % width, Math.floor(i / width)) * (radius * 2)
    if (
      position.x < 0 ||
      position.x > dimensions.x ||
      position.y < 0 ||
      position.y > dimensions.y ||
      !prime(n + 1)
    ) return acc

    return [...acc, Renderer {
      circle({ position, radius, fill: '#0F4761' })
    }]
  }

  return (context) => Array(length)
    .fill(0)
    .reduce(getSpiralIndices, [])
    .reduce(toRenderFunction, [])
    .map((fn) => context.apply(fn))
}