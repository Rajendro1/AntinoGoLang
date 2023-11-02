
---

# Blogging Platform API

This API allows you to perform CRUD operations on blog posts.

## Base URL

The base URL for the API is:

```
http://<your-server-address>:<port>
```

Replace `http://localhost:8080` and `8080` with your server's IP address and the port you are running your Gin application on.

## API Endpoints

### Get All Posts

- **URL**: `/posts`
- **Method**: `GET`
- **Response**: A list of all blog posts.
- **Status Codes**:
    - `200 OK`: Successfully retrieved posts.
    - `404 Not Found`: No posts available.

### Get a Specific Post

- **URL**: `/post`
- **Method**: `GET`
- **Query Parameters**:
    - `id`: ID of the post to retrieve.
- **Response**: The requested blog post.
- **Status Codes**:
    - `200 OK`: Successfully retrieved the post.
    - `400 Bad Request`: Invalid input.
    - `404 Not Found`: Post not found.

### Create a New Post

- **URL**: `/post`
- **Method**: `POST`
- **Request Body** (JSON):
    - `title`: Title of the post.
    - `content`: Content of the post.
    - `author`: Author of the post.
- **Response**: The created blog post.
- **Status Codes**:
    - `201 Created`: Successfully created the post.
    - `400 Bad Request`: Invalid input.
    - `500 Internal Server Error`: Failed to create the post.

### Update an Existing Post

- **URL**: `/post`
- **Method**: `PUT`
- **Request Body** (JSON):
    - `id`: ID of the post to update.
    - `title`: New title of the post.
    - `content`: New content of the post.
    - `author`: New author of the post.
- **Response**: The updated blog post.
- **Status Codes**:
    - `200 OK`: Successfully updated the post.
    - `400 Bad Request`: Invalid input.
    - `404 Not Found`: Post not found.
    - `500 Internal Server Error`: Failed to update the post.

### Delete a Post

- **URL**: `/post`
- **Method**: `DELETE`
- **Query Parameters**:
    - `id`: ID of the post to delete.
- **Response**: A message indicating the result of the operation.
- **Status Codes**:
    - `200 OK`: Successfully deleted the post.
    - `400 Bad Request`: Invalid input.
    - `404 Not Found`: Post not found.
    - `500 Internal Server Error`: Failed to delete the post.

---

This README provides a basic overview of your API endpoints. Be sure to update the URLs and status codes as needed based on your implementation.