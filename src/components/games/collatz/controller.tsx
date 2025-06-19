"use client"
import { ChangeEventHandler } from "react";
import useCollatzStateStore from "./store";

export type CollatzProps = {
  changeRotation: ChangeEventHandler<HTMLInputElement>
}

export default function Collatz() {
  const state = useCollatzStateStore();

  return (
    <>
      <p className="mb-4">
        The Collatz conjecture states that preforming the operation
        &quot;multiply n by 3 and add 1 if it is odd, or half it if is even&quot; will
        always eventually reach the final value of 1 (before looping
        indefinitely between 1, 4 and 2). This has been shown to be true for
        many values of n, but the general conjecture remains unproven. This
        tree diagram represents the possible sets of operations a sequence
        can take to get from its starting value to 1 (which is specified as
        the root node) for a set list of values of n (from 2 to N). A right
        turn in the direction facing away from 1 (the root node) indicates
        &quot;n -&gt;2n&quot;, and a left turn indicates &quot;n -&gt; (2n-1)/3&quot;.
      </p>

      <div className="mt-2 flex-col md:flex-row items-center">
        <div className="grid grid-cols-2 lg:grid-cols-1 gap-4">
          <label className="block md:text-sm text-xs text-stone-900 dark:text-white">
            Branching Angle (0° to 90°)
          </label>
          <input
            id="angle-range"
            onChange={(e) => {
              state.changeAngle(+e.target.value)
            }}
            type="range"
            min="0"
            max={Math.PI / 2}
            step={Math.PI / 180}
            value={state.angle}
            className="w-full h-2 mt-1  bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"
          ></input>
          <label className="block md:text-sm text-xs text-stone-900 dark:text-white">
            Line Length (0 to 7)
          </label>
          <input
            id="length-range"
            onChange={(e) => {
              state.changeR(+e.target.value)
            }}
            type="range"
            min="0"
            max="7"
            value={state.r}
            step="any"
            className="w-full h-2 mt-2 bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"
          ></input>
          <label className="block md:text-sm text-xs text-stone-900 dark:text-white">
            N (0 to 4000)
          </label>
          <input
            id="n-range"
            onChange={(e) => {
              state.changeN(Math.floor(+e.target.value * 4000))
            }}
            type="range"
            min="0"
            max="1"
            step="any"
            value={state.n / 4000}
            className="w-full h-2 mt-2 bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"
          ></input>
          <label className="block md:text-sm text-xs text-stone-900 dark:text-white">
            Opacity (0 to 1)
          </label>
          <input
            id="alpha-range"
            onChange={(e) => {
              state.changeAlpha(+e.target.value)
            }}
            type="range"
            min="0"
            max="1"
            step="any"
            value={state.alpha}
            className="w-full h-2 mt-2 bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"
          ></input>
          <label className="block md:text-sm text-xs text-stone-900 dark:text-white">
            Translate X
          </label>
          <input
            id="root-x-range"
            onChange={(e) => {
              state.changeRootX(Math.floor(+e.target.value * 200))
            }}
            type="range"
            min="0"
            max="1"
            step="any"
            value={state.rootX / 200}
            className="w-full h-2 mt-2 bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"
          ></input>
          <label className="block md:text-sm text-xs text-stone-900 dark:text-white">
            Translate Y
          </label>
          <input
            id="root-y-range"
            onChange={(e) => {
              state.changeRootY(250 - Math.floor(250 * (+e.target.value)))
            }}
            type="range"
            min="0"
            max="1"
            step="any"
            value={1 - state.rootY / 250}
            className="w-full h-2 mt-2 bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"
          ></input>
          <label className="block md:text-sm text-xs text-stone-900 dark:text-white">
            Global Rotation
          </label>
          <input
            id="rotation-range"
            onChange={(e) => {
              state.changeRotation(+e.target.value)
            }}
            type="range"
            min={0}
            max={2 * Math.PI}
            step={Math.PI / 180}
            value={state.rotation}
            className="w-full h-2 mt-2  bg-stone-200 rounded-lg appearance-none cursor-pointer dark:bg-stone-700"
          ></input>
        </div>
      </div>
    </>
  );
}
