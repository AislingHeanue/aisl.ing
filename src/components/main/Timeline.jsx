import React from "react";
import timeline from "../../data/timeline";
import TimelineItem from "./TimelineItem";
import Title from "../common/Title";

function Timeline() {
  return (
    <div className="flex flex-col md:flex-row justify-center my-20">
      <div className="w-full lg:w-5/6">
        <Title>Education and Work</Title>
        {timeline.map((item) => (
          <TimelineItem
            key={item.title}
            year={item.year}
            title={item.title}
            location={item.location}
            GPA={item.GPA}
            details={item.details}
          />
        ))}
      </div>
    </div>
  );
}

export default Timeline;
