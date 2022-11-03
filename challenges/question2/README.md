# Challenge Summary

We are in the business of selling Starwars themed Legos. There is a database which stores data related to all the Lego products and its inventories. The database consists of two tables with the following schema:

**products**
- product_id (int): unique order identifier
- name (string): name of the product

**inventories**
- product_id (string): unique product identifier
- inventory (int): the count of the product inventory
- updated_at (datetime): the datetime when the inventory was updates

## Assumptions

- All inventory counts are whole numbers

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

## Tasks

Write a query to find the Lego sets which has the lowest inventory in our database and order them in descending order of updated_at.