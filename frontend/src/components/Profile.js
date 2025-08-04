import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './Profile.css';

function Profile() {
  const [profile, setProfile] = useState(null);
  const [gameHistory, setGameHistory] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    fetchProfile();
    fetchGameHistory();
  }, []);

  const fetchProfile = async () => {
    try {
      const response = await axios.get('/profile');
      setProfile(response.data);
    } catch (err) {
      setError('Failed to fetch profile');
    }
  };

  const fetchGameHistory = async () => {
    try {
      const response = await axios.get('/game/history');
      setGameHistory(response.data);
    } catch (err) {
      setError('Failed to fetch game history');
    } finally {
      setLoading(false);
    }
  };

  const formatTime = (seconds) => {
    const mins = Math.floor(seconds / 60);
    const secs = seconds % 60;
    return `${mins}:${secs.toString().padStart(2, '0')}`;
  };

  const formatDate = (dateString) => {
    return new Date(dateString).toLocaleDateString();
  };

  if (loading) {
    return <div className="loading">Loading profile...</div>;
  }

  return (
    <div className="profile-container">
      <h2>👤 User Profile</h2>
      
      {error && <div className="error-message">{error}</div>}

      {profile && (
        <div className="profile-stats">
          <div className="stat-card">
            <div className="stat-number">{profile.total_points}</div>
            <div className="stat-label">Total Points</div>
          </div>
          
          <div className="stat-card">
            <div className="stat-number">{profile.games_played}</div>
            <div className="stat-label">Games Played</div>
          </div>
          
          <div className="stat-card">
            <div className="stat-number">
              {profile.games_played > 0 ? Math.round(profile.total_points / profile.games_played) : 0}
            </div>
            <div className="stat-label">Average Score</div>
          </div>
          
          <div className="stat-card">
            <div className="stat-number">{profile.username}</div>
            <div className="stat-label">Username</div>
          </div>
        </div>
      )}

      <div className="game-history">
        <h3>📊 Recent Games</h3>
        
        {gameHistory.length === 0 ? (
          <div className="no-data">
            <p>No games played yet. Start your first game!</p>
          </div>
        ) : (
          <div className="history-table-container">
            <table className="history-table">
              <thead>
                <tr>
                  <th>Date</th>
                  <th>Mode</th>
                  <th>Difficulty</th>
                  <th>Score</th>
                  <th>Time</th>
                  <th>Status</th>
                </tr>
              </thead>
              <tbody>
                {gameHistory.map((game) => (
                  <tr key={game.id}>
                    <td>{formatDate(game.created_at)}</td>
                    <td>
                      <span className={`mode-badge ${game.mode}`}>
                        {game.mode === 'play' ? '🎯 Play' : '📚 Learn'}
                      </span>
                    </td>
                    <td>
                      <span className={`difficulty-badge ${game.puzzle?.difficulty}`}>
                        {game.puzzle?.difficulty?.charAt(0).toUpperCase() + game.puzzle?.difficulty?.slice(1)}
                      </span>
                    </td>
                    <td className="score">{game.score}</td>
                    <td className="time">{formatTime(game.time_seconds)}</td>
                    <td>
                      <span className={`status-badge ${game.completed ? 'completed' : 'incomplete'}`}>
                        {game.completed ? '✅ Completed' : '⏳ Incomplete'}
                      </span>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </div>
  );
}

export default Profile; 