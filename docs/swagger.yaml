basePath: /api
definitions:
  controllers.LoginPayload:
    properties:
      email:
        type: string
      password:
        type: string
    required:
    - email
    - password
    type: object
  controllers.UserDetails:
    properties:
      email:
        type: string
      name:
        type: string
      password:
        type: string
    required:
    - email
    - name
    - password
    type: object
host: localhost:8080
info:
  contact:
    email: support@swagger.io
    name: API Support
    url: http://demo.com/support
  description: Create  Go REST API with JWT Authentication in Gin Framework
  license:
    name: MIT
    url: https://opensource.org/licenses/MIT
  termsOfService: demo.com
  title: Swagger JWT API
  version: "1.0"
paths:
  /protected/profile:
    get:
      operationId: GetUserByToken
      parameters:
      - description: Authorization header using the Bearer scheme
        in: header
        name: Authorization
        required: true
        type: string
      produces:
      - application/json
      responses:
        "200":
          description: Success
          schema:
            type: string
        "400":
          description: Error
          schema:
            type: string
      summary: Get User By Token
      tags:
      - User
  /public/login:
    post:
      consumes:
      - application/json
      description: Login
      operationId: LoginUser
      parameters:
      - description: Login
        in: body
        name: EnterDetails
        required: true
        schema:
          $ref: '#/definitions/controllers.LoginPayload'
      responses:
        "200":
          description: Success
          schema:
            type: string
        "400":
          description: Error
          schema:
            type: string
      summary: Login User
      tags:
      - User
  /public/signup:
    post:
      consumes:
      - application/json
      description: Signin
      operationId: SignupUser
      parameters:
      - description: Signin
        in: body
        name: EnterDetails
        required: true
        schema:
          $ref: '#/definitions/controllers.UserDetails'
      responses:
        "200":
          description: Success
          schema:
            type: string
        "400":
          description: Error
          schema:
            type: string
      summary: Signup User
      tags:
      - User
schemes:
- http
- https
securityDefinitions:
  ApiKeyAuth:
    in: header
    name: Authorization
    type: apiKey
  BasicAuth:
    type: basic
swagger: "2.0"
