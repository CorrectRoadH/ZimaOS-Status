import { LineChart, Line } from 'recharts';
import { useCPUUsageHistory } from "@/lib/api";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

const RenderLineChart = () => {

    const {data} = useCPUUsageHistory();
    return (
      <LineChart width={400} height={400} data={data}>
        <Line type="monotone" dataKey="percent" stroke="#8884d8" strokeWidth={2}/>
      </LineChart>
    )
  };
  
export default function Page() {
    return (
        <div className="flex w-screen h-screen bg-black text-white">
            <div className='flex flex-col w-1/2 gap-5 h-screen m-auto'>
                CPU Usage History

                <Tabs defaultValue="account" className="w-[400px]">
                  <TabsList>
                    <TabsTrigger value="min">15 Mins</TabsTrigger>
                    <TabsTrigger value="hour">1 Hour</TabsTrigger>
                    <TabsTrigger value="day">1 Days</TabsTrigger>
                    <TabsTrigger value="week">1 Week</TabsTrigger>
                  </TabsList>

                  <TabsContent value="min">Make changes to your account here.</TabsContent>
                  <TabsContent value="hour">Change your password here.</TabsContent>
                  <TabsContent value="day">Change your password here.</TabsContent>
                  <TabsContent value="week">Change your password here.</TabsContent>
                </Tabs>

                <div>
                    {/* <RenderLineChart /> */}
                </div>
            </div>
        </div>
      )
}

