# 🏥 FizyoBul — AI-Assisted Physiotherapist Discovery Platform

🇹🇷 Türkiye'nin 81 ilinde fizyoterapist aramanızı sağlayan modern, tam-yığın web uygulaması.  
🇬🇧 A modern full-stack web application that enables users to discover physiotherapists across all 81 cities in Turkey.

---

## 📸 Screenshots

### 🗺️ Home & Map Selection
![Home](./fizyo-platform/screenshots/images1.PNG)

### 📋 Physiotherapist List
![List](./fizyo-platform/screenshots/images2.PNG)

### 📄 Detail Page
![Detail](./fizyo-platform/screenshots/images3.PNG)

---

## 🚀 Features / Özellikler

### 🇬🇧 Features
- Interactive map of all 81 cities using **Leaflet**
- City selection via map click or search input
- Displays **5 physiotherapists per city**
- Detailed profile page (specialties, phone, location)
- Smooth UX with **skeleton loading**
- Fully **responsive design**

### 🇹🇷 Özellikler
- **Leaflet harita** ile 81 ilin interaktif gösterimi  
- Harita veya arama ile il seçimi  
- Her il için **5 fizyoterapist listesi** - Detay sayfası (uzmanlık alanları, telefon, konum)  
- **Skeleton loading** ile akıcı deneyim  
- **Mobil uyumlu responsive tasarım**

---

## 🛠️ Tech Stack / Teknolojiler

| Layer | Technology |
|------|-----------|
| Frontend | Vue.js 3 (Options API), Vue Router |
| HTTP Client | Axios |
| Map | Leaflet.js + OpenStreetMap |
| Backend | Go (net/http) |
| Data | In-memory fake dataset (405 records) |

---

## 📁 Project Structure / Proje Yapısı

```text
fizyo-platform/
├── frontend/
├── backend/
├── README.md
└── screenshots/
```

---

## ⚙️ Setup & Run / Kurulum ve Çalıştırma

### Requirements / Gereksinimler
- Go 1.21+
- Node.js 16+
- npm

### 🔧 Backend (Go)
```bash
cd fizyo-platform/backend
go mod tidy
go run main.go
```
➡️ Runs on: `http://localhost:8080`

### 💻 Frontend (Vue.js)
```bash
cd fizyo-platform/frontend
npm install
npm run serve
```
➡️ Runs on: `http://localhost:5173`

---

## 🔌 API Reference

### Get Physiotherapists by City
```http
GET /api/physiotherapists?city={city}
```

### Get Physiotherapist Detail
```http
GET /api/physiotherapists/{id}
```

---

## 🧠 AI-Assisted Development Workflow

🇬🇧 This project was built using a structured AI-assisted development workflow powered by Claude.  
AI was not only used for code generation, but as part of a controlled and traceable development system.

The development process includes:
- Starting from a detailed Markdown-based project specification  
- Using Claude to generate initial code structures  
- Maintaining full control over architecture and decisions  
- Documenting reasoning (why/how) in a separate structure  
- Tracking all updates through versioned changelog files (v2.0.0, v3.0.0, v4.0.0)  
- Creating a final summary document to capture the entire evolution of the project  

AI was used as a collaborative development tool, not as a replacement for engineering decisions.

🇹🇷 Bu proje, Claude kullanılarak yapılandırılmış bir AI destekli geliştirme süreci ile oluşturulmuştur.  
AI yalnızca kod üretimi için değil, kontrollü ve izlenebilir bir geliştirme sistemi kurmak için kullanılmıştır.

Geliştirme süreci:
- Markdown tabanlı proje tanımı ile başlatıldı  
- Claude ile başlangıç kodları üretildi  
- Tüm mimari ve teknik kararlar geliştirici tarafından kontrol edildi  
- Neden ve nasıl yapıldığını açıklayan reasoning dokümantasyonu oluşturuldu  
- Tüm güncellemeler versiyonlanmış changelog dosyalarında tutuldu (v2.0.0, v3.0.0, v4.0.0)  
- Projenin tamamını özetleyen final dokümanı oluşturuldu  

AI bu projede bir araçtan öte, yapılandırılmış geliştirme sürecinin bir parçası olarak kullanılmıştır.

---

## 📂 AI Workflow Structure

```text
AI/
├── changelog/
│   ├── v2.0.0.md
│   ├── v3.0.0.md
│   └── v4.0.0.md
├── proje-ozeti/
│   └── OZET.md
├── markdown.md
└── proje-aciklama.md
```

---

## 🔮 Future Improvements / Gelecek Geliştirmeler

- Filtering by specialties  
- Appointment system  
- User reviews & ratings  
- Database integration  

---

## 👨‍💻 Author

**Ahmet Eren Yaman** Frontend Developer & Physiotherapist  

GitHub: [https://github.com/ahmternym](https://github.com/ahmternym)