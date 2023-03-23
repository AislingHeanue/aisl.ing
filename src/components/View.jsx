import React from "react";
import Lines from "./Lines"

function View({r,angle,n}) {
    let linesAndCircles = Lines(r,Math.round(1000*angle)/1000,Math.floor(n*4000),75,200)
    return (
        <div className = "lg:w-3/4 md:w-full sm:w-full flex items-center justrify-center flex-col text-center pt-20 pb-6">
           <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="white"
                viewBox="0 0 200 250"
                // strokeWidth="0.2"

                // stroke="white dark:orange"
                className="w-200 h-250 stroke-stone-800 dark:stroke-white"
		    >
            <rect  width="200" height="250" strokeWidth="2" strokeOpacity="1" fillOpacity="0.01"/>

            {linesAndCircles["lines"].map(points => (
                <polyline points={points} fill="none" strokeWidth="0.1" strokeOpacity="0.7" />
            ))}
            
		</svg>
        </div>
    )
}

export default View;