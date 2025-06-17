import CanvasView from "../../components/games/canvasView";
import Layout from "../../components/layout/layout";
import { ControllerClient } from "./controllerClient";
import { GameClient } from "./gameClient";

export default function Page() {
  return (
    <Layout scrollable={false}>
      <CanvasView
        title="Collatz Tree Visualizer"
        source="https://github.com/AislingHeanue/aisl.ing/tree/master/app/components/collatz"
        controller={<ControllerClient />}
        game={<GameClient />}
      />
    </Layout>
  )
}
