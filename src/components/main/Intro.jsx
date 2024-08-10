import React from "react";

function Intro() {
  return (
    <div className="flex items-center justrify-center flex-col text-center pt-20 pb-6">
      <h1 className="text-4xl md:text-6xl dark:text-white mb:1 md:mb-3 font-bold">
        Aisling Heanue
      </h1>
      <p className="text-base md:text-xl mb-3 font-medium">
        Software Developer
      </p>
      <p className="max-w-3xl mb-6">
        Currently working working as a Cloud Engineer for HP Enterprise. I have
        strong Golang and Python programming skills, as well as experience with
        Bash, Java and Javascript. I am always looking for more opportunities to
        learn more and try new coding challenges.
      </p>
      <p>
        These are some of my projects. The first two of these are demos which
        can be run in the browser.
      </p>
    </div>
  );
}

export default Intro;
