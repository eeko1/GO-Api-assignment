# Article API made with GO/Golang

This api has basic CRUD operations. Below you can see screenshots from each type of request (GET, POST, PUT & DELETE).

## localhost:5000/api/articles

GET Request

<img src="screenshots/get-request-ss.png" alt="get" width="500"/>

POST Request

<img src="screenshots/post-request-ss.png" alt="post" width="500"/>

PUT/Update Request

<img src="screenshots/put-request-ss.png" alt="put" width="500"/>

DELETE Request

<img src="screenshots/delete-request-ss.png" alt="delete" width="500"/>

## Documentation

The API documentation was done with SwaggerEditor.

```yaml
# openapi3_0.yaml

openapi: 3.0.0
info:
  title: Articles API
  description: A simple API to manage articles.
  version: 1.0.0
servers:
  - url: http://localhost:5000
    description: Local server
paths:
  /api/articles:
    get:
      summary: Get all articles
      description: Retrieve a list of all articles in the database.
      tags:
        - Articles
      responses:
        "200":
          description: A list of articles.
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/Article"
    post:
      summary: Create a new article
      description: Create a new article and store it in the database.
      tags:
        - Articles
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/ArticleInput"
      responses:
        "201":
          description: Article added successfully.
          content:
            application/json:
              schema:
                type: object
                properties:
                  message:
                    type: string
                    example: Article added successfully
        "400":
          description: Invalid input or missing fields.
          content:
            application/json:
              schema:
                type: object
                properties:
                  error:
                    type: string
                    example: Title cannot be empty
  /api/articles/{id}:
    put:
      summary: Update an article
      description: Update an existing article by ID.
      tags:
        - Articles
      parameters:
        - in: path
          name: id
          required: true
          description: The article ID to update.
          schema:
            type: string
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/ArticleInput"
      responses:
        "200":
          description: Article updated successfully.
          content:
            application/json:
              schema:
                type: object
                properties:
                  message:
                    type: string
                    example: Article updated successfully
                  modified_count:
                    type: integer
                    example: 1
        "400":
          description: Invalid article ID or bad request.
          content:
            application/json:
              schema:
                type: object
                properties:
                  error:
                    type: string
                    example: Invalid article ID
        "404":
          description: Article not found.
          content:
            application/json:
              schema:
                type: object
                properties:
                  error:
                    type: string
                    example: Article not found
    delete:
      summary: Delete an article
      description: Delete an article by ID.
      tags:
        - Articles
      parameters:
        - in: path
          name: id
          required: true
          description: The article ID to delete.
          schema:
            type: string
      responses:
        "200":
          description: Article deleted successfully.
          content:
            application/json:
              schema:
                type: object
                properties:
                  message:
                    type: string
                    example: Article deleted successfully
        "400":
          description: Invalid article ID.
          content:
            application/json:
              schema:
                type: object
                properties:
                  error:
                    type: string
                    example: Invalid article ID

components:
  schemas:
    Article:
      type: object
      properties:
        _id:
          type: string
          example: 64d5f8f9e1b2c3d4e5f6a7b8
        title:
          type: string
          example: Sample Article
        author:
          type: string
          example: John Doe
        content:
          type: string
          example: This is the content of the article.
        publish_date:
          type: string
          format: date-time
          example: 2024-10-01T10:00:00Z
    ArticleInput:
      type: object
      properties:
        title:
          type: string
          example: Does Utah HC win Stanley Cup?
        author:
          type: string
          example: Ice hockey fan
        content:
          type: string
          example: I think it's possible that it happens.
        publish_date:
          type: string
          format: date-time
          example: 2024-10-01T10:00:00Z
```
