## [Como criar uma REST API completa do zero com GO | Golang tutorial - iniciante](https://www.youtube.com/watch?v=3p4mpId_ZU8)

### Creating Table
```sql
CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  product_name VARCHAR(50) NOT NULL,
  price NUMERIC(10, 2) NOT NULL
)

SELECT * FROM products;

INSERT INTO products (product_name, price) VALUES ('PlayStation', 100);
INSERT INTO products (product_name, price) VALUES ('XBOX', 100);
INSERT INTO products (product_name, price) VALUES ('Nintendo Switch', 100);
```