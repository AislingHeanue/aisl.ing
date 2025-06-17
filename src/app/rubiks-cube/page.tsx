import CanvasView from "../../components/games/canvasView";
import Layout from "../../components/layout/layout";
import { ControllerClient } from "./controllerClient";
import { GameClient } from "./gameClient";

export default function Page() {
  return (
    <Layout scrollable={false}>
      <CanvasView
        title="Rubik's Cube"
        source="https://github.com/AislingHeanue/aisl.ing/tree/master/wasm-demo"
        controller={<ControllerClient />}
        game={<GameClient />}
      />
    </Layout>
  )
}
