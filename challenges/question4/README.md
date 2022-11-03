# Challenge Summary

We are in the business of selling Star Wars themed Legos. There is a database which stores data related to all the Lego products and its inventories. The database consists of two tables with the following schema.

**products**
- product_id (int): unique product identifier
- name (string): name of the product

**inventories**
- product_id (string): unique product identifier
- inventory (int): the count of the product inventory available
- updated_at (datetime): the datetime when the inventory was updated

We need to build a REST API to expose some basic functionality to the client to retrieve, create, and update product and inventory records.

## Assumptions

1. We do not have to use a database for this exercise, instead, a very simple in-memory datastore is provided for you (see datastore.go for the interface and helpful usage example).
2. The test cases will automatically load the following seed data into the datastore and pass them to your server.
3. A basic router is provided that will forward requests to the relevant handler in your server, you just need to parse the request path and body (as needed), prepare and send an appropriate response.
4. An API spec is provided that shows all the operations you need to implement. Pay careful attention to the request methods and response codes. The tests will cover every error scenario included in the spec.

## Sample Data
**products**

product_id|name
---|---
1|R2D2
2|Millennium Falcon
3|Darth Vader
4|Yoda

**inventories**

product_id|inventory|updated_at
---|---|---
1|100|2022-01-01
2|200|2022-01-02
3|50|2022-01-03
4|10|2022-01-04

## API Spec
The API should support the following operations

**Get all inventories**

URL Path: `/inventories`

Sample Request: `GET /inventories`

Sample Response:

```json
[
  {
    "product_id": 1,
    "name": "R2D2",
    "inventory": 100,
    "updated_at": "2022-01-01T00:00:00Z"
  },
  {
    "product_id": 2,
    "name": "Millenium Falcon",
    "inventory": 200,
    "updated_at": "2022-01-02T00:00:00Z"
  }
]
```

Errors: N/A

**Get one inventory**

URL path: `/inventories/{product_id}`

Sample Request: `GET /inventories/1`

Sample response:

```json
{
    "product_id": 1,
    "name": "R2D2",
    "inventory": 100,
    "updated_at": "2022-01-01T00:00:00Z"
}
```

Errors:

- 404 (product not found): this status can also be used for invalid IDs

**Create a new inventory record**

URL path: `/inventories/{product_id}`

Sample Request: `PUT /inventories/5`

```json
{
    "product_id": 5,
    "name": "C-3PO",
    "inventory": 40,
    "updated_at": "2022-09-21T00:00:00Z"
}
```

Sample Response:

```json
{
   "product_id": 5,
   "name": "C-3PO",
   "inventory": 40,
   "updated_at": "2022-09-21T00:00:00Z"
}
```

Note: a successful create should return a 201 (created) status code.

Errors:

- 400 (Bad Request): malformed request (including mismatch between ID in URL and request body)
- 409 (Conflict): product with the same id already exists

**Update an existing inventory**

URL path: `/inventories/{product_id}`

Sample Request:

POST `/inventories/5`

```json
{
   "inventory_adjustment": -3
}
```

Response:

```json
{
   "product_id": 5,
   "name": "C-3PO",
   "inventory": 37,
   "updated_at": "2022-09-21T00:10:00Z"
}
```

Errors:

- 400 (Bad Request): malformed request. In addition to malformed JSON, 400 should be returned requested update would cause inventory to drop below zero.
- 404 (Not found): product not found

## Tasks
- Complete the methods in server.go to load data from the provided datastore and render appropriate response.
- Ensure that your server passes all the supplied tests.