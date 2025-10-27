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
    // className="border-2 relative border-stone-900 dark:border-white rounded-md overflow-hidden"
    >
      <div className="border-stone-900 dark:border-white border-2 rounded-md mb-3 h-106">
        <div className="w-full relative h-52 cursor-pointer">
          <Image
            src={imgUrl}
            fill={true}
            alt="portfolio"
            className="object-cover"
          />
        </div>
        <div className="w-full p-4">
          <h3 className="text-lg md:text-xl mb-2 md:mb-3 dark:text-white font-semibold line-clamp-2">
            {title}
          </h3>
          <p className="text-md mb-2 md:mb-3 dark:text-white line-clamp-3">
            {summary}
          </p>
          <p className="flex flex-wrap gap-2 flex-row items-center justify-start text-xs md:text-sm absolute bottom-6.5">
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
      </div>
    </a>
  );
}
