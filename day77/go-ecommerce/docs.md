# GO-ECOMMERCE API DOCUMENTATION

## SIGN UP

`POST /api/v1/signup`

Used to register a new user.

```curl
curl --location 'localhost:5050/api/v1/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "kevindoe",
    "email": "kevin@mail.com",
    "password": "kevinlovesgolang"
}'
```

Response:
#### Success - 200 OK:
```curl
{
    "message": "account creation is successful"
}
```

#### Error - 400 Bad Request:
```curl
{
    "error": "Empty username field."
}
```

## LOGIN

`POST /api/v1/login`

Logs a user in and assigns a JWT token

```curl
curl --location 'localhost:5050/api/v1/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "kevindoe",
    "password": "kevinlovesgolang"
}'
```

### RESPONSE:
#### Success - 200 OK:
```
{
    "msg": "Login successful."
}
```

#### ERROR - 401 UNAUTHORIZED:
```
{
    "error": "User is not authorized"
}
```

## LOGOUT
`GET /api/v1/logout`
Logs out users and deletes cookie which is used for auth.

```
curl --location 'localhost:5050/api/v1/logout'
```
### RESPONSE
#### Success 200 OK:
`"Logout success."`


## Add Product
`POST /api/v1/product` <br>
Adds new product to the database

```
curl --location 'localhost:5050/api/v1/product' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Air Jordan 6",
    "description": "air jordan gives you wings.",
    "category": "shoes",
    "color": "green",
    "product_size": 43,
    "price": 799.9,
    "quantity": 7
}'
```
### RESPONSE:
#### Success 200 OK:
```
{
    "message": "Product added successful.",
    "product_details": {
        "id": 0,
        "name": "Air Jordan 6",
        "description": "air jordan gives you wings.",
        "sku": "SHS-GRN-0043",
        "category": "shoes",
        "category_id": 4,
        "color": "green",
        "product_size": 43,
        "price": 799.9,
        "quantity": 7,
        "created_at": "2025-01-26T22:24:46.846810539+01:00",
        "updated_at": "2025-01-26T22:24:46.846810762+01:00"
    }
}
```

#### ERROR 400 BAD REQUEST:
Missing important product details
```
curl --location 'localhost:5050/api/v1/product' \
--header 'Content-Type: application/json' \
--data '{
    "name": "",
    "description": "air jordan gives you wings.",
    "category": "shoes",
    "color": "green",
    "product_size": 43,
    "price": 799.9,
    "quantity": 7
}'
```
```
{
    "error": "Name field can not be empty."
}
```
```
{
    "error": "Color field can not be empty."
}
```

## Find Product By Product Category Name

finds products with the passed parameter name

`curl --location 'localhost:5050/api/v1/product/shoes'`
### RESPONSE
#### Success 200 OK
```
[
    {
        "id": 4,
        "name": "Black Airforce One",
        "description": "shoes for GOATS. Signed by micheal Jordan himself.",
        "sku": "SHS-BLK-0043",
        "category": "",
        "category_id": 4,
        "color": "black",
        "product_size": 43,
        "price": 1001.2,
        "quantity": 16,
        "created_at": "2025-01-23T23:50:17Z",
        "updated_at": "2025-01-25T15:55:56Z"
    },
    {
        "id": 8,
        "name": "Oxford",
        "description": "The perfect loafers for gentle men.",
        "sku": "SHS-BLK-0041",
        "category": "",
        "category_id": 4,
        "color": "black",
        "product_size": 41,
        "price": 0,
        "quantity": 6,
        "created_at": "2025-01-24T13:34:36Z",
        "updated_at": "2025-01-24T13:34:36Z"
    },
    {
        "id": 10,
        "name": "Mens XXL",
        "description": "Official gucci wears for men.",
        "sku": "SHS-BLK-0040",
        "category": "",
        "category_id": 4,
        "color": "black",
        "product_size": 40,
        "price": 800.2,
        "quantity": 4,
        "created_at": "2025-01-24T21:03:00Z",
        "updated_at": "2025-01-25T09:56:12Z"
    },
    {
        "id": 21,
        "name": "X trainers",
        "description": "Best trainers for your daily jogging and style. Perfect for gym rats.",
        "sku": "SHS-BLK-0001",
        "category": "",
        "category_id": 4,
        "color": "black",
        "product_size": 1,
        "price": 7001.2,
        "quantity": 21,
        "created_at": "2025-01-25T08:03:10Z",
        "updated_at": "2025-01-25T08:03:10Z"
    },
    {
        "id": 24,
        "name": "Air Jordan 6",
        "description": "air jordan gives you wings.",
        "sku": "SHS-GRN-0043",
        "category": "",
        "category_id": 4,
        "color": "green",
        "product_size": 43,
        "price": 799.9,
        "quantity": 7,
        "created_at": "2025-01-26T21:24:47Z",
        "updated_at": "2025-01-26T21:24:47Z"
    }
]
```
#### ERROR 404 NO PRODUCT FOUND
`curl --location 'localhost:5050/api/v1/product/BAGS'`

Product with category name was not found
```
{
    "error": "Oops! Nothing Found."
}
```