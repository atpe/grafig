// @ts-check

import eslint from '@eslint/js'
import tseslint from 'typescript-eslint'

export default tseslint.config(
  eslint.configs.recommended,
  tseslint.configs.recommended,
  tseslint.configs.strict,
  tseslint.configs.stylistic,
  [
    { ignores: ['./**', '!src/**'] },
    {
      rules: {
        '@typescript-eslint/explicit-function-return-type': 'error',
      },
    },
  ]
)
