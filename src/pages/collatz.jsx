import React, { useEffect, useState } from "react";
import Footer from "../components/Footer";
import ThemeSwitcher from "../components/ThemeSwitcher.jsx";
import CollatzApp from "../components/CollatzApp";

const Collatz = () => {
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
      <div className="bg-white dark:bg-stone-900 dark:text-stone-300 text-stone-900 min-h-screen font-inter ">
        <ThemeSwitcher theme={theme} setTheme={setTheme} />
        <CollatzApp theme={theme} />
      </div>
    </>
  );
};

export default Collatz;
