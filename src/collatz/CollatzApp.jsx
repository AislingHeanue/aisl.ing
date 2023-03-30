import React, {useEffect, useState} from 'react'
import View from '../components/View';
import Title from '../components/Title';
import Footer from "../components/Footer"

function App() {
	const [theme, setTheme] = useState("dark");
	const [r,setR] = useState(2.5)
	const [angle,setAngle] = useState(0.15)
	const [n,setN] = useState(800/4000)
	const [alpha,setAlpha] = useState(0.7)
	const [rotation,setRotation] = useState(Math.PI/2)
	const [rootx,setRootx] = useState(75)
	const [rooty,setRooty] = useState(200)

	useEffect(() => {
		setTheme("dark"); 
	}, []);


	const handleThemeSwitch = () => {
		setTheme(theme === "dark" ? "light" : "dark");
	};

	const changeRotation = (e) => {
		setRotation(e.target.value)
	}

	const changeAlpha = (e) => {
		setAlpha(e.target.value)
	}

	const changeAngle = (e) => {
		setAngle(e.target.value)
	};

	const changeLength = (e) => {
		setR(e.target.value)
	};

	const changeN = (e) => {
		setN(e.target.value)
	};

	const changeRootx = (e) => {
		setRootx(Math.floor(200*e.target.value))
	}

	const changeRooty = (e) => {
		setRooty(250-Math.floor(250*e.target.value))
	}


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
        className="fixed p-2 z-10 right-20 bottom-4 md:top-4 md:bottom-auto bg-purple-300 dark:bg-orange-300 rk:bg-w text-lg p-1 rounded-md"
      >
        {theme === "dark"?sun:moon}
      </button>
      <div className="bg-white dark:bg-stone-900 dark:text-stone-300 text-stone-900 min-h-screen font-inter font-medium ">
        <div className="grid lg:grid-cols-2 md:grid-cols-1 w-full md:w-5/6 mx-auto gap-4">
			<div className=''>
				<Title>Collatz Tree Visualiser</Title>
				<button tpye="button" className='mb-2 fort-semibold text-white bg-stone-900 dark:bg-white dark:text-stone-900 rounded ml-1 pl-2 pr-2'><a href='/'>Back to homepage</a></button>
				<p className='mb-4'>The Collatz conjecture states that preformng the operation "multiply n by 3 and add 1 if it is odd, or half it if is even"
					will always eventually reach the final value of 1 (before looping indefinitely between 1, 4 and 2). This has been shown to be true for many values of n, but the general conjecture
					remains unproven. This tree diagram
					represents the possible sets of operations a sequence can take to get from its starting value to 1 (which is specified as the root node) for a set list
					of values of n (from 2 to N). A right turn in the direction facing away from 1 (the root node) indicates "n -&gt;2n", and a left turn indicates "n -&gt; (2n-1)/3".
				</p>
				
				<div className="mt-2 w-full flex-col md:flex-row items-center">
					<div className="grid grid-cols-2 lg:grid-cols-1 gap-4">
						<label className="block md:text-sm text-xs text-stone-900 dark:text-white dark:text-white">Branching Angle (0° to 90°)</label>
						<input id="angle-range" onChange={changeAngle} type="range" min="0" max={Math.PI/2} step = {Math.PI/180} value = {angle} className="w-full h-2 mt-1  bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"></input>
						<label className="block md:text-sm text-xs text-stone-900 dark:text-white dark:text-white">Line Length (0 to 7)</label>
						<input id="length-range" onChange={changeLength} type="range" min="0" max="7" value = {r} step="any" className="w-full h-2 mt-2 bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"></input>
						<label className="block md:text-sm text-xs text-stone-900 dark:text-white dark:text-white">N (0 to 4000)</label>
						<input id="n-range" onChange={changeN} type="range" min="0" max="1" step="any" value={n} className="w-full h-2 mt-2 bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"></input>
						<label className="block md:text-sm text-xs text-stone-900 dark:text-white dark:text-white">Opacity (0 to 1)</label>
						<input id="alpha-range" onChange={changeAlpha} type="range" min="0" max="1" step="any" value={alpha} className="w-full h-2 mt-2 bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"></input>
						<label className="block md:text-sm text-xs text-stone-900 dark:text-white dark:text-white">Translate X</label>
						<input id="alpha-range" onChange={changeRootx} type="range" min="0" max="1" step="any" value={rootx/200} className="w-full h-2 mt-2 bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"></input>
						<label className="block md:text-sm text-xs text-stone-900 dark:text-white dark:text-white">Translate Y</label>
						<input id="alpha-range" onChange={changeRooty} type="range" min="0" max="1" step="any" value={1-rooty/250} className="w-full h-2 mt-2 bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"></input>
						<label className="block md:text-sm text-xs text-stone-900 dark:text-white dark:text-white">Global Rotation</label>
						<input id="angle-range" onChange={changeRotation} type="range" min={0} max={2*Math.PI} step = {Math.PI/180} value = {rotation} className="w-full h-2 mt-2  bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"></input>
					</div>
				</div>
			</div>
			<div>
				<View r={r} angle={angle} n={n} alpha={alpha} rotation = {rotation} rootx={rootx} rooty={rooty}/>  
			</div>
			   
   
        </div>

		<Footer/>  
      </div>
    </>
  )
}

export default App
