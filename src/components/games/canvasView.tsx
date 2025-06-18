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
      <div className="lg:h-screen lg:w-screen lg:flex">
        <div className="grid lg:grid-cols-3 md:grid-cols-1 lg:w-2/3 md:w-5/6 mx-auto gap-4">
          <div className="col-span-1 grid grid-cols-3 h-fit">
            <div className="col-span-2">
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
            <div className="lg:mb-30 lg:scrollbar lg:scrollbar-thumb-stone-600 lg:scrollbar-track-stone-700 lg:pr-3 lg:overflow-y-auto col-span-3 align-top" >
              {controller}
            </div>
          </div>
          <div className="flex aspect-square lg:mt-20 lg:pb-30 lg:pr-30 lg:col-span-2 lg:max-h-screen">
            {game}
          </div>
        </div>
      </div >
    </>
  );
};
