# E-Vote GO-Project

## Overview
The E-Vote project is an online voting system built with the Gin framework in Go. It provides features for user management, candidate management, division-based elections, and voting with results per division. The project also incorporates CORS middleware for secure cross-origin requests.

## Model Database
![Screenshot 2024-12-08 195550](https://github.com/user-attachments/assets/632b78a5-8028-4332-a79a-c83cd9c170fb)

---

## Features

### User Management
- **Login**: Authenticate users and provide access tokens.
  - Endpoint: `POST /login`
- **Create User**: Add new users to the system.
  - Endpoint: `POST /users`
- **Get Users**: Retrieve a list of all users.
  - Endpoint: `GET /users`
- **Update User**: Modify user details.
  - Endpoint: `PUT /users/:id`
- **Delete User**: Remove a user from the system.
  - Endpoint: `DELETE /users/:id`

### Candidate Management
- **Create Candidate**: Add new candidates to the system.
  - Endpoint: `POST /candidates`
- **Get Candidates**: Retrieve a list of all candidates.
  - Endpoint: `GET /candidates`
- **Update Candidate**: Modify candidate details.
  - Endpoint: `PUT /candidates/:id`
- **Delete Candidate**: Remove a candidate from the system.
  - Endpoint: `DELETE /candidates/:id`

### Division Management
- **Create Division**: Add new voting divisions.
  - Endpoint: `POST /divisions`
- **Get Divisions**: Retrieve a list of all divisions.
  - Endpoint: `GET /divisions`
- **Get Candidates by Division**: Retrieve candidates specific to a division.
  - Endpoint: `GET /divisions/:id/candidates`
- **Update Division**: Modify division details.
  - Endpoint: `PUT /divisions/:id`
- **Delete Division**: Remove a division from the system.
  - Endpoint: `DELETE /divisions/:id`

### Voting System
- **Vote**: Submit a user's vote for a specific division.
  - Endpoint: `POST /vote/div-:divisions.id/:users.id`
- **Get User Votes**: Retrieve votes submitted by a user.
  - Endpoint: `GET /vote/:users.id`
- **Get Division Results**: Retrieve voting results for a specific division.
  - Endpoint: `GET /results/:division_id`

### Cross-Origin Resource Sharing (CORS)
CORS middleware ensures secure interaction between the backend API and the frontend client.
- Allowed Origin: `http://localhost:3000`
- Supported Methods: `GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`
- Supports Credentials: `true`

---

## Getting Started

### Prerequisites
- [Go](https://golang.org/) installed on your system.
- A frontend client (e.g., React) running at `http://localhost:3000`.

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/e-vote.git
   cd e-vote
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the server:
   ```bash
   go run main.go
   ```
4. Access the API at `http://localhost:8080`.

5. Run the front-end
   ```bash
   cd fe-vote
   npm install
   npm start
   ```
6. Access the API at `http://localhost:3000`.

---

## API Routes Summary

### User Routes
| Method | Endpoint        | Description        |
|--------|-----------------|--------------------|
| OPTIONS| `/login`        | Handle preflight requests for login. |
| POST   | `/login`        | Authenticate users. |
| POST   | `/users`        | Create a new user.  |
| GET    | `/users`        | Retrieve all users. |
| PUT    | `/users/:id`    | Update user details. |
| DELETE | `/users/:id`    | Delete a user.      |

### Candidate Routes
| Method | Endpoint        | Description        |
|--------|-----------------|--------------------|
| POST   | `/candidates`   | Add a new candidate. |
| GET    | `/candidates`   | Retrieve all candidates. |
| PUT    | `/candidates/:id` | Update candidate details. |
| DELETE | `/candidates/:id` | Delete a candidate. |

### Division Routes
| Method | Endpoint                        | Description                    |
|--------|---------------------------------|--------------------------------|
| GET    | `/divisions`                   | Retrieve all divisions.        |
| GET    | `/divisions/:id/candidates`    | Retrieve candidates by division. |
| POST   | `/divisions`                   | Add a new division.            |
| PUT    | `/divisions/:id`               | Update division details.       |
| DELETE | `/divisions/:id`               | Delete a division.             |

### Voting Routes
| Method | Endpoint                        | Description                    |
|--------|---------------------------------|--------------------------------|
| GET    | `/vote/:users.id`              | Retrieve votes by user.        |
| POST   | `/vote/div-:divisions.id/:users.id` | Submit a vote.            |
| GET    | `/results/:division_id`        | Retrieve division results.     |

---

## License
This project is licensed under the [MIT License](LICENSE).

## Contributors
- [Chalvin Reza](https://github.com/chalvnrlv)

