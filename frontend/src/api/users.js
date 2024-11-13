import axios from "axios";

export async function fetchProfile() {
    const token = localStorage.getItem('access_token');

    try {
        const response = await axios.get('http://localhost:8080/api/authentication/profile', {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        });

        //console.log('Datos del perfil:', response.data);
        return response.data; // Devuelve los datos del perfil para usarlos en el componente
    } catch (error) {
        console.error('Error al obtener el perfil:', error.response?.data || error.message);
        return null; // Devuelve null si hay un error
    }
}

export async function fetchDomains() {
    const token = localStorage.getItem('access_token');
    const response = await axios.get(`http://localhost:8080/api/domain/getDomains`, {
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
    return response.data.domains;
}

export async function deleteDomain(domainId) {
    const token = localStorage.getItem('access_token');
    await axios.delete(`http://localhost:8080/api/domain/delete`, {
        headers: {
            Authorization: `Bearer ${token}`,
        },
        data: { domain_id: domainId },
    });
}

export async function addDomain(domainData) {
    const token = localStorage.getItem('access_token');
    const response = await axios.post(`http://localhost:8080/api/domain/add`, domainData, {
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
    return response.data;
}
