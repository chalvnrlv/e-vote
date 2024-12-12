import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';

function AdminDashboard() {
  const navigate = useNavigate();
  const [users, setUsers] = useState([]);
  const [divisions, setDivisions] = useState([]);
  const [candidates, setCandidates] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        // Fetch Users, Divisions, and Candidates
        const usersResponse = await fetch('http://localhost:8080/users');
        const usersData = await usersResponse.json();
        setUsers(usersData.users);

        const divisionsResponse = await fetch('http://localhost:8080/divisions');
        const divisionsData = await divisionsResponse.json();
        setDivisions(divisionsData); // No need to change this part

        const candidatesResponse = await fetch('http://localhost:8080/candidates');
        const candidatesData = await candidatesResponse.json();
        setCandidates(candidatesData.candidates);
      } catch (err) {
        console.error('Error fetching data:', err);
      }
    };
    fetchData();
  }, []);

  const handleDeleteUser = async (id) => {
    try {
      const response = await fetch(`http://localhost:8080/users/${id}`, {
        method: 'DELETE',
      });

      if (response.ok) {
        setUsers(users.filter(user => user.id !== id));
      } else {
        console.error('Error deleting user');
      }
    } catch (err) {
      console.error('Error:', err);
    }
  };

  const handleDeleteCandidate = async (id) => {
    try {
      const response = await fetch(`http://localhost:8080/candidates/${id}`, {
        method: 'DELETE',
      });

      if (response.ok) {
        setCandidates(candidates.filter(candidate => candidate.id !== id));
      } else {
        console.error('Error deleting candidate');
      }
    } catch (err) {
      console.error('Error:', err);
    }
  };

  const handleDeleteDivision = async (id) => {
    try {
      const response = await fetch(`http://localhost:8080/divisions/${id}`, {
        method: 'DELETE',
      });

      if (response.ok) {
        setDivisions(divisions.filter(division => division.id !== id));
      } else {
        console.error('Error deleting division');
      }
    } catch (err) {
      console.error('Error:', err);
    }
  };

  return (
    <div>
      <h1>Admin Dashboard</h1>

      {/* Users Section */}
      <section>
        <h2>Users</h2>
        <button onClick={() => navigate('/admin/create-user')}>Create User</button>
        <ul>
          {users.map(user => (
            <li key={user.id}>
              {user.name} ({user.identity_number})
              <button onClick={() => navigate(`/admin/edit-user/${user.id}`)}>Edit</button>
              <button onClick={() => handleDeleteUser(user.id)}>Delete</button>
            </li>
          ))}
        </ul>
      </section>

      {/* Divisions Section */}
      <section>
        <h2>Divisions</h2>
        <button onClick={() => navigate('/admin/create-division')}>Create Division</button>
        <ul>
          {divisions.map(division => (
            <li key={division.id}>
              {division.division}
              <button onClick={() => navigate(`/admin/edit-division/${division.id}`)}>Edit</button>
              <button onClick={() => handleDeleteDivision(division.id)}>Delete</button>
            </li>
          ))}
        </ul>
      </section>

      {/* Candidates Section */}
      <section>
        <h2>Candidates</h2>
        <button onClick={() => navigate('/admin/create-candidate')}>Create Candidate</button>
        <ul>
          {candidates.map(candidate => (
            <li key={candidate.id}>
              {candidate.name} ({candidate.division}) {/* Update here to use 'division' */}
              <button onClick={() => navigate(`/admin/edit-candidate/${candidate.id}`)}>Edit</button>
              <button onClick={() => handleDeleteCandidate(candidate.id)}>Delete</button>
            </li>
          ))}
        </ul>
      </section>
    </div>
  );
}

export default AdminDashboard;
