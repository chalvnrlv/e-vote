import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

function Login() {
  const [identityNumber, setIdentityNumber] = useState('5025221054');
  const [password, setPassword] = useState('loginyok');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json',
        },
        body: JSON.stringify({ 
          Identity_Number: identityNumber, 
          password 
        }),
      });
  
      console.log('Response status:', response.status);
      console.log('Response headers:', Object.fromEntries(response.headers.entries()));
  
      if (!response.ok) {
        const text = await response.text();
        console.error('Error response text:', text);
        throw new Error(text || 'Login failed');
      }
      
      const data = await response.json();
  
      // Redirect
      navigate(data.redirectURL);
    } catch (err) {
      console.error('Fetch error:', err);
      setError(err.message);
    }
  };

  return (
    <div>
      <h1>Login</h1>
      <form onSubmit={handleSubmit}>
        <input 
          type="text" 
          id="identityNumber"
          value={identityNumber} 
          onChange={(e) => setIdentityNumber(e.target.value)} 
          placeholder="Identity Number"
        />
        <input 
          type="password" 
          id="password"
          value={password} 
          onChange={(e) => setPassword(e.target.value)} 
          placeholder="Password"
        />
        <button type="submit">Login</button>
      </form>
      {error && <p style={{ color: 'red' }}>{error}</p>}
    </div>
  );
}

export default Login;