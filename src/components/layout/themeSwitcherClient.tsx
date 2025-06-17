"use client"

import dynamic from 'next/dynamic'

const ThemeSwitcher = dynamic(() => import('./themeSwitcher'), { ssr: false })

export default function ThemeSwitcherClient() {
  return <ThemeSwitcher />
}
