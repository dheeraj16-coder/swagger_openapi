<p align="center">
  <img src="https://img.icons8.com/color/96/globe--v1.png" alt="Country Explorer Logo" width="80"/>
</p>

<h1 align="center">🌍 Country Explorer</h1>

<p align="center">
  <strong>Discover the nations across the globe</strong><br>
  A full-stack application to search and explore country data by name, capital, currency, language, and more.
</p>

<p align="center">
  <a href="https://ip5g8umkrk.us-east-1.awsapprunner.com">
    <img src="https://img.shields.io/badge/🚀_Live_Demo-Click_Here-blue?style=for-the-badge" alt="Live Demo"/>
  </a>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go"/>
  <img src="https://img.shields.io/badge/React-20232A?style=flat-square&logo=react&logoColor=61DAFB" alt="React"/>
  <img src="https://img.shields.io/badge/TypeScript-007ACC?style=flat-square&logo=typescript&logoColor=white" alt="TypeScript"/>
  <img src="https://img.shields.io/badge/Docker-2496ED?style=flat-square&logo=docker&logoColor=white" alt="Docker"/>
  <img src="https://img.shields.io/badge/AWS-232F3E?style=flat-square&logo=amazon-aws&logoColor=white" alt="AWS"/>
  <img src="https://img.shields.io/badge/GitHub_Actions-2088FF?style=flat-square&logo=github-actions&logoColor=white" alt="GitHub Actions"/>
</p>

<p align="center">
  <img src="https://github.com/dheeraj16-coder/swagger_openapi/actions/workflows/deploy.yml/badge.svg" alt="Build Status"/>
  <img src="https://img.shields.io/badge/License-MIT-green.svg" alt="License"/>
</p>

---

## 📸 Screenshots

<p align="center">
  <img src="images/image copy.png" alt="Search by Name" width="80%"/>
</p>

<p align="center">
  <img src="images/image copy 3.png" alt="Search by Capital" width="80%"/>
</p>

<p align="center">
  <img src="images/image.png" alt="Search by Language" width="80%"/>
</p>

---

## ✨ Features

| Feature | Description |
|---------|-------------|
| 🔍 **Multi-Search** | Search countries by name, capital, language, currency, or country code |
| 🎯 **Field Filtering** | Select specific fields to display (population, flags, region, etc.) |
| 🌙 **Dark Mode** | Beautiful dark theme UI |
| 📱 **Responsive** | Works on desktop and mobile |
| ⚡ **Fast** | Optimized API responses |
| 🏥 **Health Checks** | Production-ready health endpoint |

---

## 🏗️ Architecture

```
┌─────────────────┐         ┌─────────────────┐         ┌─────────────────┐
│                 │         │                 │         │                 │
│  React Frontend │ ──────▶ │   Go Backend    │ ──────▶ │ RestCountries   │
│  (TypeScript)   │         │   (Gin)         │         │ External API    │
│                 │         │                 │         │                 │
└─────────────────┘         └─────────────────┘         └─────────────────┘
        │                           │
        │                           │
        ▼                           ▼
┌─────────────────┐         ┌─────────────────┐
│     Nginx       │         │  OpenAPI Spec   │
│  (Reverse Proxy)│         │  (Code Gen)     │
└─────────────────┘         └─────────────────┘
        │
        ▼
┌─────────────────────────────────────────────┐
│              AWS App Runner                  │
│         (Auto-scaling, HTTPS)               │
└─────────────────────────────────────────────┘
```

---

## 🛠️ Tech Stack

### Backend
| Technology | Purpose |
|------------|---------|
| **Go** | Backend language |
| **Gin** | HTTP web framework |
| **OpenAPI Generator** | Auto-generated server stubs & client |

### Frontend
| Technology | Purpose |
|------------|---------|
| **React** | UI library |
| **TypeScript** | Type-safe JavaScript |
| **Vite** | Build tool |
| **OpenAPI Generator** | Auto-generated API client |

### DevOps
| Technology | Purpose |
|------------|---------|
| **Docker** | Containerization |
| **AWS App Runner** | Deployment & hosting |
| **Amazon ECR** | Container registry |
| **GitHub Actions** | CI/CD pipeline |
| **Nginx** | Reverse proxy |

---

## 📡 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/health` | Health check |
| `GET` | `/v3.1/all` | Get all countries |
| `GET` | `/v3.1/name/:name` | Search by country name |
| `GET` | `/v3.1/capital/:capital` | Search by capital city |
| `GET` | `/v3.1/currency/:currency` | Search by currency code |
| `GET` | `/v3.1/lang/:language` | Search by language |
| `GET` | `/v3.1/alpha/:code` | Search by country code (2 or 3 letter) |
| `GET` | `/v3.1/region/:region` | Search by region |
| `GET` | `/v3.1/subregion/:subregion` | Search by subregion |
| `GET` | `/v3.1/translation/:translation` | Search by translation |
| `GET` | `/v3.1/independent` | Get independent countries |

### Query Parameters

All endpoints support the `fields` query parameter for filtering:

```bash
GET /v3.1/name/india?fields=name,capital,population,flags
```

---

## 🚀 Getting Started

### Prerequisites

- Docker & Docker Compose
- Go 1.21+ (for local development)
- Node.js 18+ (for local development)

### Run Locally

```bash
# Clone the repository
git clone https://github.com/dheeraj16-coder/swagger_openapi.git
cd swagger_openapi

# Start with Docker Compose
docker-compose up --build

# Access the app
open http://localhost
```

### Run Without Docker

**Backend:**
```bash
cd my-go-backend
go mod tidy
go run .
# Server runs on http://localhost:8080
```

**Frontend:**
```bash
cd my-ui-project
npm install
npm run dev
# App runs on http://localhost:5173
```

---

## 🔧 OpenAPI Code Generation

This project uses **API-first development** with OpenAPI specification.

### Generate Go Client
```bash
openapi-generator generate \
  -i rest-countries-api.yaml \
  -g go \
  -o restcountries-go-client
```

### Generate TypeScript Client
```bash
openapi-generator generate \
  -i rest-countries-api.yaml \
  -g typescript-axios \
  -o my-ui-project/src/api-client
```

---

## 📦 Deployment

The app auto-deploys to AWS on every push to `main`:

```
Push to main → GitHub Actions → Build Docker Images → Push to ECR → Deploy to App Runner
```

### Manual Deployment

```bash
# Build images for AMD64 (required for AWS)
docker buildx build --platform linux/amd64 -t countries-backend -f my-go-backend/Dockerfile .
docker buildx build --platform linux/amd64 -t countries-frontend -f my-ui-project/Dockerfile .

# Tag and push to ECR
docker tag countries-backend:latest <account-id>.dkr.ecr.us-east-1.amazonaws.com/countries-backend:latest
docker push <account-id>.dkr.ecr.us-east-1.amazonaws.com/countries-backend:latest
```

---

## 🗺️ Roadmap

- [x] Core search functionality
- [x] Docker containerization
- [x] AWS deployment
- [x] CI/CD pipeline
- [x] Health check endpoint
- [ ] Redis caching
- [ ] Circuit breaker pattern
- [ ] Rate limiting
- [ ] Prometheus metrics
- [ ] Grafana dashboard

---

## 📁 Project Structure

```
swagger_openapi/
├── my-go-backend/           # Go backend service
│   ├── go/                  # API handlers
│   ├── Dockerfile
│   └── main.go
├── my-ui-project/           # React frontend
│   ├── src/
│   │   ├── components/
│   │   └── api-client/      # Generated TypeScript client
│   ├── Dockerfile
│   └── nginx.conf
├── restcountries-go-client/ # Generated Go client
├── rest-countries-api.yaml  # OpenAPI specification
├── docker-compose.yml
└── .github/workflows/       # CI/CD
```

---

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 👤 Author

**Dheeraj Sai**

- GitHub: [@dheeraj16-coder](https://github.com/dheeraj16-coder)
- LinkedIn: [Your LinkedIn](https://www.linkedin.com/in/dheerajsai16)

---

<p align="center">
  Made with ❤️ and ☕
</p>

<p align="center">
  <a href="https://ip5g8umkrk.us-east-1.awsapprunner.com">
    <img src="https://img.shields.io/badge/Try_it_Live-🌍-blue?style=for-the-badge" alt="Live Demo"/>
  </a>
</p>