# Kafka Producer-Consumer Example

This project demonstrates a simple Kafka producer-consumer setup using Go and Docker. The producer sends messages to a Kafka topic, and the consumer reads these messages from the same topic.

## Project Structure

The project is structured as follows:

- `docker-compose.yml`: This file contains the configuration for the Docker services, including Kafka, the producer, and the consumer.
- `producer.go`: This is the Go code for the Kafka producer. It sends messages to a Kafka topic.
- `consumer.go`: This is the Go code for the Kafka consumer. It reads messages from the Kafka topic.

## How It Works

1. The producer establishes a connection to the Kafka broker and sends messages to a specified topic.
2. The consumer, which is also connected to the Kafka broker, reads these messages from the same topic.

## Running the Project Locally

To run this project on your local machine, follow these steps:

1. Ensure you have Docker and Docker Compose installed on your machine.
2. Clone the repository to your local machine.
3. Navigate to the project directory.
4. Run the following command to start the services:


```bash
docker-compose up --build
```

This command builds the Docker images for the producer and consumer (if they haven't been built already), and starts the services as defined in the `docker-compose.yml` file.

You can create multiple instances of the producer and consumer by running the following command:

```bash
docker-compose up --scale producer=2 --scale consumer=2
```

## Example

Here's an example of what you might see in the logs when running this project:

```bash
producer_1  | sent message: Hello Kafka
consumer_1  | message at offset 1:  = Hello Kafka
producer_1  | sent message: Hello Kafka
consumer_1  | message at offset 2:  = Hello Kafka
```

This shows the producer sending messages ("Hello Kafka") to the Kafka topic, and the consumer reading these messages from the topic.