import React from 'react';
import { Navigate } from 'react-router-dom';

function ProtectedRoute({ children }) {
    const token = localStorage.getItem('access_token');

    if (!token) {
        // Si no hay token, redirige a la página de inicio de sesión
        return <Navigate to="/login" replace />;
    }

    // Si hay token, renderiza el componente deseado
    return children;
}

export default ProtectedRoute;
