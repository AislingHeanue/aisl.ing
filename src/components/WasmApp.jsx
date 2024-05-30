import { useState, useEffect } from "react";

const WasmApp = () => {
  console.log("loading 2");
  useEffect(() => {
    console.log("loading 1");
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
};
export default WasmApp;
