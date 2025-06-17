'use client'

import dynamic from 'next/dynamic'

const Controller = dynamic(() => import('../../components/games/collatz/controller'), { ssr: false })

export function ControllerClient() {
  return <Controller />
}
