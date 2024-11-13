import * as React from 'react';
import * as ReactDOM from 'react-dom/client';
import Root from './routes/root.jsx'
import {
    createBrowserRouter,
    RouterProvider,
} from 'react-router-dom'
import './index.css'
import ErrorPage from "./error-page.jsx";
import LoginPage from "./routes/login";
import ProtectedRoute from "./protectedRoute";
import PublicRoute from "./publicRoute";
import Home from "./routes/home";

const router = createBrowserRouter([
    {
        path: '/',
        element: <PublicRoute><Root /></PublicRoute>,
        errorElement: <ErrorPage />,
    },
    {
        path: '/home',
        element: <Home />,
    },
    {
        path: '/login',
        element: <PublicRoute><LoginPage /></PublicRoute>
    }
])

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
