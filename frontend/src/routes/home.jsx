import { useEffect, useState } from 'react';
import { fetchProfile } from '../api/users.js';
import { useNavigate } from 'react-router-dom';

export default function Home() {
    const [data, setData] = useState(null);
    const navigate = useNavigate();

    // Función para cerrar sesión
    const handleLogout = () => {
        localStorage.removeItem('access_token'); // Elimina el token de localStorage
        navigate('/'); // Redirige a la página de inicio
    };

    useEffect(() => {
        async function getData() {
            const profileData = await fetchProfile();
            setData(profileData);
        }

        getData();
    }, []);

    if (!data) {
        return <p>Cargando...</p>; // Muestra un mensaje de carga mientras los datos se obtienen
    }

    return (
        <div>
            <p>HOME</p>
            <p>Nombre: {data.first_name}</p>
            <p>Apellido: {data.last_name}</p>
            <p>Email: {data.email}</p>

            {/* Botón para cerrar sesión */}
            <button onClick={handleLogout} className="btn btn-danger">
                Cerrar sesión
            </button>
        </div>
    );
}
