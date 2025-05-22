import FigScript, { Canvas } from '@grafig/lib' 

import { reduceChar } from './fenstring'
import { Tile } from './components'

const start = 'rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1'

const params = new URLSearchParams(window.location.search)
const fenstr = params.get('fenstr') ?? start

// renders the board to the canvas
Canvas({ parent: document.body }) {
  background('#0F4761')
  apply(
    ...fenstr
      .split(' ')[0] // get piece placement segment
      .split('') // get chars
      .reduce(reduceChar, [[], (0, 0)])[0] // get tile data
      .map(Tile) // render tiles
  )
}()
