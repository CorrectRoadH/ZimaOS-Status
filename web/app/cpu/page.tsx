"use client";
import { LineChart, Line, XAxis, YAxis } from 'recharts';
import { useCPUUsageHistory } from "@/lib/api";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

const RenderLineChart = () => {

    const {data} = useCPUUsageHistory();
    return (
      <LineChart width={800} height={400} data={data}>
          <XAxis dataKey="timestamp" />
          <YAxis />
        <Line type="monotone" dataKey="percent" stroke="#8884d8" strokeWidth={2}/>
      </LineChart>
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
                    <RenderLineChart />
                  </TabsContent>
                  <TabsContent value="hour">
                    <RenderLineChart />
                  </TabsContent>
                  <TabsContent value="day">
                    <RenderLineChart />
                  </TabsContent>
                  <TabsContent value="week">
                    <RenderLineChart />
                  </TabsContent>
                </Tabs>

                <div>
                    {/* <RenderLineChart /> */}
                </div>
            </div>
        </div>
      )
}

