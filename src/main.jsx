import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import MainPage from "./components/main/MainPage.jsx";
import CollatzPage from "./components/collatz/CollatzPage.jsx";
import RubiksCubePage from "./components/rubiks-cube/RubiksCubePage.jsx";
import "./styles/tailwind.css";

// flash of unstyled content preventer
document.documentElement.classList.add("dark");
ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <Router>
      <Routes>
        <Route exact path="/" element={<MainPage />} />
        <Route path="/collatz" element={<CollatzPage />} />
        <Route path="/rubiks-cube" element={<RubiksCubePage />} />
      </Routes>
    </Router>
  </React.StrictMode>
);
