# hot_news API

This API is a platform for managing news using Golang and MySQL.

## Key Features

- **User Management:** Authentication features such as sign up, sign in, sign out, get current user, and update user
- **Article Management:** CRUD operations for articles, including creation, reading, updating, and deletion.
- **Category Management:** Endpoints for managing article categories.
- **Comment Management:** Facilities to add, retrieve, update, and delete comments on articles.
- **Like Management:** Facilities to add and delete likes on articles.

## Technologies Used
- **Golang Libraries:** Utilizes `httprouter` for routing HTTP requests efficiently, `gorm` as the ORM library for interacting with MySQL databases, `viper` for configuration management, and `jwt-go` for JWT (JSON Web Token) authentication.
- **MySQL Database:** Uses MySQL to store users, articles, categories, comments and likes.
- **JSON Data Format:** API communicates data using JSON format for seamless integration and exchange.

## Installation and Usage

### Requirements

- Go (Golang) 1.22 or newer.
- MySQL Server.
- Postman or similar software to test the API.

### Installation Steps

1. **Clone the Repository**

   ```bash
   git clone https://github.com/your-username/hot_news.git
   cd hot_news
   ```

2. **Database Configuration**

    - Create a MySQL database with the name hot_news.
    - Change the database connection configuration in the config/config.go or config/database.go files according to your MySQL settings. Make sure to adjust your Username and Password if necessary.

3. **Database Migration**

    - Database migration tool installation (if not already installed):
        ```bash
        go get -tags 'mysql' -u github.com/golang-migrate/migrate/cmd/migrate@latest
        ```

    - Run migration to create an initial schema:
        ```bash 
        migrate -path=db/migration -database="mysql://username:password@tcp(localhost:3306)/hot_news" up
        ```
        Change the username and password according to your MySQL credentials.   

4. **Running Applications**

    Once the migration is complete, run the application with the command: 

    ```bash
    go run .
    ```

    The application will run at http://localhost:3080 by default.

## API Documentation

For more information about using the API, take a look at the open api file hotNews.json.

---
Thank you for visiting the hot_news API project! Don't hesitate to contact us if you have any questions or feedback.
