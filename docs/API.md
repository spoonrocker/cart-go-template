## API Endpoints

### Create a new cart

URL: /carts  
METHOD: POST.  
BODY:
```json
{}
```
RESPONSE:
```json
Status: 201 Created
{
    "id": 1,
    "items": []
}
```

### Get a cart

URL: /carts/cart_id  
EXAMPLE: /carts/1  
METHOD: GET.  
RESPONSE:  
```
Status: 200 Ok
{
    "id": 1,
    "items": []
}
```

### Create a cart item

URL: /carts/cart_id/items  
EXAMPLE: /carts/1/items  
METHOD: POST.  
BODY:
```json
{
    "product": "name_of_the_product",
    "quantity": 5
}
```
RESPONSE:
```
Status: 201 Created
{
    "id": 1,
    "cart_id": 1,
    "quantity": 5,
    "product": "name_of_the_product"
}
```

### Delete a cart item

URL: /carts/cart_id/items /item_id  
EXAMPLE: /carts/1/items/1  
METHOD: DELETE.  
RESPONSE:
```
Status: 204 No Content
```