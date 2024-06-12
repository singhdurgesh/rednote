# Rednote

## Description
Rednote is a Go-based backend project designed to provide a comprehensive set of user authentication methods, real-time chat and location tracking APIs, job/task processing, and a structured deployment pipeline. This repository aims to offer a robust and scalable backend solution for modern web and mobile applications.

## Features
### User Registration Methods
- **Using Username/Password** - Done
- **Using OTP on Phone/Email** - Done
- **Using Google SSO** - Done
- Using Apple SSO

### User Login Methods
- **Using Username/Password** - Done
- **Using OTP on Phone/Email** - Done
- **Using Google SSO** - Done
- Using Apple SSO

### User Password Reset
- Using OTP on Phone/Email

### Settings CRUD API
- **Accessible only for Admin User** - Done

### Real-Time Features
- Real-Time Chat Application APIs
- Real-Time Location Tracking APIs

### Database Management
- **Versioned Database Migration with Goose Package** - Done

### Job/Task Processing
- **RabbitMQ Integration** - Done
- **Machinery Package Integration** - Done
- **AsyncTask Module** - Done
  - **Notification Task Example** - Done
- AWS SQS Integration

### Notification Service
- **Abstract and Implementation of a Service** - Done

### Deployment Pipelines Setup
- **Docker Setup**
  - **Dockerfile Creation** - Done
  - **Docker Compose File** - Done
- **Github Pipelines**
  - **Workflow Setup** - Done
  - Workflow Deployment to AWS EC2

## TODO
- Containerized Deployments Guide
- Deployments without Downtime (Kubernetes | EKS)
- Setup Wiki for Deployment and Documentation
  - Exposing New APIs
  - Building New Data Models
  - Building New Services

## Getting Started

### Prerequisites
- Go 1.21+
- Docker
- Docker Compose
- RabbitMQ
- AWS Account (for SQS integration)
- Makefile

### Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/singhdurgesh/rednote.git
    cd rednote
    ```

2. Setup environment variables:
    Create a `.env` file in the root directory and populate it with necessary environment variables.
    ```sh
    cp .env.example .env
    cp configs/environments/config.docker.yaml.example configs/environments/config.docker.yaml
    cp configs/environments/config.yaml.example configs/environments/config.local.yaml
    ```

3. Build and run the application using Docker Compose:
    ```sh
    docker-compose up --build
    ```

4. Run Database Migration
     ```sh
    make migrate
    ```

5. Verify the Application Server
    ```sh
    curl localhost:8080/public/ping
    ```

### Deployment
#### Docker Setup
- Ensure Docker and Docker Compose are installed on the deployment server.
- Use the provided `Dockerfile` and `docker-compose.yml` for containerized deployments.

#### Github Pipelines
- Workflows are set up for CI/CD using Github Actions. Update the workflows as needed for your specific AWS EC2 deployment.

### Documentation
- Detailed documentation for API endpoints, data models, and services will be available in the Wiki section of this repository.

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request for review.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact
For any inquiries or issues, please contact [singhdurgesh403@gmail.com](mailto:singhdurgesh403@gmail.com).
