'use client'

import dynamic from 'next/dynamic'

const View = dynamic(() => import('../../components/games/collatz/view'), { ssr: false })


export function GameClient() {
  return <View />
}
