function createCanvasElement_withoutParent_withDefaultDimensions() {
  Canvas()
}

function createCanvasElement_withoutParent_withNonDefaultDimensions() {
  Canvas({ dimensions: { x: 100, y: 100 } })
}

function createCanvasElement_withBodyParent_withMatchingDimensions() {
  Canvas({ parent: document.body })
}

function createCanvasElement_withBodyParent_withDifferingDimensions() {
  Canvas({ parent: document.body, dimensions: { x: 100, y: 100 } })
}

function createCanvasElement_withElementParent_withMatchingDimensions() {
  Canvas({ parent: document.getElementById('element-id') })
}

function renderCanvasElement_withBackground() {
  Canvas({}, (context) => context.background('black'))
}

function renderCanvasElement_withBackground_withOutlinedShapes() {
  Canvas({}, (context) => {
    context.background('black')
    context.square({
      position: { x: 0, y: 0 },
      dimensions: { x: 50, y: 50 },
      stroke: 'rgb(255, 255, 255)',
    })
    context.circle({
      position: { x: 75, y: 25 },
      radius: 25,
      stroke: 'rgb(255, 255, 255)',
    })
    context.square({
      position: { x: 50, y: 50 },
      dimensions: { x: 50, y: 50 },
      stroke: 'rgb(255, 255, 255)',
    })
    context.circle({
      position: { x: 25, y: 75 },
      radius: 25,
      stroke: 'rgb(255, 255, 255)',
    })
  })
}

function renderCanvasElement_withRandomCircleWithinBounds() {
  const radius = 10
  const random = { x: Math.random(), y: Math.random() }

  Canvas({}, (context) => {
    context.circle({
      position: add(mult(sub(context.dimensions, radius * 2), random), radius),
      radius,
      fill: 'rgb(0, 0, 0)',
    })
  })
}

function renderCanvasElement_withDynamicTasklist() {
  function TasklistItem(position, text, done, color) {
    return (context) => {
      context.square({
        position: sub(position, { x: 0, y: 20 }),
        size: 20,
        fill: done ? color : 'rgba(0, 0, 0, 0)',
        stroke: color,
      })
      context.text({
        text,
        position: add(position, { x: 25, y: 0 }),
        fill: color,
        stroke: 'rgb(255, 255, 255)',
      })
    }
  }

  const tasks = [
    ['some overdue task', false, 'rgb(250, 0, 0)'],
    ['some important task', false, 'rgb(250, 250, 0)'],
    ['some completed task', true, 'rgb(100, 250, 0)'],
    ['some abandoned task', true, 'rgb(200, 200, 200)'],
  ]

  Canvas({}, (context) => {
    context.apply(
      ...tasks.map((task, i) => {
        const center = div(context.dimensions, 2)
        const position = add(center, { x: 0, y: (i - tasks.length / 2) * 50 })
        TasklistItem(position, ...task)
      })
    )
  })
}
