const API_URL = 192.168.20.233;

/** @type {import('next').NextConfig} */
const nextConfig = {
	async rewrites() {
		return [
			{
				source: '/v2/*',
				destination: `${API_URL}/*`,
			},
		]
	},
}

module.exports = nextConfig
