import { Input } from "./components/ui/input";
import { Button } from "./components/ui/button";

function App() {
  return (
    <div>
      <h1>URL shortener</h1>

      <div className="flex w-full max-w-sm items-center gap-2">
        <Input type="text" placeholder="Long url" />
        <Button type="submit" variant="outline">
          Shorten
        </Button>
      </div>
    </div>
  );
}

export default App;
