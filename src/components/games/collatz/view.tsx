"use client"
import Lines from "./lines"
import useCollatzStateStore from "./store"

export default function View() {
  const state = useCollatzStateStore()
  console.log(state.n)
  const lines = Lines(state);
  const alpha = state.alpha

  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="white"
      viewBox="0 0 250 250"
      className="flex aspect-square stroke-stone-800 dark:stroke-white"
    >
      <rect
        width="250"
        height="250"
        strokeWidth="2"
        strokeOpacity="1"
        fillOpacity="0.01"
      />

      {lines.map((points, index) => (
        <polyline
          key={index}
          points={points.toString()}
          fill="none"
          strokeWidth="0.1"
          strokeOpacity={alpha}
        />
      ))}
    </svg>
    // </div>
  );
}

