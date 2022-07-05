This is a sample application to be used to create traffic inside a RabbitMQ cluster. It currently connects to a specific URI: `amqp://guest:guest@rabbitmq.sample-apps.svc.cluster.local:5672`.

In order to build your docker images. use the following commands from the project root:

`docker build -f ./src/consumer/Dockerfile -t gaantunes/go-amqp-consumer:latest .`

`docker build -f ./src/publisher/Dockerfile -t gaantunes/go-amqp-publisher:latest .`
