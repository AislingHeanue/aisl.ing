import React, { useEffect, useState } from "react";
import View from "./View";
import Title from "../common/Title";

function Collatz({ theme }) {
  const [r, setR] = useState(2.5);
  const [angle, setAngle] = useState(0.15);
  const [n, setN] = useState(800 / 4000);
  const [alpha, setAlpha] = useState(0.7);
  const [rotation, setRotation] = useState(Math.PI / 2);
  const [rootx, setRootx] = useState(75);
  const [rooty, setRooty] = useState(200);

  const changeRotation = (e) => {
    setRotation(e.target.value);
  };

  const changeAlpha = (e) => {
    setAlpha(e.target.value);
  };

  const changeAngle = (e) => {
    setAngle(e.target.value);
  };

  const changeLength = (e) => {
    setR(e.target.value);
  };

  const changeN = (e) => {
    setN(e.target.value);
  };

  const changeRootx = (e) => {
    setRootx(Math.floor(200 * e.target.value));
  };

  const changeRooty = (e) => {
    setRooty(250 - Math.floor(250 * e.target.value));
  };

  return (
    <>
      <div className="grid lg:grid-cols-2 md:grid-cols-1 w-full lg:w-8/12 md:w-5/6 mx-auto gap-4">
        <div className="">
          <Title>Collatz Tree Visualiser</Title>
          <button
            type="button"
            className="mb-2 fort-semibold text-white bg-stone-900 dark:bg-white dark:text-stone-900 rounded ml-1 pl-2 pr-2"
          >
            <a href="/">Back to homepage</a>
          </button>
          <p className="mb-4">
            The Collatz conjecture states that preforming the operation
            "multiply n by 3 and add 1 if it is odd, or half it if is even" will
            always eventually reach the final value of 1 (before looping
            indefinitely between 1, 4 and 2). This has been shown to be true for
            many values of n, but the general conjecture remains unproven. This
            tree diagram represents the possible sets of operations a sequence
            can take to get from its starting value to 1 (which is specified as
            the root node) for a set list of values of n (from 2 to N). A right
            turn in the direction facing away from 1 (the root node) indicates
            "n -&gt;2n", and a left turn indicates "n -&gt; (2n-1)/3".
          </p>

          <div className="mt-2 flex-col md:flex-row items-center">
            <div className="grid grid-cols-2 lg:grid-cols-1 gap-4">
              <label className="block md:text-sm text-xs text-stone-900 dark:text-white">
                Branching Angle (0° to 90°)
              </label>
              <input
                id="angle-range"
                onChange={changeAngle}
                type="range"
                min="0"
                max={Math.PI / 2}
                step={Math.PI / 180}
                value={angle}
                className="w-full h-2 mt-1  bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"
              ></input>
              <label className="block md:text-sm text-xs text-stone-900 dark:text-white">
                Line Length (0 to 7)
              </label>
              <input
                id="length-range"
                onChange={changeLength}
                type="range"
                min="0"
                max="7"
                value={r}
                step="any"
                className="w-full h-2 mt-2 bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"
              ></input>
              <label className="block md:text-sm text-xs text-stone-900 dark:text-white">
                N (0 to 4000)
              </label>
              <input
                id="n-range"
                onChange={changeN}
                type="range"
                min="0"
                max="1"
                step="any"
                value={n}
                className="w-full h-2 mt-2 bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"
              ></input>
              <label className="block md:text-sm text-xs text-stone-900 dark:text-white">
                Opacity (0 to 1)
              </label>
              <input
                id="alpha-range"
                onChange={changeAlpha}
                type="range"
                min="0"
                max="1"
                step="any"
                value={alpha}
                className="w-full h-2 mt-2 bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"
              ></input>
              <label className="block md:text-sm text-xs text-stone-900 dark:text-white">
                Translate X
              </label>
              <input
                id="alpha-range"
                onChange={changeRootx}
                type="range"
                min="0"
                max="1"
                step="any"
                value={rootx / 200}
                className="w-full h-2 mt-2 bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"
              ></input>
              <label className="block md:text-sm text-xs text-stone-900 dark:text-white">
                Translate Y
              </label>
              <input
                id="alpha-range"
                onChange={changeRooty}
                type="range"
                min="0"
                max="1"
                step="any"
                value={1 - rooty / 250}
                className="w-full h-2 mt-2 bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"
              ></input>
              <label className="block md:text-sm text-xs text-stone-900 dark:text-white">
                Global Rotation
              </label>
              <input
                id="angle-range"
                onChange={changeRotation}
                type="range"
                min={0}
                max={2 * Math.PI}
                step={Math.PI / 180}
                value={rotation}
                className="w-full h-2 mt-2  bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"
              ></input>
            </div>
          </div>
        </div>
        <div className="">
          <View
            r={r}
            angle={angle}
            n={n}
            alpha={alpha}
            rotation={rotation}
            rootx={rootx}
            rooty={rooty}
          />
        </div>
      </div>
    </>
  );
}

export default Collatz;
