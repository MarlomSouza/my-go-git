# GitHub Repos Viewer

This project is a web application that allows users to view their GitHub repositories, organizations, and organization members. It consists of a backend service written in Go and a frontend application built with React.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Environment Variables](#environment-variables)
- [Generating GitHub OAuth Credentials](#generating-github-oauth-credentials)
- [Docker](#docker)
- [Contributing](#contributing)
- [License](#license)

## Features

- View personal and organization repositories
- View organization members
- OAuth2 authentication with GitHub
- Pagination and filtering of repositories
- Responsive design with Tailwind CSS

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.23.4 or later)
- [Node.js](https://nodejs.org/) (version 18 or later)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Installation
### Backend

1. Clone the repository:
   ```sh
   git clone https://github.com/MarlomSouza/go-git.git
   cd go-git/github-backend
   ```

2. Install Go dependencies:
   ```sh
   go mod download
   ```

3. Create a `.env` file in the `github-backend` directory and add the following environment variables:
   ```sh
   GITHUB_CLIENT_ID=your_github_client_id
   GITHUB_CLIENT_SECRET=your_github_client_secret
   GITHUB_REDIRECT_URL=http://localhost:5000/login/github/callback
   FRONTEND_REDIRECT_URL=http://localhost:3000
   PORT=5000
   ```

### Frontend

1. Navigate to the `git-app` directory:
   ```sh
   cd ../git-app
   ```

2. Install Node.js dependencies:
   ```sh
   npm install
   ```

3. Create a `.env` file in the `git-app` directory and add the following environment variable:
   ```sh
   REACT_APP_API_URL=http://localhost:5000
   ```

## Usage

### Running Locally

#### Backend
Start the backend server:
   ```sh
   go run cmd/main.go
   ```

Alternatively, you can use [Air](https://github.com/cosmtrek/air) for live reloading:
   ```sh
   air
   ```

#### Frontend
Start the frontend development server:
   ```sh
   npm start
   ```

2. Open your browser and navigate to [http://localhost:3000](http://localhost:3000).

### Running with Docker
Ensure Docker and Docker Compose are installed.

Build and start the services:
   ```sh
   docker-compose up --build
   ```

## Environment Variables

The following environment variables are used in the project:

### Backend
- `GITHUB_CLIENT_ID`: GitHub OAuth client ID
- `GITHUB_CLIENT_SECRET`: GitHub OAuth client secret
- `GITHUB_REDIRECT_URL`: GitHub OAuth redirect URL
- `FRONTEND_REDIRECT_URL`: URL to redirect to after successful login
- `PORT`: Port for the backend server (default: 5000)

### Frontend
- `REACT_APP_API_URL`: URL of the backend API (default: http://localhost:5000)

## Generating GitHub OAuth Credentials

To generate the GitHub OAuth credentials, follow these steps:

1. Go to [GitHub Developer Settings](https://github.com/settings/developers).
2. Click on "New OAuth App".
3. Fill in the required fields:
   - **Application name**: Your application name
   - **Homepage URL**: `http://localhost:3000`
   - **Authorization callback URL**: `http://localhost:5000/login/github/callback`
4. Click "Register application".
5. Copy the `Client ID` and `Client Secret` and add them to your `.env` file as `GITHUB_CLIENT_ID` and `GITHUB_CLIENT_SECRET`.

## Docker

The project includes a `docker-compose.yml` file to run the backend and frontend services in Docker containers.

### Building and Running
Build and start the services:
   ```sh
   docker-compose up --build
   ```

Open your browser and navigate to [http://localhost:3000](http://localhost:3000).

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.