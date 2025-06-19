import CanvasView from "../../components/games/canvasView";
import Collatz from "../../components/games/collatz/controller";
import View from "../../components/games/collatz/view";
import Layout from "../../components/layout/layout";

export default function Page() {
  return (
    <Layout scrollable={false}>
      <CanvasView
        title="Collatz Tree Visualizer"
        source="https://github.com/AislingHeanue/aisl.ing/tree/master/src/app/components/games/collatz"
        controller={<Collatz />}
        game={<View />}
      />
    </Layout>
  )
}
