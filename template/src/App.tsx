import { useEffect, useState } from "react"

declare global {
  interface Window {
    injectValue: (value: any) => void
  }
}

interface DataType {
  name: string
  createdAt: string
}

function App() {
  const [data, setData] = useState<DataType>()
  useEffect(() => {
    window.injectValue = (value: any) => {
      setData(value)
    }
  }, [])

  return (
    <div>
       <div>
        Hello, {data?.name}
       </div>
    </div>
  )
}

export default App
