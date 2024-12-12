import React, { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';

function EditUserForm() {
  const { id } = useParams();
  const navigate = useNavigate();

  const [user, setUser] = useState({});

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const response = await fetch(`http://localhost:8080/users/${id}`);
        const data = await response.json();
        setUser(data);
      } catch (err) {
        console.error('Error fetching user:', err);
      }
    };
    fetchUser();
  }, [id]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch(`http://localhost:8080/users/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(user),
      });

      if (response.ok) {
        navigate('/admin');
      } else {
        console.error('Failed to update user');
      }
    } catch (err) {
      console.error('Error:', err);
    }
  };

  return (
    <div>
      <h1>Edit User</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Name"
          value={user.name || ''}
          onChange={(e) => setUser({ ...user, name: e.target.value })}
          required
        />
        <input
          type="text"
          placeholder="Identity Number"
          value={user.identity_number || ''}
          onChange={(e) => setUser({ ...user, identity_number: e.target.value })}
          required
        />
        <input
          type="password"
          placeholder="Password"
          value={user.password || ''}
          onChange={(e) => setUser({ ...user, password: e.target.value })}
        />
        <select
          value={user.role_id || 2}
          onChange={(e) => setUser({ ...user, role_id: Number(e.target.value) })}
        >
          <option value={1}>Admin</option>
          <option value={2}>User</option>
        </select>
        <button type="submit">Update User</button>
      </form>
    </div>
  );
}

export default EditUserForm;
