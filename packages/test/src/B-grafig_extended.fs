function createCanvasElement_withoutParent_withDefaultDimensions() {
  Canvas()
}

function createCanvasElement_withoutParent_withNonDefaultDimensions() {
  Canvas({ dimensions: (100, 100) })
}

function createCanvasElement_withBodyParent_withMatchingDimensions() {
  Canvas({ parent: document.body })
}

function createCanvasElement_withBodyParent_withDifferingDimensions() {
  Canvas({ parent: document.body, dimensions: (100, 100) })
}

function createCanvasElement_withElementParent_withMatchingDimensions() {
  Canvas({ parent: document.getElementById('element-id') })
}

function renderCanvasElement_withBackground() {
  Canvas { background('black') }
}

function renderCanvasElement_withBackground_withOutlinedShapes() {
  Canvas {
    background('black')
    square({
      position: (0, 0),
      dimensions: (50, 50),
      stroke: 'rgb(255, 255, 255)',
    })
    circle({
      position: (75, 25),
      radius: 25,
      stroke: 'rgb(255, 255, 255)',
    })
    square({
      position: (50, 50),
      dimensions: (50, 50),
      stroke: 'rgb(255, 255, 255)',
    })
    circle({
      position: (25, 75),
      radius: 25,
      stroke: 'rgb(255, 255, 255)',
    })
  }
}

function renderCanvasElement_withRandomCircleWithinBounds() {
  const radius = 10
  const random = (Math.random(), Math.random())

  Canvas {
    circle({
      position: (it.dimensions - radius * 2) * random + radius,
      radius,
      fill: 'rgb(0, 0, 0)',
    })
  }
}

function renderCanvasElement_withDynamicTasklist() {
  function TasklistItem(position, text, done, color) {
    return Renderer {
      square({
        position: position - (0, 20),
        size: 20,
        fill: done ? color : 'rgba(0, 0, 0, 0)',
        stroke: color,
      })
      text({
        text,
        position: position + (25, 0),
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

  Canvas {
    apply(
      ...tasks.map((task, i) => {
        const center = context.dimensions / 2
        const position = center + (0, (i - tasks.length / 2) * 50)
        TasklistItem(position, ...task)
      })
    )
  }
}