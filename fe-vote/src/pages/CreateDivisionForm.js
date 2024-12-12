import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

function CreateDivisionForm() {
  const navigate = useNavigate();
  const [name, setName] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch('http://localhost:8080/divisions', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name }),
      });

      if (response.ok) {
        navigate('/admin');
      } else {
        console.error('Failed to create division');
      }
    } catch (err) {
      console.error('Error:', err);
    }
  };

  return (
    <div>
      <h1>Create Division</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Division Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
        />
        <button type="submit">Create Division</button>
      </form>
    </div>
  );
}

export default CreateDivisionForm;
