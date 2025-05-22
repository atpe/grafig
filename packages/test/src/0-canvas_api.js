function createCanvasElement_withoutParent_withDefaultDimensions() {
  const canvas = document.createElement('canvas')

  const context = canvas.getContext('2d')
  if (context == null) {
    throw new Error('cannot use canvas with null context')
  }
}

function createCanvasElement_withoutParent_withNonDefaultDimensions() {
  const canvas = document.createElement('canvas')

  const context = canvas.getContext('2d')
  if (context == null) {
    throw new Error('cannot use canvas with null context')
  }

  canvas.width = 100
  canvas.height = 100
}

function createCanvasElement_withBodyParent_withMatchingDimensions() {
  const canvas = document.createElement('canvas')

  const context = canvas.getContext('2d')
  if (context == null) {
    throw new Error('cannot use canvas with null context')
  }

  canvas.width = document.body.clientWidth
  canvas.height = document.body.clientHeight
  document.body.addChild(canvas)
}

function createCanvasElement_withBodyParent_withDifferingDimensions() {
  const canvas = document.createElement('canvas')

  const context = canvas.getContext('2d')
  if (context == null) {
    throw new Error('cannot use canvas with null context')
  }

  canvas.width = 100
  canvas.height = 100
  document.body.addChild(canvas)
}

function createCanvasElement_withElementParent_withMatchingDimensions() {
  const canvas = document.createElement('canvas')

  const context = canvas.getContext('2d')
  if (context == null) {
    throw new Error('cannot use canvas with null context')
  }

  const element = document.getElementById('element-id')
  if (context == null) {
    throw new Error('cannot use element as it does not exist')
  }

  canvas.width = element.clientWidth
  canvas.height = element.clientHeight
  element.addChild(canvas)
}

function renderCanvasElement_withBackground() {
  const canvas = document.createElement('canvas')

  const context = canvas.getContext('2d')
  if (context == null) {
    throw new Error('cannot use canvas with null context')
  }

  context.fillStyle = 'rgb(0, 0, 0)'
  context.beginPath()
  context.fillRect(0, 0, canvas.width, canvas.height)
  context.closePath()
}

function renderCanvasElement_withBackground_withOutlinedShapes() {
  const canvas = document.createElement('canvas')

  const context = canvas.getContext('2d')
  if (context == null) {
    throw new Error('cannot use canvas with null context')
  }

  context.fillStyle = 'rgb(0, 0, 0)'

  context.beginPath()
  context.fillRect(0, 0, canvas.width, canvas.height)
  context.closePath()

  context.fillStyle = 'rgba(0, 0, 0, 0)'
  context.strokeStyle = 'rgb(255, 255, 255)'

  context.beginPath()
  context.strokeRect(0, 0, 50, 50)
  context.closePath()

  context.beginPath()
  context.ellipse(75, 25, 25, 25, 0, 0, Math.PI * 2)
  context.stroke()
  context.closePath()

  context.beginPath()
  context.ellipse(25, 75, 25, 25, 0, 0, Math.PI * 2)
  context.stroke()
  context.closePath()

  context.beginPath()
  context.strokeRect(50, 50, 50, 50)
  context.closePath()
}

function renderCanvasElement_withRandomCircleWithinBounds() {
  const canvas = document.createElement('canvas')

  const context = canvas.getContext('2d')
  if (context == null) {
    throw new Error('cannot use canvas with null context')
  }

  const radius = 10

  const x = radius + Math.random() * (canvas.width - radius * 2)
  const y = radius + Math.random() * (canvas.height - radius * 2)

  context.fillStyle = 'rgb(255, 255, 255)'
  context.beginPath()
  context.ellipse(x, y, radius, radius, 0, 0, Math.PI * 2)
  context.fill()
  context.closePath()
}

function renderCanvasElement_withDynamicTasklist() {
  function TasklistItem(x, y, text, done, color) {
    return (context) => {
      context.fillStyle = done ? color : 'rgba(0, 0, 0, 0)'
      context.strokeStyle = color
      context.beginPath()
      context.strokeRect(x, y - 20, 20, 20)
      context.closePath()

      context.strokeStyle = 'rgb(0, 0, 0)'
      context.fillStyle = color
      context.beginPath()
      context.fillText(text, x + 25, y)
      context.closePath()
    }
  }

  const canvas = document.createElement('canvas')

  const context = canvas.getContext('2d')
  if (context == null) {
    throw new Error('cannot use canvas with null context')
  }

  const x = canvas.width / 2
  const y = canvas.height / 2

  const tasks = [
    ['some overdue task', false, 'rgb(250, 0, 0)'],
    ['some important task', false, 'rgb(250, 250, 0)'],
    ['some completed task', true, 'rgb(100, 250, 0)'],
    ['some abandoned task', true, 'rgb(200, 200, 200)'],
  ]

  tasks.forEach((task, i) =>
    TasklistItem(x, y + (i - tasks.length / 2) * 50, ...task)(context)
  )
}
