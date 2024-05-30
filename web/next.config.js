// @ts-check

const isProduction = process.env.NODE_ENV === 'production';

 
/** @type {import('next').NextConfig} */
const nextConfig = {
  // dev proxy
  rewrites: async () => {
    return [
      {
        source: '/:path*',
        destination: 'http://10.0.0.83/:path*',
      },
    ]
  },

  basePath: isProduction ? '/modules/zimaos_status' : '',
  assetPrefix: isProduction ? '/modules/zimaos_status/' : '',

  distDir: '../raw/usr/share/casaos/www/modules/zimaos_status',
  output: 'export',
	/* config options here */
}
   
module.exports = nextConfig