import { apiUrl } from "@/consts";
import type { Url } from "@/models/url";
import axios from "axios";

async function createShortUrl(longUrl: string): Promise<void> {
  return await axios.post(`${apiUrl}/url/shorten`, { url: longUrl });
}

async function listUrls(): Promise<Url[]> {
  return (await axios.get(`${apiUrl}/url/list`)).data;
}

export { createShortUrl, listUrls };
