# 📊 Go Prometheus Example

This project is a simple demonstration of **Observability** in Go using [Prometheus](https://prometheus.io/) and [prometheus/client_golang](https://github.com/prometheus/client_golang).  
It starts an HTTP server, registers metrics, and exposes them on `/metrics` for Prometheus scraping.

---

## 🚀 Features

- Count HTTP requests using a **Counter**
- Simple message response on `/`
- Expose metrics on `/metrics`
- Built-in Go runtime metrics (goroutines, memory usage, GC stats)

---

## 📦 Installation & Run

### 1. Clone the repository

```bash
git clone https://github.com/your-username/go-prometheus-example.git
cd go-prometheus-example
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Run the server

```bash
go run main.go
```

---

## 🔗 Usage

👉 Root endpoint:

```bash
curl http://localhost:8080/
```

Output:

```
Hello, Abl Dev 🚀
```

👉 Metrics endpoint:

```bash
curl http://localhost:8080/metrics
```

Sample output:

```
# HELP http_requests_total Total number of HTTP requests
# TYPE http_requests_total counter
http_requests_total 5
```

---

## 📊 Integration with Prometheus

Minimal Prometheus configuration to scrape metrics:

```yaml
scrape_configs:
  - job_name: "go-app"
    static_configs:
      - targets: ["localhost:8080"]
```

Once Prometheus is running, you can explore the metrics via its UI or build dashboards in [Grafana](https://grafana.com/).

---

## 🛠️ Next Steps

- Add **CounterVec** with labels (e.g., path, status code)
- Use **Histogram** for request latency monitoring
- Integrate **Tracing** with OpenTelemetry for full observability
