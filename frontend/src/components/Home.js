import React, { useContext } from 'react';
import { Link } from 'react-router-dom';
import { AuthContext } from '../context/AuthContext';
import './Home.css';

function Home() {
  const { user } = useContext(AuthContext);

  return (
    <div className="home">
      <div className="hero">
        <h1>🎮 Welcome to Sudoku Game</h1>
        <p>Challenge yourself with puzzles of varying difficulty levels</p>
        <p>Compete with others on the leaderboard or learn at your own pace</p>
      </div>

      <div className="features">
        <div className="feature-card">
          <h3>🎯 Play Mode</h3>
          <p>Competitive gameplay with scoring and leaderboards</p>
          <ul>
            <li>Timer runs from start to finish</li>
            <li>No hints allowed</li>
            <li>Auto-solve disqualifies from scoring</li>
            <li>+10 points per correct number</li>
          </ul>
          {user ? (
            <Link to="/game" className="btn btn-primary">Start Playing</Link>
          ) : (
            <Link to="/login" className="btn btn-primary">Login to Play</Link>
          )}
        </div>

        <div className="feature-card">
          <h3>📚 Learn Mode</h3>
          <p>Educational mode for learning Sudoku techniques</p>
          <ul>
            <li>Use hints and auto-solver</li>
            <li>Step through solutions</li>
            <li>See human-style reasoning</li>
            <li>No timer pressure</li>
          </ul>
          {user ? (
            <Link to="/game" className="btn btn-secondary">Start Learning</Link>
          ) : (
            <Link to="/register" className="btn btn-secondary">Register to Learn</Link>
          )}
        </div>

        <div className="feature-card">
          <h3>🏆 Leaderboards</h3>
          <p>Compete with players worldwide</p>
          <ul>
            <li>Fastest time rankings</li>
            <li>Highest score rankings</li>
            <li>Filter by difficulty</li>
            <li>Real-time updates</li>
          </ul>
          <Link to="/leaderboard" className="btn btn-success">View Leaderboard</Link>
        </div>
      </div>

      {!user && (
        <div className="cta">
          <h2>Ready to start?</h2>
          <p>Join thousands of players and improve your Sudoku skills</p>
          <div className="cta-buttons">
            <Link to="/register" className="btn btn-primary">Create Account</Link>
            <Link to="/login" className="btn btn-secondary">Sign In</Link>
          </div>
        </div>
      )}
    </div>
  );
}

export default Home; 