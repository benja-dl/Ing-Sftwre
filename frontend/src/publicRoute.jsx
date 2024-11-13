import React from 'react';
import { Navigate } from 'react-router-dom';

function PublicRoute({ children }) {
    const token = localStorage.getItem('access_token');

    if (token) {
        // Si hay token, el usuario ya está autenticado, redirige a una ruta protegida (por ejemplo, `/home`)
        return <Navigate to="/home" replace />;
    }

    // Si no hay token, permite el acceso a la ruta pública
    return children;
}

export default PublicRoute;
