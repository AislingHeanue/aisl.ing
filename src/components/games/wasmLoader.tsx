"use client"
/* eslint-disable react-hooks/exhaustive-deps */
import { useEffect, useRef } from "react";

declare global {
  function resizeCanvas(scale: number, square: boolean): { height: number, width: number }
  function setupMultipleCanvases(
    webglCanvas: HTMLCanvasElement,
    zoomCanvas: HTMLCanvasElement,
    customHeight: number,
    customWidth: number
  ): {
    gl: WebGLRenderingContext,
    zoomContext: CanvasRenderingContext2D,
    displayContext: CanvasRenderingContext2D
  }
  function setupCanvas(
    customHeight: number,
    customWidth: number
  ): {
    gl: WebGLRenderingContext
  }
  function canvasEventListener(event: string, listener: EventListener): void
}

export default function WasmCanvas({ game }: { game: string }) {
  let wasmLoaded = false;
  const canvasRef = useRef<HTMLCanvasElement>(null);

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


  function resizeCanvas(scale: number, square: boolean) {
    const mainCanvas = canvasRef.current!
    const style = window.getComputedStyle(mainCanvas)
    const displayedHeight = +style.height.split("px")[0]
    const displayedWidth = +style.width.split("px")[0]

    let height = displayedHeight
    let width = displayedWidth

    if (scale != 0) {
      const pixelRatio = window.devicePixelRatio
      height = Math.floor(pixelRatio * displayedHeight)
      width = Math.floor(pixelRatio * displayedWidth)
    }

    if (square) {
      const minDimension = Math.min(width, height)
      height = minDimension
      width = minDimension
    }

    mainCanvas.height = height
    mainCanvas.width = width

    return { height: height, width: width }
  }

  function setupMultipleCanvases(
    webglCanvas: HTMLCanvasElement,
    zoomCanvas: HTMLCanvasElement,
    customHeight: number,
    customWidth: number
  ) {
    const mainCanvas = canvasRef.current!

    webglCanvas.height = customHeight
    webglCanvas.width = customWidth

    zoomCanvas.height = customHeight
    zoomCanvas.width = customWidth

    const gl = webglCanvas.getContext("webgl", { "alpha": false })
    const zoomCtx = zoomCanvas.getContext("2d", { "alpha": false })
    const displayCtx = mainCanvas.getContext("2d", { "alpha": false })
    displayCtx!.imageSmoothingEnabled = false

    gl?.viewport(0, 0, customWidth, customHeight)

    return { gl: gl!, zoomContext: zoomCtx!, displayContext: displayCtx! }
  }

  function setupCanvas() {
    const mainCanvas = canvasRef.current!
    const gl = mainCanvas.getContext("webgl", { "alpha": false })
    gl?.viewport(0, 0, mainCanvas.width, mainCanvas.height)

    return { gl: gl! }
  }

  function canvasEventListener(event: string, listener: EventListener) {
    const mainCanvas = canvasRef.current!
    mainCanvas.addEventListener(event, listener)
  }

  useEffect(() => {
    window.resizeCanvas = resizeCanvas
    window.setupMultipleCanvases = setupMultipleCanvases
    window.setupCanvas = setupCanvas
    window.canvasEventListener = canvasEventListener
    if (!wasmLoaded) {
      wasmLoaded = true;
      loadWasm()
    }
  }, []);


  return (
    <>
      <canvas
        className="border-4 border-stone-300 dark:border-stone-600 rendering-pixelated max-w-full box-content"
        ref={canvasRef}
      />
    </>
  );
}

