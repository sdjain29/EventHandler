# Event Handler - MVP

## Introduction

This document provides an overview of the Minimum Viable Product (MVP) for an Event Routing and Delivery System that meets the specified requirements. The system is designed to fan-out events to multiple destinations while handling failures, ensuring durability, maintaining order, and providing at-least-once delivery.

## Requirements

The following key requirements are addressed by the MVP:

- **Fan-out and Failure Handling**: Distribute events to multiple destinations while handling downstream failures.
- **Durability**: Store ingested events in Redis to maintain data even during system crashes.
- **At-least-once Delivery**: Implement retries to ensure event delivery even in case of destination failures.
- **Retry Backoff and Limit**: Use a backoff algorithm for message retries, and eventually drain events from the system after a certain number of attempts.
- **Maintaining Order**: Guarantee First-In-First-Out (FIFO) order for events delivered to each destination.
- **Delivery Isolation**: Prevent delays or failures in one destination's delivery from affecting others.

## Architecture Overview

The MVP architecture consists of the following components:

1. **Event Ingestion**: Accepts incoming events and stores them in Redis.
2. **Event Routing**: Routes events from Redis to different destinations based on configuration.
3. **Retry Manager**: Manages event retry logic, applying backoff and limiting retries.
4. **Destination Delivery**: Handles delivery of events to individual destinations.
5. **Redis**: Acts as the persistent storage for events.

## Components

### Event Ingestion
Incoming events are accepted and stored in Redis, ensuring durability.

### Event Routing
Events are fetched from Redis and routed to multiple destinations based on configuration. If a destination is unavailable, the event is handed over to the Retry Manager.

### Retry Manager
Manages the retry process for events that fail to be delivered to destinations. Implements backoff and retry limiting. After a maximum number of retries, the event is removed from the system.

### Destination Delivery
Handles the delivery of events to individual destinations. Retried events are prioritized, ensuring at-least-once delivery.

### Redis
Serves as the persistent storage for events, supporting durability.

## Data Flow

1. Events are ingested and stored in Redis.
2. The Event Routing component fetches events from Redis and sends them to respective destinations.
3. If delivery to a destination fails, the event is passed to the Retry Manager.
4. The Retry Manager applies backoff and retry limits while attempting to deliver events.
5. If the retry limit is exceeded, the event is removed from the system.

## Retry Mechanism

The Retry Manager employs an exponential backoff algorithm to retry events, gradually increasing the time between retries. This prevents overwhelming the destination and network with frequent retries.

## Maintaining Order

Events have event timestamp attached to it which will set order once delivery to destination is done.

## Delivery Isolation

Failures or delays in delivering events to one destination do not impact the ingestion process or delivery to other destinations. Each destination's delivery is independent.

## Running the System

To run the Event Routing and Delivery System MVP, follow these steps:

1. Ensure you have Docker and Docker Compose installed on your system.
2. Clone this repository to your local machine.
3. Open a terminal and navigate to the main folder of the cloned repository.
4. Run the following command to start the system using Docker Compose:

```sh
docker-compose up
```

## Test Cases

| Case No. | Name                                   | Description                                          | Method                  |
|----------|----------------------------------------|------------------------------------------------------|-------------------------|
| 1        | Happy Flow                             | Event should get ingested and delivered to Destination | Store in Redis, Delivery to Destination |
| 2        | Ingest Event                           | Test if an incoming event is successfully ingested   | Ingest and Redis Update |
| 3        | Event Fan-out                          | Test if events are correctly routed to destinations  | Fetch from Redis, Route to Destinations |
| 4        | Retry on Destination Failure           | Test if the system retries on destination failure   | Retry Mechanism          |
| 5        | Retry Backoff Algorithm                | Test if the retry backoff algorithm is applied      | Backoff Algorithm       |
| 6        | Max Retry Limit                        | Test if events are removed after max retries        | Retry Count Limit       |
| 7        | Delivery Isolation                     | Test if one destination's failure doesn't affect others | Isolated Delivery       |
| 8        | Durability on System Crash             | Test if ingested events survive a system crash      | Redis Persistence       |
| 9        | At-least-once Delivery                 | Test if events are delivered at least once          | Retries and Acknowledgment |
| 10       | Running the System with Docker Compose | Test running the system using Docker Compose        | Command Line            |
| 11       | Invalid Event Ingestion                | Test if invalid events are rejected                | Validation Logic        |
| 12       | Concurrent Event Ingestion             | Test simultaneous ingestion of multiple events     | Parallel Ingestion      |
| 13       | Event Routing Configuration            | Test different event routing configurations        | Configuration Settings |
| 14       | Event Retry Configuration              | Test different event retry configurations          | Configuration Settings |
| 15       | Delivery to Multiple Destinations      | Test fan-out to multiple destinations              | Multiple Destinations  |
| 16       | Successful Event Delivery              | Test successful event delivery to a destination    | Successful Delivery     |
| 17       | Event Delivery Retry                   | Test if the system retries event delivery          | Retry Mechanism         |
| 18       | Delivery Failure Handling              | Test handling of delivery failures                | Error Handling          |
| 19       | System Stability                       | Test system stability over extended operation     | Long-duration Testing  |
| 20       | Integration with Kafka                 | Test integration of Kafka for event processing     | Future Enhancement     |
| 21       | Security and Authentication            | Test security measures for incoming events        | Future Enhancement     |
| 22       | Event Processing Component             | Test event processing between ingestion and delivery | Future Enhancement     |

## Future Enhancements

In the future, the system can be enhanced by introducing the following features:

- **Kafka Integration**: Add Kafka for improved event processing and handling high-throughput scenarios in destination delivery side.
- **Security and Authentication**: Implement security measures for incoming events to ensure data integrity.
- **Event Processing**: Introduce a component to process events between ingestion and delivery for data enrichment or transformation.

This MVP provides a solid foundation for the Event Routing and Delivery System, with the potential for further advancements to meet evolving needs.

For questions, concerns, or feedback, please contact.
