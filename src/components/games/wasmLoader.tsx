/* eslint-disable react-hooks/exhaustive-deps */
import { useEffect } from "react";

export default function WasmCanvas({ game }: { game: string }) {
  let wasmLoaded = false;
  useEffect(() => {
    if (!wasmLoaded) {
      wasmLoaded = true;
      const loadWasm = async () => {
        // Load wasm_exec.js dynamically
        const wasmPromise = new Promise((resolve, reject) => {
          const script = document.createElement("script");
          script.src = "../../wasm_exec.js";
          script.onload = resolve;
          script.onerror = reject;
          document.body.appendChild(script);
        })
        const codeBufferPromise = fetch("/demo.wasm").then((response) => response.arrayBuffer());
        const buffer = await Promise.all([wasmPromise, codeBufferPromise]).then((values) => (values[1]))

        const go = new Go()
        go.argv = [game];
        const { instance } = await WebAssembly.instantiate(
          buffer,
          go.importObject
        );
        go.run(instance);
      };
      loadWasm();
    }
  }, []);

  return (
    <>
      <canvas
        id="wasm-canvas"
        className="grow border-4 max-h-full max-w-full overflow-auto border-stone-300 dark:border-stone-600 rendering-pixelated"
      />
    </>
  );
};
