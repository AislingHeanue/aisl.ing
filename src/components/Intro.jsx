import React from "react";

function Intro() {
    return (
        <div className = "flex items-center justrify-center flex-col text-center pt-20 pb-6">
            <h1 className="text-4xl md:text-6xl dark:text-white mb:1 md:mb-3 font-bold">Aisling Heanue</h1>
            <p className="text-base md:text-xl mb-3 font-medium">Software Developer</p>
            <p className="text-sm max-w-3xl mb-6 font-bold">
                Having completed my master's in Computer Science at University College Dublin, I am now working as a cloud engineer for HP Enterprise. 
                I have strong Python, Java and NodeJS programming skills, as well as experience with Go, MATLAB and Bash scripting. 
                I am always looking for more opportunities to learn more and try new coding challenges.
            </p>
        </div>
    )
}

export default Intro;