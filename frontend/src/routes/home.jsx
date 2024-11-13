import { useEffect, useState } from 'react';
import { fetchProfile } from '../api/users.js';

export default function Home() {
    const [data, setData] = useState(null);

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
        </div>
    );
}
