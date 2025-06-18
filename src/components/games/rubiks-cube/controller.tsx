import { useEffect, useState } from "react";
import Button from "./button";

// wasm will hook into an existing JS canvas in the document calling it, this canvas
// should be called wasmCanvas
export default function RubiksCube() {
  const [shiftHeld, setShiftHeld] = useState(false);

  useEffect(() => {
    addEventListener("keydown", (e) => e.key == "Shift" && setShiftHeld(true));
    addEventListener("keyup", (e) => e.key == "Shift" && setShiftHeld(false));
  })


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
      <div className="grid grid-cols-5 w-full h-fit mx-auto gap-1">
        <div className="col-span-1">
          <Button id="u" shiftHeld={shiftHeld} />
          <Button id="d" shiftHeld={shiftHeld} />
          <Button id="e" shiftHeld={shiftHeld} />
          <Button id="kilt" name="Kilt" shiftHeld={shiftHeld} />
        </div>
        <div className="col-span-1">
          <Button id="r" shiftHeld={shiftHeld} />
          <Button id="l" shiftHeld={shiftHeld} />
          <Button id="m" shiftHeld={shiftHeld} />
          <Button id="s-flip" name="Superflip" shiftHeld={shiftHeld} />
        </div>
        <div className="col-span-1">
          <Button id="f" shiftHeld={shiftHeld} />
          <Button id="b" shiftHeld={shiftHeld} />
          <Button id="s" shiftHeld={shiftHeld} />
          <Button id="cubecube" name="Cube in a Cube" shiftHeld={shiftHeld} />
        </div>
        <div className="col-span-1">
          <Button id="x" lowercase={true} shiftHeld={shiftHeld} />
          <Button id="y" lowercase={true} shiftHeld={shiftHeld} />
          <Button id="z" lowercase={true} shiftHeld={shiftHeld} />
          <Button id="checker" name="Checker Board" shiftHeld={shiftHeld} />
        </div>
        <div className="col-span-1">
          <Button id="shuffle" name="Shuffle" shiftHeld={shiftHeld} />
          <Button id="reset" name="Reset" shiftHeld={shiftHeld} />
          <Button id="resetcam" name="Reset Camera" shiftHeld={shiftHeld} />
          <Button id="snake" name="Snake" shiftHeld={shiftHeld} />
        </div>
      </div>
    </>
  );
};
