import { LineChart, Line } from 'recharts';
import { useCPUUsageHistory } from "./api";

const RenderLineChart = () => {

    const {data} = useCPUUsageHistory();
    return (
      <LineChart width={400} height={400} data={data}>
        <Line type="monotone" dataKey="percent" stroke="#8884d8" strokeWidth={2}/>
      </LineChart>
    )
  };
  
const CPUPage = () => {
    return (
        <div className="flex w-screen h-screen bg-black text-white">
            <div className='flex flex-col w-1/2 gap-5 h-screen m-auto'>
                CPU Usage History
                <div>
                    <RenderLineChart />
                </div>
            </div>
        </div>
      )
}

export default CPUPage;