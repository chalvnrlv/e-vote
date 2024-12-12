import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Login from './components/Login';
import Vote from './components/Vote';
import Division from './components/Division';
import AdminDashboard from './pages/AdminDashboard';
import CreateUserForm from './pages/CreateUserForm';
import EditUserForm from './pages/EditUserForm';
import CreateDivisionForm from './pages/CreateDivisionForm';
import EditDivisionForm from './pages/EditDivisionForm';
import CreateCandidateForm from './pages/CreateCandidateForm';
import EditCandidateForm from './pages/EditCandidateForm';



function App() {

  return (
    <Router>
       <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/vote/:userId" element={<Vote />} />
        <Route path="/divisions/:divisionId" element={<Division />} />
        <Route path="/admin/dashboard" element={<AdminDashboard />} />
        <Route path="/admin/create-user" element={<CreateUserForm />} />
        <Route path="/admin/edit-user/:id" element={<EditUserForm />} />
        <Route path="/admin/create-division" element={<CreateDivisionForm />} />
        <Route path="/admin/edit-division/:id" element={<EditDivisionForm />} />
        <Route path="/admin/create-candidate" element={<CreateCandidateForm />} />
        <Route path="/admin/edit-candidate/:id" element={<EditCandidateForm />} />
      </Routes>
    </Router>
  );
}

export default App;