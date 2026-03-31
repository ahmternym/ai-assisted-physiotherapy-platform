# 🧠 Türkiye Fizyoterapist Bulma Platformu

##  Proje Tanımı

Bu proje, kullanıcıların Türkiye'nin 81 ilinde fizyoterapistleri kolayca bulmasını sağlayan modern, şık ve sade bir web uygulamasıdır.

Kullanıcılar, ana sayfada yer alan Türkiye haritası üzerinden bir il seçerek o bölgedeki fizyoterapistleri görüntüleyebilir ve detaylı bilgilere ulaşabilir.

---

##  Amaç

* Kullanıcıların bulunduğu şehirde fizyoterapist bulmasını kolaylaştırmak
* Basit ve kullanıcı dostu bir arayüz sunmak
* Modern frontend + backend mimarisi ile ölçeklenebilir bir yapı oluşturmak

---

##  Kullanıcı Akışı

### 1. Karşılama Ekranı

* Modern ve sade bir UI
* Türkiye haritası
* İl seçimi yapılabilecek alan

### 2. İl Seçimi

* Kullanıcı haritadan veya listeden bir il seçer
* Seçilen ile göre veri çekilir

### 3. Fizyoterapist Listeleme

Seçilen ile göre kullanıcıya **5 adet fizyoterapist kartı** gösterilir.

#### Kart İçeriği:

* Ad Soyad
* İl / İlçe
* Telefon Numarası

### 4. Fizyoterapist Detay Sayfası

Kullanıcı karta tıkladığında detay sayfasına yönlendirilir.

#### Detay İçeriği:

* Ad Soyad
* İl / İlçe
* Telefon Numarası
* Uzmanlık Alanları (örnek: spor rehabilitasyonu, ortopedik rehabilitasyon vb.)

---

##  Tasarım Prensipleri

* Modern ve minimal UI
* Responsive tasarım (mobil uyumlu)
* Kullanıcı dostu navigasyon
* Hızlı ve sade deneyim

---

## ⚙️ Teknolojiler

### Frontend

* Vue.js (Vue CLI)
* Options API
* CSS / SCSS (isteğe bağlı Tailwind vb.)

### Backend

* Golang (Go)
* REST API mimarisi

---

## 🔌 API Yapısı

### 1. İl Bazlı Fizyoterapist Listeleme

**GET /api/physiotherapists?city={city}**

#### Response (örnek):

```json
[
  {
    "id": 1,
    "name": "Ahmet Yılmaz",
    "city": "Ankara",
    "district": "Çankaya",
    "phone": "0555 555 55 55"
  }
]
```

---

### 2. Fizyoterapist Detay

**GET /api/physiotherapists/{id}**

#### Response (örnek):

```json
{
  "id": 1,
  "name": "Ahmet Yılmaz",
  "city": "Ankara",
  "district": "Çankaya",
  "phone": "0555 555 55 55",
  "specialties": [
    "Ortopedik Rehabilitasyon",
    "Spor Yaralanmaları"
  ]
}
```

---

## 🗃️ Veri Yapısı (Fake Data)

Backend tarafında tüm veriler başlangıçta **fake veri** olacaktır.

### Örnek Model (Go struct):

```go
type Physiotherapist struct {
    ID          int      `json:"id"`
    Name        string   `json:"name"`
    City        string   `json:"city"`
    District    string   `json:"district"`
    Phone       string   `json:"phone"`
    Specialties []string `json:"specialties"`
}
```

---

##  Proje Yapısı

### Frontend (Vue)

```
src/
 ├── components/
 │    ├── MapSelector.vue
 │    ├── PhysiotherapistCard.vue
 │
 ├── views/
 │    ├── Home.vue
 │    ├── List.vue
 │    ├── Detail.vue
 │
 ├── services/
 │    └── api.js
```

---

### Backend (Go)

```
/backend
 ├── main.go
 ├── handlers/
 ├── models/
 ├── routes/
```

---

##  Gelecek Geliştirmeler

* Filtreleme (uzmanlık alanına göre)
* Randevu sistemi
* Kullanıcı yorumları

---

## ✅ Sonuç

Bu proje, basit ama güçlü bir yapı ile Türkiye genelinde fizyoterapist bulmayı kolaylaştıran modern bir platform sunmayı hedeflemektedir.
