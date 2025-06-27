# Env Crafter

**Env Crafter** is a cross-platform desktop app that lets developers visually build and manage Docker-based development environments — without writing a single line of `docker-compose.yaml`.

<p align="center">
  <img src="https://your-screenshot-url" width="600" alt="Env Crafter UI Preview" />
</p>

---

## 🚀 Why Env Crafter?

Spinning up dev environments with Docker should be fast and painless. Whether you're a junior developer or a seasoned engineer, `docker-compose` can be repetitive, hard to maintain, or error-prone. **Env Crafter** solves that by giving you:

- A clean **visual interface** to add and configure services (PHP, PostgreSQL, Redis, Node.js, etc).
- Smart generation of production-ready `docker-compose.yaml` files.
- Cross-platform support (Windows, macOS, Linux).
- Future integrations: LLM-based assistant, pro templates, stack sharing.

---

## ✨ Features

- 🐳 Visual Docker Compose builder
- ⚙️ Support for multiple runtime stacks
- 📦 Prebuilt templates for common use cases
- 💾 Save & reuse environment configs
- 📁 One-click export to `docker-compose.yaml`
- 🔜 Upcoming: Plugin system, Pro features, AI assistant

---

## 🧑‍💻 Built With

- [Go](https://go.dev/) — core logic and backend
- [Wails](https://wails.io/) — for cross-platform GUI
- [Docker](https://www.docker.com/) — environment engine
- [YAML](https://yaml.org/) — config serialization

---

## 📥 Getting Started

> 🚧 **This project is under active development.**  
> A public alpha release will be available in **August 2025**.

In the meantime:

```bash
git clone https://github.com/yourusername/env-crafter.git
cd env-crafter
# Follow Wails setup instructions
wails dev
