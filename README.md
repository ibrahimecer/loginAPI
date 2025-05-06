# Giriş Sistemi Projesi

Bu proje, öğretmen ve öğrenciler için ayrı giriş ve kayıt sistemi içeren bir kullanıcı yönetim sistemidir.  
Öğrenciler kayıt olabilir, notlarını görebilir; öğretmenler ise tüm öğrencilerin notlarını görüntüleyip güncelleyebilir.

---

## 🎯 Özellikler

- ✅ Öğrenciler sisteme **kayıt olabilir** ve giriş yapabilir
- ✅ Öğrenciler sadece **kendi notlarını** görebilir
- ✅ Öğretmenler sisteme sadece giriş yapabilir
- ✅ Öğretmenler **tüm öğrencilerin notlarını görebilir**
- ✅ Notlar **düzenlenebilir** ve **ortalama hesaplanır**
- ✅ Öğrenci detay sayfasında ad, soyad, e-posta gibi bilgiler yer alır
- ✅ Frontend tamamen HTML + JS ile hazırlandı (Vanilla JS)

---

## 🧠 Kullanılan Teknolojiler

| Teknoloji | Açıklama |
|-----------|----------|
| **Go (Golang)** | Backend API için |
| **Gin Framework** | HTTP router |
| **GORM** | ORM (veritabanı işlemleri için) |
| **PostgreSQL** | Veritabanı yönetimi |
| **HTML & JavaScript** | Frontend arayüz |
| **bcrypt** | Şifreleri güvenli saklama |
| **Docker (isteğe bağlı)** | Servisleri container olarak çalıştırmak için

---

## 🗂️ Proje Yapısı

├── config/ # Veritabanı bağlantısı
├── controller/ # Auth & CRUD işlemleri
├── models/ # GORM modelleri (User, Student, etc.)
├── routes/ # API yönlendirme
├── static/ # HTML dosyaları (giriş, kayıt, panel)
├── main.go # Ana uygulama
---

## 🚀 Nasıl Başlatılır?

### 1. PostgreSQL kurulu olmalı
- Veritabanı adı: `login_api`
- Kullanıcı: `postgres`
- Şifre: `postgres`

### 2. Projeyi çalıştır(Go kurulu olmalı)
```bash
go run main.go

http://localhost:4569/ adresi üzerinden erişim sağlayabilirsiniz.
