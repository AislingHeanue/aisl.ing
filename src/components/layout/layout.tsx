
import { PropsWithChildren } from "react";
import '../../style/tailwind.css'
import Footer from "./footer";
import ThemeSwitcherClient from "./themeSwitcherClient";


export default function Layout({ children, scrollable }: PropsWithChildren<{ scrollable: boolean }>) {
  return (
    <>
      <div className={`bg-white dark:bg-stone-900 dark:text-stone-300 text-stone-900 min-h-screen font-inter lg:justify-center ${!scrollable ? "lg:h-screen lg:flex-row lg:flex lg:overflow-hidden" : ""}`} >
        <ThemeSwitcherClient />
        <main className={` ${!scrollable ? "lg:pb-25 lg:overflow-hidden" : ""}`}>
          {children}
        </main>
        <Footer absolute={!scrollable} />
      </div >
    </>
  );
};
