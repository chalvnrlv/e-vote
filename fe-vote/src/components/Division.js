import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';


function Division() {
  const { divisionId } = useParams();
  const navigate = useNavigate();  // Import useNavigate
  const [candidates, setCandidates] = useState([]);
  const [selectedCandidate, setSelectedCandidate] = useState(null);
  const [userId, setUserId] = useState(null); // Add userId state
  const [voteMessage, setVoteMessage] = useState('');

  useEffect(() => {
    const fetchUserId = async () => {
        const storedUserId = localStorage.getItem('userId'); // Retrieve userId from localStorage
        setUserId(storedUserId); // Store user id in state
    }

    fetchUserId();

    const fetchCandidates = async () => {
      try {
        const response = await fetch(`http://localhost:8080/divisions/${divisionId}`);
        if (!response.ok) {
          throw new Error('Failed to fetch candidates');
        }
        const data = await response.json();
        setCandidates(data);
      } catch (error) {
        console.error('Error fetching candidates:', error);
      }
    };
    fetchCandidates();
  }, [divisionId]);


  const handleCandidateSelect = (candidateId) => {
    setSelectedCandidate(candidateId);
  };

  const handleSubmitVote = async () => {
    if (!selectedCandidate || !userId) return; // Do nothing if no candidate is selected

    try {
      const response = await fetch(`http://localhost:8080/vote/div-${divisionId}/${userId}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ Candidate_ID: selectedCandidate }),
      });

      if (!response.ok) {
        const data = await response.json(); // Get error message from response
        throw new Error(data.error || 'Vote submission failed');
      }

      setVoteMessage('Vote submitted successfully!');
      navigate(`/vote/${userId}`);

    } catch (error) {
      console.error('Error submitting vote:', error);
      setVoteMessage(`Error: ${error.message}`); // Display error message
    }
  };



  return (
    <div>
      <h1>Candidates for Division {divisionId}</h1>
      <ul>
        {candidates.map(candidate => (
          <li key={candidate.id}>
            <input 
              type="radio" 
              name="candidate" 
              value={candidate.id} 
              onChange={() => handleCandidateSelect(candidate.id)} 
            />
            {candidate.name}
          </li>
        ))}
      </ul>
      <button onClick={handleSubmitVote} disabled={!selectedCandidate}>Submit Vote</button>
      {/* Display vote message */}
      {voteMessage && <p>{voteMessage}</p>}
    </div>
  );
}

export default Division;