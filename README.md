# GRPC Demo


## Architecture
```mermaid
flowchart TD
    Client[Client] -->|HTTP/REST| Gateway[API Gateway]
    Gateway -->|gRPC| UserService[User Service<br>:50052<br>/user.v1.UserService]
    Gateway -->|gRPC| TaskService[Task Service<br>:50051<br>/task.v1.TaskService]
    
    subgraph User Service API
        UserService -->|UserInfo| UserInfo[/UserInfo<br>user_id -> user details/]
        UserService -->|UserCreate| UserCreate[/UserCreate<br>email, username -> user_id/]
    end
    
    subgraph Task Service API
        TaskService -->|TaskInfo| TaskInfo[/TaskInfo<br>task_id -> task details/]
        TaskService -->|TaskCreate| TaskCreate[/TaskCreate<br>name, description -> task_id/]
    end
```

## Run
```
docker compose up [--detach] --build
```

## Test
User service
```
grpcurl -plaintext -d '{"user_id": "123"}' localhost:50052 user.v1.UserService/UserInfo
grpcurl -plaintext -d '{"email": "new@example.com", "username": "newuser"}' localhost:50052 user.v1.UserService/UserCreate
```

Task service
```
grpcurl -plaintext -d '{"name": "My New Task", "description": "This is a task grpcurl -plaintext -d '{"task_id": "123e4567-e89b-12d3-a456-426614174000"}' localhost:50051 task.v1.TaskService/TaskInfo
```
