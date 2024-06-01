import { useState, useEffect } from "react";
import Title from "./Title";
import Footer from "./Footer";

// wasm will hook into an existing JS canvas in the document calling it, this canvas
// should be called wasmCanvas
const WasmApp = () => {
  const [red, setRed] = useState(255);
  const [green, setGreen] = useState(0);
  const [blue, setBlue] = useState(255);

  useEffect(() => {
    const loadWasm = async () => {
      // Load wasm_exec.js dynamically
      await new Promise((resolve, reject) => {
        const script = document.createElement("script");
        script.src = "../../wasm_exec.js";
        script.onload = resolve;
        script.onerror = reject;
        document.body.appendChild(script);
      });

      const go = new Go();
      const response = await fetch("/demo.wasm");
      console.log(response);
      const buffer = await response.arrayBuffer();
      const { instance } = await WebAssembly.instantiate(
        buffer,
        go.importObject
      );
      go.run(instance);
    };
    loadWasm();
  }, []);

  return (
    <>
      <div className="h-full w-full lg:flex justify-center">
        <div className="grid lg:grid-cols-3 md:grid-cols-1 w-full h-fit lg:w-11/12 lg:pl-[11rem] md:w-5/6 mx-auto gap-4 ">
          <div className="col-span-1">
            <Title>Go Wasm Demo</Title>
            <button
              type="button"
              className="mb-2 fort-semibold text-white bg-stone-900 dark:bg-white dark:text-stone-900 rounded ml-1 pl-2 pr-2"
            >
              <a href="/">Back to homepage</a>
            </button>
            <p className="mb-4">
              Demonstrate some basic functionality of using Go to create images
              in the browser.
            </p>

            <div className="mt-2 flex-col md:flex-row items-center">
              <div className="grid grid-cols-2 lg:grid-cols-1 gap-4">
                <label className="block md:text-sm text-xs text-stone-900 dark:text-white">
                  Red
                </label>
                <input
                  id="red"
                  onChange={(e) => setRed(e.target.value)}
                  type="range"
                  min="0"
                  max="255"
                  step="1"
                  value={red}
                  className="w-full h-2 mt-1  bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"
                ></input>
              </div>
            </div>
            <div className="mt-2 flex-col md:flex-row items-center">
              <div className="grid grid-cols-2 lg:grid-cols-1 gap-4">
                <label className="block md:text-sm text-xs text-stone-900 dark:text-white">
                  Green
                </label>
                <input
                  id="green"
                  onChange={(e) => setGreen(e.target.value)}
                  type="range"
                  min="0"
                  max="255"
                  step="1"
                  value={green}
                  className="w-full h-2 mt-1  bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"
                ></input>
              </div>
            </div>
            <div className="mt-2 flex-col md:flex-row items-center">
              <div className="grid grid-cols-2 lg:grid-cols-1 gap-4">
                <label className="block md:text-sm text-xs text-stone-900 dark:text-white">
                  Blue
                </label>
                <input
                  id="blue"
                  onChange={(e) => setBlue(e.target.value)}
                  type="range"
                  min="0"
                  max="255"
                  step="1"
                  value={blue}
                  className="w-full h-2 mt-1  bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"
                ></input>
              </div>
            </div>
          </div>
          <div className="lg:col-span-2 lg:max-h-screen">
            <div className="flex flex-col max-h-full aspect-square lg:pt-[5rem] lg:pr-[11rem] lg:pb-[6rem]">
              <canvas
                id="wasm-canvas"
                className="w-100 flex-grow border-4 max-h-full max-w-full overflow-auto border-stone-300 dark:border-stone-600"
              />
            </div>
          </div>
        </div>
        <div className="lg:absolute lg:bottom-0 ">
          <Footer />
        </div>
      </div>
    </>
  );
};
export default WasmApp;
