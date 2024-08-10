import React from "react";
import Lines from "./Lines";

function View({ r, angle, n, alpha, rotation, rootx, rooty }) {
  let linesAndCircles = Lines(
    r,
    Math.round(1000 * angle) / 1000,
    Math.floor(n * 4000),
    rootx,
    rooty,
    rotation
  );

  return (
    <div className="w-full flex items-center justify-center flex-col text-center lg:mt-20 pb-6">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="white"
        viewBox="0 0 200 250"
        className="w-200 h-250 stroke-stone-800 dark:stroke-white"
      >
        <rect
          width="200"
          height="250"
          strokeWidth="2"
          strokeOpacity="1"
          fillOpacity="0.01"
        />

        {linesAndCircles["lines"].map((points) => (
          <polyline
            points={points}
            fill="none"
            strokeWidth="0.1"
            strokeOpacity={alpha}
          />
        ))}
      </svg>
    </div>
  );
}

export default View;
