"use client"
import { useUsage } from "@/lib/api";
import { useZimaAuth } from "@/lib/auth";
import axios from "axios";
import { useEffect } from "react";

interface PerformanceCardProps {
  title: string;
  value: string;
}

function PerformanceCard({title,value}:PerformanceCardProps) {
  return (
    <div className='flex flex-col w-48 h-64 bg-slate-800 rounded-2xl p-2'>
      <div className='mx-auto'>{title}</div>
      <div className='m-auto font-black text-6xl'>
        {value}
      </div>
    </div>
  )
}
export default function Home() {

  useEffect(() => {
    const zimaAuthAxios = axios.create();
    useZimaAuth(zimaAuthAxios);  
    zimaAuthAxios.get('http://10.0.0.83/')
  }, []);

  const { data, isLoading } = useUsage();
  return (
    <div className="flex w-screen h-screen bg-black text-white">
      <div className='flex flex-col w-1/2 gap-5 h-screen m-auto'>
        <div className="font-bold text-7xl">
          Performance
        </div>
        
        <div className='flex gap-4 flex-wrap'>
          <PerformanceCard title='CPU' value={`${Math.ceil(data?.cpu?.percent||0)}%`} />

          <PerformanceCard title='Memory' value={`${Math.ceil(data?.memory?.percent||0)}%`} />
        </div>
      </div>
    </div>
  );
}
