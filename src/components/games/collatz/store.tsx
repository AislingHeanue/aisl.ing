import { create } from 'zustand'

export interface State {
  angle: number
  r: number
  n: number
  alpha: number
  rootX: number
  rootY: number
  rotation: number
  changeAngle: (v: number) => void
  changeR: (v: number) => void
  changeN: (v: number) => void
  changeAlpha: (v: number) => void
  changeRootX: (v: number) => void
  changeRootY: (v: number) => void
  changeRotation: (v: number) => void
}

const useCollatzStateStore = create<State>()((set) => ({
  angle: 0.15,
  r: 2.5,
  n: 800,
  alpha: 0.7,
  rootX: 100,
  rootY: 200,
  rotation: Math.PI / 2,
  changeAngle: (v: number) => set((_) => ({
    angle: v,
  })),
  changeR: (v: number) => set((_) => ({
    r: v,
  })),
  changeN: (v: number) => set((_) => ({
    n: v,
  })),
  changeAlpha: (v: number) => set((_) => ({
    alpha: v,
  })),
  changeRootX: (v: number) => set((_) => ({
    rootX: v,
  })),
  changeRootY: (v: number) => set((_) => ({
    rootY: v,
  })),
  changeRotation: (v: number) => set((_) => ({
    rotation: v,
  })),
}))

export default useCollatzStateStore
