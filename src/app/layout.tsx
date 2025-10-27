import { PropsWithChildren } from "react";
import { cookies } from "next/headers";

export default async function RootLayout({ children }: PropsWithChildren) {
  // TODO: next/fonts

  const cookieStore = await cookies();
  const isDark = cookieStore.get("theme")?.value !== "light";
  return (
    <html lang="en" className={isDark ? "dark" : ""}>
      <head>
        <title>Aisling&apos;s Coding Portfolio</title>
      </head>
      <body>
        <div id="root" >{children}</div>
      </body>
    </html >
  )
}
