import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './Leaderboard.css';

function Leaderboard() {
  const [leaderboard, setLeaderboard] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [filters, setFilters] = useState({
    difficulty: '',
    type: 'score'
  });

  useEffect(() => {
    fetchLeaderboard();
  }, [filters]);

  const fetchLeaderboard = async () => {
    setLoading(true);
    setError('');
    
    try {
      const params = new URLSearchParams();
      if (filters.difficulty) params.append('difficulty', filters.difficulty);
      if (filters.type) params.append('type', filters.type);
      
      const response = await axios.get(`/leaderboard?${params}`);
      setLeaderboard(response.data);
    } catch (err) {
      setError('Failed to fetch leaderboard');
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
    return <div className="loading">Loading leaderboard...</div>;
  }

  return (
    <div className="leaderboard-container">
      <h2>🏆 Leaderboard</h2>
      
      <div className="leaderboard-filters">
        <div className="filter-group">
          <label>Difficulty:</label>
          <select 
            value={filters.difficulty} 
            onChange={(e) => setFilters(prev => ({ ...prev, difficulty: e.target.value }))}
          >
            <option value="">All Difficulties</option>
            <option value="easy">Easy</option>
            <option value="medium">Medium</option>
            <option value="hard">Hard</option>
          </select>
        </div>
        
        <div className="filter-group">
          <label>Sort by:</label>
          <select 
            value={filters.type} 
            onChange={(e) => setFilters(prev => ({ ...prev, type: e.target.value }))}
          >
            <option value="score">Highest Score</option>
            <option value="time">Fastest Time</option>
          </select>
        </div>
      </div>

      {error && <div className="error-message">{error}</div>}

      {leaderboard.length === 0 ? (
        <div className="no-data">
          <p>No leaderboard data available</p>
        </div>
      ) : (
        <div className="leaderboard-table-container">
          <table className="leaderboard-table">
            <thead>
              <tr>
                <th>Rank</th>
                <th>Player</th>
                <th>Difficulty</th>
                {filters.type === 'score' ? (
                  <>
                    <th>Score</th>
                    <th>Time</th>
                  </>
                ) : (
                  <>
                    <th>Time</th>
                    <th>Score</th>
                  </>
                )}
                <th>Date</th>
              </tr>
            </thead>
            <tbody>
              {leaderboard.map((entry, index) => (
                <tr key={index} className={index < 3 ? `rank-${index + 1}` : ''}>
                  <td className="rank">#{index + 1}</td>
                  <td className="player">{entry.username}</td>
                  <td className="difficulty">
                    <span className={`difficulty-badge ${entry.difficulty}`}>
                      {entry.difficulty.charAt(0).toUpperCase() + entry.difficulty.slice(1)}
                    </span>
                  </td>
                  {filters.type === 'score' ? (
                    <>
                      <td className="score">{entry.score}</td>
                      <td className="time">{formatTime(entry.time_seconds)}</td>
                    </>
                  ) : (
                    <>
                      <td className="time">{formatTime(entry.time_seconds)}</td>
                      <td className="score">{entry.score}</td>
                    </>
                  )}
                  <td className="date">{formatDate(entry.completed_at)}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}

export default Leaderboard; 