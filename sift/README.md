# 🧮 Go MathOps — Test Project for [sift](https://github.com/timtatt/sift)

[🇮🇷 فارسی بخوانید](./README.fa.md)

---

A simple Go project with basic math operations and unit tests — built to test the **sift TUI test viewer**.

---

## 🎯 Purpose

This project was created to test the [sift](https://github.com/timtatt/sift) tool.  
`sift` is a terminal user interface (TUI) that visualizes Go test output in a clean, interactive way.

Official sift repository: [github.com/timtatt/sift](https://github.com/timtatt/sift)

---

## ⚙️ Installation

```bash
git clone https://github.com/abolfazlalz/abldev.git
cd abldev/sift
go mod tidy
```

## ▶️ Run Tests

```bash
go test -v -json ./... | sift
```


