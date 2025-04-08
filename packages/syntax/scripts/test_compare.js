const { readFileSync, writeFileSync } = require('fs')

const matrix = [
  [
    'Path',
    'Type',
    'Name',
    'Line',
    'SLOC',
    'LLOC',
    'Params',
    'Cyclomatic Complexity',
    'Cyclomatic Density',
    'Halstead Operators Distinct',
    'Halstead Operators Total',
    'Halstead Operands Distinct',
    'Halstead Operands Total',
    'Halstead Vocabulary',
    'Halstead Length',
    'Halstead Volume',
    'Halstead Difficulty',
    'Halstead Effort',
    'Halstead Time',
    'Halstead Bugs',
  ],
  ...importReports('control', controlReducer),
  ...importReports('test', testReducer),
]

writeFileSync('test/out.csv', matrix.reduce(matrixReducer, ''))

//

function importReports(path, reducer) {
  const data = JSON.parse(readFileSync(`./test/${path}.json`).toString())
  const reports = data instanceof Array ? data : data.reports
  return reports.reduce(reducer, [])
}

function controlReducer(acc, report) {
  const reportMapper =
    (path) =>
    ({ name, line, sloc, params, cyclomatic, cyclomaticDensity, halstead }) =>
      [
        path.replace(process.cwd() + '/', ''),
        'control',
        name,
        line,
        sloc.physical,
        sloc.logical,
        params,
        cyclomatic,
        cyclomaticDensity,
        halstead.operators.distinct,
        halstead.operators.total,
        halstead.operands.distinct,
        halstead.operands.total,
        halstead.vocabulary,
        halstead.length,
        halstead.volume,
        halstead.difficulty,
        halstead.effort,
        halstead.time,
        halstead.bugs,
      ]
  return [...acc, ...report.functions.map(reportMapper(report.path))]
}

function testReducer(acc, report) {
  if (!report.functions) return acc
  const reportMapper = ({
    file,
    function: name,
    line,
    sloc,
    lloc,
    params,
    cyclomatic,
    halstead,
  }) => [
    file,
    'test',
    name,
    line,
    sloc,
    lloc,
    params,
    cyclomatic.complexity,
    cyclomatic.density,
    halstead.operatorsUnique,
    halstead.operatorsTotal,
    halstead.operandsUnique,
    halstead.operandsTotal,
    halstead.vocabulary,
    halstead.length,
    halstead.volume,
    halstead.difficulty,
    halstead.effort,
    halstead.time,
    halstead.bugs,
  ]
  return [...acc, ...report.functions.map(reportMapper)]
}

function matrixReducer(acc, row) {
  return `${acc}${row.reduce((acc, col) => `${acc}, ${col}`)}\n`
}
