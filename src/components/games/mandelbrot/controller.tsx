"use client"

import { useEffect, useState } from "react";
import { InlineMath, BlockMath } from 'react-katex';
import 'katex/dist/katex.min.css';

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
      <p className="mb-4">The Mandelbrot set a 2 Dimensional set of complex numbers (<InlineMath math="c" />) for which the following equation will never diverge:</p>
      <BlockMath math="z_n = \begin{cases}0\;\;\;\;\;\;\;\;\;\;\;\;\;\,n = 0\\z_{n-1}^2 + c \;\;\; n>0\\\end{cases}" />
      <p>For example, if <InlineMath math="c" /> is <InlineMath math="-1" />, this results in the sequence <InlineMath math="0,-1,0,-1 \ldots " /> which
        never escapes this small region of the graph, so the point <InlineMath math="-1 + 0i" /> is coloured black on the graph. The remaining points are typically
        coloured based on how many iterations it takes for their sequence to escape some arbitrary distance from the origin of this graph (such as <InlineMath math="|z| > 2" />). The teal-coloured
        region comprises points whose sequences grow very quickly, so they take very few iterations. Zooming in to the edge of the mandelbrot set will reveal an intricate fractal pattern that only
        reveals more and more complex patterns at higher zoom levels.
      </p>
      <p>
        Use the mouse scroll wheel or pinch gestures to zoom into the image at different points. The most interesting regions are those one the edge of any black region.
      </p>
    </>
  );
};
