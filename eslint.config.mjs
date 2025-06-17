import { FlatCompat } from '@eslint/eslintrc'

const compat = new FlatCompat({
	baseDirectory: import.meta.dirname,
})

const eslintConfig = [
	...compat.config({
		extends: ['next', 'next/typescript'],
		rules: {
			'@typescript-eslint/no-unused-vars': 'off',
			'no-unused-vars': 'off',
		},
	}),
]

export default eslintConfig
