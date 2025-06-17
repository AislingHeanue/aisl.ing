import React, { PropsWithChildren } from "react";

export default function RootLayout({ children }: PropsWithChildren) {
  // TODO: add icon.png and/or favicon.ico to the app folder
  // TODO: next/fonts
  return (
    <html lang="en" suppressHydrationWarning={true}>
      <head>
        <title>Aisling&apos;s Coding Portfolio</title>
        <script dangerouslySetInnerHTML={{
          __html: `
            document.documentElement.classList.toggle(
              "dark",
              localStorage.theme === "dark" || (!("theme" in localStorage)),
            );
          `,
        }} />
      </head>
      <body>
        <div id="root" >{children}</div>
      </body>
    </html >
  )
}
