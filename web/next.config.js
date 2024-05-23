// @ts-check
 
/** @type {import('next').NextConfig} */
const nextConfig = {
  basePath: '/modules/zimaos_status', // 将 '/myapp' 替换为你的子路径
  assetPrefix: '/modules/zimaos_status/',
  distDir: '../raw/usr/share/casaos/www/modules/zimaos_status',
  output: 'export',
	/* config options here */
}
   
module.exports = nextConfig