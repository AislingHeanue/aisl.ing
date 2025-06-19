import CanvasView from "../../components/games/canvasView";
import Controller from "../../components/games/rubiks-cube/controller";
import WasmCanvas from "../../components/games/wasmLoader";
import Layout from "../../components/layout/layout";

export default function Page() {
  return (
    <Layout scrollable={false}>
      <CanvasView
        title="Rubik's Cube"
        source="https://github.com/AislingHeanue/aisl.ing/tree/master/wasm-demo"
        controller={<Controller />}
        game={<WasmCanvas game="rubiks" />}
      />
    </Layout>
  )
}
