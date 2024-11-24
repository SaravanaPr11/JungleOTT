OTT App Project
A lightweight and scalable Over-The-Top (OTT) platform designed to provide seamless video streaming services. The platform allows users to sign up, log in, browse available content, and stream videos. Admins can upload and manage the content effectively.

Tech Stack
Backend: Golang (Gin Framework, GORM)
Frontend: React.js or Next.js
Database: PostgreSQL
Authentication: JSON Web Tokens (JWT)
Storage: AWS S3 (for video uploads)
API Management: RESTful APIs
DevOps: Docker (optional), CI/CD pipelines (future goal)
Initial Goals
Develop an MVP with core functionality for users and admins.
Provide a responsive and intuitive user experience.
Ensure scalability and performance using modern backend technologies.
Enable secure and efficient content management.
Core Features for MVP
Users
Sign-up/Login: Users can register and log in using email and password.
Browse Videos: Explore a catalog of videos by categories, genres, or tags.
Watch a Video: Stream videos with adaptive bitrate for smooth playback.
Admins
Upload Videos: Upload and store video content securely.
Manage Content: Add, update, or remove video metadata like title, description, and tags.
Project Setup
Backend Setup
Prerequisites:

Install Go (latest stable version).
Install PostgreSQL.
Install Postman (optional, for API testing).
Steps:

Clone the repository:
bash
Copy code
git clone <repository_url>
cd ott-project/backend
Install dependencies:
bash
Copy code
go mod tidy
Set up the .env file:
makefile
Copy code
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=ott_database
JWT_SECRET=your_secret_key
Run the backend:
bash
Copy code
go run main.go
Test the API:
Use Postman or curl to test endpoints like http://localhost:8080/api/v1.