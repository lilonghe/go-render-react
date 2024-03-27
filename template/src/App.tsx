import { useEffect, useState } from "react"

declare global {
  interface Window {
    injectData: DataType
  }
}

interface DataType {
  name: string
  createdAt: string
}

function App() {
  const [data, setData] = useState<DataType>()

  useEffect(() => {
    setTimeout(() => {
      setData(window.injectData);
    }, 300);
  }, []);

  return (
    <div>
       <div>
        Hello, {data?.name}
       </div>
    </div>
  )
}

export default App
