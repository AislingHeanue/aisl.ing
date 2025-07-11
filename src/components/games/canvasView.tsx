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
      <div className="lg:w-[59.5%] sm:w-11/12 max-sm:w-full mx-auto lg:h-full">
        <div className="grid lg:grid-cols-3 lg:grid-rows-9 md:grid-cols-1 gap-x-4 h-full pt-2">
          <div className="col-span-1 grid grid-cols-5 lg:grid-rows-2 h-full lg:row-span-1 overflow-y-auto scrollbar scrollbar-thumb-stone-600 scrollbar-track-stone-700">
            <div className="col-span-3 h-fit overflow-x-auto scrollbar scrollbar-thumb-stone-600 scrollbar-track-stone-700 overflow-y-clip">
              <Title>{title}</Title>
            </div>
            <div className="col-span-2 mt-2 lg:row-span-1">
              <Link href="/">
                <button
                  type="button"
                  className="mb-2 fort-semibold text-white bg-stone-600 dark:bg-white dark:text-stone-900 rounded w-full"
                >
                  Back Home
                </button>
              </Link>
              <a
                href={source}
                target="_blank"
                rel="noopener"
              >
                <button
                  type="button"
                  className="mb-2 fort-semibold text-white bg-stone-600 dark:bg-white dark:text-stone-900 rounded pl-2 pr-2 w-full"
                >
                  Source Code
                </button>
              </a>
            </div>
          </div >
          <div className=" w-full col-span-1 col-start-1 h-full lg:row-span-10">
            <div className="h-full w-full lg:scrollbar lg:scrollbar-thumb-stone-600 lg:scrollbar-track-stone-700 lg:overflow-x-hidden lg:pr-3 lg:overflow-y-auto col-span-3" >
              {controller}
            </div>
          </div>
          <div className="lg:col-span-2 lg:row-span-8 lg:row-start-2 lg:col-start-2 flex w-full max-h-full aspect-square">
            {game}
          </div>
        </div>
      </div>
    </>
  );
};
