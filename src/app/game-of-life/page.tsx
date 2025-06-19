import CanvasView from "../../components/games/canvasView";
import Controller from "../../components/games/game-of-life/controller";
import WasmCanvas from "../../components/games/wasmLoader";
import Layout from "../../components/layout/layout";

export default function Page() {
  return (
    <Layout scrollable={false}>
      <CanvasView
        title="Conway's Game of Life"
        source="https://github.com/AislingHeanue/aisl.ing/tree/master/wasm-demo"
        controller={<Controller />}
        game={<WasmCanvas game="life" />}
      />
    </Layout>
  )
}
