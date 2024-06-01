import React, { useEffect, useState } from "react";
import ThemeSwitcher from "../components/ThemeSwitcher";
import WasmApp from "../components/WasmApp";
import Footer from "../components/Footer";

const Wasm = () => {
  const [theme, setTheme] = useState("dark");
  useEffect(() => {
    if (theme === "dark") {
      document.documentElement.classList.add("dark");
    } else {
      document.documentElement.classList.remove("dark");
    }
  }, [theme]);

  return (
    <>
      <div className="bg-white dark:bg-stone-900 dark:text-stone-300 text-stone-900 min-h-screen font-inter lg:justify-center lg:flex">
        <script src="../public/wasm_exec.js"></script>
        <ThemeSwitcher theme={theme} setTheme={setTheme} />
        <WasmApp theme={theme} />
      </div>
    </>
  );
};

export default Wasm;
