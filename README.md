<div align="center">
  <img src="https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg" width="100">
</div>

# Fenlie - Split Bill API Service

Fenlie (分列), which means "split" in Chinese, is a powerful RESTful API service designed to facilitate split bill transactions. As developers, we provide this API for clients to integrate into their applications, enabling seamless transaction processing and bill splitting functionality.

## 🚀 Features

- 👥 User Management (Employee and Admin)
- 🏢 Company Registration (Internal access only)
- 🔑 Secure Client Key Generation
- 📦 Product and Category Management
- 💳 Transaction Processing
- 📧 Split Bill with Email Notifications
- 💰 Payment Gateway Integration with [Brick](https://www.onebrick.io/)

## 💻 Tech Stack

<div align="center">

![Go](https://img.shields.io/badge/-Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/-Gin-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![GORM](https://img.shields.io/badge/-GORM-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/-PostgreSQL-336791?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/-Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![JWT](https://img.shields.io/badge/-JWT-000000?style=for-the-badge&logo=json-web-tokens&logoColor=white)
![Clean Architecture](https://img.shields.io/badge/-Clean%20Architecture-FF6C37?style=for-the-badge&logo=clean-architecture&logoColor=white)

</div>

## 🏁 Getting Started

1. Clone the repository
   ```
   git clone https://github.com/your-username/fenlie.git
   ```

2. Navigate to the project directory
   ```
   cd fenlie
   ```

3. Build and run with Docker
   ```
   docker-compose up --build
   ```

4. The API will be available at `http://localhost:8080`

## 📘 API Documentation

API documentation is available via Postman. Import the following files into your Postman application:

- Collection: [Postmant Collection](docs/finpro-fenlie.postman_collection.json)
- Environment: [Env](docs/fenlie.postman_environment.json)

## 🗄️ Database Schema

<img src="docs/Fenlie ERD.drawio.svg"/>

<!-- ## 🧪 Testing

To run the unit tests:

```
go test ./...
``` -->

## 🚢 Deployment

Fenlie is containerized using Docker for easy deployment. Ensure you have Docker and Docker Compose installed on your deployment environment.

<br/>
<br/>
<div align="center">

  [![Website](https://img.shields.io/badge/-My_Website-black?style=for-the-badge&logo=google-chrome&logoColor=white)](https://lutfikhoir.com/)
  [![Instagram](https://img.shields.io/badge/-Instagram-E4405F?style=for-the-badge&logo=instagram&logoColor=white)](https://www.instagram.com/lutfi.khoirudin/)
  [![YouTube](https://img.shields.io/badge/-YouTube-FF0000?style=for-the-badge&logo=youtube&logoColor=white)](https://www.youtube.com/@lutfikhoir2502)
  [![Twitter](https://img.shields.io/badge/-Twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white)](https://x.com/LutfiKhoirudin)
  [![LinkedIn](https://img.shields.io/badge/-LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/lutfi-khoir-632524235/)
</div>