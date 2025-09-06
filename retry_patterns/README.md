# Retry Pattern

خیلی وقت‌ها یه درخواست به خاطر خطای موقت (مثل timeout یا قطعی شبکه) fail میشه. بجای اینکه همون لحظه شکست بخوره، می‌تونیم با Retry Pattern دوباره امتحان کنیم. البته باهوش! نه بی‌نهایت و نه پشت سر هم.

## روش های معروف Retry Pattern

### ۱. روش Fixed Delay

هر بار یک فاصله زمانی ثابت تکرار میکنه

**مثال در گولنگ:**

```go
// run every 2 seconds
base := 2
sleep := time.Duration(base*time.Second)
time.Sleep(sleep)
```

### ۲. روش Expontial Backoff

هر بار فاصله بیشتر می شود
نتایج مشابه 1s -> 2s -> 4s -> 8s در هر تکرار

**مثال در گولنگ:**

```go
// i تعداد تکرار کد است
sleep := time.Duration(1<<i) * time.Second // 1s, 2s, 4s...
time.Sleep(sleep)
```

### ۳. روش ترکیب Expontial + Jitter

همون روش Backoff ولی با کمی random تا ترافیک زیاد نشه (تا باعث نشه همه کلاینت ها همزمان درخواست ارسال کنند)

**مثال در گولنگ**

```go
// i تعداد تکرار کد است
base := 1<<i
jitter := rand.Intn(1000)
sleep := time.Duration(base * 1000+jitter) * time.Millisecond
time.Sleep(sleep)
```

> **نکته:**
>
> ۱. اجرای retry فقط برای خطا های موقت خوبه مثل timeout و network error
> ۲. برای خطاهای منطقی و از سمت کاربر مثل 404 یا unauthorized فایده نداره
> ۳. بهتره که توی API Client Library این متد رو پساده سازی کنی تا همه جا قابل استفاده بشه
