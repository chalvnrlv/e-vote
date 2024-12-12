import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';

function Vote() {
  const { userId } = useParams();
  const [voteStatus, setVoteStatus] = useState(null);
  const [divisions, setDivisions] = useState([]);
  const [candidates, setCandidates] = useState([]);
  const [selectedCandidate, setSelectedCandidate] = useState(null);
  const [selectedDivision, setSelectedDivision] = useState(null); // Track selected division
  const [votedDivisions, setVotedDivisions] = useState([]); // Track divisions the user has voted for
  const [results, setResults] = useState([]); // Store results for each division
  const navigate = useNavigate();

  // Fetch vote status and divisions when the component mounts
  useEffect(() => {
    const fetchVoteStatus = async () => {
      try {
        const response = await fetch(`http://localhost:8080/vote/${userId}`);
        if (!response.ok) {
          throw new Error('Failed to fetch vote status');
        }
        const data = await response.json();
        setVoteStatus(data);
        setVotedDivisions(data.votedDivisions || []); // Ensure votedDivisions is an empty array if undefined
      } catch (err) {
        console.error(err);
      }
    };

    const fetchDivisions = async () => {
      try {
        const response = await fetch('http://localhost:8080/divisions');
        if (!response.ok) {
          throw new Error('Failed to fetch divisions');
        }
        const data = await response.json();
        setDivisions(data);
      } catch (err) {
        console.error(err);
      }
    };

    fetchVoteStatus();
    fetchDivisions();
  }, [userId]);

  // Fetch candidates for a selected division
  const fetchCandidates = async (divisionId) => {
    try {
      const response = await fetch(`http://localhost:8080/divisions/${divisionId}/candidates`);
      if (!response.ok) {
        throw new Error('Failed to fetch candidates');
      }
      const data = await response.json();
      setCandidates(data);
      setSelectedDivision(divisionId); // Set the selected division
    } catch (err) {
      console.error(err);
    }
  };

  // Handle vote submission
  const handleVoteSubmit = async () => {
    if (!selectedCandidate) {
      alert('Please select a candidate to vote for.');
      return;
    }

    const voteData = { Candidate_ID: selectedCandidate };

    try {
      const response = await fetch(`http://localhost:8080/vote/div-${selectedDivision}/${userId}`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(voteData),
      });

      if (!response.ok) {
        const errorData = await response.json();
        alert(`Error: ${errorData.error}`);
        return;
      }

      const updatedData = await response.json();
      setVoteStatus(updatedData);
      setVotedDivisions((prev) => [...prev, selectedDivision]); // Add division to voted list
      alert('Vote submitted successfully!');

      // Check if all divisions are voted
      if (updatedData.votedDivisions.length === divisions.length) {
        fetchResults(); // Fetch results if all divisions are voted
      } else {
        navigate(`/vote/${userId}`); // Reload the vote page
      }
    } catch (err) {
      console.error(err);
      alert('Error casting vote.');
    }
  };

  // Handle candidate selection
  const handleCandidateSelect = (candidateId) => {
    setSelectedCandidate(candidateId);
  };

  // Fetch results for all divisions
  const fetchResults = async () => {
    try {
      const resultsData = [];
      for (const division of divisions) {
        const response = await fetch(`http://localhost:8080/results/${division.id}`);
        if (!response.ok) {
          throw new Error('Failed to fetch results');
        }
        const data = await response.json();
        resultsData.push(data);
      }
      setResults(resultsData);
    } catch (err) {
      console.error(err);
    }
  };

  if (!voteStatus || !divisions.length) {
    return <div>Loading...</div>;
  }

  const allDivisionsVoted = voteStatus.message === 'You have voted in all divisions';

  return (
    <div>
      <h1>Vote</h1>
      <p>{voteStatus.message}</p>

      {allDivisionsVoted ? (

        <div>
          <h3>Voting Results</h3>
          {results.length > 0 ? (
            <div>
              {results.map((result, index) => (
                <div key={divisions[index].id}>
                  <h4>{divisions[index].division}</h4>
                  <ul>
                    {result.candidates.map((candidate) => (
                      <li key={candidate.id}>
                        {candidate.name}: {candidate.vote_count}
                      </li>
                    ))}
                  </ul>
                </div>
              ))}
            </div>
          ) : (
            <div>Loading results...</div>
          )}
        </div>
      ) : (
        // If not all divisions are voted, show divisions that the user hasn't voted for
        <div>
          <h2>Divisions</h2>
          <ul>
            {divisions
              .filter((division) => !votedDivisions.includes(division.id)) // Only show unvoted divisions
              .map((division) => (
                <li key={division.id} onClick={() => fetchCandidates(division.id)}>
                  {division.division}
                </li>
              ))}
          </ul>
        </div>
      )}

      {candidates.length > 0 && (
        <div>
          <h3>Select a Candidate</h3>
          <ul>
            {candidates.map((candidate) => (
              <li
                key={candidate.id}
                onClick={() => handleCandidateSelect(candidate.id)}
                style={{
                  cursor: 'pointer',
                  backgroundColor: selectedCandidate === candidate.id ? 'lightgray' : 'transparent',
                }}
              >
                {candidate.name}
              </li>
            ))}
          </ul>

          <button onClick={handleVoteSubmit}>Submit Vote</button>
        </div>
      )}
    </div>
  );
}

export default Vote;
