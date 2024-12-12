import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

function CreateCandidateForm() {
  const navigate = useNavigate();
  const [divisions, setDivisions] = useState([]);
  const [name, setName] = useState('');
  const [divisionId, setDivisionId] = useState('');

  useEffect(() => {
    const fetchDivisions = async () => {
      try {
        const response = await fetch('http://localhost:8080/divisions');
        const data = await response.json();
        setDivisions(data); // Assuming the response gives an array of divisions
      } catch (err) {
        console.error('Error fetching divisions:', err);
      }
    };
    fetchDivisions();
  }, []);

  const handleSubmit = async (event) => {
    event.preventDefault();
    const candidate = {
      name: name,
      division_id: divisionId,
    };

    try {
      const response = await fetch('http://localhost:8080/candidates', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(candidate),
      });

      if (response.ok) {
        navigate('/admin');
      } else {
        console.error('Error creating candidate');
      }
    } catch (err) {
      console.error('Error:', err);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <label>
        Candidate Name:
        <input
          type="text"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
        />
      </label>

      <label>
        Division:
        <select
          value={divisionId}
          onChange={(e) => setDivisionId(e.target.value)}
          required
        >
          <option value="">Select a Division</option>
          {divisions.map((division) => (
            <option key={division.id} value={division.id}>
              {division.division}
            </option>
          ))}
        </select>
      </label>

      <button type="submit">Create Candidate</button>
    </form>
  );
}

export default CreateCandidateForm;
