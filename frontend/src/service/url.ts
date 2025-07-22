import type { Url } from "@/models/url";
import axios from "axios";

async function createShortUrl(longUrl: string): Promise<void> {
  return await axios.post("http://localhost:8000/url/shorten", { url: longUrl });
}

async function listUrls(): Promise<Url[]> {
  return (await axios.get("http://localhost:8000/url/list")).data;
}

export { createShortUrl, listUrls };
