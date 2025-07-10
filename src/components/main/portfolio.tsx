import portfolio from "../../data/portfolio";
import PortfolioItem from "./portfolioItem";
import ExpandableGrid from "./expandableGrid";


export default function Portfolio() {
  return (
    <div className="flex flex-col md:flex-row items-center justify-center">
      {/* <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4"> */}
      <ExpandableGrid items={
        portfolio.map((project) => (
          <PortfolioItem
            key={project.title}
            title={project.title}
            summary={project.summary}
            imgUrl={project.imgURL}
            stack={project.stack}
            link={project.link}
            internal={project.internal}
          />
        ))
      } />
    </div>
    // </div >
  );
}

