import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Main from "./pages/main.jsx";
import Collatz from "./pages/collatz.jsx";
import Wasm from "./pages/wasm.jsx";
import "./styles/tailwind.css";

// console.log("loading 4");
// flash of unstyled content preventer
document.documentElement.classList.add("dark");
ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <Router>
      <Routes>
        <Route exact path="/" element={<Main />} />
        <Route path="/collatz" element={<Collatz />} />
        <Route path="/wasm" element={<Wasm />} />
      </Routes>
    </Router>
  </React.StrictMode>
);
