# BWA Crowdfunding REST API

## Description
This project is a copy of the back-end from the **crowdfunding** project available at [https://github.com/fauzan264/crowdfunding/tree/master/backend](https://github.com/fauzan264/crowdfunding/tree/master/backend), adapted and implemented using **Laravel 11.37.0**. The project involves building a secure and scalable REST API with **JWT** for authentication and **PostgreSQL** as the database solution.

## Learning Flow

### API Development with Laravel
The back-end is built using **Laravel 11.37.0**, leveraging its routing, middleware, and Eloquent ORM for efficient database interaction. The goal is to build an efficient API that handles various crowdfunding operations.

### Authentication with JWT
Implement **JWT** (JSON Web Tokens) for secure, token-based user authentication, allowing users to access protected endpoints after successful login.

### PostgreSQL as the Database
We use **PostgreSQL** for data management, offering a reliable and robust database solution for storing crowdfunding-related data.

### Personal Improvisations
This implementation includes custom optimizations and added features, such as secure JWT handling, and improvements to API performance and user experience.

## Technologies Used

- **Laravel 11.37.0**: The back-end framework for building the REST API.
- **JWT (JSON Web Token)**: For secure authentication and authorization of users.
- **PostgreSQL**: Version 13.x for managing and storing the application's data.

## Project Requirements

### Software Requirements

- **PHP (v8.1 or later)**
  - Required for running the Laravel application.

- **PostgreSQL**
  - The project uses PostgreSQL as the database for easy data management.

- **Node.js (Optional)**
  - If you're using Laravel for front-end tools (e.g., Mix for asset compilation), **Node.js** might be needed. However, if you're only building an API with Laravel, you don't need **Node.js** unless you're integrating front-end assets.

## Installation and Setup
### 1. Clone the repository:
```bash
    git clone https://github.com/fauzan264/crowdfunding-rest-api.git
```

### 2. Navigate to the project directory:
```bash
    cd crowdfunding-rest-api
```

### 3. Install project dependencies using Composer:
```bash
    composer install
```

### 4. Set up the .env file by copying .env.example:
```bash
    cp .env.example .env
```

### 5. Generate the application key:
```bash
php artisan key:generate
```

### 6. Set up PostgreSQL database details in .env:
Edit the .env file and update the database connection details:
```plaintext
    DB_CONNECTION=pgsql
    DB_HOST=127.0.0.1
    DB_PORT=5432
    DB_DATABASE=your_database
    DB_USERNAME=your_username
    DB_PASSWORD=your_password
```

### 7. Run the migrations to set up the database:
```bash
    php artisan migrate
```

### 8. Optionally, seed the database with sample data:
```bash
    php artisan db:seed
```

### 9. Start the development server:
```bash
    php artisan serve
```