# APK Remover Bot 🤖

Telegram guruhiga yuborilgan `.apk` fayllarni avtomatik o'chiradigan bot.

---

## Talablar

- Go 1.21+
- Telegram Bot Token ([BotFather](https://t.me/BotFather) orqali oling)

---

## O'rnatish va ishga tushirish

### 1. Bog'liqliklarni yuklab olish
```bash
go mod tidy
```

### 2. Token ni muhit o'zgaruvchisi sifatida set qiling

**Linux / macOS:**
```bash
export TELEGRAM_BOT_TOKEN="your_bot_token_here"
```

**Windows:**
```cmd
set TELEGRAM_BOT_TOKEN=your_bot_token_here
```

### 3. Botni ishga tushiring
```bash
go run main.go
```

yoki build qilib:
```bash
go build -o bot .
./bot
```

---

## Docker bilan ishlatish

```bash
docker build -t bot .
docker run -d -e TELEGRAM_BOT_TOKEN="your_token_here" bot
```

---

## Botni guruhga qo'shish va sozlash

1. **BotFather da botni yarating** va tokenni oling.

2. **Muhim sozlama** — BotFather da quyidagi buyruqni yuboring:
   ```
   /setprivacy → @your_bot_name → Disable
   ```
   > Bu sozlama bo'lmasa bot guruh xabarlarini ko'ra olmaydi!

3. **Botni guruhga admin sifatida qo'shing** va quyidagi ruxsatlarni bering:
   - ✅ **Delete messages** (xabarlarni o'chirish)
   - Boshqa ruxsatlar ixtiyoriy

---

## Ishlash mantiqi

```
Yangi xabar keldi
       ↓
Guruhdan kelganmi? → YO'Q → Ignore
       ↓ HA
Document/Fayl bormi? → YO'Q → Ignore
       ↓ HA
.apk kengaytmali? → YO'Q → Ignore
       ↓ HA
Xabarni o'chir ✅
```

---

## Xususiyatlar

- ✅ Faqat `.apk` fayllarni o'chiradi
- ✅ Boshqa fayl turlarini (pdf, zip, mp3 va h.k.) tegmaydi
- ✅ Oddiy matnli xabarlarni tegmaydi
- ✅ Kichik/katta harf farqsiz (`FILE.APK`, `file.apk` — bari o'chiriladi)
- ✅ Bir nechta guruhlarda ishlaydi