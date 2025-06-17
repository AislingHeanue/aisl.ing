'use client'

import dynamic from 'next/dynamic'

const WasmCanvas = dynamic(() => import('../../components/games/wasmLoader'), { ssr: false })

export function GameClient() {
  return <WasmCanvas game="life" />
}
