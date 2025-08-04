@echo off
echo 🎮 Setting up Sudoku Game Project...

REM Check if Go is installed
go version >nul 2>&1
if errorlevel 1 (
    echo ❌ Go is not installed. Please install Go 1.21 or higher.
    pause
    exit /b 1
)

REM Check if Node.js is installed
node --version >nul 2>&1
if errorlevel 1 (
    echo ❌ Node.js is not installed. Please install Node.js 16 or higher.
    pause
    exit /b 1
)

@REM REM Check if PostgreSQL is installed
@REM psql --version >nul 2>&1
@REM if errorlevel 1 (
@REM     echo ❌ PostgreSQL is not installed. Please install PostgreSQL 12 or higher.
@REM     pause
@REM     exit /b 1
@REM )

echo ✅ Prerequisites check passed!

REM Install Go dependencies
echo 📦 Installing Go dependencies...
go mod tidy

REM Create .env file if it doesn't exist
if not exist .env (
    echo 📝 Creating .env file...
    copy env.example .env
    echo ⚠️  Please edit .env file with your database credentials before continuing.
)

REM Install frontend dependencies
echo 📦 Installing frontend dependencies...
cd frontend
call npm install
cd ..

echo 🎉 Setup completed!
echo.
echo Next steps:
echo 1. Edit .env file with your database credentials
echo 2. Create PostgreSQL database: createdb sudoku
echo 3. Run database seeding: go run cmd/seed/main.go
echo 4. Start backend: go run main.go
echo 5. Start frontend: cd frontend ^&^& npm start
echo.
echo Happy coding! 🚀
pause 