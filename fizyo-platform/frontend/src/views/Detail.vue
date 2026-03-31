<template>
  <div class="detail-view">
    <!-- Topbar -->
    <div class="detail-topbar">
      <button class="back-btn" @click="$router.go(-1)">← Geri</button>
      <div class="topbar-brand">🏥 FizyoBul</div>
    </div>

    <div v-if="loading" class="skeleton-detail">
      <div class="skel skel-avatar-lg"></div>
      <div class="skel skel-line w50"></div>
      <div class="skel skel-line w35"></div>
    </div>

    <div v-else-if="error" class="state-box">
      <span>⚠️</span>
      <p>{{ error }}</p>
    </div>

    <div v-else-if="physio" class="detail-container">
      <!-- Profil Kartı -->
      <div class="profile-card">
        <div class="profile-bg"></div>
        <div class="profile-avatar" :style="{ background: avatarGradient }">
          {{ initials }}
        </div>
        <h2 class="profile-name">{{ physio.name }}</h2>
        <p class="profile-sub">Fizyoterapist</p>

        <div class="profile-chips">
          <span class="chip chip-location">
            📍 {{ physio.district }}, {{ physio.city }}
          </span>
        </div>
      </div>

      <!-- Bilgi Kutuları -->
      <div class="info-grid">
        <div class="info-box">
          <div class="info-box-icon">📞</div>
          <div>
            <div class="info-box-label">Telefon</div>
            <div class="info-box-value">{{ physio.phone }}</div>
          </div>
        </div>
        <div class="info-box">
          <div class="info-box-icon">🏙️</div>
          <div>
            <div class="info-box-label">İl</div>
            <div class="info-box-value">{{ physio.city }}</div>
          </div>
        </div>
        <div class="info-box">
          <div class="info-box-icon">📌</div>
          <div>
            <div class="info-box-label">İlçe</div>
            <div class="info-box-value">{{ physio.district }}</div>
          </div>
        </div>
      </div>

      <!-- Uzmanlık Alanları -->
      <div class="specialties-card" v-if="physio.specialties && physio.specialties.length">
        <h3 class="section-title">
          <span class="section-icon">🎯</span>
          Uzmanlık Alanları
        </h3>
        <div class="specialty-grid">
          <div
            v-for="(spec, i) in physio.specialties"
            :key="i"
            class="specialty-item"
          >
            <span class="spec-bullet">✓</span>
            {{ spec }}
          </div>
        </div>
      </div>

      <!-- CTA Alanı -->
      <div class="cta-card">
        <div class="cta-info">
          <h4>Randevu Almak İster Misiniz?</h4>
          <p>Direkt iletişim için aşağıdaki butonu kullanabilirsiniz</p>
        </div>
        <a :href="'tel:' + physio.phone.replace(/\s/g,'')" class="call-btn">
          📞 Şimdi Ara
        </a>
      </div>
    </div>
  </div>
</template>

<script>
import api from '@/services/api'

const GRADIENTS = [
  'linear-gradient(135deg,#3b82f6,#1d4ed8)',
  'linear-gradient(135deg,#8b5cf6,#6d28d9)',
  'linear-gradient(135deg,#ec4899,#be185d)',
  'linear-gradient(135deg,#10b981,#047857)',
  'linear-gradient(135deg,#f59e0b,#b45309)',
  'linear-gradient(135deg,#06b6d4,#0e7490)',
]

export default {
  name: 'DetailView',
  data() {
    return { physio: null, loading: false, error: null }
  },
  computed: {
    initials() {
      if (!this.physio) return ''
      return this.physio.name.split(' ').map(n => n[0]).join('').substring(0, 2).toUpperCase()
    },
    avatarGradient() {
      if (!this.physio) return GRADIENTS[0]
      return GRADIENTS[this.physio.id % GRADIENTS.length]
    }
  },
  created() { this.fetchDetail() },
  methods: {
    async fetchDetail() {
      this.loading = true
      this.error = null
      try {
        const res = await api.getById(this.$route.params.id)
        this.physio = res.data
      } catch {
        this.error = 'Fizyoterapist bilgileri yüklenemedi.'
      } finally {
        this.loading = false
      }
    }
  }
}
</script>

<style scoped>
.detail-view {
  min-height: 100vh;
  background: #f0f4ff;
}

/* Topbar */
.detail-topbar {
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

.topbar-brand { font-weight: 800; color: #1e3a5f; font-size: 1rem; }

/* Skeleton */
.skeleton-detail {
  max-width: 520px;
  margin: 2rem auto;
  padding: 2rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
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
.skel-avatar-lg { width: 90px; height: 90px; border-radius: 50%; }
.skel-line { height: 14px; }
.w50 { width: 50%; }
.w35 { width: 35%; }

/* State */
.state-box {
  text-align: center;
  padding: 3rem;
  color: #64748b;
}

/* Container */
.detail-container {
  max-width: 560px;
  margin: 0 auto;
  padding: 2rem 1.2rem;
  display: flex;
  flex-direction: column;
  gap: 1.2rem;
}

/* Profil Kartı */
.profile-card {
  background: white;
  border-radius: 20px;
  padding: 2.5rem 2rem 2rem;
  text-align: center;
  position: relative;
  overflow: hidden;
  border: 1.5px solid #e8eef6;
  box-shadow: 0 4px 24px rgba(30,58,95,0.08);
}

.profile-bg {
  position: absolute;
  top: 0; left: 0; right: 0;
  height: 90px;
  background: linear-gradient(135deg, #0f2a4e, #1d4ed8);
}

.profile-avatar {
  position: relative;
  width: 88px;
  height: 88px;
  border-radius: 50%;
  color: white;
  font-weight: 800;
  font-size: 2rem;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 1rem;
  border: 4px solid white;
  box-shadow: 0 4px 20px rgba(0,0,0,0.18);
  letter-spacing: 0.03em;
}

.profile-name {
  font-size: 1.4rem;
  font-weight: 800;
  color: #0f172a;
  margin: 0 0 0.3rem;
}

.profile-sub {
  font-size: 0.82rem;
  color: #94a3b8;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.06em;
  margin: 0 0 1rem;
}

.profile-chips { display: flex; justify-content: center; flex-wrap: wrap; gap: 0.5rem; }

.chip {
  font-size: 0.8rem;
  font-weight: 600;
  border-radius: 99px;
  padding: 0.3rem 0.9rem;
}

.chip-location {
  background: #f0f4ff;
  color: #1d4ed8;
  border: 1px solid #c7d7fe;
}

/* Info Grid */
.info-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 0.8rem;
}

.info-box {
  background: white;
  border-radius: 14px;
  padding: 1rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  gap: 0.5rem;
  border: 1.5px solid #e8eef6;
  box-shadow: 0 2px 10px rgba(30,58,95,0.05);
}

.info-box-icon { font-size: 1.4rem; }
.info-box-label { font-size: 0.7rem; color: #94a3b8; text-transform: uppercase; letter-spacing: 0.06em; font-weight: 600; }
.info-box-value { font-size: 0.85rem; color: #1e293b; font-weight: 700; margin-top: 0.2rem; }

/* Uzmanlık */
.specialties-card {
  background: white;
  border-radius: 20px;
  padding: 1.5rem;
  border: 1.5px solid #e8eef6;
  box-shadow: 0 2px 12px rgba(30,58,95,0.05);
}

.section-title {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 1rem;
  font-weight: 800;
  color: #1e3a5f;
  margin: 0 0 1rem;
}

.section-icon { font-size: 1.1rem; }

.specialty-grid {
  display: flex;
  flex-direction: column;
  gap: 0.6rem;
}

.specialty-item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  font-size: 0.88rem;
  color: #334155;
  font-weight: 500;
  background: #f8faff;
  border-radius: 10px;
  padding: 0.65rem 1rem;
  border: 1px solid #e8eef6;
}

.spec-bullet {
  width: 22px; height: 22px;
  background: linear-gradient(135deg, #3b82f6, #1d4ed8);
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 0.7rem;
  font-weight: 800;
  flex-shrink: 0;
}

/* CTA */
.cta-card {
  background: linear-gradient(135deg, #0f2a4e, #1d4ed8);
  border-radius: 20px;
  padding: 1.5rem;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  flex-wrap: wrap;
  color: white;
}

.cta-info h4 { margin: 0 0 0.3rem; font-size: 0.95rem; font-weight: 700; }
.cta-info p { margin: 0; font-size: 0.78rem; opacity: 0.7; }

.call-btn {
  display: inline-flex;
  align-items: center;
  gap: 0.4rem;
  background: white;
  color: #1d4ed8;
  font-weight: 700;
  font-size: 0.9rem;
  border-radius: 12px;
  padding: 0.65rem 1.4rem;
  text-decoration: none;
  transition: all 0.2s;
  white-space: nowrap;
}

.call-btn:hover {
  background: #eff6ff;
  transform: scale(1.03);
}

@media (max-width: 480px) {
  .info-grid { grid-template-columns: repeat(2, 1fr); }
  .cta-card { flex-direction: column; text-align: center; }
}
</style>
