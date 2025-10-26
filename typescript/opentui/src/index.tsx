import { render } from "@opentui/react";
import { App } from "./App.tsx";

await render(<App />, { exitOnCtrlC: true });
