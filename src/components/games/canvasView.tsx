"use client"
import { ReactNode } from "react";
import Title from "../layout/title"
import Link from "next/link";

export type CanvasViewProps = {
  title: string
  source: string
  controller: ReactNode
  game: ReactNode
}
// wasm will hook into an existing JS canvas in the document calling it, this canvas
// should be called wasmCanvas
export default function CanvasView({ title, source, controller, game }: CanvasViewProps) {
  return (
    <>
      <div className="lg:w-[58.8%] md:w-5/6 mx-auto lg:h-full"> <div className="grid lg:grid-cols-3 lg:grid-rows-10 md:grid-cols-1 gap-4 flex h-full pt-4">
        <div className="col-span-1 grid grid-cols-3 lg:grid-rows-10 h-full lg:row-span-10">
          <div className="col-span-2 h-fit lg:row-span-1">
            <Title>{title}</Title>
          </div>
          <div className="col-span-1 h-fit">
            <button
              type="button"
              className="mb-2 fort-semibold text-white bg-stone-900 dark:bg-white dark:text-stone-900 rounded ml-1 pl-2 pr-2"
            >
              <Link href="/">Back Home</Link>
            </button>
            <button
              type="button"
              className="mb-2 fort-semibold text-white bg-stone-900 dark:bg-white dark:text-stone-900 rounded ml-1 pl-2 pr-2"
            >
              <a
                href={source}
                target="_blank"
                rel="noopener"
              >Source Code</a>
            </button>
          </div>
          <div className=" w-full col-span-3 h-full overflow-hidden lg:row-span-9 bg-stone-600">
            <div className="h-full lg:scrollbar lg:scrollbar-thumb-stone-600 lg:scrollbar-track-stone-700 lg:overflow-auto lg:pr-3 lg:overflow-y-auto col-span-3" >
              {controller}
            </div>
          </div>
        </div>
        <div className="lg:col-span-2 lg:row-span-1"></div>
        <div className="lg:col-span-2 lg:row-span-9 flex w-full max-h-full aspect-square">
          {game}
        </div>
      </div>
      </div >
    </>
  );
};
