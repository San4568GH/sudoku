# 🎮 Sudoku Game

A comprehensive Sudoku game with competitive play modes, educational features, and leaderboards. Built with Go backend and React frontend.

## 🚀 Features

### 🎯 Game Modes
- **Play Mode (Competitive)**: Timed gameplay with scoring and leaderboards
- **Learn Mode (Educational)**: Hints, auto-solver, and step-by-step guidance

### 🏆 Competitive Features
- Real-time timer that runs continuously
- Scoring system (+10 points per correct number)
- Auto-solve disqualifies from leaderboards
- No hints allowed in competitive mode
- Leaderboards for fastest time and highest scores

### 📚 Educational Features
- Hint system with human-style reasoning
- Auto-solver with step-by-step explanations
- No timer pressure
- Perfect for learning Sudoku techniques

### 👤 User System
- User registration and authentication
- JWT-based security
- User profiles with statistics
- Game history tracking

## 🛠️ Tech Stack

### Backend
- **Language**: Go 1.21+
- **Framework**: Chi router
- **Database**: PostgreSQL with GORM
- **Authentication**: JWT with bcrypt password hashing
- **CORS**: Enabled for frontend communication

### Frontend
- **Framework**: React 18
- **Routing**: React Router DOM
- **HTTP Client**: Axios
- **Styling**: CSS with modern glassmorphism design
- **State Management**: React Context API

## 📋 Prerequisites

- Go 1.21 or higher
- Node.js 16 or higher
- PostgreSQL 12 or higher
- Git

## 🚀 Quick Start

### 1. Clone the Repository
```bash
git clone <repository-url>
cd sudoku
```

### 2. Set Up Database
```bash
# Create PostgreSQL database
createdb sudoku

# Or using psql
psql -c "CREATE DATABASE sudoku;"
```

### 3. Configure Environment
```bash
# Copy environment example
cp env.example .env

# Edit .env file with your database credentials
DATABASE_URL=host=localhost user=postgres password=your_password dbname=sudoku port=5432 sslmode=disable
JWT_SECRET=your-super-secret-jwt-key-change-in-production
PORT=8080
```

### 4. Set Up Backend
```bash
# Install Go dependencies
go mod tidy

# Run database migrations and seed data
go run cmd/seed/main.go

# Start the backend server
go run main.go
```

The backend will be available at `http://localhost:8080`

### 5. Set Up Frontend
```bash
# Navigate to frontend directory
cd frontend

# Install dependencies
npm install

# Start the development server
npm start
```

The frontend will be available at `http://localhost:3000`

## 📁 Project Structure

```
sudoku/
├── main.go                 # Main server entry point
├── go.mod                  # Go module file
├── env.example             # Environment variables template
├── cmd/
│   └── seed/
│       └── main.go         # Database seeding script
├── internal/
│   ├── auth/               # Authentication service
│   │   ├── service.go      # Auth business logic
│   │   └── middleware.go   # JWT middleware
│   ├── handlers/           # HTTP handlers
│   │   ├── auth.go         # Auth endpoints
│   │   ├── game.go         # Game endpoints
│   │   └── puzzle.go       # Puzzle endpoints
│   ├── models/             # Database models
│   │   ├── user.go         # User model
│   │   ├── puzzle.go       # Puzzle model
│   │   └── game_result.go  # Game result model
│   └── sudoku/             # Sudoku game logic
│       └── service.go      # Game algorithms and validation
└── frontend/               # React frontend
    ├── package.json
    ├── public/
    ├── src/
    │   ├── components/     # React components
    │   ├── context/        # React context
    │   ├── App.js          # Main app component
    │   └── index.js        # Entry point
    └── README.md
```

## 🎮 API Endpoints

### Authentication
- `POST /auth/register` - User registration
- `POST /auth/login` - User login
- `GET /profile` - Get user profile (protected)
- `PUT /profile` - Update user profile (protected)

### Game Management
- `POST /game/start` - Start new game (protected)
- `POST /game/submit` - Submit completed game (protected)
- `POST /game/hint` - Get hint for cell (protected)
- `POST /game/solve` - Auto-solve puzzle (protected)
- `GET /game/history` - Get user game history (protected)

### Puzzles & Leaderboards
- `GET /puzzles` - Get available puzzles
- `GET /leaderboard` - Get leaderboard rankings

## 🎯 Game Rules

### Play Mode (Competitive)
- Timer starts when game begins and runs continuously
- No hints or auto-solve allowed
- +10 points for each correct number placed
- Wrong numbers give no points and no feedback
- Auto-solve disqualifies from leaderboards
- Only one leaderboard entry per puzzle per user

### Learn Mode (Educational)
- No timer pressure
- Hints available with reasoning
- Auto-solver with step-by-step explanations
- No scoring or leaderboard impact
- Perfect for learning techniques

## 🔧 Development

### Running Tests
```bash
# Backend tests
go test ./...

# Frontend tests
cd frontend
npm test
```

### Database Migrations
The application uses GORM auto-migration. Tables are created automatically when the server starts.

### Adding New Puzzles
Use the seeding script to add new puzzles:
```bash
go run cmd/seed/main.go
```

## 🚀 Deployment

### Backend Deployment
1. Build the Go binary:
   ```bash
   go build -o sudoku main.go
   ```

2. Set up environment variables in production
3. Run the binary with proper database configuration

### Frontend Deployment
1. Build the React app:
   ```bash
   cd frontend
   npm run build
   ```

2. Serve the `build` folder with a web server

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📝 License

This project is licensed under the MIT License.

## 🆘 Support

If you encounter any issues or have questions, please open an issue on the repository.

---

**Happy Sudoku Solving! 🎮** 