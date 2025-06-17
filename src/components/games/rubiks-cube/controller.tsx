import { useState } from "react";

// wasm will hook into an existing JS canvas in the document calling it, this canvas
// should be called wasmCanvas
export default function RubiksCube() {
  const [shiftHeld, setShiftHeld] = useState(false);

  addEventListener("keydown", (e) => e.key == "Shift" && setShiftHeld(true));
  addEventListener("keyup", (e) => e.key == "Shift" && setShiftHeld(false));

  return (
    <>
      <p className="mb-4">
        A Rubik&apos;s Cube which I made in WebAssembly so I could learn how to
        render 3D objects with WebGL. To prevent glitchy animations from
        pressing the buttons way too fast, I made it so animations are
        queued up and executed asynchronously, in such a way that they run
        in parallel where possible, and sequentially when they would
        conflict with each other.
      </p>
      <p>
        Drag on the canvas to rotate the cube, and use the buttons or
        keyboard keys to make a move. Hold Shift to make anti-clockwise
        moves. The bottom row of buttons preforms some algorithms which
        look best when the cube is in a solved position.
      </p>
      <div className="grid grid-cols-5 w-5/6 h-fit mx-auto gap-1">
        <div className="col-span-1">
          <div className="aspect-square">
            <div className="w-full h-3/4 mr-1 mb-1 bg-stone-300 dark:bg-stone-500">
              <button id="u" className="w-full h-full hover:bg-stone-400">
                U{shiftHeld ? "'" : ""}
              </button>
            </div>
            <div className="w-full h-3/4 mr-1 mb-1 bg-stone-300 dark:bg-stone-500">
              <button id="d" className="w-full h-full hover:bg-stone-400">
                D{shiftHeld ? "'" : ""}
              </button>
            </div>
            <div className="w-full h-3/4 mr-1 mb-1 bg-stone-300 dark:bg-stone-500">
              <button id="e" className="w-full h-full hover:bg-stone-400">
                E{shiftHeld ? "'" : ""}
              </button>
            </div>
            <div className="w-full h-3/4 mr-1 mb-1 mt-5 bg-stone-300 dark:bg-stone-500">
              <button
                id="kilt"
                className="w-full h-full hover:bg-stone-400"
              >
                Kilt
              </button>
            </div>
          </div>
        </div>
        <div className="col-span-1">
          <div className="aspect-square">
            <div className="w-full h-3/4 mr-1 mb-1 bg-stone-300 dark:bg-stone-500">
              <button id="r" className="w-full h-full hover:bg-stone-400">
                R{shiftHeld ? "'" : ""}
              </button>
            </div>
            <div className="w-full h-3/4 mr-1 mb-1 bg-stone-300 dark:bg-stone-500">
              <button id="l" className="w-full h-full hover:bg-stone-400">
                L{shiftHeld ? "'" : ""}
              </button>
            </div>
            <div className="w-full h-3/4 mr-1 mb-1 bg-stone-300 dark:bg-stone-500">
              <button id="m" className="w-full h-full hover:bg-stone-400">
                M{shiftHeld ? "'" : ""}
              </button>
            </div>
            <div className="w-full h-3/4 mr-1 mb-1 mt-5 bg-stone-300 dark:bg-stone-500">
              <button
                id="superflip"
                className="w-full h-full hover:bg-stone-400"
              >
                Super Flip
              </button>
            </div>
          </div>
        </div>
        <div className="col-span-1">
          <div className="aspect-square">
            <div className="w-full h-3/4 mr-1 mb-1 bg-stone-300 dark:bg-stone-500">
              <button id="f" className="w-full h-full hover:bg-stone-400">
                F{shiftHeld ? "'" : ""}
              </button>
            </div>
            <div className="w-full h-3/4 mr-1 mb-1 bg-stone-300 dark:bg-stone-500">
              <button id="b" className="w-full h-full hover:bg-stone-400">
                B{shiftHeld ? "'" : ""}
              </button>
            </div>
            <div className="w-full h-3/4 mr-1 mb-1 bg-stone-300 dark:bg-stone-500">
              <button id="s" className="w-full h-full hover:bg-stone-400">
                S{shiftHeld ? "'" : ""}
              </button>
            </div>
            <div className="w-full h-3/4 mr-1 mb-1 mt-5 bg-stone-300 dark:bg-stone-500">
              <button
                id="cubecube"
                className="w-full h-full hover:bg-stone-400"
              >
                Cube in a Cube
              </button>
            </div>
          </div>
        </div>
        <div className="col-span-1">
          <div className="aspect-square">
            <div className="w-full h-3/4 mr-1 mb-1 bg-stone-300 dark:bg-stone-500">
              <button id="x" className="w-full h-full hover:bg-stone-400">
                x{shiftHeld ? "'" : ""}
              </button>
            </div>
            <div className="w-full h-3/4 mr-1 mb-1 bg-stone-300 dark:bg-stone-500">
              <button id="y" className="w-full h-full hover:bg-stone-400">
                y{shiftHeld ? "'" : ""}
              </button>
            </div>
            <div className="w-full h-3/4 mr-1 mb-1 bg-stone-300 dark:bg-stone-500">
              <button id="z" className="w-full h-full hover:bg-stone-400">
                z{shiftHeld ? "'" : ""}
              </button>
            </div>
            <div className="w-full h-3/4 mr-1 mb-1 mt-5 bg-stone-300 dark:bg-stone-500">
              <button
                id="checkerboard"
                className="w-full h-full hover:bg-stone-400"
              >
                Checker Board
              </button>
            </div>
          </div>
        </div>
        <div className="col-span-1">
          <div className="aspect-square">
            <div className="w-full h-3/4 mr-1 mb-1 bg-stone-300 dark:bg-stone-500">
              <button
                id="shuffle"
                className="w-full h-full hover:bg-stone-400"
              >
                Shuffle
              </button>
            </div>
            <div className="w-full h-3/4 mr-1 mb-1 bg-stone-300 dark:bg-stone-500">
              <button
                id="reset"
                className="w-full h-full hover:bg-stone-400"
              >
                Reset
              </button>
            </div>
            <div className="w-full h-3/4 mr-1 mb-1 bg-stone-300 dark:bg-stone-500">
              <button
                id="reset-camera"
                className="w-full h-full hover:bg-stone-400"
              >
                Reset Camera
              </button>
            </div>
            <div className="w-full h-3/4 mr-1 mb-1 mt-5 bg-stone-300 dark:bg-stone-500">
              <button
                id="snake"
                className="w-full h-full hover:bg-stone-400"
              >
                Snake
              </button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};
