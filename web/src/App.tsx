import { useUsage } from "./api";

interface PerformanceCardProps {
  title: string;
  value: string;
}

function PerformanceCard({title,value}:PerformanceCardProps) {
  return (
    <div className='flex flex-col w-48 h-64 bg-slate-800 rounded-2xl'>
      <div className='mx-auto'>{title}</div>
      <div className='m-auto font-black text-6xl'>
        {value}
      </div>
    </div>
  )
}

function App() {

  const { data, isLoading } = useUsage();
  return (
    <div className="flex w-screen h-screen bg-black text-white">
      <div className='flex flex-col w-1/2 gap-5 h-screen m-auto'>
        <div>
          性能
        </div>
        
        <div className='flex gap-4 flex-wrap'>
          <PerformanceCard title='CPU占用率' value={`${data?.cpu}%`} />

          <PerformanceCard title='内存占用率' value={`${data?.memory}%`} />
        </div>
      </div>
    </div>
  );
}

export default App;
