export type TimelineProps = {
  year: string
  title: string
  location: string
  GPA: string
  details: string
}

export default function TimelineItem({ year, title, location, GPA, details }: TimelineProps) {
  return (
    <ol className="flex flex-col md:flex-row relative border-l border-stone-200 dark:border-stone-700">
      <li className="mb-10 ml-4">
        <div className="absolute w-3 h-3 bg-stone-200 rounded-full mt-1.5 -left-1.5 border border-white dark:border-stone-900 dark:bg-stone-700" />
        <div className="flex flex-wrap gap-4 flex-row items-center justify-start text-xs md:text-sm">
          <span className="inline-block px-2 py-1 fort-semibold text-white bg-stone-900 dark:bg-white dark:text-stone-900 rounded-md">
            {year}
          </span>
          <h3 className="text-lg font-semibold text-stone-900 dark:text-white">
            {title}
          </h3>
          <p className="my-1 text-md font-normal leading-none text-stone-500 dark:text-stone-400">
            {location}
          </p>
          <p className="my-1 text-md font-normal leading-none text-stone-500 dark:text-stone-400">
            {GPA}
          </p>
        </div>
        <p className="my-2 text-base font-normal text-stone-600 dark:text-stone-300">
          {details}
        </p>
      </li>
    </ol>
  );
}
