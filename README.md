
# Project Setup

This document provides a step-by-step guide to set up the project, including database configuration, mock data generation, and cleanup instructions.

---

## Prerequisites

1. **Docker**: Ensure Docker is installed to set up PostgreSQL.
2. **Go**: Install Go (version 1.18 or higher recommended).
3. **Database Migration Tool**: Install `migrate` or a similar tool for managing database migrations.

---

## Step 1: Set Up PostgreSQL with Docker

Run the following command to start a PostgreSQL container:

```bash
docker run --name postgres-container -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=postgres -p 5432:5432 -d postgres
```

---

## Step 2: Database Configuration

Add the following details to your configuration:

```yaml
database:
  db: 'postgres'
  dsn: "postgres://postgres:123456@localhost:5432/postgres?sslmode=disable"
  tables:
    - 'users'
    - 'orders'
    - 'products'
    - 'stock'
    - 'payments'

  outFile: 'gen.go'  # All models combined into a single file
  outPath: './model/query' # The output directory
  modelPkgName: '../../../../modules/entities/dbmodels'
  fieldWithTypeTag: true
  fieldNullable: true
  fieldWithIndexTag: true
  withUnitTest: false
  onlyModel: true
  fieldSignable: false

```
### Step 2.1: Database Schema

Run Script gen models
```
cd .\pkg\databases\migrations\
gentool -c "./gorm-gen-table.tool.yaml"
```

---

## Step 3: Database Schema

Run the following SQL script to create the tables:

```sql
-- Users Table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Products Table
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Stock Table
CREATE TABLE stock (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    quantity INT NOT NULL CHECK (quantity >= 0),
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);

-- Payments Table
CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    product_id INT NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);

-- Orders Table
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);
```
---

## Step 4: Generate Mock Data

Run the following SQL script to generate mock data:

```sql
-- Insert mock data into Users table
INSERT INTO users (name, email)
VALUES 
('Alice Smith', 'alice@example.com'),
('Bob Johnson', 'bob@example.com'),
('Charlie Brown', 'charlie@example.com');

-- Insert mock data into Products table
INSERT INTO products (name, description, price)
VALUES 
('Laptop', 'A high-end gaming laptop', 1200.00),
('Smartphone', 'A smartphone with the latest features', 800.00),
('Headphones', 'Noise-cancelling headphones', 150.00);

-- Insert mock data into Stock table
INSERT INTO stock (product_id, quantity)
VALUES 
(1, 50),
(2, 100),
(3, 200);

-- Generate 1000 mock orders
WITH users AS (
    SELECT 1 AS id UNION ALL SELECT 2 UNION ALL SELECT 3
), 
products AS (
    SELECT 1 AS id, 'Laptop' AS name UNION ALL 
    SELECT 2, 'Smartphone' UNION ALL 
    SELECT 3, 'Headphones'
)
INSERT INTO orders (user_id, product_id, quantity, order_date)
SELECT 
    (SELECT id FROM users OFFSET FLOOR(RANDOM() * 3) LIMIT 1),
    (SELECT id FROM products OFFSET FLOOR(RANDOM() * 3) LIMIT 1),
    FLOOR(RANDOM() * 5) + 1,
    NOW() - (RANDOM() * (INTERVAL '30 day'))
FROM generate_series(1, 1000);
```

---

## Step 5: Drop Tables Script

If you need to clean up the database, use the following script:

```sql
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS payments;
DROP TABLE IF EXISTS stock;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS users;
```

---

## Step 6: Run the API

Run the Go Fiber API:

```
cd app && go run main.go
```

---

## License

This project is licensed under the MIT License.
