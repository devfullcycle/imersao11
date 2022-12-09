import axios from "axios";

export const httpAdmin = axios.create({
  baseURL: process.env.NEXT_PUBLIC_ADMIN_API_URL,
});

export const fetcherStats = (url: string) =>
  httpStats.get(url).then((res) => res.data);

export const httpStats = axios.create({
  baseURL: process.env.NEXT_PUBLIC_STATS_API_URL,
});
