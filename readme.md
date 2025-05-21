# Task Service

This project is a microservice for managing tasks, using Go, PostgreSQL, Kafka, and Zookeeper. It is designed to be run locally using Docker Compose for all dependencies.

## 1. Docker Compose Setup

This will start Postgres, Zookeeper, Kafka, and the Task Service app:

```sh
docker-compose up --build
```

- The API will be available at [http://localhost:8080](http://localhost:8080)
- Postgres will be available at `localhost:5432` (from your host)
- Kafka will be available at `localhost:9092` (from your host)

To stop and remove all containers and volumes:

```sh
docker-compose down -v
```

## 2. Endpoints

To interact with the Task Service via POSTMAN, you can use the following endpoints:

- `POST   /createTask` — Create a new task
   - curl -X POST -H "Content-Type: application/json" -d '{"title":"Write blog","status":"Pending"}' http://localhost:8080/tasks

- `GET    /tasks`      — List tasks
   - curl "http://localhost:8080/tasks?status=Pending&limit=5&offset=0"

- `PUT    /tasks/{id}` — Update a task
   - curl -X PUT http://localhost:8080/tasks/1 \
    -H "Content-Type: application/json" \
    -d '{"title": "Updated Task Title", "status": "Completed"}'
    
- `DELETE /tasks/{id}` — Delete a task
   - curl -X DELETE http://localhost:8080/tasks/1

## 3. Kafka Integration

- The Task Service consumes messages from the `tasks` topic in Kafka.
- You can publish task events (e.g., create, update) to Kafka from other microservices.

NOTE: This kafka consumer is for demonstration purposes only. As there is no producer producing any messages, the consumer will not receive any messages.

## 4. Notes

- Zookeeper and Kafka are managed by Docker Compose; **no local installation is needed**.
- Make sure Docker Desktop is running on your machine.
- For any issues, check logs with `docker-compose logs <service>`.
  - Eg: `docker-compose logs app`
  - `docker-compose logs kafka`
- Update and Delete Api's have lock in place to avoid conflicts

## 5. System Design:

If I add another microservice (like a User Service), the services can communicate using message queues (Kafka). For example, the User Service can publish events to the `tasks` topic, and the Task Service will automatically consume and process those events based on their type. This keeps things decoupled and scalable. REST or gRPC are also options, but Kafka is already integrated and works well for async workflows.

To scale task service horizontally, I can run multiple instances of the Task Service behind a load balancer. Since the service is stateless and uses Postgres and Kafka, each instance can handle requests or consume messages independently. This way, the system can handle more load just by adding more containers. Update and Delete Api's have lock in place to avoid conflicts.

## 6. TODOS:

- Unit/Integration Tests
