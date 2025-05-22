function createCanvasElement_withoutParent_withDefaultDimensions() {
  new p5((context) => {
    context.setup = () => {
      context.createCanvas()
      context.noLoop()
    }
  })
}

function createCanvasElement_withoutParent_withNonDefaultDimensions() {
  new p5((context) => {
    context.setup = () => {
      context.createCanvas(100, 100)
      context.noLoop()
    }
  })
}

function createCanvasElement_withBodyParent_withMatchingDimensions() {
  const { clientWidth, clientHeight } = document.body
  new p5((context) => {
    context.setup = () => {
      context.createCanvas(clientWidth, clientHeight)
      context.noLoop()
    }
  }, document.body)
}

function createCanvasElement_withBodyParent_withDifferingDimensions() {
  new p5((context) => {
    context.setup = () => {
      context.createCanvas(100, 100)
      context.noLoop()
    }
  }, document.body)
}

function createCanvasElement_withElementParent_withMatchingDimensions() {
  const element = document.getElementById('element-id')
  if (context == null) {
    throw new Error('cannot use element as it does not exist')
  }

  const { clientWidth, clientHeight } = element

  new p5((context) => {
    context.setup = () => {
      context.createCanvas(clientWidth, clientHeight)
      context.noLoop()
    }
  }, element)
}

function renderCanvasElement_withBackground() {
  new p5((context) => {
    context.setup = () => {
      context.createCanvas()
      context.noLoop()
    }
    context.draw = () => context.background(0)
  })
}

function renderCanvasElement_withBackground_withOutlinedShapes() {
  new p5((context) => {
    context.setup = () => {
      context.createCanvas()
      context.noLoop()
    }
    context.draw = () => {
      context.background(0)
      context.noFill()
      context.stroke(255)
      context.square(0, 0, 50)
      context.circle(75, 25, 25)
      context.circle(25, 75, 25)
      context.square(50, 50, 50)
    }
  })
}

function renderCanvasElement_withRandomCircleWithinBounds() {
  const radius = 10

  new p5((context) => {
    context.setup = () => {
      context.createCanvas()
      context.noLoop()
    }
    context.draw = () => {
      const position = context
        .createVector(context.width, context.height)
        .sub(radius * 2)
        .mult(context.random2D())
        .add(radius)

      context.fill(255)
      context.circle(position.x, position.y, radius)
    }
  })
}

function renderCanvasElement_withDynamicTasklist() {
  function TasklistItem(x, y, text, done, color) {
    return (context) => {
      if (done) context.fill(color)
      else context.noFill()
      context.stroke(color)
      context.square(x, y - 20, 20)

      context.stroke('rgb(0, 0, 0)')
      context.fill(color)
      context.text(text, x + 25, y)
    }
  }

  const tasks = [
    ['some overdue task', false, 'rgb(250, 0, 0)'],
    ['some important task', false, 'rgb(250, 250, 0)'],
    ['some completed task', true, 'rgb(100, 250, 0)'],
    ['some abandoned task', true, 'rgb(200, 200, 200)'],
  ]

  new p5((context) => {
    context.setup = () => {
      context.createCanvas()
      context.noLoop()
    }
    context.draw = () => {
      const center = context.createVector(context.width, context.height).div(2)
      tasks.forEach((task, i) =>
        TasklistItem(
          center.x,
          center.y + (i - tasks.length / 2) * 50,
          ...task
        )(context)
      )
    }
  })
}
