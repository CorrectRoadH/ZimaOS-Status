import { LineChart, Line } from 'recharts';
import { useCPUUsage } from "./api";

// const data = [{name: 'Page A', uv: 400, pv: 2400, amt: 2400}];

const RenderLineChart = () => {

  const {data} = useCPUUsage();
  return (
    <LineChart width={400} height={400} data={data}>
      <Line type="monotone" dataKey="percent" stroke="#8884d8" />
    </LineChart>
  )
};

function App() {
  return (
    <div className="">
      CPU占用率
      <RenderLineChart />
    </div>
  );
}

export default App;
