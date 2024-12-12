import React, { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';

function EditDivisionForm() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [division, setDivision] = useState({ name: '' });

  useEffect(() => {
    const fetchDivision = async () => {
      try {
        const response = await fetch(`http://localhost:8080/divisions/${id}`);
        const data = await response.json();
        setDivision(data);
      } catch (err) {
        console.error('Error fetching division:', err);
      }
    };
    fetchDivision();
  }, [id]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch(`http://localhost:8080/divisions/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(division),
      });

      if (response.ok) {
        navigate('/admin');
      } else {
        console.error('Failed to update division');
      }
    } catch (err) {
      console.error('Error:', err);
    }
  };

  return (
    <div>
      <h1>Edit Division</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Division Name"
          value={division.name}
          onChange={(e) => setDivision({ ...division, name: e.target.value })}
          required
        />
        <button type="submit">Update Division</button>
      </form>
    </div>
  );
}

export default EditDivisionForm;
