"use client"

import { useEffect, useState } from "react";

// wasm will hook into an existing JS canvas in the document calling it, this canvas
// should be called wasmCanvas
export default function Controller() {
  const [_, setShiftHeld] = useState(false);

  useEffect(() => {
    addEventListener("keydown", (e) => e.key == "Shift" && setShiftHeld(true));
    addEventListener("keyup", (e) => e.key == "Shift" && setShiftHeld(false));
  })

  return (
    <>
      <p className="mb-4">Game of Life description</p>
      <p>Another paragraph </p>
    </>
  );
};
