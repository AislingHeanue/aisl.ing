import Image from "next/image"

export type PortfolioProps = {
  title: string
  summary: string
  imgUrl: string
  stack: string[]
  link: string
  internal: boolean
}

export default function PortfolioItem({ title, summary, imgUrl, stack, link, internal }: PortfolioProps) {
  return (
    <a
      href={link}
      target={internal ? "_self" : "_blank"}
      rel={internal ? "" : "noopener"}
      className="border-2 border-stone-900 dark:border-white rounded-md overflow-hidden"
    >
      <Image
        src={imgUrl}
        alt="portfolio"
        className="w-full h-52 object-cover cursor-pointer"
      />
      <div className="w-full p-4">
        <h3 className="text-lg md:text-xl mb-2 md:mb-3 dark:text-white font-semibold">
          {title}
        </h3>
        <p className="text-md mb-2 md:mb-3 dark:text-white">
          {summary}
        </p>
        <p className="flex flex-wrap gap-2 flex-row items-center justify-start text-xs md:text-sm">
          {stack.map((item) => (
            <span
              key={item}
              className="inline-block px-2 py-1 dark:text-white font-semibold border-2 border-stone-900 dark:border-white rounded-md"
            >
              {item}
            </span>
          ))}
        </p>
      </div>
    </a>
  );
}
