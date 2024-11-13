import { useEffect, useState } from 'react';
import { fetchProfile, fetchDomains, deleteDomain, addDomain } from '../api/users.js';
import { useNavigate } from 'react-router-dom';

export default function Home() {
    const [data, setData] = useState(null);
    const [domains, setDomains] = useState([]);
    const [newDomain, setNewDomain] = useState({ name: '', url: '' });
    const navigate = useNavigate();

    useEffect(() => {
        async function getData() {
            const profileData = await fetchProfile();
            setData(profileData);

            const domainData = await fetchDomains();
            setDomains(domainData);
        }

        getData();
    }, []);

    const handleDelete = async (domainId) => {
        try {
            await deleteDomain(domainId);
            setDomains(domains.filter(domain => domain.id !== domainId));
        } catch (error) {
            console.error("Error al eliminar el dominio:", error);
        }
    };

    const handleLogout = () => {
        localStorage.removeItem('access_token');
        navigate('/');
    };

    const handleAddDomain = async (e) => {
        e.preventDefault();
        try {
            // Agregar dominio y recargar la lista de dominios
            await addDomain(newDomain);
            const updatedDomains = await fetchDomains();
            setDomains(updatedDomains);
            setNewDomain({ name: '', url: '' });
        } catch (error) {
            console.error("Error al agregar dominio:", error);
        }
    };

    if (!data) {
        return <p>Cargando...</p>;
    }

    return (
        <div>
            <h2>Bienvenido {data.first_name} {data.last_name}</h2>
            <p>Lista de dominios:</p>
            <ul>
                {domains.map((domain) => (
                    <li key={domain.id}>
                        <span>{domain.name} - {domain.url}</span>
                        <button onClick={() => handleDelete(domain.id)}>Eliminar</button>
                    </li>
                ))}
            </ul>

            <h3>Añadir un nuevo dominio</h3>
            <form onSubmit={handleAddDomain}>
                <input
                    type="text"
                    placeholder="Nombre del dominio"
                    value={newDomain.name}
                    onChange={(e) => setNewDomain({ ...newDomain, name: e.target.value })}
                    required
                />
                <input
                    type="text"
                    placeholder="URL del dominio"
                    value={newDomain.url}
                    onChange={(e) => setNewDomain({ ...newDomain, url: e.target.value })}
                    required
                />
                <button type="submit">Agregar Dominio</button>
            </form>

            <button onClick={handleLogout}>Cerrar sesión</button>
        </div>
    );
}
