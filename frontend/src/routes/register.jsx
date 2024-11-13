import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

export default function Register() {
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const navigate = useNavigate();

    // Validación de email
    const validateEmail = (email) => {
        const emailRegex = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/;
        return emailRegex.test(email);
    };

    // Validación de la contraseña
    const validatePassword = (password) => {
        const passwordRegex = /^(?=.*[A-Z])(?=.*[!@#$%^&*])[A-Za-z\d!@#$%^&*]{8,}$/;
        return passwordRegex.test(password);
    };

    // Manejo del formulario
    async function handleRegister(e) {
        e.preventDefault();
        // Validamos el email
        if (!validateEmail(email)) {
            setError("El correo electrónico no tiene un formato válido.");
            return;
        }

        // Validamos la contraseña
        if (!validatePassword(password)) {
            setError("La contraseña debe tener al menos 8 caracteres, una mayúscula y un símbolo.");
            return;
        }

        // Si las validaciones son exitosas, intentamos registrar al usuario
        try {
            const requestBody = { firstName, lastName, email, password };
            const response = await axios.post('http://localhost:8080/api/authentication/signup', requestBody);
            console.log(response);
            navigate('/'); // Redirigir al usuario después del registro
        } catch (error) {
            console.log(error);
            setError("Hubo un error al crear la cuenta. Intenta de nuevo.");
        }
    }

    return (
        <div className="container" style={{ marginTop: "10vh" }}>
            <form onSubmit={handleRegister}>
                <h2>Crea una cuenta para poder manejar tus dominios</h2>
                <p>Bienvenido</p>

                {/* Primer nombre */}
                <div className="mb-3">
                    <label htmlFor="firstName" className="form-label">First Name :</label>
                    <input
                        type="text"
                        className="form-control"
                        id="firstName"
                        value={firstName}
                        onChange={e => setFirstName(e.target.value)}
                    />
                </div>

                {/* Apellido */}
                <div className="mb-3">
                    <label htmlFor="lastName" className="form-label">Last Name :</label>
                    <input
                        type="text"
                        className="form-control"
                        id="lastName"
                        value={lastName}
                        onChange={e => setLastName(e.target.value)}
                    />
                </div>

                {/* Correo electrónico */}
                <div className="mb-3">
                    <label htmlFor="email" className="form-label">Email address :</label>
                    <input
                        type="email"
                        className="form-control"
                        id="email"
                        value={email}
                        onChange={e => setEmail(e.target.value)}
                    />
                </div>

                {/* Contraseña */}
                <div className="mb-3">
                    <label htmlFor="password" className="form-label">Password :</label>
                    <input
                        type="password"
                        className="form-control"
                        id="password"
                        value={password}
                        onChange={e => setPassword(e.target.value)}
                    />
                </div>

                {/* Mensaje de error */}
                {error && <div className="alert alert-danger">{error}</div>}

                {/* Botón para enviar el formulario */}
                <button type="submit" className="btn btn-primary">Register</button>

                {/* Información de usuario demo */}
                <p><br/>Demo User: <br/>Email: user@example.com <br/>Password: password12345</p>
            </form>
        </div>
    );
}
