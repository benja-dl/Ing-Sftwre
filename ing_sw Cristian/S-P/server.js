// Importar las dependencias necesarias
const express = require('express');
const cors = require('cors');
const app = express();
const port = 5000; // Puedes cambiar el puerto si es necesario

// Habilitar CORS para todas las rutas
app.use(cors());

// Definir una ruta básica
app.get('/', (req, res) => {
    res.send('¡Servidor CORS habilitado!');
});

// Ruta para probar la latencia
app.get('/latency', (req, res) => {
    const start = Date.now();
    setTimeout(() => {
        const end = Date.now();
        const latency = end - start;
        res.json({ latency: latency });  // Regresar la latencia en milisegundos
    }, 500);  // Simulamos un retraso de 500 ms
});

// Escuchar en el puerto 5000
app.listen(port, () => {
    console.log(`Servidor corriendo en http://localhost:${port}`);
});
