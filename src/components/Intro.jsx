import React from "react";

function Intro() {
    return (
        <div className = "flex items-center justrify-center flex-col text-center pt-20 pb-6">
            <h1 className="text-4xl md:text-6xl dark:text-white mb:1 md:mb-3 font-bold">Aisling Heanue</h1>
            <p className="text-base md:text-xl mb-3 font-medium">Computer Science Student</p>
            <p className="text-sm max-w-3xl mb-6 font-bold">
                I am studying Computer Science at University College Dublin, and I am on track to graduate this September. 
                I have strong Python and Java programming skills, as well as experience with JavaScript, MATLAB and Bash scripting.
                I am currently taking part in a Software Engineering internship program with Adaptemy.
            </p>
        </div>
    )
}

export default Intro;