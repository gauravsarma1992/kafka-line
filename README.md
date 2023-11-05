# Kafka Pipeline

Kafka Pipeline is a framework which allows users to define
the stream of events from one Kafka topic to another after
performing different types of transformations.

## Why is this needed?
Kafka is the de-facto choice for systems to share their
knowledge with other systems. Kafka connect is a good
example of how different types of transformations are
performed between Sink and Source connectors.
As the data flows in and out of different systems, the
major change is the ability of engineers to perform
different operations with the output which can be a
different Kafka topic chained together.
Performing these operations at scale and allow engineers
a simple DSL to perform these operations would reduce
a lot of intermediate services.

## How does this work?
Users can define the source from where the message should
be read, the operation to perform and the destination where
it should be sent next.
The sourcce and destination can either be another system, or
another Kafka topic, which would allow users to perform complex
pipelined operations at scale while providing transactional
guarantees.
