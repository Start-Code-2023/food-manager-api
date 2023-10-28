# Food Manager Application

The Food Manager Application is a Go-based server application for managing food items and users. It provides RESTful API endpoints for adding, retrieving, updating, and deleting food items and user documents in a Firebase Firestore database.

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (1.16 or higher)
- Firebase Firestore account
- Firebase Admin SDK credentials (serviceAccountKey.json)
- Postman or any API testing tool

### Installing

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/food-manager.git

2. Navigate to the project directory:
    ```bash
    cd food-manager
   
3. Install project dependencies:
   ```bash
   go mod tidy

4. Set up Firebase Admin SDK credentials:

Place your Firebase Admin SDK credentials (serviceAccountKey.json) in the project root directory.

5. Configure Firebase Firestore:

Replace the Firebase Firestore configuration values in your code.

6. Run the application
    ```bash
   go run main.go

## API Endpoints

### Create 
* URL: /create
* Method: POST
* Description: Create a new user's food list document.

### Get 
* URL: /get?userID={userID}
* Queries:
  * ?userID={userID}
  * ?userID={userID}&foodID={foodID}
* Method: GET
* Description: Retrieve the food list for a specific user, or retrieve a specific food item for a user.

### Add 
* URL: /add
* Method: POST
* Description: Add food items to a user's food list.

### Delete 
* URL: /delete
* Method: DELETE
* Description: Delete a user's food list document.

### Remove 
* URL: /remove?userID={userID}&foodID={foodID}&quantity={quantity}
* Method: DELETE
* Description: Remove food items from a user's food list.

## Deployment
Deploy the application to a production server, and ensure that your Firebase Admin SDK credentials are properly configured.

## Built With
* Go - The programming language used
* Firebase Firestore - The database used
* Postman - API testing tool