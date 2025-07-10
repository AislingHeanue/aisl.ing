"use client"
import React, { ReactNode, useRef, useState } from 'react';

const ExpandableGrid = ({ items }: { items: ReactNode[] }) => {
  const [height, setHeight] = useState(0 as number | 'auto')
  const contentRef = useRef<HTMLDivElement>(null);


  const toggleHeight = () => {
    const el = contentRef.current;
    if (!el) return;

    if (height == 0) {
      setHeight(el.scrollHeight)
    } else {
      setHeight(0)
    }
  }



  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4" id="heightContainer">
      {items.map((item, index) => {
        return (
          <div
            ref={contentRef}
            style={{
              minHeight: height,
            }}
            key={index}
            className="
            transition-[min-height] ease-in-out duration-300
            border-2 relative border-stone-900 dark:border-white rounded-md overflow-hidden 
            h-0
            [&:nth-child(-n+4)]:md:max-lg:h-auto
            [&:nth-child(-n+3)]:lg:h-auto
            [&:nth-child(-n+2)]:sm:max-md:h-auto
            "
          >
            {item}
          </div>
        )
      })}
      < button
        className="col-start-1 col-span-1 md:col-span-2 lg:col-start-2 lg:col-span-1 bg-stone-500 hover:bg-stone-400 text-white font-bold py-2 px-4 rounded-4xl"
        onClick={() => toggleHeight()}
      >
        {height != 0 ? 'Show Fewer' : 'Show More'}
      </button>
    </div >
  );
};

export default ExpandableGrid;
