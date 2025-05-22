export default async () => {
  return {
    verbose: true,
    coverageProvider: 'v8',
    coverageReporters: ['json-summary', 'text', 'lcov'],
    coveragePathIgnorePatterns: ['index.ts'],
  }
}
