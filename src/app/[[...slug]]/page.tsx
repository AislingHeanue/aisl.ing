import Layout from "../../components/layout/layout";
import Intro from "../../components/main/intro";
import Portfolio from "../../components/main/portfolio";
import Timeline from "../../components/main/timeline";

export function generateStaticParams() {
  return [{ slug: [''] }]
}

export default function Page() {
  return (
    <Layout scrollable={true}>
      <div className="max-w-5xl w-11/12 mx-auto">
        <Intro />
        <Portfolio />
        <Timeline />
      </div>
    </Layout >
  )
}

