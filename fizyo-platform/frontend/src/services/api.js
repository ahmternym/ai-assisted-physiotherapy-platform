import axios from 'axios'

// vue.config.js'deki devServer.proxy ayarı sayesinde
// '/api/...' istekleri otomatik olarak 'http://localhost:8080/api/...' adresine yönlendirilir.
// Bu şekilde CORS hatası almadan çalışır.
// Production build'da nginx veya benzeri bir reverse proxy aynı işi üstlenir.
const api = axios.create({
  baseURL: '/api',
  timeout: 8000,
  headers: { 'Content-Type': 'application/json' }
})

export default {
  // İle göre fizyoterapist listesi (max 5)
  getByCity(city) {
    return api.get('/physiotherapists', { params: { city } })
  },

  // ID'ye göre fizyoterapist detayı
  getById(id) {
    return api.get(`/physiotherapists/${id}`)
  }
}
