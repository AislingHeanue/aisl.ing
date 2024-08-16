import { useState, useEffect } from "react";
import Title from "../common/Title";
import Footer from "../common/Footer";
import WasmCanvas from "../common/WasmCanvas";

// wasm will hook into an existing JS canvas in the document calling it, this canvas
// should be called wasmCanvas
const GameOfLife = () => {
  // const [dimension, setDimension] = useState(3);

  const [shiftHeld, setShiftHeld] = useState(false);

  addEventListener("keydown", (e) => e.key == "Shift" && setShiftHeld(true));
  addEventListener("keyup", (e) => e.key == "Shift" && setShiftHeld(false));

  return (
    <>
      <div className="h-full w-full lg:flex justify-center">
        <div className="grid lg:grid-cols-3 md:grid-cols-1 w-full h-fit lg:w-11/12 lg:pl-[11rem] md:w-5/6 mx-auto gap-4 ">
          <div className="col-span-1 lg:max-h-0">
            {" "}
            {/* Prevents a really dumb overscroll issue that I do not understand */}
            <Title>Conway's Game of Life</Title>
            <button
              type="button"
              className="mb-2 fort-semibold text-white bg-stone-900 dark:bg-white dark:text-stone-900 rounded ml-1 pl-2 pr-2"
            >
              <a href="/">Back to homepage</a>
            </button>
            <p className="mb-4">Game of Life description</p>
            <p>Another paragraph</p>
          </div>

          {/* <div className="mt-2 flex-col md:flex-row items-center">
              <div className="grid grid-cols-2 lg:grid-cols-1 gap-4">
                <label className="block md:text-sm text-xs text-stone-900 dark:text-white">
                  Cube dimension
                </label>
                <input
                  id="dimension"
                  onChange={(e) => setDimension(e.target.value)}
                  type="range"
                  min="1"
                  max="10"
                  step="1"
                  value={dimension}
                  className="w-full h-2 mt-1  bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"
                ></input>
              </div>
            </div> */}
          <div className="lg:col-span-2 lg:max-h-screen">
            <WasmCanvas game="life" />
          </div>
        </div>
        <div className="lg:absolute lg:bottom-0 ">
          <Footer />
        </div>
      </div>
    </>
  );
};
export default GameOfLife;
