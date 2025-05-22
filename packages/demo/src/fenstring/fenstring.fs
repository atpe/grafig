import FigScript from '@grafig/lib'

// creates piece data from char
export function createPiece(char) {
  const type = char.toUpperCase()
  return { type, white: char === type }
}

// creates tile data from position and char
export function createTile(position, char) {
  const piece = char ? createPiece(char) : undefined
  const white = (position.x + position.y) % 2 === 0
  return { position, white, piece }
}

// converts chars into an array
export function reduceChar([tiles, position], char) {
  // if '/', go to next row
  if (char === '/') return [tiles, position + (-8, 1)]

  // if number, add empty tiles
  const n = parseInt(char)
  if (!isNaN(n)) {
    const empty = Array(n)
      .fill(0)
      .map((_, i) => createTile(position + (i, 0)))
    return [[...tiles, ...empty], position + (n, 0)]
  }

  // if char, add tile with piece
  return [[...tiles, createTile(position, char)], position + (1, 0)]
}
