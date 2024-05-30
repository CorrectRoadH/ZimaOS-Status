"use client";
import { AreaChart, Area, XAxis, YAxis } from 'recharts';
import { useCPUUsageHistory } from "@/lib/api";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

interface RenderLineChartProps{
  duration: number;
}
const RenderLineChart = ({duration}:RenderLineChartProps) => {

    const {data} = useCPUUsageHistory(duration);
    return (
      <AreaChart
        width={800}
        height={400}
        data={data}
        margin={{
          top: 10,
          right: 30,
          left: 0,
          bottom: 0,
        }}
      >
          {/* <XAxis dataKey="timestamp" /> */}
          <YAxis />
          <Area type="monotone" dataKey="percent" stroke="#8884d8" fill="#8884d8" />
      </AreaChart>
    )
  };
  
export default function Page() {
    return (
        <div className="flex w-screen h-screen bg-black text-white">
            <div className='flex flex-col w-1/2 gap-5 h-screen m-auto'>
                CPU Usage History

                <Tabs defaultValue="account" className="w-full">

                  <TabsList  className="w-full"> 
                    <TabsTrigger value="min">15 Mins</TabsTrigger>
                    <TabsTrigger value="hour">1 Hour</TabsTrigger>
                    <TabsTrigger value="day">1 Days</TabsTrigger>
                    <TabsTrigger value="week">1 Week</TabsTrigger>
                  </TabsList>

                  <TabsContent value="min">
                    <RenderLineChart duration={15} />
                  </TabsContent>
                  <TabsContent value="hour">
                    <RenderLineChart duration={60} />
                  </TabsContent>
                  <TabsContent value="day">
                    <RenderLineChart duration={720} />
                  </TabsContent>
                  <TabsContent value="week">
                    <RenderLineChart duration={5040} />
                  </TabsContent>
                </Tabs>

                <div>
                  <div>1 Min Avg Load</div> 
                  <div>3 Min Avg Load</div> 
                  <div>5 Min Avg Load</div> 
                </div>
            </div>
        </div>
      )
}

