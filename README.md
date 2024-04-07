# Apache Kafka Project - Orders

1. **Setup Docker:**
   - Install Docker by following the instructions on this site: [Docker Installation Guide](https://docs.docker.com/get-docker/).

2. **Start Containers:**
   - Navigate to the project root and run the command `docker-compose up` to start the containers.

3. **Run Producer:**
   - Open a terminal inside the `producer` package.
   - Run `go run main.go` to start the producer, which sends orders in JSON format to a Kafka topic.

4. **Run Consumer:**
   - Open another terminal inside the `consumer` package.
   - Run `go run main.go` to start the consumer, which reads the messages (orders) from the Kafka topic.

5. **View Output:**
   - You should now see the producer sending each order and the consumer receiving them.

