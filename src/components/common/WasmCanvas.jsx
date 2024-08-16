import { useState, useEffect } from "react";

// wasm will hook into an existing JS canvas in the document calling it, this canvas
// should be called wasmCanvas
const WasmCanvas = ({ game }) => {
  // const [dimension, setDimension] = useState(3);
  let wasmLoaded = false;

  useEffect(() => {
    if (!wasmLoaded) {
      wasmLoaded = true;
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
        go.argv = [game];
        go.run(instance);
      };
      loadWasm();
    }
  }, []);

  return (
    <>
      <div className="flex flex-col max-h-full aspect-square lg:pt-[5rem] lg:pr-[11rem] lg:pb-[6rem]">
        <canvas
          id="wasm-canvas"
          className="w-100 flex-grow border-4 max-h-full max-w-full overflow-auto border-stone-300 dark:border-stone-600 rendering-pixelated"
        />
      </div>
    </>
  );
};
export default WasmCanvas;
