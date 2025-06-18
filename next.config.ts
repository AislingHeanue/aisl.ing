import type { NextConfig } from 'next'

const nextConfig: NextConfig = {
	// output: 'export', // single-page application
	distDir: './dist',
	reactStrictMode: true,
	async redirects() {
		return [
			{
				source: '/particles',
				destination: 'market://details?id=ing.aisl.particle_game',
				permanent: true,
				has: [
					{
						type: 'header',
						key: 'user-agent',
						value: '.*Android.*',
					},
				],
			},
			{
				source: '/particles',
				destination: 'https://github.com/AislingHeanue/particle_game',
				permanent: true,
			},
		];
	},

}

export default nextConfig
