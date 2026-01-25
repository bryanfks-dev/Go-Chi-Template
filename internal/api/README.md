# API

## Overview

This directory contains the implementation of the API layer for the Chi service. It defines the HTTP routes, handlers, and middleware used to process incoming requests and send responses.

## Concepts

The API layer is build with DDD (Domain-Driven Design) principles in mind.
Each component has a specific responsibility, promoting maintainability and scalability. It separates concerns into different components:

- `route.go`: Defines the HTTP routes and associates them with their respective handlers.
- `delivery/`: Contains the handler functions that process requests and generate responses.
- `usecase/`: Contains the business logic for handling various operations requested via the API.
- `repository/`: Contains the data access layer for interacting with databases or other external services.
- `middleware/`: Contains middleware functions for tasks such as logging, authentication, and request validation.
- `domain/`: Contains domain-specific data structures used by the API (e.g. entities, errors, value objects, and other non-user input).
- `data/`: Contains data transfer objects (DTOs) and schemas used for request and response payloads.
- `utils/`: Utility functions used across the API layer.

## Getting Started

To get started with the API layer, follow these steps:

1. Define your routes in `route.go`.
2. Implement the handler functions in the `delivery/` directory.
3. Implement the business logic in the `usecase/` directory.
4. Set up data access in the `repository/` directory.
5. Add any necessary middleware in the `middleware/` directory.
6. Define your domain models in the `domain/` directory.
7. Create DTOs and schemas in the `data/` directory.
8. Use utility functions from the `utils/` directory only as needed.
9. Register the API routes with the main application router (`/cmd/http/route.go`, then choose either it's a public or internal API).
10. Test the API endpoints to ensure they work as expected.
11. Create a swagger documentation for the API endpoints if needed.
