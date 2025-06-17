export default function Intro() {
  return (
    <div className="flex items-center justrify-center flex-col text-center pt-20 pb-6">
      <h1 className="text-4xl md:text-6xl dark:text-white mb:1 md:mb-3 font-bold">
        Aisling Heanue
      </h1>
      <p className="text-base md:text-xl mb-3 font-medium">
        Software Developer
      </p>
      <p className="max-w-5xl mb-6">
        Currently working working as a Cloud Engineer for HP Enterprise.
        I have experience with a wide variety of languages and am always eager to learn more.
        I have experience writing and deploying Go services to an AWS production environment using
        Docker and Kubernetes, as well as Terraform and other IaC tools.
      </p>
    </div>
  );
}
