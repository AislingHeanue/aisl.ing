import React, { useEffect, useState } from "react";
import Intro from "./Intro.jsx";
import Timeline from "./Timeline.jsx";
import Portfolio from "./Portfolio.jsx";
import Footer from "../common/Footer.jsx";
import ThemeSwitcher from "../common/ThemeSwitcher.jsx";

const MainPage = () => {
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

export default MainPage;
