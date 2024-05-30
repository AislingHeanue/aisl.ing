import React, { useEffect, useState } from "react";
import WasmApp from "../components/WasmApp";

const Wasm = () => {
  const [theme, setTheme] = useState("dark");
  useEffect(() => {
    if (theme === "dark") {
      document.documentElement.classList.add("dark");
    } else {
      document.documentElement.classList.remove("dark");
    }
  }, [theme]);
  console.log("loading 3");

  return (
    <>
      <script src="../public/wasm_exec.js"></script>
      <WasmApp theme={theme} />
    </>
  );
};

export default Wasm;
