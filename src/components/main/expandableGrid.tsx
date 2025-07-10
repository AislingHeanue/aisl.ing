"use client"
import React, { ReactNode, useEffect, useRef, useState } from 'react';

const updateItems = () => {
  if (typeof window !== 'undefined') {
    if (window.innerWidth >= 1024) { // lg breakpoint: 1024px and up
      return 3
    } else if (window.innerWidth >= 768) { // md breakpoint: 768px and up
      return 4
    } else {
      return 2
    }
  }
  return 0
}

const ExpandableGrid = ({ items }: { items: ReactNode[] }) => {
  const [height, setHeight] = useState(0 as number | 'auto')
  const [initialItems, setInitialItems] = useState(3)
  const [isOpen, setIsOpen] = useState(false);
  const contentRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    setInitialItems(updateItems())
    window.addEventListener("resize", () => {
      setInitialItems(updateItems())
    }
    )
  }, [])

  const toggleHeight = () => {
    const el = contentRef.current;
    if (!el) return;
    setIsOpen(!isOpen)

    if (!isOpen) {
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
              height: index < initialItems ? 'auto' : height,
              transition: 'height 300ms ease',
            }}
            key={index}
            className="border-2 relative border-stone-900 dark:border-white rounded-md overflow-hidden"
          >
            {item}
          </div>
        )
      })}
      < button
        className="col-start-1 col-span-1 md:col-span-2 lg:col-start-2 lg:col-span-1 bg-stone-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-4xl"
        onClick={() => toggleHeight()}
      >
        {height != 0 ? 'Show Fewer' : 'Show More'}
      </button>
    </div >
  );
};

export default ExpandableGrid;
