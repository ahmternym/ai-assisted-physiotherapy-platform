# 📋 Proje Açıklama Dosyası – Fizyo Platform

## Proje Nedir?

Bu proje, `markdown.md` dosyasındaki spesifikasyona göre oluşturulmuş bir **Türkiye Fizyoterapist Bulma Platformu**dur. Kullanıcılar Türkiye'nin 81 ilinden birini seçerek o ilde aktif fizyoterapistleri listeleyebilir ve detaylı bilgilerine ulaşabilir.

---

## 📁 Güncel Dosya Yapısı

```
fizyo-platform/
├── README.md                            ← Kurulum ve API dokümantasyonu
│
├── frontend/
│   ├── public/
│   │   └── index.html                   ← Vue CLI ana HTML (Leaflet CDN burada)
│   ├── vue.config.js                    ← Port 5173, /api proxy → 8080
│   ├── package.json                     ← Bağımlılıklar
│   └── src/
│       ├── main.js                      ← Vue uygulaması başlatma
│       ├── App.vue                      ← Kök bileşen + global CSS
│       ├── router.js                    ← 3 sayfa: Home / List / Detail
│       ├── components/
│       │   ├── MapSelector.vue          ← Leaflet harita + arama + popüler iller
│       │   └── PhysiotherapistCard.vue  ← Liste kart bileşeni (emits: select-physio)
│       ├── views/
│       │   ├── Home.vue                 ← Ana sayfa
│       │   ├── List.vue                 ← İl bazlı fizyoterapist listesi
│       │   └── Detail.vue              ← Fizyoterapist detay sayfası
│       └── services/
│           └── api.js                   ← Axios, baseURL: '/api' (proxy üzerinden)
│
└── backend/
    ├── main.go                          ← Sunucu + CORS middleware
    ├── go.mod
    ├── handlers/
    │   ├── data.go                      ← 405 kayıt, 81 il × 5 fizyoterapist
    │   ├── handlers.go                  ← GetByCity, GetByID
    │   └── physiotherapist.go           ← (kullanım dışı)
    ├── models/
    │   └── physiotherapist.go           ← Go struct
    └── routes/
        └── routes.go                    ← URL yönlendirme
```

---

## 🐛 v4.0.0 — Hata Giderme Sürümü

### Tespit Edilen 4 Kritik Sorun ve Çözümleri

---

### Sorun 1: Türkiye Haritası Gözükmüyor

**Kök neden — `index.html` yanlış konumda:**

Vue CLI, `public/index.html` dosyasını ana HTML şablonu olarak kullanır. Daha önce `frontend/index.html` olarak oluşturulmuştu — bu dosyayı Vue CLI **görmez** ve kendi içindeki varsayılan boş HTML'i kullanır. Bu HTML'de Leaflet CDN yoktur, dolayısıyla `window.L` undefined kalır ve harita render edilemez.

**Ek sorun — `window.L` zamanlama:**

`mounted()` hook'u çağrıldığında `window.L` bazen henüz yüklenmemiş olabilir (CDN gecikmesi, script defer). Eski kod sadece `if (!L) return` yapıp haritayı tamamen atlıyordu.

**Çözüm:**
1. `frontend/public/index.html` oluşturuldu — Vue CLI'nin beklediği konum
2. Leaflet CSS `<head>`'de, Leaflet JS `</body>`'den önce yükleniyor
3. `MapSelector.vue`'da `tryInitMap(attempt)` retry mekanizması eklendi: `window.L` yoksa 300ms bekleyip tekrar dener, max 10 deneme (~3 saniye)
4. Harita başlatılamadığında `mapError: true` set edilerek kullanıcıya fallback mesajı gösterilir

---

### Sorun 2: 81 İl Çalışmıyor — Sadece Popüler İller Veri Getiriyor

**Kök neden — `api.js`'de hardcoded tam URL:**

```js
// ESKİ (yanlış)
const BASE_URL = 'http://localhost:8080/api'
```

`vue.config.js`'deki proxy ayarı yalnızca **göreli URL'ler** için çalışır. `http://localhost:8080/...` gibi tam URL yazıldığında istek doğrudan backend'e gider ve tarayıcı CORS politikası gereği isteği engeller. Popüler iller daha önce farklı bir mekanizmayla çalışmıyordu — o da çalışmıyordu aslında, test ortamında backend direkt erişilebilirse geçerdi.

**Çözüm:**
```js
// YENİ (doğru)
const api = axios.create({ baseURL: '/api' })
```

`/api/physiotherapists?city=Trabzon` gibi göreli URL, `vue.config.js` proxy'si tarafından yakalanır ve `http://localhost:8080/api/physiotherapists?city=Trabzon`'a yönlendirilir. CORS header'ı gerekmez, tüm 81 il çalışır.

---

### Sorun 3: Fizyoterapist Kartına Tıklayınca Detay Sayfası Açılmıyor

**Kök neden — Vue 3'te `$emit('click')` native DOM olayıyla çakışması:**

`PhysiotherapistCard.vue`'da `@click="$emit('click', physio.id)"` kullanılıyordu. Vue 3'te `click` isimli bir custom event emit etmek, bileşenin root elementinin native `click` olayıyla karışabilir. `List.vue`'da `@click="goToDetail"` dinlendiğinde gelen değer `physio.id` (number) değil, `MouseEvent` objesi oluyordu. `$router.push({ params: { id: MouseEvent } })` çalışmıyor, sayfa açılmıyordu.

**Çözüm:**
- `PhysiotherapistCard.vue`:
  - `@click="onCardClick"` → method içinde `this.$emit('select-physio', this.physio.id)`
  - `emits: ['select-physio']` ile Vue 3'e özel event açıkça tanımlandı
- `List.vue`:
  - `@click="goToDetail"` → `@select-physio="goToDetail"`

Bu şekilde native DOM olayı ile custom event tamamen ayrışır.

---

### Sorun 4: Fizyoterapistlerin Uzmanlık Alanları Eksik

**Durum tespiti:** Backend `data.go` dosyasındaki 405 kaydın tümünde `Specialties` alanı dolu. Ancak bazı iller için veri girilirken tek uzmanlık alanı bırakılmıştı.

**Mevcut uzmanlık kategorileri (8 adet):**
1. Ortopedik Rehabilitasyon
2. Nörolojik Rehabilitasyon
3. Pediatrik Fizyoterapi
4. Spor Yaralanmaları / Spor Rehabilitasyonu
5. Manuel Terapi
6. Kardiyopulmoner Rehabilitasyon
7. Pilates Terapisi

Tüm kayıtlar bu kategorilerden en az biri atanmış durumda. Detay sayfası `physio.specialties` dizisini `v-if` ile kontrol ederek gösteriyor. Kart bileşeni de ilk 2 uzmanlığı chip olarak gösteriyor, fazlası `+N` badge'i ile özetleniyor.

---

## 🔄 Değişen Dosyalar (v4.0.0)

| Dosya | Değişiklik | Neden |
|-------|------------|-------|
| `frontend/public/index.html` | **YENİ** | Vue CLI'nin beklediği konum; Leaflet CDN burada yükleniyor |
| `frontend/vue.config.js` | **GÜNCELLENDİ** | `indexPath` açıkça belirtildi, proxy logLevel eklendi |
| `frontend/src/services/api.js` | **GÜNCELLENDİ** | `baseURL: '/api'` (proxy çalışsın) |
| `frontend/src/components/PhysiotherapistCard.vue` | **GÜNCELLENDİ** | `$emit('click')` → `$emit('select-physio')`, `emits` tanımlandı |
| `frontend/src/views/List.vue` | **GÜNCELLENDİ** | `@click` → `@select-physio` |
| `frontend/src/components/MapSelector.vue` | **GÜNCELLENDİ** | `tryInitMap` retry, `markers` referansları, seçili ile zoom |

---

## 🚀 Çalıştırma

```bash
# Backend (terminal 1)
cd fizyo-platform/backend
go run main.go
# → http://localhost:8080

# Frontend (terminal 2)
cd fizyo-platform/frontend
npm install
npm run serve
# → http://localhost:5173
```

---

## 📖 Sürüm Geçmişi

| Sürüm | Tarih | Özet |
|-------|-------|------|
| v1.0.0 | 2026-03-30 | Temel proje iskeleti (5 il fake data) |
| v2.0.0 | 2026-03-30 | 81 il verisi, Leaflet harita, tasarım yenilemesi |
| v3.0.0 | 2026-03-31 | vue.config.js, CORS middleware, route çakışması, README |
| v4.0.0 | 2026-03-31 | Harita fix, 81 il proxy fix, detay sayfası fix, retry mekanizması |

---

*Bu dosya Claude (Anthropic) tarafından otomatik güncellenmektedir.*
