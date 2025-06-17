import { PropsWithChildren } from "react";

export default function Title({ children }: PropsWithChildren) {
  return (
    <h1
      className="text-2xl font-bold underline underline-offset-8 decoration-4 mb-5 text-stone-900 dark:text-white"
    >
      {" "}
      {children}
      {" "}
    </h1>
  );
}
