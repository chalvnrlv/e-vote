import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

function CreateUserForm() {
  const navigate = useNavigate();

  const [name, setName] = useState('');
  const [identityNumber, setIdentityNumber] = useState('');
  const [password, setPassword] = useState('');
  const [roleID, setRoleID] = useState(2);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch('http://localhost:8080/users', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name, identity_number: identityNumber, password, role_id: roleID }),
      });

      if (response.ok) {
        navigate('/admin');
      } else {
        console.error('Failed to create user');
      }
    } catch (err) {
      console.error('Error:', err);
    }
  };

  return (
    <div>
      <h1>Create User</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
        />
        <input
          type="text"
          placeholder="Identity Number"
          value={identityNumber}
          onChange={(e) => setIdentityNumber(e.target.value)}
          required
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <select
          value={roleID}
          onChange={(e) => setRoleID(Number(e.target.value))}
        >
          <option value={1}>Admin</option>
          <option value={2}>User</option>
        </select>
        <button type="submit">Create User</button>
      </form>
    </div>
  );
}

export default CreateUserForm;
