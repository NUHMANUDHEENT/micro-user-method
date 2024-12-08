Here’s the updated and formatted **README.md** file, including instructions for using Docker Compose, API routes, and details on testing the microservices:

---

# **Microservices with Go: User Management and Method Handling**

## **Overview**
This project implements two Go microservices with gRPC communication:  
- **Microservice 1**: Handles user management (CRUD operations) with PostgreSQL and Redis.  
- **Microservice 2**: Executes parallel and sequential methods with database interactions.

---

## **Features**
### **Microservice 1**:  
- Create, retrieve, update, and delete user records.  
- Caches user data in Redis for faster access.  

### **Microservice 2**:  
- Executes **Method 1** sequentially and **Method 2** in parallel.  
- Communicates with **Microservice 1** via gRPC to fetch user data.

---

## **Prerequisites**
- **Go** (1.19 or higher)  
- **Docker** and **Docker Compose**  
- **PostgreSQL**  
- **Redis**  
- Optional: **Kubernetes** and `kubectl`  

---

## **Quick Setup**

### **Step 1: Clone Repository**
```bash
git clone https://github.com/NUHMANUDHEENT/micro-user-method.git
cd micro-user-method
```

### **Step 2: Run with Docker Compose**
Use the provided `docker-compose.yml` to spin up all services with one command:  
```bash
docker-compose up --build
```

This will start the following services:  
- **Microservice 1**: `http://localhost:8080`  HTTP port
- **Microservice 2**: `http://localhost:50060` Grpc port 
- **PostgreSQL**: Running on `localhost:5432`  
- **Redis**: Running on `localhost:6379`

---

## **API Endpoints**

### **Microservice 1**: User Management API
The User Management microservice provides the following REST endpoints:

| Method | Endpoint         | Description                   |
|--------|------------------|-------------------------------|
| POST   | `/user/create`   | Create a new user.            |
| GET    | `/user/:id`      | Retrieve user by ID.          |
| GET    | `/user/list`     | List all users.               |
| PUT    | `/user/:id`      | Update user details.          |
| DELETE | `/user/:id`      | Delete user by ID.            |
| GET    | `/user/methods`  | List all user names.          |

---

### **Microservice 2**: Method Execution API
The Method Execution microservice supports the following operations:

| Method | Endpoint | Description                                                 |
|--------|----------|-------------------------------------------------------------|
| GET   | `/methods` | Accepts a `method` (1 or 2) and `waitTime` (seconds).  |

#### **Behavior**:
- **Method 1**: Executes sequentially.  
- **Method 2**: Executes in parallel.

---

## **Testing the APIs**

1. **Create a User**  
   - Endpoint: `POST /user/create`  
   - Sample Request:  
     ```json
     {
       "name": "John Doe",
       "email": "john.doe@example.com",
       "phone":985940495
     }
     ```

2. **Retrieve a User by ID**  
   - Endpoint: `GET /user/:id`  
   - Example: `GET /user/1`

3. **Update User Details**  
   - Endpoint: `PUT /user/:id`  
   - Sample Request:  
     ```json
     {
       "name": "Jane Doe",
       "email": "jane.doe@example.com",
       "age": 32
     }
     ```

4. **Delete a User by ID**  
   - Endpoint: `DELETE /user/:id`

5. **List All Users**  
   - Endpoint: `GET /user/list`

6. **Execute a Method**  
   - Endpoint: `POST /method`  
   - Sample Request:  
     ```json
     {
       "method": 1,
       "waitTime": 10
     }
     ```
   - Methods:
     - **Method 1**: Executes tasks sequentially.  
     - **Method 2**: Executes tasks in parallel.  


---

## **Future Enhancements**
- Add support for Kubernetes deployment.  
- Enhance caching with Redis for additional optimizations.  
- Extend gRPC communication for additional services.

---
