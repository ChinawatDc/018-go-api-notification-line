# 📲 018-go-api-notification-line
## Go API – ส่งแจ้งเตือนด้วย LINE OA (Messaging API)

โปรเจกต์นี้เป็นส่วนหนึ่งของคอร์ส **Go API Backend Series**  
ใช้สำหรับสร้างระบบ **แจ้งเตือน / แชทบอท / SOS / Admin Notification** ผ่าน **LINE Official Account (Messaging API)**

---

## ✅ คุณสมบัติหลัก
- รับ Webhook จาก LINE OA
- ตรวจสอบ X-Line-Signature (ความปลอดภัย)
- Reply Message (ตอบผู้ใช้)
- Push Message (แจ้งเตือนผู้ดูแล / Admin)
- รองรับการทำ Rich Menu / Postback ต่อยอด
- โครงสร้าง Production-ready

---

## 📁 โครงสร้างโปรเจกต์

```
018-go-api-notification-line
├─ go.mod
├─ .env.example
├─ cmd/api/main.go
├─ internal
│  ├─ config
│  │  └─ config.go
│  ├─ lineoa
│  │  ├─ client.go
│  │  ├─ signature.go
│  │  ├─ models.go
│  │  └─ service.go
│  └─ httpserver
│     ├─ router.go
│     └─ handlers.go
└─ README.md
```

---

## 🧰 Tech Stack
- Go 1.22+
- Gin Web Framework
- LINE Messaging API
- HMAC SHA256 (Signature Verify)
- REST API

---

## 🚀 การติดตั้งและใช้งาน

### 1️⃣ ตั้งค่า Go Module
```bash
go mod edit -module=github.com/ChinawatDc/018-go-api-notification-line
go mod tidy
```

---

### 2️⃣ ตั้งค่า Environment
```bash
Copy-Item .env.example .env
```

ตัวอย่าง `.env`

```env
APP_PORT=8080

LINE_CHANNEL_SECRET=xxxxxxxxxxxxxxxx
LINE_CHANNEL_ACCESS_TOKEN=xxxxxxxxxxxxxxxx

# optional (userId / groupId)
LINE_ADMIN_TO=Uxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

---

### 3️⃣ รันเซิร์ฟเวอร์
```bash
go run cmd/api/main.go
```

ผลลัพธ์:
```
server running :8080
```

---

## 🔗 LINE Developers Setup

1. ไปที่ https://developers.line.biz/
2. สร้าง Provider และ Messaging API Channel
3. คัดลอก
   - Channel Secret
   - Channel Access Token
4. ตั้งค่า Webhook URL:
```
https://<your-domain>/api/line/webhook
```
5. เปิด Use Webhook = ON

> Local แนะนำใช้ ngrok  
> `ngrok http 8080`

---

## 📮 API Endpoints

### 🔹 Health Check
```
GET /api/health
```

---

### 🔹 LINE Webhook
```
POST /api/line/webhook
```

LINE จะยิงเข้ามาอัตโนมัติ

---

### 🔹 แจ้งเตือนผู้ดูแล (Admin Push)
```
POST /api/notify/admin
```

```json
{
  "message": "🚨 แจ้งเตือนจากระบบ"
}
```

---

### 🔹 Push ไป user / group
```
POST /api/notify/push
```

```json
{
  "to": "Uxxxxxxxxxxxx",
  "message": "Hello from Go API"
}
```

---

## 🧪 ตัวอย่างการใช้งานใน LINE

| คำสั่ง | ผลลัพธ์ |
|------|--------|
| ping | pong |
| sos มีเหตุฉุกเฉิน | แจ้ง admin |
| ข้อความอื่น | ตอบกลับอัตโนมัติ |

---

## 🎓 เหมาะสำหรับ
- ระบบแจ้งเตือน (Notification)
- ระบบ SOS / ร้องเรียน
- Chatbot องค์กร / หมู่บ้าน
- เชื่อม n8n / Backend อื่น

---

## 📌 Production Tips
- Reply ให้เร็ว (ภายใน timeout LINE)
- งานหนัก → Push แทน
- เก็บ log event ลง DB ได้
- ต่อ Rich Menu + Postback ได้ทันที

---

## 👨‍💻 Author
**Chinawat Daochai**  
Go Backend / Fullstack Developer  
Go API Course Series
# 018-go-api-notification-line
