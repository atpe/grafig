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
  const mapper =
    ({ path } = report) =>
    (report) => {
      const r = report.aggregate ?? report
      return [
        path.replace(process.cwd() + '/', ''),
        'control',
        report.name ?? 'program',
        r.line,
        r.sloc.physical,
        r.sloc.logical,
        r.params,
        r.cyclomatic,
        r.cyclomaticDensity / 100,
        r.halstead.operators.distinct,
        r.halstead.operators.total,
        r.halstead.operands.distinct,
        r.halstead.operands.total,
        r.halstead.vocabulary,
        r.halstead.length,
        r.halstead.volume,
        r.halstead.difficulty,
        r.halstead.effort,
        r.halstead.time,
        r.halstead.bugs,
      ]
    }

  return report.functions
    ? [...acc, mapper()(report), ...report.functions.map(mapper(report))]
    : [...acc, mapper()(report())]
}

function testReducer(acc, report) {
  console.log(report)

  const mapper = (report) => {
    return [
      report.file,
      'test',
      report.function,
      report.line,
      report.sloc,
      report.lloc,
      report.params,
      report.cyclomatic.complexity,
      report.cyclomatic.density,
      report.halstead.operatorsUnique,
      report.halstead.operatorsTotal,
      report.halstead.operandsUnique,
      report.halstead.operandsTotal,
      report.halstead.vocabulary,
      report.halstead.length,
      report.halstead.volume,
      report.halstead.difficulty,
      report.halstead.effort,
      report.halstead.time,
      report.halstead.bugs,
    ]
  }
  return report.functions
    ? report.functions.reduce(testReducer, [...acc, mapper(report)])
    : [...acc, mapper(report)]
}

function matrixReducer(acc, row) {
  return `${acc}${row.reduce((acc, col) => `${acc}, ${col}`)}\n`
}
