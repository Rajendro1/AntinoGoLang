
---

# Blogging Platform API

This API allows you to perform CRUD operations on blog posts.

Certainly! Below is an updated README documentation that includes the steps to start the server using Docker Compose:

---

# Blogging Platform API

This API allows you to perform CRUD operations on blog posts.

## Setup and Running the API

### Prerequisites

- Docker and Docker Compose installed on your machine.

### Steps to Start the Server

1. **Create `.env` File:**  
   Create an `.env` file in the root of the project and add the following content:

    ```env
    DB_HOST=db
    DB_USER=root
    DB_PASSWORD=my-secret-pw
    DB_NAME=mydb
    ```

2. **Build and Run with Docker Compose:**  
   Navigate to the project directory and run the following command:

    ```sh
    sudo docker-compose up --build
    ```

   This command will build your Go application and MySQL as Docker containers and start them.

3. **Access the API:**  
   The API will be accessible at:

    ```
    http://localhost:8080
    ```

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