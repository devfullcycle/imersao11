import axios from "axios";

export const httpAdmin = axios.create({
  baseURL: "http://localhost:8000",
});
