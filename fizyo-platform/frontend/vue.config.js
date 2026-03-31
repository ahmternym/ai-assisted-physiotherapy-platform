const { defineConfig } = require('@vue/cli-service')

module.exports = defineConfig({
  transpileDependencies: true,

  // public/index.html'i Vue CLI'ye tanıt
  indexPath: 'index.html',

  devServer: {
    port: 5173,
    // /api/* isteklerini backend'e (Go) proxy'le — böylece CORS sorunu olmaz
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        logLevel: 'debug'
      }
    }
  }
})
