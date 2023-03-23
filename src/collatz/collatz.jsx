import React from 'react'
import ReactDOM from 'react-dom/client'
import '../styles/tailwind.css'
import App from "./CollatzApp"

document.documentElement.classList.add("dark")
ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <App/>
  </React.StrictMode>,
)
