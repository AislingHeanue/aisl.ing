import CanvasView from "../../components/games/canvasView";
import Controller from "../../components/games/mandelbrot/controller";
import WasmCanvas from "../../components/games/wasmLoader";
import Layout from "../../components/layout/layout";

export default function Page() {
  return (
    <Layout scrollable={false}>
      <CanvasView
        title="Mandelbrot Set"
        source="https://github.com/AislingHeanue/aisl.ing/tree/master/wasm-demo"
        controller={<Controller />}
        game={<WasmCanvas game="mandelbrot" />}
      />
    </Layout>
  )
}
