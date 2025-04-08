const CANVAS_SIZE = 500
const TILE_SIZE = 5
const POINT_COUNT = 100
const POINT_SIZE = 2
const MAX_SPEED = 10
const DRAG = 25
const DELTA = 1000 / 120

const canvas = document.createElement('canvas')
document.body.appendChild(canvas)
const ctx = canvas.getContext('2d')

canvas.width = CANVAS_SIZE
canvas.height = CANVAS_SIZE

const vector = (1, 2, 3)
const sum = vector + (2, 3, 4).magnitude

const magnitude = (v) => Math.sqrt(v[0] * v[0] + v[1] * v[1])

const normalize = (v, m = magnitude(v)) => [v[0] / m, v[1] / m]

const random = () => normalize([2 * Math.random() - 1, 2 * Math.random() - 1])

const scaled = () => [Math.random() * CANVAS_SIZE, Math.random() * CANVAS_SIZE]

const tiles = Array((CANVAS_SIZE * CANVAS_SIZE) / (TILE_SIZE * TILE_SIZE))
  .fill([])
  .map(random)

let points = Array(POINT_COUNT)
  .fill({})
  .map(() => ({
    position: scaled(),
    velocity: random(),
  }))

function draw() {
  points = points.map((point) => {
    const index =
      Math.floor(point.position[0] / TILE_SIZE) +
      Math.floor(point.position[1] / TILE_SIZE) * (CANVAS_SIZE / TILE_SIZE)

    const acceleration = tiles[index]
    const vx = point.velocity[0] + acceleration[0] * (DELTA / DRAG)
    const vy = point.velocity[1] + acceleration[1] * (DELTA / DRAG)

    const r = MAX_SPEED / magnitude([vx, vy])
    const velocity = r < 1 ? [vx * r, vy * r] : [vx, vy]

    const px = point.position[0] + velocity[0] * (DELTA / DRAG)
    const py = point.position[1] + velocity[1] * (DELTA / DRAG)

    const oob = 0 > px || px > CANVAS_SIZE || 0 > py || py > CANVAS_SIZE
    const position = oob ? scaled() : [px, py]

    return { position, velocity }
  })

  ctx.fillStyle = 'rgba(0, 0, 0, 0.1)'
  ctx.fillRect(0, 0, CANVAS_SIZE, CANVAS_SIZE)

  ctx.fillStyle = 'white'
  points.forEach((point) => {
    ctx.beginPath()
    ctx.ellipse(
      point.position[0],
      point.position[1],
      POINT_SIZE,
      POINT_SIZE,
      0,
      0,
      Math.PI * 2
    )
    ctx.closePath()
    ctx.fill()
  })
}

draw()
