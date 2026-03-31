<template>
  <div class="list-view">
    <!-- Üst Bar -->
    <div class="list-topbar">
      <button class="back-btn" @click="$router.push('/')">
        <span>←</span> Ana Sayfa
      </button>
      <div class="topbar-brand">🏥 FizyoBul</div>
    </div>

    <!-- Başlık Alanı -->
    <div class="list-hero">
      <div class="city-badge">📍 {{ city }}</div>
      <h2>Fizyoterapistler</h2>
      <p v-if="!loading && physiotherapists.length">
        {{ city }} ilinde <strong>{{ physiotherapists.length }}</strong> uzman bulundu
      </p>
    </div>

    <div class="list-body">
      <!-- Yükleniyor: skeleton -->
      <div v-if="loading" class="skeleton-list">
        <div v-for="i in 5" :key="i" class="skeleton-card">
          <div class="skel skel-avatar"></div>
          <div class="skel-info">
            <div class="skel skel-line w70"></div>
            <div class="skel skel-line w50"></div>
            <div class="skel skel-line w40"></div>
          </div>
        </div>
      </div>

      <!-- Hata -->
      <div v-else-if="error" class="state-box error-box">
        <span class="state-icon">⚠️</span>
        <h3>Bir Hata Oluştu</h3>
        <p>{{ error }}</p>
        <button class="retry-btn" @click="fetchList">Tekrar Dene</button>
      </div>

      <!-- Boş -->
      <div v-else-if="physiotherapists.length === 0" class="state-box empty-box">
        <span class="state-icon">🔍</span>
        <h3>Fizyoterapist Bulunamadı</h3>
        <p>{{ city }} iline ait kayıt henüz mevcut değil.</p>
        <button class="retry-btn" @click="$router.push('/')">Başka İl Seç</button>
      </div>

      <!-- Kart Listesi -->
      <div v-else class="card-list">
        <!--
          @select-physio: PhysiotherapistCard'ın emit ettiği özel event.
          Eski @click kullanımı Vue 3'te native click ile çakışıyor ve
          id yerine MouseEvent objesi geliyordu — bu yüzden detay sayfası açılmıyordu.
        -->
        <PhysiotherapistCard
          v-for="p in physiotherapists"
          :key="p.id"
          :physio="p"
          @select-physio="goToDetail"
        />
      </div>
    </div>
  </div>
</template>

<script>
import PhysiotherapistCard from '@/components/PhysiotherapistCard.vue'
import api from '@/services/api'

export default {
  name: 'ListView',
  components: { PhysiotherapistCard },

  data() {
    return {
      physiotherapists: [],
      loading: false,
      error: null
    }
  },

  computed: {
    city() { return this.$route.params.city }
  },

  created() { this.fetchList() },

  methods: {
    async fetchList() {
      this.loading = true
      this.error = null
      try {
        const res = await api.getByCity(this.city)
        this.physiotherapists = res.data
      } catch (err) {
        console.error('fetchList hatası:', err)
        this.error = 'Veriler yüklenemedi. Backend çalışıyor mu? (go run main.go)'
      } finally {
        this.loading = false
      }
    },

    goToDetail(id) {
      // id: PhysiotherapistCard'dan gelen number
      this.$router.push({ name: 'Detail', params: { id: String(id) } })
    }
  }
}
</script>

<style scoped>
.list-view { min-height: 100vh; background: #f0f4ff; }

.list-topbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0.9rem 1.5rem;
  background: rgba(255,255,255,0.9);
  backdrop-filter: blur(12px);
  border-bottom: 1px solid rgba(59,130,246,0.1);
  position: sticky;
  top: 0;
  z-index: 50;
}

.back-btn {
  display: flex;
  align-items: center;
  gap: 0.4rem;
  background: none;
  border: 1.5px solid #e2e8f0;
  color: #475569;
  border-radius: 10px;
  padding: 0.45rem 1rem;
  cursor: pointer;
  font-size: 0.88rem;
  font-weight: 600;
  transition: all 0.2s;
}
.back-btn:hover { background: #f1f5f9; border-color: #3b82f6; color: #3b82f6; }

.topbar-brand { font-weight: 800; font-size: 1rem; color: #1e3a5f; }

.list-hero {
  background: linear-gradient(135deg, #0f2a4e 0%, #1d4ed8 100%);
  padding: 2rem 1.5rem;
  text-align: center;
  color: white;
}

.city-badge {
  display: inline-block;
  background: rgba(255,255,255,0.18);
  border: 1px solid rgba(255,255,255,0.3);
  border-radius: 99px;
  padding: 0.3rem 1rem;
  font-size: 0.8rem;
  font-weight: 600;
  margin-bottom: 0.7rem;
}

.list-hero h2 { font-size: 1.7rem; font-weight: 900; margin: 0 0 0.4rem; }
.list-hero p { font-size: 0.9rem; opacity: 0.8; margin: 0; }
.list-hero p strong { opacity: 1; }

.list-body { max-width: 700px; margin: 0 auto; padding: 1.8rem 1.2rem; }

/* Skeleton */
.skeleton-list { display: flex; flex-direction: column; gap: 0.8rem; }
.skeleton-card {
  display: flex; align-items: center; gap: 1rem;
  background: white; border-radius: 16px; padding: 1rem 1.2rem;
  border: 1.5px solid #e8eef6;
}
.skel {
  background: linear-gradient(90deg, #f1f5f9 25%, #e2e8f0 50%, #f1f5f9 75%);
  background-size: 200% 100%;
  animation: shimmer 1.4s infinite;
  border-radius: 8px;
}
@keyframes shimmer {
  0% { background-position: 200% 0; }
  100% { background-position: -200% 0; }
}
.skel-avatar { width: 54px; height: 54px; border-radius: 50%; flex-shrink: 0; }
.skel-info { flex: 1; display: flex; flex-direction: column; gap: 0.5rem; }
.skel-line { height: 12px; }
.w70 { width: 70%; }
.w50 { width: 50%; }
.w40 { width: 40%; }

/* State box */
.state-box {
  text-align: center; background: white; border-radius: 20px;
  padding: 3rem 2rem; border: 1.5px solid #e8eef6;
  box-shadow: 0 4px 20px rgba(30,58,95,0.06);
}
.state-icon { font-size: 2.5rem; display: block; margin-bottom: 1rem; }
.state-box h3 { font-size: 1.1rem; color: #1e293b; margin: 0 0 0.5rem; }
.state-box p { font-size: 0.88rem; color: #64748b; margin: 0 0 1.5rem; }

.retry-btn {
  background: #1d4ed8; color: white; border: none;
  border-radius: 10px; padding: 0.6rem 1.5rem;
  font-size: 0.9rem; font-weight: 600; cursor: pointer;
  transition: background 0.2s;
}
.retry-btn:hover { background: #1e40af; }

.card-list { display: flex; flex-direction: column; gap: 0.8rem; }
</style>
