# 🎮 SudokuPro - Play and Learn Sudoku

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

![Home](demoimages/home.png)

![Modes](demoimages/mode.png)

![Step solve](demoimages/hint.png)

![Full Solve](demoimages/solve.png)

![Profile](demoimages/profile.png)

![Leaderboard](demoimages/leaderboard.png)


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
- **Styling**: CSS with modern glassmorphism design, Mantine UI for components and Stylings
- **State Management**: React Context API

## 📋 Prerequisites

- Go 1.21 or higher
- Node.js 16 or higher
- PostgreSQL 12 or higher
- Git

## 🚀 Quick Start

### Option 1: Automated Setup (Recommended)

#### For Windows:
```bash
# Clone the repository
git clone https://github.com/sarthakmehndiratta/sudoku.git
cd sudoku

# Run the automated setup script
setup.bat
```

#### For Linux/macOS:
```bash
# Clone the repository
git clone https://github.com/sarthakmehndiratta/sudoku.git
cd sudoku

# Make setup script executable and run it
chmod +x setup.sh
./setup.sh
```

### Option 2: Manual Setup

#### 1. Clone the Repository
```bash
git clone https://github.com/sarthakmehndiratta/sudoku.git
cd sudoku
```

#### 2. Install Prerequisites

**Go (Backend):**
- Download from [golang.org](https://golang.org/dl/)
- Verify: `go version`

**Node.js (Frontend):**
- Download from [nodejs.org](https://nodejs.org/)
- Verify: `node --version` and `npm --version`

**PostgreSQL:**
- **Windows**: Download from [postgresql.org](https://www.postgresql.org/download/windows/)
- **macOS**: `brew install postgresql`
- **Ubuntu/Debian**: `sudo apt install postgresql postgresql-contrib`
- **CentOS/RHEL**: `sudo yum install postgresql postgresql-server`

#### 3. Set Up PostgreSQL Database

**Start PostgreSQL service:**
```bash
# Windows (if installed as service, it should start automatically)
# macOS
brew services start postgresql

# Linux
sudo systemctl start postgresql
sudo systemctl enable postgresql
```

**Create database and user:**
```bash
# Connect to PostgreSQL as superuser
psql -U postgres

# Create database
CREATE DATABASE sudoku;

# Create a user (optional but recommended)
CREATE USER sudoku_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE sudoku TO sudoku_user;

# Exit PostgreSQL
\q
```

#### 4. Configure Environment
```bash
# Copy environment example
cp env.example .env

# Edit .env file with your database credentials
# Use your favorite text editor or:
# Windows:
notepad .env

# macOS/Linux:
nano .env
```

**Example .env content:**
```env
DATABASE_URL=host=localhost user=postgres password=your_password dbname=sudoku port=5432 sslmode=disable
JWT_SECRET=your-super-secret-jwt-key-change-in-production
PORT=8080
```

#### 5. Set Up Backend
```bash
# Install Go dependencies
go mod tidy

# Run database migrations and seed data
go run cmd/seed/main.go

# Start the backend server
go run main.go
```

The backend will be available at `http://localhost:8080`

#### 6. Set Up Frontend
```bash
# Navigate to frontend directory
cd frontend

# Install dependencies
npm install

# Start the development server
npm start
```

The frontend will be available at `http://localhost:3000`

## 🔧 Troubleshooting

### Common Issues

#### 1. Database Connection Issues
**Error**: `failed to connect to database: connection refused`
**Solution**: 
- Ensure PostgreSQL is running
- Check your `.env` file has correct credentials
- Verify database exists: `psql -U postgres -d sudoku`

#### 2. Port Already in Use
**Error**: `listen tcp :8080: bind: address already in use`
**Solution**:
```bash
# Windows
netstat -ano | findstr :8080
taskkill /PID <PID> /F

# macOS/Linux
lsof -i :8080
kill -9 <PID>
```

#### 3. Frontend Build Issues
**Error**: `npm install` fails
**Solution**:
```bash
# Clear npm cache
npm cache clean --force

# Delete node_modules and reinstall
rm -rf node_modules package-lock.json
npm install
```

#### 4. Go Module Issues
**Error**: `go: module not found`
**Solution**:
```bash
# Clean module cache
go clean -modcache

# Reinstall dependencies
go mod tidy
```

#### 5. CORS Issues
**Error**: Frontend can't connect to backend
**Solution**: Ensure backend is running on port 8080 and frontend on port 3000

### Environment Variables

Make sure your `.env` file contains:
```env
DATABASE_URL=host=localhost user=postgres password=your_actual_password dbname=sudoku port=5432 sslmode=disable
JWT_SECRET=your-super-secret-jwt-key-change-in-production
PORT=8080
```

### Database Setup Verification

Test your database connection:
```bash
# Test connection
psql -U postgres -d sudoku -c "SELECT version();"

# Check if tables exist
psql -U postgres -d sudoku -c "\dt"
```

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

## 🚀 Running the Application

### Start Backend
```bash
# In the project root directory
go run main.go
```
Backend will start on `http://localhost:8080`

### Start Frontend
```bash
# In a new terminal, navigate to frontend directory
cd frontend
npm start
```
Frontend will start on `http://localhost:3000`

### Testing the Application

1. **Open your browser** and go to `http://localhost:3000`
2. **Register a new account** or login with existing credentials
3. **Start a new game**:
   - Choose between Play Mode (competitive) or Learn Mode (educational)
   - Select difficulty: Easy, Medium, or Hard
   - Click "Start Game"
4. **Play the game**:
   - Click on empty cells to select them
   - Use number keys (1-9) or click number buttons
   - Use hints and auto-solve in Learn Mode
   - Submit your solution when done

### API Testing

You can test the API endpoints directly:

```bash
# Register a new user
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"password123"}'

# Login
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123"}'

# Get puzzles (public endpoint)
curl http://localhost:8080/puzzles
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
