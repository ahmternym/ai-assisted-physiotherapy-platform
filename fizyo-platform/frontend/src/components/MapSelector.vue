<template>
  <div class="map-selector">
    <div class="selector-intro">
      <h2>Türkiye'de Fizyoterapist Bul</h2>
      <p>Haritadan bir il seçin veya arama kutusunu kullanın</p>
    </div>

    <!-- Leaflet Harita -->
    <div class="map-wrapper">
      <div id="turkey-map" class="leaflet-map"></div>
      <!-- Harita yüklenemezse fallback mesajı -->
      <div v-if="mapError" class="map-fallback">
        🗺️ Harita yüklenemedi. Aşağıdan il seçebilirsiniz.
      </div>
      <div v-if="hoveredCity" class="map-hover-label">{{ hoveredCity }}</div>
    </div>

    <!-- Arama kutusu -->
    <div class="search-row">
      <div class="search-box">
        <span class="search-icon">🔍</span>
        <input
          v-model="searchQuery"
          type="text"
          placeholder="İl ara... (örn: Trabzon)"
          @focus="showDropdown = true"
          @blur="onBlur"
        />
        <button v-if="searchQuery" class="clear-btn" @mousedown.prevent="clearSearch">✕</button>
      </div>
      <div v-if="showDropdown && filteredCities.length" class="dropdown">
        <div
          v-for="city in filteredCities"
          :key="city"
          class="dropdown-item"
          @mousedown.prevent="selectCity(city)"
        >
          <span class="city-dot"></span>
          {{ city }}
        </div>
      </div>
    </div>

    <!-- Popüler İller -->
    <div class="quick-cities">
      <span class="quick-label">Popüler İller:</span>
      <button
        v-for="city in popularCities"
        :key="city"
        class="quick-btn"
        :class="{ active: selectedCity === city }"
        @click="selectCity(city)"
      >
        {{ city }}
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'MapSelector',

  data() {
    return {
      selectedCity: '',
      searchQuery: '',
      showDropdown: false,
      hoveredCity: '',
      map: null,
      markers: {},      // { cityName: L.CircleMarker } — seçili marker'ı yeşile çevirmek için
      mapError: false,

      popularCities: ['Ankara', 'İstanbul', 'İzmir', 'Bursa', 'Antalya', 'Konya', 'Gaziantep'],

      cities: [
        'Adana','Adıyaman','Afyonkarahisar','Ağrı','Amasya','Ankara','Antalya','Artvin',
        'Aydın','Balıkesir','Bilecik','Bingöl','Bitlis','Bolu','Burdur','Bursa',
        'Çanakkale','Çankırı','Çorum','Denizli','Diyarbakır','Edirne','Elazığ',
        'Erzincan','Erzurum','Eskişehir','Gaziantep','Giresun','Gümüşhane','Hakkari',
        'Hatay','Isparta','Mersin','İstanbul','İzmir','Kars','Kastamonu','Kayseri',
        'Kırklareli','Kırşehir','Kocaeli','Konya','Kütahya','Malatya','Manisa',
        'Kahramanmaraş','Mardin','Muğla','Muş','Nevşehir','Niğde','Ordu','Rize',
        'Sakarya','Samsun','Siirt','Sinop','Sivas','Tekirdağ','Tokat','Trabzon',
        'Tunceli','Şanlıurfa','Uşak','Van','Yozgat','Zonguldak','Aksaray','Bayburt',
        'Karaman','Kırıkkale','Batman','Şırnak','Bartın','Ardahan','Iğdır','Yalova',
        'Karabük','Kilis','Osmaniye','Düzce'
      ],

      cityCoords: {
        'Adana':[37.0,35.32],'Adıyaman':[37.76,38.27],'Afyonkarahisar':[38.76,30.54],
        'Ağrı':[39.72,43.05],'Amasya':[40.65,35.83],'Ankara':[39.93,32.85],
        'Antalya':[36.89,30.71],'Artvin':[41.18,41.82],'Aydın':[37.84,27.85],
        'Balıkesir':[39.65,27.89],'Bilecik':[40.15,29.98],'Bingöl':[38.89,40.50],
        'Bitlis':[38.40,42.11],'Bolu':[40.74,31.61],'Burdur':[37.72,30.29],
        'Bursa':[40.20,29.06],'Çanakkale':[40.15,26.41],'Çankırı':[40.60,33.61],
        'Çorum':[40.55,34.96],'Denizli':[37.77,29.09],'Diyarbakır':[37.91,40.22],
        'Edirne':[41.67,26.56],'Elazığ':[38.68,39.22],'Erzincan':[39.75,39.49],
        'Erzurum':[39.90,41.27],'Eskişehir':[39.78,30.52],'Gaziantep':[37.07,37.38],
        'Giresun':[40.91,38.39],'Gümüşhane':[40.46,39.48],'Hakkari':[37.57,43.74],
        'Hatay':[36.20,36.16],'Isparta':[37.76,30.56],'Mersin':[36.80,34.64],
        'İstanbul':[41.01,28.95],'İzmir':[38.42,27.14],'Kars':[40.61,43.10],
        'Kastamonu':[41.38,33.78],'Kayseri':[38.73,35.49],'Kırklareli':[41.73,27.22],
        'Kırşehir':[39.15,34.17],'Kocaeli':[40.76,29.92],'Konya':[37.87,32.49],
        'Kütahya':[39.42,29.99],'Malatya':[38.36,38.31],'Manisa':[38.62,27.43],
        'Kahramanmaraş':[37.59,36.94],'Mardin':[37.31,40.74],'Muğla':[37.22,28.36],
        'Muş':[38.73,41.49],'Nevşehir':[38.62,34.71],'Niğde':[37.97,34.68],
        'Ordu':[40.98,37.88],'Rize':[41.02,40.52],'Sakarya':[40.69,30.43],
        'Samsun':[41.29,36.33],'Siirt':[37.93,41.95],'Sinop':[42.02,35.15],
        'Sivas':[39.75,37.02],'Tekirdağ':[40.98,27.51],'Tokat':[40.31,36.55],
        'Trabzon':[41.00,39.72],'Tunceli':[39.11,39.55],'Şanlıurfa':[37.16,38.80],
        'Uşak':[38.68,29.41],'Van':[38.49,43.38],'Yozgat':[39.82,34.81],
        'Zonguldak':[41.46,31.80],'Aksaray':[38.37,34.04],'Bayburt':[40.26,40.23],
        'Karaman':[37.18,33.22],'Kırıkkale':[39.85,33.51],'Batman':[37.88,41.13],
        'Şırnak':[37.52,42.46],'Bartın':[41.63,32.34],'Ardahan':[41.11,42.70],
        'Iğdır':[39.92,44.04],'Yalova':[40.65,29.27],'Karabük':[41.21,32.62],
        'Kilis':[36.72,37.12],'Osmaniye':[37.07,36.25],'Düzce':[40.84,31.16]
      }
    }
  },

  computed: {
    filteredCities() {
      if (!this.searchQuery) return this.cities.slice(0, 8)
      return this.cities
        .filter(c => c.toLowerCase().includes(this.searchQuery.toLowerCase()))
        .slice(0, 8)
    }
  },

  mounted() {
    // Vue bileşeni DOM'a eklendikten sonra haritayı başlat.
    // window.L henüz yüklenmediyse (CDN geç gelebilir) 300ms bekleyip tekrar dene.
    this.$nextTick(() => {
      this.tryInitMap(0)
    })
  },

  beforeUnmount() {
    if (this.map) {
      this.map.remove()
      this.map = null
    }
  },

  methods: {
    // Leaflet CDN geç yüklenirse diye retry mekanizması (max 10 deneme = 3sn)
    tryInitMap(attempt) {
      if (window.L) {
        this.initMap()
        return
      }
      if (attempt >= 10) {
        console.warn('Leaflet yüklenemedi, harita gösterilmiyor.')
        this.mapError = true
        return
      }
      setTimeout(() => this.tryInitMap(attempt + 1), 300)
    },

    initMap() {
      const L = window.L
      const container = document.getElementById('turkey-map')
      if (!container) return

      // Daha önce başlatılmışsa temizle (HMR senaryosu)
      if (this.map) {
        this.map.remove()
        this.map = null
      }

      this.map = L.map('turkey-map', {
        center: [39.0, 35.5],
        zoom: 5,
        zoomControl: true,
        scrollWheelZoom: false,
        // Türkiye dışına çıkılmasın
        maxBounds: [[34.0, 24.0], [43.0, 45.0]],
        maxBoundsViscosity: 0.8
      })

      L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '© <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a>',
        maxZoom: 12
      }).addTo(this.map)

      // 81 il için marker ekle
      Object.entries(this.cityCoords).forEach(([cityName, coords]) => {
        const marker = L.circleMarker(coords, {
          radius: 6,
          fillColor: '#3b82f6',
          color: '#1d4ed8',
          weight: 1.5,
          opacity: 1,
          fillOpacity: 0.85
        }).addTo(this.map)

        marker.bindTooltip(cityName, {
          permanent: false,
          direction: 'top',
          className: 'city-tooltip'
        })

        marker.on('click', () => this.selectCity(cityName))

        marker.on('mouseover', () => {
          this.hoveredCity = cityName
          marker.setStyle({ fillColor: '#f59e0b', radius: 9, fillOpacity: 1 })
        })

        marker.on('mouseout', () => {
          this.hoveredCity = ''
          const isSelected = this.selectedCity === cityName
          marker.setStyle({
            fillColor: isSelected ? '#10b981' : '#3b82f6',
            radius: isSelected ? 9 : 6,
            fillOpacity: 0.85
          })
        })

        // Marker referansını sakla — seçim değişince güncelle
        this.markers[cityName] = marker
      })
    },

    selectCity(city) {
      // Önceki seçili marker'ı sıfırla
      if (this.selectedCity && this.markers[this.selectedCity]) {
        this.markers[this.selectedCity].setStyle({
          fillColor: '#3b82f6',
          radius: 6,
          fillOpacity: 0.85
        })
      }

      this.selectedCity = city
      this.searchQuery = city
      this.showDropdown = false

      // Yeni seçili marker'ı yeşile çevir
      if (this.markers[city]) {
        this.markers[city].setStyle({
          fillColor: '#10b981',
          radius: 9,
          fillOpacity: 1
        })
        // Haritayı seçili ile merkez
        if (this.map && this.cityCoords[city]) {
          this.map.setView(this.cityCoords[city], 7, { animate: true })
        }
      }

      this.$emit('city-selected', city)
    },

    clearSearch() {
      this.searchQuery = ''
      this.selectedCity = ''
      this.showDropdown = false
    },

    onBlur() {
      setTimeout(() => { this.showDropdown = false }, 150)
    }
  }
}
</script>

<style scoped>
.map-selector { padding: 0; }

.selector-intro {
  text-align: center;
  margin-bottom: 1.5rem;
}
.selector-intro h2 {
  font-size: 1.6rem;
  font-weight: 800;
  color: #1e3a5f;
  margin: 0 0 0.4rem;
}
.selector-intro p {
  color: #64748b;
  font-size: 0.95rem;
}

/* Harita */
.map-wrapper {
  position: relative;
  border-radius: 16px;
  overflow: hidden;
  box-shadow: 0 4px 24px rgba(0,0,0,0.10);
  margin-bottom: 1.5rem;
  border: 1px solid #e2e8f0;
  background: #e8f0fe;  /* tile yüklenene kadar görünür zemin */
  min-height: 320px;
}

.leaflet-map {
  width: 100%;
  height: 320px;
}

.map-fallback {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f0f4ff;
  color: #64748b;
  font-size: 0.9rem;
  font-weight: 600;
  text-align: center;
  padding: 1rem;
}

.map-hover-label {
  position: absolute;
  bottom: 12px;
  left: 12px;
  background: rgba(30,58,95,0.92);
  color: white;
  padding: 0.35rem 0.85rem;
  border-radius: 99px;
  font-size: 0.85rem;
  font-weight: 600;
  pointer-events: none;
  z-index: 1000;
  backdrop-filter: blur(4px);
}

/* Arama */
.search-row {
  position: relative;
  margin-bottom: 1.2rem;
}

.search-box {
  display: flex;
  align-items: center;
  background: #fff;
  border: 2px solid #e2e8f0;
  border-radius: 12px;
  padding: 0 1rem;
  transition: border-color 0.2s;
  box-shadow: 0 1px 6px rgba(0,0,0,0.05);
}
.search-box:focus-within {
  border-color: #3b82f6;
  box-shadow: 0 0 0 4px rgba(59,130,246,0.12);
}
.search-icon { font-size: 1rem; margin-right: 0.5rem; }
.search-box input {
  flex: 1;
  border: none;
  outline: none;
  padding: 0.75rem 0;
  font-size: 0.95rem;
  color: #1e293b;
  background: transparent;
}
.clear-btn {
  background: none;
  border: none;
  color: #94a3b8;
  cursor: pointer;
  font-size: 0.85rem;
  padding: 0.2rem 0.4rem;
  border-radius: 4px;
  transition: color 0.2s;
}
.clear-btn:hover { color: #ef4444; }

.dropdown {
  position: absolute;
  top: calc(100% + 4px);
  left: 0; right: 0;
  background: #fff;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  box-shadow: 0 8px 32px rgba(0,0,0,0.12);
  z-index: 999;
  overflow: hidden;
}
.dropdown-item {
  display: flex;
  align-items: center;
  gap: 0.6rem;
  padding: 0.7rem 1rem;
  cursor: pointer;
  font-size: 0.9rem;
  color: #334155;
  transition: background 0.15s;
}
.dropdown-item:hover { background: #f0f7ff; }
.city-dot {
  width: 8px; height: 8px;
  border-radius: 50%;
  background: #3b82f6;
  flex-shrink: 0;
}

/* Popüler iller */
.quick-cities {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 0.5rem;
}
.quick-label {
  font-size: 0.8rem;
  color: #94a3b8;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.04em;
  white-space: nowrap;
}
.quick-btn {
  background: #f1f5f9;
  border: 1.5px solid #e2e8f0;
  color: #475569;
  border-radius: 99px;
  padding: 0.35rem 0.9rem;
  font-size: 0.82rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}
.quick-btn:hover {
  background: #eff6ff;
  border-color: #93c5fd;
  color: #1d4ed8;
}
.quick-btn.active {
  background: #1d4ed8;
  border-color: #1d4ed8;
  color: white;
}
</style>
