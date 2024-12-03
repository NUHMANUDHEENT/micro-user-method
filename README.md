# Microservices with Go: User Management and Method Handling

## Overview
This project implements two Go microservices with gRPC communication:
- **Microservice 1**: Handles user management (CRUD) with PostgreSQL and Redis.
- **Microservice 2**: Executes parallel and sequential methods with database interactions.

## Features
- **Microservice 1**:
  - Create, retrieve, update, and delete user records.
  - Caches user data in Redis for faster access.
- **Microservice 2**:
  - Executes `Method 1` sequentially and `Method 2` in parallel.
  - Communicates with Microservice 1 via gRPC to fetch user data.

## Prerequisites
- Go (1.19+), Docker, Kubernetes, PostgreSQL, Redis, `kubectl`

## Quick Setup

### **Clone Repository**
```bash
git clone https://github.com/NUHMANUDHEENT/micro-user-method.git
cd micro-user-method
