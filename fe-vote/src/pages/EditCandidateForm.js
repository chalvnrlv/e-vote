import React, { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';

function EditCandidateForm() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [candidate, setCandidate] = useState({ name: '', division_id: '' });
  const [divisions, setDivisions] = useState([]);

  useEffect(() => {
    const fetchCandidate = async () => {
      try {
        const response = await fetch(`http://localhost:8080/candidates/${id}`);
        const data = await response.json();
        setCandidate(data);
      } catch (err) {
        console.error('Error fetching candidate:', err);
      }
    };

    const fetchDivisions = async () => {
      try {
        const response = await fetch('http://localhost:8080/divisions');
        const data = await response.json();
        setDivisions(data);
      } catch (err) {
        console.error('Error fetching divisions:', err);
      }
    };

    fetchCandidate();
    fetchDivisions();
  }, [id]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch(`http://localhost:8080/candidates/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(candidate),
      });

      if (response.ok) {
        navigate('/admin');
      } else {
        console.error('Failed to update candidate');
      }
    } catch (err) {
      console.error('Error:', err);
    }
  };

  return (
    <div>
      <h1>Edit Candidate</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Candidate Name"
          value={candidate.name}
          onChange={(e) => setCandidate({ ...candidate, name: e.target.value })}
          required
        />
        <select
          value={candidate.division_id}
          onChange={(e) => setCandidate({ ...candidate, division_id: e.target.value })}
          required
        >
          <option value="">Select Division</option>
          {divisions.map((division) => (
            <option key={division.id} value={division.id}>
              {division.division}
            </option>
          ))}
        </select>
        <button type="submit">Update Candidate</button>
      </form>
    </div>
  );
}

export default EditCandidateForm;
