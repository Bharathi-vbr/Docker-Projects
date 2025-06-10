
# 🧮 Multi-Stage Go Docker App - Calculator

This project is a simple terminal-based calculator written in Go, built using a **multi-stage Dockerfile**. It demonstrates how to efficiently build and package Go applications into lightweight Docker images using Alpine.

---

## 📂 Project Structure

├── calculator.go # Main Go application (calculator)
├── go.mod # Go module definition
├── Dockerfile # Multi-stage Dockerfile using 'scratch'
└── README.md # Project documentation
---

## 🚀 Features

- Basic calculator that supports: `+`, `-`, `*`, and `/`
- Accepts expressions like: `4 + 5` or `9 / 3`
- Interactive input in terminal
- Exits when you type `exit`
- Built with a multi-stage Docker build for small image size

---

## 🐳 Docker Instructions

### 🔧 Build the Docker Image

docker build -t go-calculator-app .


### ▶️ Run the Calculator in a Container

docker run -it --rm go-calculator-app

-it: Enables interactive terminal

--rm: Removes container after it exits

### 🧪 Example Usage

Hi Bharathi Bhumireddy, I am a calculator app ....

Enter any calculation (Example: 1 + 2 or 2 * 5 -> Please maintain spaces as shown in example): 7 * 6
Result: 42

Enter any calculation: 10 / 2
Result: 5
Enter any calculation: exit


### 🏗️ Multi-Stage Build Explanation
🔹 Stage 1: Builder
Uses golang:1.20-alpine

Compiles the statically-linked Go binary

🔹 Stage 2: Minimal Runtime
Uses the scratch base image

Contains only the final binary

No shell, no package manager — zero bloat

### ✅ Benefits of scratch
Extremely small image size

Minimal attack surface

Great for production

ℹ️ Note: Your Go binary must be statically linked. Go does this by default when no CGO dependencies are used.

