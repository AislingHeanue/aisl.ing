import React, { useEffect, useState } from "react";
import Intro from "../components/Intro";
import Timeline from "../components/Timeline";
import Portfolio from "../components/Portfolio";
import Footer from "../components/Footer";
import ThemeSwitcher from "../components/ThemeSwitcher.jsx";

const Main = () => {
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
        <div className="max-w-5xl w-11/12 mx-auto">
          <Intro />
          <Portfolio />
          <Timeline />
          <Footer />
        </div>
      </div>
    </>
  );
};

export default Main;
