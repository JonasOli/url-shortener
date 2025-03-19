import dayjs from 'dayjs';
import { cookies } from 'next/headers';

type Url = {
  id: number;
  original: string;
  short: string;
  visit_count: number;
  created_at: string;
};

export default async function Dashboard() {
  async function listShortenUrls(): Promise<Url[] | null> {
    const cookieStore = await cookies();
    const sessionId = cookieStore.get('session-id');
    const res = await fetch('http://localhost:8000/urls/list', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Cookie: `${sessionId?.name}=${sessionId?.value}`,
      },
      credentials: 'include',
      cache: 'no-cache',
    });

    return res.json();
  }

  const urls = await listShortenUrls();

  return (
    <section>
      <h1>My urls</h1>

      {urls?.length && (
        <table>
          <thead>
            <tr>
              <th>Original url</th>
              <th>Shortened url</th>
              <th>Visit count</th>
              <th>Created at</th>
            </tr>
          </thead>
          <tbody>
            {urls.map((url) => {
              return (
                <tr key={url.id}>
                  <td>{url?.original.slice(0, 25)}...</td>
                  <td>
                    <a
                      rel="stylesheet"
                      href={`http://localhost:8000/${url.short}`}
                      target="_blank"
                    >
                      {`http://localhost:8000/${url.short}`}
                    </a>
                  </td>
                  <td>{url.visit_count}</td>
                  <td>{dayjs(url.created_at).format('DD/MM/YYYY HH:mm')}</td>
                </tr>
              );
            })}
          </tbody>
        </table>
      )}
    </section>
  );
}
