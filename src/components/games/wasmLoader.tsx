"use client"
/* eslint-disable react-hooks/exhaustive-deps */
import { useEffect } from "react";

export default function WasmCanvas({ game }: { game: string }) {
  let wasmLoaded = false;

  async function loadWasm() {
    // Load wasm_exec.js dynamically
    await new Promise((resolve, reject) => {
      const script = document.createElement("script");
      script.src = "../../wasm_exec.js";
      script.onload = resolve;
      script.onerror = reject;
      document.body.appendChild(script);
    })

    const go = new Go()
    go.argv = [game];

    const { instance } = await WebAssembly.instantiateStreaming(
      fetch("/demo.wasm"),
      go.importObject
    );
    go.run(instance);
  }

  useEffect(() => {
    if (!wasmLoaded) {
      wasmLoaded = true;
      loadWasm()
    }
  }, []);

  return (
    <>
      <canvas
        id="wasm-canvas"
        className="border-4 border-stone-300 dark:border-stone-600 rendering-pixelated max-w-full"
      />
    </>
  );
};
