import { useMutation, useQuery } from "@tanstack/react-query";
import { Copy } from "lucide-react";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { Button } from "./components/ui/button";
import { Input } from "./components/ui/input";
import { apiUrl } from "./consts";
import { queryClient } from "./main";
import { createShortUrl, listUrls } from "./service/url";

function App() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<{ url: string }>();

  const mutation = useMutation({
    mutationFn: createShortUrl,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["urls"] });
    },
  });

  const { data: urls, isLoading } = useQuery({
    initialData: [],
    queryKey: ["urls"],
    queryFn: listUrls,
  });

  const handleCopy = async (textToCopy: string) => {
    try {
      await navigator.clipboard.writeText(textToCopy);
      toast("Copied to clipboard!");
    } catch (err) {
      console.error("Failed to copy: ", err);
    }
  };

  return (
    <div className="min-h-screen flex flex-col items-center justify-center">
      <h1 className="text-5xl font-bold mb-10">URL shortener</h1>

      <form
        className="flex flex-row gap-2 w-full max-w-md"
        onSubmit={handleSubmit((data) => {
          mutation.mutate(data.url);
        })}
      >
        <div className="w-full">
          <Input
            type="text"
            placeholder="Long url"
            {...register("url", { required: true })}
          />
          {errors.url && (
            <p className="text-red-700 text-sm mt-2">Url is required.</p>
          )}
        </div>
        <Button type="submit" variant="outline">
          Shorten
        </Button>
      </form>

      {isLoading ? (
        <p>loading...</p>
      ) : (
        <table className="mt-10 w-200 table-auto text-lg text-left rtl:text-right text-gray-500 dark:text-gray-400">
          <thead>
            <tr>
              <th>Original URL</th>
              <th>Short URL</th>
              <th>Created At</th>
              {/* <th>Click Count</th> */}
            </tr>
          </thead>
          <tbody>
            {urls?.map((url) => (
              <tr key={url.short_code}>
                <td>{url.original_url.slice(0, 25)}...</td>
                <td className="flex place-content-between">
                  <a
                    href={`${apiUrl}/${url.short_code}`}
                    target="_blank"
                    rel="noopener noreferrer"
                  >
                    {url.short_code}
                  </a>
                  <Button
                    variant="secondary"
                    size="icon"
                    className="size-8 cursor-pointer mr-10"
                    type="button"
                    onClick={() => {
                      handleCopy(`${apiUrl}/${url.short_code}`);
                    }}
                  >
                    <Copy />
                  </Button>
                </td>
                <td>{new Date(url.created_at).toLocaleString()}</td>
                {/* <td>{url.click_count}</td> */}
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}

export default App;
