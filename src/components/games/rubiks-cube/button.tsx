export default function Button({ id, shiftHeld, name, lowercase }: { id: string, shiftHeld: boolean, name?: string, lowercase?: boolean }) {
  return (
    <div className={`w-full h-12 ${name ? "text-sm" : "text-lg"} mr-1 mb-1 bg-stone-300 dark:bg-stone-500`}>
      <button id={id} className="w-full h-full hover:bg-stone-400">
        {
          name ||
          ((lowercase ? id[0] : id[0].toUpperCase()) + id.substring(1) + (shiftHeld ? "'" : ""))}
      </button>
    </div>
  )
}
