import React, {useEffect, useState} from 'react'
import View from '../components/View';
import Title from '../components/Title';
import Footer from "../components/Footer"

function App() {
	const [theme, setTheme] = useState("dark");
	const [r,setR] = useState(2.5)
	const [angle,setAngle] = useState(0.15)
	const [n,setN] = useState(800/4000)

	useEffect(() => {
		setTheme("dark"); 
	}, []);


	const handleThemeSwitch = () => {
		setTheme(theme === "dark" ? "light" : "dark");
	};

	const changeAngle = (e) => {
		setAngle(e.target.value)
	};

	const changeLength = (e) => {
		setR(e.target.value)
	};

	const changeN = (e) => {
		setN(e.target.value)
	};

	useEffect(() => {
		if (theme === "dark") {
			document.documentElement.classList.add("dark");
		} else {
			document.documentElement.classList.remove("dark");
		}
	}, [theme]);




  const sun = (
		<svg
			xmlns="http://www.w3.org/2000/svg"
			fill="none"
			viewBox="0 0 24 24"
			strokeWidth={1.5}
			stroke="currentColor"
			className="w-6 h-6"
		>
			<path
				strokeLinecap="round"
				strokeLinejoin="round"
				d="M12 3v2.25m6.364.386l-1.591 1.591M21 12h-2.25m-.386 6.364l-1.591-1.591M12 18.75V21m-4.773-4.227l-1.591 1.591M5.25 12H3m4.227-4.773L5.636 5.636M15.75 12a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0z"
			/>
		</svg>
	);

	const moon = (
		<svg
			xmlns="http://www.w3.org/2000/svg"
			fill="none"
			viewBox="0 0 24 24"
			strokeWidth={1.5}
			stroke="white"
			className="w-6 h-6"
		>
			<path
				strokeLinecap="round"
				strokeLinejoin="round"
				d="M21.752 15.002A9.718 9.718 0 0118 15.75c-5.385 0-9.75-4.365-9.75-9.75 0-1.33.266-2.597.748-3.752A9.753 9.753 0 003 11.25C3 16.635 7.365 21 12.75 21a9.753 9.753 0 009.002-5.998z"
			/>
		</svg>
	);
  return (
    <>
      <button
        type = "button"
        onClick={handleThemeSwitch}
        className="fixed p-2 z-10 right-20 top-4 bg-purple-300 dark:bg-orange-300 rk:bg-w text-lg p-1 rounded-md"
      >
        {theme === "dark"?sun:moon}
      </button>
      <div className="bg-white dark:bg-stone-900 dark:text-stone-300 text-stone-900 min-h-screen font-inter font-medium ">
        <div className="max-w-5xl w-11/12 mx-auto">
			<Title>Collatz Tree Visualiser</Title>
			<p><a href='/'>Back to homepage</a></p>
			<p>Todo: Add a README here.</p>
			<label for="angle-range" className="block mb-2 text-sm text-stone-900 dark:text-white dark:text-white">Angle (0 to 0.3)</label>
			<input id="angle-range" onChange={changeAngle} type="range" min="0" max="0.3" step = "any" value = {angle} className="w-1/2 h-2 bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"></input>
			<label for="length-range" className="block mb-2 text-smtext-stone-900 dark:text-white dark:text-white">Line length (0 to 5)</label>
			<input id="length-range" onChange={changeLength} type="range" min="0" max="5" value = {r} step="any" className="w-1/2 h-2 bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"></input>
			<label for="n-range" className="block mb-2 text-s text-stone-900 dark:text-white dark:text-white">N (0 to 1600)</label>
			<input id="n-range" onChange={changeN} type="range" min="0" max="1" step="any" value={n} className="w-1/2 h-2 bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"></input>
			<View r={r} angle={angle} n={n}/>     
			<Footer/>     
        </div>
      </div>
    </>
  )
}

export default App
