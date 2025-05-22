# FigScript

When using `grafig`, code is able to be written in an extended version of vanilla JavaScript called _FigScript_. This allows the use of syntax commonly available in other applications of vector graphics such as operator overloads and more.

> The use of this syntax relies on the default `FigScript` export from `@grafig/lib` being imported into any _FigScript_ file - e.g.
>
> ```js
> import FigScript from '@grafig/lib'
> ```

## Reasoning

The `<canvas>` element provides rudimentary apis, that create complex and poorly maintainable code. Libraries alone can help, but extending the existing syntax of JavaScript allows more domain specific functionality.

## Syntax

Given that the _FigScript_ syntax extends that of JavaScript, any vanilla javascript is usable by default. The `@grafig/syntax` package uses _ANTLR_ to parse any JavaScript and, by extension, _FigScript_ code it encounters.

### Vector Literals

#### `(number, number, number?)`

#### Usage

The `@grafig/lib` package defines vectors using the following type:

```ts
type Vec = {
  x: number,
  y: number,
  z: number?,
}
```

Any occurances of a vector literal are equivalent to this and may be used as such.

#### Example

```js
const one = 1

const vector2d = (1e3, one / 200)
console.log(vector2d.x) // 1000
console.log(vector2d.y) // 0.005
console.log(vector2d.z) // undefined

const vector3d = (
  vector2d.x, 
  vector2d.x, 
  Math.random() * (1-1)
)
console.log(vector3d.x) // 1000
console.log(vector3d.y) // 1000
console.log(vector3d.z) // 0
```

### Operator Overloads

#### `vector + [number | vector]`

#### `vector - [number | vector]`

#### `vector / [number | vector]`

#### `vector * [number | vector]`

#### Usage

For any of the operators above, if the first argument given is a vector, it will apply the vector calculation to the operands as defined in the `@grafig/lib` package.

#### Example

```js
const vector = (1, 2, 3) + 5
console.log(vector.x) // 6
console.log(vector.y) // 7
console.log(vector.z) // 8

const divided = vector / (6, 7, 8)
console.log(divided.x) // 1
console.log(divided.y) // 1
console.log(divided.z) // 1
```

### Swizzling

#### `vector.xx`

#### `vector.xy`

#### `vector.xz`

#### `vector.yx`

#### `...`

#### `vector.zzy`

#### `vector.zzz`

#### Usage

To create a new vector by _swizzling_, any combination of `x`, `y`, or `z` can be used to copy the component values from the target vector into a new 2D or 3D vector.

#### Example

```js
const vector = (1, 2, 3)
console.log(vector.x) // 1
console.log(vector.y) // 2
console.log(vector.z) // 3

const swizzled = vector.yzx
console.log(swizzled.x) // 2
console.log(swizzled.y) // 3
console.log(swizzled.z) // 1
```

### Trailing Lambdas

#### `fn(...args) { ... }`

#### Usage

To simplify function calls, if the last argument is itself a function with a single parameter, each statement is interpreted as a method call upon that parameter. While not as comprehensive as their Kotlin counterparts, they still provide a benefit in their early form.

#### Example

```ts
// for the given type 'Dog' and function 'interact'
interface Dog {
  giveTreat()
  giveTreats(n: number)
  rubBelly()
  throwBall()
}
function interact(dog: Dog, fn: (dog: Dog) => void) => void {
  ...
}

// the 'interact' function can then be called as follows:
interact(dog) {
  giveTreat()
  rubBelly()
  giveTreats(2)
  throwBall()
  giveTreat()
}
```
