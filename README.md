# 🧠 Go Subnetting Tool

A simple yet powerful Go-based subnet planner that helps visualize, split, and assign IP subnets (e.g. `192.168.0.0/16` into multiple `/24` networks) for different areas like offices, cafes, and labs. Great for learning and real-world IP planning!

---

## 🚀 Features

- 📌 Define a base CIDR block and split it into smaller subnets (e.g., `/16` to `/24`)
- 🏷️ Assign subnets to physical/logical areas like Office, Lab, Studio
- 🔍 Calculate total subnets, usable hosts, broadcast addresses
- 📊 Display output in a readable table format
- 💡 Extendable: Can be integrated with pinging, port scanning, or visualization tools

---

## 📦 Installation

Make sure you have Go installed: [Download Go](https://golang.org/dl/)

```bash
git clone https://github.com/your-username/go-subnet-planner.git
cd go-subnet-planner
go run main.go
