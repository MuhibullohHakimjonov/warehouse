# Product Management System

Это система управления продуктами на основе Go, которая использует PostgreSQL с расширением `ltree` для управления продуктами, полками и их иерархическим расположением.

## Features

- **Product Management**: Create and search for products by name.
- **Shelf Management**: Create shelves with hierarchical paths using the `ltree` data type.
- **Validation**: Input validation for required fields and `ltree` path format.
- **Logging**: Request logging using `chi` middleware.

## Prerequisites

- Go 1.20 or higher
- PostgreSQL 12 or higher
- `ltree` extension enabled in PostgreSQL

## Setup

### 1. Clone the Repository

```bash
git clone git@github.com:MuhibullohHakimjonov/warehouse.git
cd cmd/server