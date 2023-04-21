# Blog

This is a simple blogging service provide **REST API** for consumer to build their own blogging platform on demand.

<!--- If we have only one grouop/collection, then no need for the "ungrouped" heading -->

## Variables

| Key         | Value            | Type   |
| ----------- | ---------------- | ------ |
| blog-server | www.blog-api.com | string |
| username    | user             | string |
| password    | pass             | string |
| limit       | 20               | string |

## Endpoints

- [Users](#users)
  1. [Create user](#1-create-user)
     - [Invalid username/password](#i-example-request-invalid-usernamepassword)
     - [Empty body](#ii-example-request-empty-body)
     - [Validation error](#iii-example-request-validation-error)
     - [Success](#iv-example-request-success)
  1. [Fetch user](#2-fetch-user)
     - [Success](#i-example-request-success)
     - [Invalid user id](#ii-example-request-invalid-user-id)
     - [Invalid username/password](#iii-example-request-invalid-usernamepassword)
  1. [Fetch users](#3-fetch-users)
     - [Success](#i-example-request-success-1)
     - [Invalid user id](#ii-example-request-invalid-user-id-1)
     - [Invalid username/password](#iii-example-request-invalid-usernamepassword-1)
  1. [Update user](#4-update-user)
     - [Success](#i-example-request-success-2)
     - [Validation error](#ii-example-request-validation-error)
     - [Empty request body](#iii-example-request-empty-request-body)
     - [Invalid username/password](#iv-example-request-invalid-usernamepassword)
     - [Invalid user id](#v-example-request-invalid-user-id)
  1. [Delete user](#5-delete-user)
     - [Success](#i-example-request-success-3)
     - [Invalid user id](#ii-example-request-invalid-user-id-2)
     - [Invalid username/password](#iii-example-request-invalid-usernamepassword-2)
- [Articles](#articles)
  1. [Create article](#1-create-article)
     - [Validation error](#i-example-request-validation-error)
     - [Invalid username/password](#ii-example-request-invalid-usernamepassword)
     - [Success](#iii-example-request-success)
  1. [List articles](#2-list-articles)
     - [Invalid token](#i-example-request-invalid-token)
     - [Success](#ii-example-request-success)
     - [Without filtering](#iii-example-request-without-filtering)
- [Users/V2](#usersv2)
  1. [Create user](#1-create-user-1)
     - [Validation Error](#i-example-request-validation-error-1)
     - [Success](#ii-example-request-success-1)
     - [Invalid username/password](#iii-example-request-invalid-usernamepassword-3)
  1. [Update user](#2-update-user)
     - [Invalid username/password](#i-example-request-invalid-usernamepassword-1)
     - [Success](#ii-example-request-success-2)
- [Ungrouped](#ungrouped)
  1. [Health check](#1-health-check)
     - [Error](#i-example-request-error)
     - [Success](#ii-example-request-success-3)

---

## Users

`Users` directory contains all the user related APIs. For authentication these apis requrie `BasicAuth`

### 1. Create user

Create user use `JSON` payload to create a user

**_Endpoint:_**

```bash
Method: POST
Type: RAW
URL: {{blog-server}}/v1/users
```

**_Body:_**

```js
{
	"name": "Captain Jack Sparrow",
	"bio": "Captain Jack Sparrow is a fictional character and the main protagonist of the Pirates of the Caribbean film series. The character was created by screenwriters Ted Elliott and Terry Rossio and is portrayed by Johnny Depp"
}
```

**_More example Requests/Responses:_**

##### I. Example Request: Invalid username/password

##### I. Example Response: Invalid username/password

```js
{
    "message": "Unauthorized attempt"
}
```

**_Status Code:_** 401

<br>

##### II. Example Request: Empty body

##### II. Example Response: Empty body

```js
{
    "message": "Invalid request body"
}
```

**_Status Code:_** 400

<br>

##### III. Example Request: Validation error

**_Body:_**

```js
{
	"name": "",
	"bio": ""
}
```

##### III. Example Response: Validation error

```js
{
    "errors": {
        "bio": [
            "Bio can not be empty"
        ],
        "name": [
            "Name can not be empty"
        ]
    },
    "message": "Validation error"
}
```

**_Status Code:_** 422

<br>

##### IV. Example Request: Success

**_Body:_**

```js
{
	"name": "Tom Hanks",
	"bio": "Thomas Jeffrey Hanks is an American actor and filmmaker. Known for both his comedic and dramatic roles, Hanks is one of the most popular and recognizable film stars worldwide, and is widely regarded as an American cultural icon."
}
```

##### IV. Example Response: Success

```js
{
    "data": {
        "id": 1,
        "name": "Tom Hanks",
        "bio": "Thomas Jeffrey Hanks is an American actor and filmmaker. Known for both his comedic and dramatic roles, Hanks is one of the most popular and recognizable film stars worldwide, and is widely regarded as an American cultural icon."
    },
    "message": "User created successfully"
}
```

**_Status Code:_** 201

<br>

### 2. Fetch user

Fetch a single user using `id`

**_Endpoint:_**

```bash
Method: GET
Type:
URL: {{blog-server}}/v1/users/{{id}}
```

**_More example Requests/Responses:_**

##### I. Example Request: Success

##### I. Example Response: Success

```js
{
    "data": {
        "id": 1,
        "name": "Tom Hanks",
        "bio": "Thomas Jeffrey Hanks is an American actor and filmmaker. Known for both his comedic and dramatic roles, Hanks is one of the most popular and recognizable film stars worldwide, and is widely regarded as an American cultural icon."
    }
}
```

**_Status Code:_** 200

<br>

##### II. Example Request: Invalid user id

##### II. Example Response: Invalid user id

```js
{
    "message": "Invalid user id"
}
```

**_Status Code:_** 400

<br>

##### III. Example Request: Invalid username/password

##### III. Example Response: Invalid username/password

```js
{
    "message": "Unauthorized attempt"
}
```

**_Status Code:_** 401

<br>

### 3. Fetch users

Fetch list of users using `pagination`

**_Endpoint:_**

```bash
Method: GET
Type:
URL: {{blog-server}}/v1/users
```

**_Query params:_**

| Key   | Value | Description |
| ----- | ----- | ----------- |
| page  | 1     |             |
| limit | 20    |             |

**_More example Requests/Responses:_**

##### I. Example Request: Success

**_Query:_**

| Key   | Value | Description                                                             |
| ----- | ----- | ----------------------------------------------------------------------- |
| page  | 1     | Page numer is a `integer` which represents the page your are requesting |
| limit | 20    | Limit it the maximum number of users listing in the result.             |

##### I. Example Response: Success

```js
{
    "data": [
        {
            "id": 1,
            "name": "Tom Hanks",
            "bio": "Thomas Jeffrey Hanks is an American actor and filmmaker. Known for both his comedic and dramatic roles, Hanks is one of the most popular and recognizable film stars worldwide, and is widely regarded as an American cultural icon."
        },
        {
            "id": 2,
            "name": "Captain Jack Sparrow",
            "bio": "Captain Jack Sparrow is a fictional character and the main protagonist of the Pirates of the Caribbean film series. The character was created by screenwriters Ted Elliott and Terry Rossio and is portrayed by Johnny Depp"
        }
    ]
}
```

**_Status Code:_** 200

<br>

##### II. Example Request: Invalid user id

##### II. Example Response: Invalid user id

```js
{
    "message": "Invalid user id"
}
```

**_Status Code:_** 400

<br>

##### III. Example Request: Invalid username/password

##### III. Example Response: Invalid username/password

```js
{
    "message": "Unauthorized attempt"
}
```

**_Status Code:_** 401

<br>

### 4. Update user

Update user use `JSON` payload to create a user

**_Endpoint:_**

```bash
Method: PUT
Type: RAW
URL: {{blog-server}}/v1/users/{{id}}
```

**_Body:_**

```js
{
	"name": "Captain Jack Sparrow",
	"bio": "Captain Jack Sparrow is a fictional character and the main protagonist of the Pirates of the Caribbean film series. The character was created by screenwriters Ted Elliott and Terry Rossio and is portrayed by Johnny Depp"
}
```

**_More example Requests/Responses:_**

##### I. Example Request: Success

**_Body:_**

```js
{
	"name": "Captain Jack Sparrow",
	"bio": "Captain Jack Sparrow is a fictional character and the main protagonist of the Pirates of the Caribbean film series. The character was created by screenwriters Ted Elliott and Terry Rossio and is portrayed by Johnny Depp"
}
```

##### I. Example Response: Success

```js
{
    "data": {
        "id": 1,
        "name": "Captain Jack Sparrow",
        "bio": "Captain Jack Sparrow is a fictional character and the main protagonist of the Pirates of the Caribbean film series. The character was created by screenwriters Ted Elliott and Terry Rossio and is portrayed by Johnny Depp"
    },
    "message": "User updated successfully"
}
```

**_Status Code:_** 200

<br>

##### II. Example Request: Validation error

**_Body:_**

```js
{
	"name": "",
	"bio": ""
}
```

##### II. Example Response: Validation error

```js
{
    "errors": {
        "bio": [
            "Bio can not be empty"
        ],
        "name": [
            "Name can not be empty"
        ]
    },
    "message": "Validation error"
}
```

**_Status Code:_** 422

<br>

##### III. Example Request: Empty request body

##### III. Example Response: Empty request body

```js
{
    "message": "Invalid request body"
}
```

**_Status Code:_** 400

<br>

##### IV. Example Request: Invalid username/password

##### IV. Example Response: Invalid username/password

```js
{
    "message": "Unauthorized attempt"
}
```

**_Status Code:_** 401

<br>

##### V. Example Request: Invalid user id

**_Body:_**

```js
{
	"name": "Captain Jack Sparrow",
	"bio": "Captain Jack Sparrow is a fictional character and the main protagonist of the Pirates of the Caribbean film series. The character was created by screenwriters Ted Elliott and Terry Rossio and is portrayed by Johnny Depp"
}
```

##### V. Example Response: Invalid user id

```js
{
    "message": "Invalid user id"
}
```

**_Status Code:_** 400

<br>

### 5. Delete user

Delete a single user using `id`

**_Endpoint:_**

```bash
Method: DELETE
Type:
URL: {{blog-server}}/v1/users/{{id}}
```

**_More example Requests/Responses:_**

##### I. Example Request: Success

##### I. Example Response: Success

```js
{
    "message": "User deleted successfully"
}
```

**_Status Code:_** 200

<br>

##### II. Example Request: Invalid user id

##### II. Example Response: Invalid user id

```js
{
    "message": "Invalid user id"
}
```

**_Status Code:_** 400

<br>

##### III. Example Request: Invalid username/password

##### III. Example Response: Invalid username/password

```js
{
    "message": "Unauthorized attempt"
}
```

**_Status Code:_** 401

<br>

## Articles

`Articles` directory contains all the article related APIs. Use `JWT` toekn for authentications for articles API.

### 1. Create article

Create article endpoint is accept form-data to create an article with binary data like file

**_Endpoint:_**

```bash
Method: POST
Type: FORMDATA
URL: {{blog-server}}/v1/articles
```

**_Headers:_**

| Key            | Value               | Description |
| -------------- | ------------------- | ----------- |
| ExampleHeader1 | ExampleHeader1Value |             |
| ExampleHeader2 | ExampleHeader2Value |             |

**_Body:_**

| Key       | Value             | Description                                      |
| --------- | ----------------- | ------------------------------------------------ |
| author_id | 1                 | Accept `author_id` as the primary _id_ of author |
| title     | This is title one | The `title` field must be between _1-255_ chars  |
| body      | This is body one  | The `body` field must be between _1-2000_ chars  |

**_More example Requests/Responses:_**

##### I. Example Request: Validation error

**_Headers:_**

| Key            | Value                             | Description |
| -------------- | --------------------------------- | ----------- |
| ExampleHeader1 | ExampleHeader1Value               |             |
| ExampleHeader2 | ExampleHeader2Value               |             |
| Content-Type   | application/x-www-form-urlencoded |             |

**_Body:_**

| Key       | Value | Description                                      |
| --------- | ----- | ------------------------------------------------ |
| author_id | 1     | Accept `author_id` as the primary _id_ of author |

##### I. Example Response: Validation error

```js
{
    "errors": {
        "title": [
            "Title can not be empty"
        ],
        "body": [
            "Body can not be empty"
        ]
    },
    "message": "Validation error"
}
```

**_Status Code:_** 422

<br>

##### II. Example Request: Invalid username/password

##### II. Example Response: Invalid username/password

```js
{
    "message": "Unauthorized attempt"
}
```

**_Status Code:_** 401

<br>

##### III. Example Request: Success

**_Headers:_**

| Key            | Value               | Description |
| -------------- | ------------------- | ----------- |
| ExampleHeader1 | ExampleHeader1Value |             |
| ExampleHeader2 | ExampleHeader2Value |             |

**_Body:_**

| Key       | Value             | Description                                      |
| --------- | ----------------- | ------------------------------------------------ |
| author_id | 1                 | Accept `author_id` as the primary _id_ of author |
| title     | This is title one | The `title` field must be between _1-255_ chars  |
| body      | This is body one  | The `body` field must be between _1-2000_ chars  |

##### III. Example Response: Success

```js
{
    "data": {
        "id": 1,
        "author_id": 1,
        "title": "This is title one",
        "body": "This is body one"
    },
    "message": "Article created successfully"
}
```

**_Status Code:_** 201

<br>

### 2. List articles

List articles endpoint provide article listing with _filtering_, _patination_

**_Endpoint:_**

```bash
Method: GET
Type:
URL: {{blog-server}}/v1/articles
```

**_Query params:_**

| Key     | Value       | Description                                                  |
| ------- | ----------- | ------------------------------------------------------------ |
| page    | {{id}}      | Page is a `unsigned integer` which represents the page numer |
| limit   | {{limit}}   | Limit represents maximum numer of results in the response    |
| user_id | {{user_id}} | Filter the articles using _author_id_                        |

**_More example Requests/Responses:_**

##### I. Example Request: Invalid token

**_Headers:_**

| Key           | Value               | Description                         |
| ------------- | ------------------- | ----------------------------------- |
| Authorization | Bearer inalid_token | Providing invalid token cause `401` |

**_Query:_**

| Key     | Value       | Description                                                  |
| ------- | ----------- | ------------------------------------------------------------ |
| page    | {{id}}      | Page is a `unsigned integer` which represents the page numer |
| limit   | {{limit}}   | Limit represents maximum numer of results in the response    |
| user_id | {{user_id}} | Filter the articles using _author_id_                        |

##### I. Example Response: Invalid token

```js
{
    "message": "Unauthorized token"
}
```

**_Status Code:_** 401

<br>

##### II. Example Request: Success

**_Headers:_**

| Key           | Value                | Description                             |
| ------------- | -------------------- | --------------------------------------- |
| Authorization | Bearer {{jwt_token}} | Providing valid token won't cause _401_ |

**_Query:_**

| Key     | Value       | Description                                                  |
| ------- | ----------- | ------------------------------------------------------------ |
| page    | {{id}}      | Page is a `unsigned integer` which represents the page numer |
| limit   | {{limit}}   | Limit represents maximum numer of results in the response    |
| user_id | {{user_id}} | Filter the articles using _author_id_                        |

##### II. Example Response: Success

```js
{
    "data": [
        {
            "id": 3,
            "user_id": 10,
            "title": "Article 13",
            "body": "This is article three"
        },
        {
            "id": 2,
            "user_id": 10,
            "title": "Article 2",
            "body": "This is article two"
        },
        {
            "id": 1,
            "user_id": 10,
            "title": "Article 1",
            "body": "This is article one"
        }
    ]
}
```

**_Status Code:_** 200

<br>

##### III. Example Request: Without filtering

**_Headers:_**

| Key           | Value                | Description                             |
| ------------- | -------------------- | --------------------------------------- |
| Authorization | Bearer {{jwt_token}} | Providing valid token won't cause _401_ |

**_Query:_**

| Key   | Value     | Description                                                  |
| ----- | --------- | ------------------------------------------------------------ |
| page  | {{id}}    | Page is a `unsigned integer` which represents the page numer |
| limit | {{limit}} | Limit represents maximum numer of results in the response    |

##### III. Example Response: Without filtering

```js
{
    "data": [
        {
            "id": 3,
            "user_id": 10,
            "title": "Article 13",
            "body": "This is article three"
        },
        {
            "id": 2,
            "user_id": 10,
            "title": "Article 2",
            "body": "This is article two"
        },
        {
            "id": 1,
            "user_id": 10,
            "title": "Article 1",
            "body": "This is article one"
        }
    ]
}
```

**_Status Code:_** 200

<br>

## Users/V2

### 1. Create user

Create user use `JSON` payload to create a user

**_Endpoint:_**

```bash
Method: POST
Type: URLENCODED
URL: {{blog-server}}/v2/users
```

**_Body:_**

| Key  | Value                                                                                                                                                                                                                                | Description |
| ---- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ----------- |
| name | Tom Hanks                                                                                                                                                                                                                            |             |
| bio  | Thomas Jeffrey Hanks is an American actor and filmmaker. Known for both his comedic and dramatic roles, Hanks is one of the most popular and recognizable film stars worldwide, and is widely regarded as an American cultural icon. |             |

**_More example Requests/Responses:_**

##### I. Example Request: Validation Error

**_Body:_**

| Key  | Value | Description |
| ---- | ----- | ----------- |
| name |       |             |
| bio  |       |             |

##### I. Example Response: Validation Error

```js
{
    "errors": {
        "bio": [
            "Bio can not be empty"
        ],
        "name": [
            "Name can not be empty"
        ]
    },
    "message": "Validation error"
}
```

**_Status Code:_** 422

<br>

##### II. Example Request: Success

**_Body:_**

| Key  | Value                                                                                                                                                                                                                                | Description |
| ---- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ----------- |
| name | Tom Hanks                                                                                                                                                                                                                            |             |
| bio  | Thomas Jeffrey Hanks is an American actor and filmmaker. Known for both his comedic and dramatic roles, Hanks is one of the most popular and recognizable film stars worldwide, and is widely regarded as an American cultural icon. |             |

##### II. Example Response: Success

```js
{
    "data": {
        "id": 1,
        "name": "Tom Hanks",
        "bio": "Thomas Jeffrey Hanks is an American actor and filmmaker. Known for both his comedic and dramatic roles, Hanks is one of the most popular and recognizable film stars worldwide, and is widely regarded as an American cultural icon."
    },
    "message": "User created successfully"
}
```

**_Status Code:_** 201

<br>

##### III. Example Request: Invalid username/password

##### III. Example Response: Invalid username/password

```js
{
    "message": "Unauthorized attempt"
}
```

**_Status Code:_** 401

<br>

### 2. Update user

Use **urlencoded** data to update user partially

**_Endpoint:_**

```bash
Method: PATCH
Type: URLENCODED
URL: {{blog-server}}/v2/users/1
```

**_Body:_**

| Key  | Value         | Description |
| ---- | ------------- | ----------- |
| name | Mr. Tom Hanks |             |

**_More example Requests/Responses:_**

##### I. Example Request: Invalid username/password

##### I. Example Response: Invalid username/password

```js
{
    "message": "Unauthorized attempt"
}
```

**_Status Code:_** 401

<br>

##### II. Example Request: Success

**_Body:_**

| Key  | Value         | Description |
| ---- | ------------- | ----------- |
| name | Mr. Tom Hanks |             |

##### II. Example Response: Success

```js
{
    "data": {
        "id": 1,
        "name": "Mr. Tom Hanks",
        "bio": "Thomas Jeffrey Hanks is an American actor and filmmaker. Known for both his comedic and dramatic roles, Hanks is one of the most popular and recognizable film stars worldwide, and is widely regarded as an American cultural icon."
    },
    "message": "User partial update successful"
}
```

**_Status Code:_** 206

<br>

## Ungrouped

### 1. Health check

System health check endpoint provide system health status for `probe`

**_Endpoint:_**

```bash
Method: GET
Type:
URL: {{blog-server}}/
```

**_More example Requests/Responses:_**

##### I. Example Request: Error

##### I. Example Response: Error

```js
{
    "system": "Failed"
}
```

**_Status Code:_** 500

<br>

##### II. Example Request: Success

##### II. Example Response: Success

```js
{
    "system": "OK"
}
```

**_Status Code:_** 200

<br>

---

[Back to top](#blog)

> Generated at 2022-01-25 15:47:32 by [docgen](https://github.com/cksidharthan/docgen)
