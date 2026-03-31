<template>
  <div class="home-view">
    <!-- Navbar -->
    <nav class="navbar">
      <div class="nav-brand">
        <span class="brand-icon">🏥</span>
        <span class="brand-name">FizyoBul</span>
      </div>
      <div class="nav-tagline">Türkiye'nin Fizyoterapi Rehberi</div>
    </nav>

    <!-- Hero -->
    <header class="hero">
      <div class="hero-bg-shapes">
        <div class="shape shape-1"></div>
        <div class="shape shape-2"></div>
        <div class="shape shape-3"></div>
      </div>
      <div class="hero-content">
        <div class="hero-badge">
          <span>🇹🇷</span>
          <span>81 İlde Aktif</span>
        </div>
        <h1>
          Size En Yakın<br>
          <span class="hero-accent">Fizyoterapisti</span> Bulun
        </h1>
        <p class="hero-sub">
          Uzman fizyoterapistleri keşfedin, karşılaştırın ve hemen iletişime geçin.
        </p>
        <div class="hero-stats">
          <div class="stat">
            <span class="stat-num">405+</span>
            <span class="stat-label">Uzman</span>
          </div>
          <div class="stat-divider"></div>
          <div class="stat">
            <span class="stat-num">81</span>
            <span class="stat-label">İl</span>
          </div>
          <div class="stat-divider"></div>
          <div class="stat">
            <span class="stat-num">8</span>
            <span class="stat-label">Uzmanlık</span>
          </div>
        </div>
      </div>
    </header>

    <!-- Ana İçerik -->
    <main class="main-content">
      <div class="card-panel">
        <MapSelector @city-selected="onCitySelected" />
      </div>

      <!-- Özellikler -->
      <section class="features">
        <div class="feature-card">
          <span class="feat-icon">🗺️</span>
          <h4>Haritadan Seç</h4>
          <p>Türkiye haritasından ilini seçerek bölgenizdeki uzmanları bulun</p>
        </div>
        <div class="feature-card">
          <span class="feat-icon">⚡</span>
          <h4>Anında Sonuç</h4>
          <p>İl seçimi sonrası size en uygun 5 fizyoterapist anında listelenir</p>
        </div>
        <div class="feature-card">
          <span class="feat-icon">📋</span>
          <h4>Detaylı Profil</h4>
          <p>Uzmanlık alanları, iletişim bilgileri ve konum detaylarına ulaşın</p>
        </div>
      </section>
    </main>

    <footer class="site-footer">
      <p>© 2026 FizyoBul · Türkiye Fizyoterapi Platformu</p>
    </footer>
  </div>
</template>

<script>
import MapSelector from '@/components/MapSelector.vue'

export default {
  name: 'HomeView',
  components: { MapSelector },
  methods: {
    onCitySelected(city) {
      this.$router.push({ name: 'List', params: { city } })
    }
  }
}
</script>

<style scoped>
.home-view {
  min-height: 100vh;
  background: #f0f4ff;
  display: flex;
  flex-direction: column;
}

/* Navbar */
.navbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1rem 2rem;
  background: rgba(255,255,255,0.85);
  backdrop-filter: blur(12px);
  border-bottom: 1px solid rgba(59,130,246,0.1);
  position: sticky;
  top: 0;
  z-index: 100;
}

.nav-brand {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-weight: 800;
  font-size: 1.2rem;
  color: #1e3a5f;
}

.brand-icon { font-size: 1.4rem; }

.nav-tagline {
  font-size: 0.8rem;
  color: #94a3b8;
  font-weight: 500;
}

/* Hero */
.hero {
  background: linear-gradient(135deg, #0f2a4e 0%, #1d4ed8 60%, #7c3aed 100%);
  padding: 4rem 2rem 3rem;
  text-align: center;
  color: white;
  position: relative;
  overflow: hidden;
}

.hero-bg-shapes { position: absolute; inset: 0; pointer-events: none; }

.shape {
  position: absolute;
  border-radius: 50%;
  opacity: 0.08;
}
.shape-1 {
  width: 400px; height: 400px;
  background: white;
  top: -100px; right: -100px;
}
.shape-2 {
  width: 250px; height: 250px;
  background: #a78bfa;
  bottom: -80px; left: -60px;
}
.shape-3 {
  width: 150px; height: 150px;
  background: #60a5fa;
  top: 40%; left: 40%;
}

.hero-content { position: relative; z-index: 1; max-width: 640px; margin: 0 auto; }

.hero-badge {
  display: inline-flex;
  align-items: center;
  gap: 0.4rem;
  background: rgba(255,255,255,0.15);
  backdrop-filter: blur(8px);
  border: 1px solid rgba(255,255,255,0.25);
  border-radius: 99px;
  padding: 0.4rem 1.1rem;
  font-size: 0.82rem;
  font-weight: 600;
  margin-bottom: 1.2rem;
  letter-spacing: 0.03em;
}

.hero h1 {
  font-size: clamp(1.8rem, 5vw, 2.8rem);
  font-weight: 900;
  line-height: 1.2;
  margin: 0 0 1rem;
  letter-spacing: -0.02em;
}

.hero-accent {
  background: linear-gradient(90deg, #60a5fa, #a78bfa);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.hero-sub {
  font-size: 1rem;
  opacity: 0.8;
  margin: 0 0 2rem;
  line-height: 1.6;
}

.hero-stats {
  display: inline-flex;
  align-items: center;
  gap: 1.5rem;
  background: rgba(255,255,255,0.12);
  backdrop-filter: blur(8px);
  border: 1px solid rgba(255,255,255,0.2);
  border-radius: 16px;
  padding: 0.85rem 2rem;
}

.stat { text-align: center; }

.stat-num {
  display: block;
  font-size: 1.4rem;
  font-weight: 800;
  line-height: 1;
}

.stat-label {
  display: block;
  font-size: 0.72rem;
  opacity: 0.7;
  text-transform: uppercase;
  letter-spacing: 0.06em;
  margin-top: 0.2rem;
}

.stat-divider {
  width: 1px;
  height: 36px;
  background: rgba(255,255,255,0.25);
}

/* Ana İçerik */
.main-content {
  max-width: 860px;
  margin: 0 auto;
  padding: 2.5rem 1.5rem;
  flex: 1;
  width: 100%;
}

.card-panel {
  background: white;
  border-radius: 20px;
  padding: 2rem;
  box-shadow: 0 8px 40px rgba(30,58,95,0.10);
  border: 1px solid rgba(59,130,246,0.08);
  margin-bottom: 2.5rem;
}

/* Özellikler */
.features {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 1.2rem;
}

.feature-card {
  background: white;
  border-radius: 16px;
  padding: 1.5rem;
  text-align: center;
  border: 1px solid #e8eef6;
  box-shadow: 0 2px 12px rgba(30,58,95,0.05);
  transition: transform 0.2s, box-shadow 0.2s;
}

.feature-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 24px rgba(30,58,95,0.10);
}

.feat-icon { font-size: 2rem; display: block; margin-bottom: 0.75rem; }

.feature-card h4 {
  font-size: 0.95rem;
  font-weight: 700;
  color: #1e3a5f;
  margin: 0 0 0.5rem;
}

.feature-card p {
  font-size: 0.8rem;
  color: #64748b;
  margin: 0;
  line-height: 1.5;
}

/* Footer */
.site-footer {
  text-align: center;
  padding: 1.2rem;
  font-size: 0.78rem;
  color: #94a3b8;
  border-top: 1px solid #e2e8f0;
}

@media (max-width: 640px) {
  .features { grid-template-columns: 1fr; }
  .hero { padding: 2.5rem 1rem 2rem; }
  .card-panel { padding: 1.2rem; }
  .hero-stats { gap: 1rem; padding: 0.7rem 1.2rem; }
}
</style>
