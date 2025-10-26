## OpenTUI Build-First Authoring Guide (for AI + humans)

This guide teaches how to compose OpenTUI UIs quickly from primitives. It assumes your environment is already set up and focuses strictly on building screens, interactions, and patterns.

### Table of contents

- [What to build, not how to run](#what-to-build-not-how-to-run)
- [Component quick reference](#component-quick-reference)
- [Layout patterns](#layout-patterns)
- [Interactive patterns](#interactive-patterns)
- [Text and styling patterns](#text-and-styling-patterns)
- [Animation patterns](#animation-patterns)
- [Data-driven lists and collections](#data-driven-lists-and-collections)
- [Recipe library](#recipe-library)
- [Component extension (only when necessary)](#component-extension-only-when-necessary)
- [Intent → constructs mapping (AI cheatsheet)](#intent--constructs-mapping-ai-cheatsheet)
- [Quality checklist for AI](#quality-checklist-for-ai)
- [Appendix: sources referenced](#appendix-sources-referenced)

### What to build, not how to run

- Assume rendering is available; your job is to decide which components, props, and patterns to use to satisfy a UI intent.
- Prefer concise compositions using `box` for layout and small focused children.
- Favor controlled inputs and explicit focus management for interactive screens.

## Component quick reference

Short intros with micro-snippets (≤10 lines) to remember key props and usage.

### text

- Purpose: render styled text; supports inline modifiers as children.
- Essentials: `content?`, `fg?`, `bg?`, `attributes?` or place children `<span|strong|em|u|b|i|br>` inside `<text>`.

```tsx
<text>
  Plain text <strong>bold</strong> and <span fg="red">red</span>
</text>
```

### box

- Purpose: main layout container; borders, paddings, alignment.
- Essentials: `border?`, `title?`, `flexDirection?` ("row" | "column"), size (`width|height`), spacing (`padding|margin|gap`).

```tsx
<box flexDirection="row" style={{ gap: 1 }}>
  <box border title="Left" style={{ width: 24 }}>
    <text>Sidebar</text>
  </box>
  <box border title="Main" style={{ flexGrow: 1 }}>
    <text>Content</text>
  </box>
  <box border title="Right" style={{ width: 16 }} />
</box>
```

### input

- Purpose: single-line text input.
- Essentials: `placeholder?`, `focused?`, `onInput(value)`, `onSubmit(value)`.

```tsx
<box title="Name" style={{ border: true, height: 3 }}>
  <input placeholder="Type..." focused onInput={setValue} onSubmit={save} />
</box>
```

### select

- Purpose: list selection widget.
- Essentials: `options: SelectOption[]`, `focused?`, `onChange(index, option)`, sizing via surrounding `box`.

```tsx
<select
  focused
  options={[
    { name: "One", value: 1 },
    { name: "Two", value: 2 },
  ]}
  onChange={(i, opt) => setSelected(opt?.value)}
/>
```

### scrollbox

- Purpose: scrollable content area with optional styled track/arrows.
- Essentials: put content inside; control look via `style.rootOptions|wrapperOptions|viewportOptions|contentOptions|scrollbarOptions`.

```tsx
<scrollbox focused style={{ viewportOptions: { backgroundColor: "#1a1b26" } }}>
  {items.map((it, i) => (
    <text key={i}>{it}</text>
  ))}
</scrollbox>
```

### tab-select

- Purpose: tabbed selection.
- Essentials: provide `options`, handle selection change, show content based on active index.

```tsx
<tab-select options={[{ name: "Home" }, { name: "Settings" }]} />
```

### ascii-font

- Purpose: render ASCII art text with preset fonts ("tiny" | "block" | "slick" | "shade").
- Essentials: compute `width/height` via `measureText({ text, font })` and apply to style.

```tsx
<ascii-font text="ASCII" font="tiny" style={{ width: 20, height: 5 }} />
```

### inline modifiers (inside <text>)

- Use `<span>`, `<strong>`, `<em>`, `<u>`, `<b>`, `<i>`, `<br>`.

```tsx
<text>
  <strong>Bold</strong>, <em>Italic</em>, <u>Underlined</u>
  <br />
  <span fg="blue">Blue text</span>
</text>
```

## Layout patterns

### Columns and rows (flex-like)

```tsx
<box flexDirection="column" style={{ gap: 1 }}>
  <box border title="Header" style={{ height: 3 }} />
  <box flexDirection="row" style={{ gap: 1, flexGrow: 1 }}>
    <box border title="Nav" style={{ width: 24 }} />
    <box border title="Main" style={{ flexGrow: 1 }} />
  </box>
  <box border title="Footer" style={{ height: 3 }} />
</box>
```

### Split panes

```tsx
<box flexDirection="row" style={{ gap: 1 }}>
  <box border title="Left" style={{ width: "40%" }} />
  <box border title="Right" style={{ flexGrow: 1 }} />
</box>
```

### Sidebar layout

```tsx
<box flexDirection="row">
  <box border title="Sidebar" style={{ width: 24 }}>
    <text>Menu</text>
  </box>
  <box style={{ flexGrow: 1, padding: 1 }}>
    <text>Content</text>
  </box>
</box>
```

### Dashboard grid

```tsx
<box flexDirection="column" style={{ gap: 1 }}>
  <box flexDirection="row" style={{ gap: 1 }}>
    <box border style={{ flexGrow: 1, height: 5 }} />
    <box border style={{ flexGrow: 1, height: 5 }} />
  </box>
  <box border style={{ height: 8 }} />
</box>
```

Notes

- Use `flexDirection`, `flexGrow`, percent `width/height`, and `gap` to achieve flexible layouts.
- Wrap interactive content in bordered `box`es with titles to aid scannability.

## Interactive patterns

### Form with validation

```tsx
function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [focused, setFocused] = useState<"username" | "password">("username");
  const [status, setStatus] = useState<"idle" | "error" | "success">("idle");

  const submit = () =>
    setStatus(
      username === "admin" && password === "secret" ? "success" : "error"
    );

  useKeyboard((key) => {
    if (key.name === "tab")
      setFocused((f) => (f === "username" ? "password" : "username"));
    if (key.name === "return") submit();
  });

  return (
    <box style={{ border: true, padding: 1, gap: 1, flexDirection: "column" }}>
      <box title="Username" style={{ border: true, height: 3 }}>
        <input
          focused={focused === "username"}
          onInput={setUsername}
          onSubmit={submit}
        />
      </box>
      <box title="Password" style={{ border: true, height: 3 }}>
        <input
          focused={focused === "password"}
          onInput={setPassword}
          onSubmit={submit}
        />
      </box>
      <text
        style={{
          fg:
            status === "success"
              ? "green"
              : status === "error"
              ? "red"
              : "#999",
        }}
      >
        {status.toUpperCase()}
      </text>
    </box>
  );
}
```

### Keyboard-driven list (arrow navigation)

```tsx
function Menu({ items }: { items: string[] }) {
  const [index, setIndex] = useState(0);

  useKeyboard((key) => {
    if (key.name === "up") setIndex((i) => Math.max(0, i - 1));
    if (key.name === "down") setIndex((i) => Math.min(items.length - 1, i + 1));
  });

  return (
    <box style={{ border: true, padding: 1 }}>
      {items.map((label, i) => (
        <text key={label}>
          {i === index ? "> " : "  "}
          <strong>{label}</strong>
        </text>
      ))}
    </box>
  );
}
```

### Tabs

```tsx
function Tabs() {
  const [tab, setTab] = useState(0);
  const options = [{ name: "Home" }, { name: "Settings" }, { name: "About" }];
  return (
    <box style={{ gap: 1 }}>
      <tab-select options={options} focused onChange={(i) => setTab(i)} />
      <box border style={{ padding: 1 }}>
        {options[tab].name} content
      </box>
    </box>
  );
}
```

### Scrollable list with focus

```tsx
function ScrollList({ rows }: { rows: string[] }) {
  const [index, setIndex] = useState(0);
  useKeyboard((key) => {
    if (key.name === "up") setIndex((i) => Math.max(0, i - 1));
    if (key.name === "down") setIndex((i) => Math.min(rows.length - 1, i + 1));
  });
  return (
    <scrollbox focused style={{ height: 12 }}>
      {rows.map((r, i) => (
        <text key={i}>
          <span fg={i === index ? "#7aa2f7" : undefined}>{r}</span>
        </text>
      ))}
    </scrollbox>
  );
}
```

Notes

- `focused` controls which widget receives keyboard input; use state + `useKeyboard` to switch.
- Keep each interactive area isolated in its own `box` when practical.

## Text and styling patterns

### Rich text blocks

```tsx
<text>
  <strong>Bold</strong>, <em>Italic</em>, <u>Underline</u>
  <br />
  <span fg="red">Signals</span> and <span fg="blue">status</span>
</text>
```

### Style prop vs direct props

```tsx
<box backgroundColor="blue" padding={1}><text>A</text></box>
<box style={{ backgroundColor: "blue", padding: 1 }}><text>B</text></box>
```

Guidelines

- Use direct props for simple flags (`border`, `title`); use `style` for grouped visual concerns.
- Prefer `gap`, `padding`, and `margin` for spacing consistency.

## Animation patterns

### Progress bar tween

```tsx
function Progress() {
  const [pct, setPct] = useState(0);
  const timeline = useTimeline({ duration: 2000 });
  useEffect(() => {
    timeline.add(
      { pct },
      {
        pct: 100,
        duration: 2000,
        ease: "linear",
        onUpdate: (a) => setPct(a.targets[0].pct),
      }
    );
  }, []);
  return (
    <box style={{ backgroundColor: "#333" }}>
      <box
        style={{
          width: `${Math.round(pct)}%`,
          height: 1,
          backgroundColor: "#6a5acd",
        }}
      />
    </box>
  );
}
```

### Count-up number

```tsx
function CounterAnim() {
  const [n, setN] = useState(0);
  const timeline = useTimeline({ duration: 1000 });
  useEffect(() => {
    timeline.add(
      { n },
      {
        n: 100,
        duration: 1000,
        ease: "linear",
        onUpdate: (a) => setN(Math.round(a.targets[0].n)),
      }
    );
  }, []);
  return <text>{`Count: ${n}`}</text>;
}
```

### Pulse effect

```tsx
function PulseText({ label }: { label: string }) {
  const [w, setW] = useState(10);
  const timeline = useTimeline({ duration: 800, loop: true });
  useEffect(() => {
    timeline.add(
      { w },
      {
        w: 20,
        duration: 800,
        ease: "linear",
        onUpdate: (a) => setW(a.targets[0].w),
      }
    );
  }, []);
  return (
    <box style={{ width: Math.round(w), backgroundColor: "#4682b4" }}>
      <text>{label}</text>
    </box>
  );
}
```

Best practices

- Use `onUpdate` to lift animated values into state and re-render specific regions.
- Avoid animating dozens of independent values simultaneously.

## Data-driven lists and collections

### Mapping with stable keys

```tsx
<box>
  {users.map((u) => (
    <box key={u.id} border style={{ marginBottom: 1 }}>
      <text>{u.name}</text>
    </box>
  ))}
</box>
```

### Large collections with scrollbox

```tsx
<scrollbox focused style={{ height: 16 }}>
  {rows.map((row, i) => (
    <text key={i}>{row}</text>
  ))}
</scrollbox>
```

### Selection model

```tsx
function Selectable({ items }: { items: string[] }) {
  const [idx, setIdx] = useState(0);
  useKeyboard((k) => {
    if (k.name === "up") setIdx((i) => Math.max(0, i - 1));
    if (k.name === "down") setIdx((i) => Math.min(items.length - 1, i + 1));
  });
  return (
    <box>
      {items.map((t, i) => (
        <text key={t}>
          <span fg={i === idx ? "#7aa2f7" : undefined}>{t}</span>
        </text>
      ))}
    </box>
  );
}
```

## Recipe library

### Page shell (header / sidebar / content)

```tsx
function PageShell({ children }: { children: any }) {
  return (
    <box flexDirection="column" style={{ gap: 1 }}>
      <box border title="Header" style={{ height: 3 }} />
      <box flexDirection="row" style={{ gap: 1, flexGrow: 1 }}>
        <box border title="Sidebar" style={{ width: 24 }} />
        <box border title="Content" style={{ flexGrow: 1, padding: 1 }}>
          {children}
        </box>
      </box>
    </box>
  );
}
```

### Login form (adapted)

```tsx
// Based on README: Login Form
// Path: opentui/packages/react/README.md
```

```tsx
function LoginForm() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [focused, setFocused] = useState<"username" | "password">("username");
  const [status, setStatus] = useState("idle");
  const submit = () =>
    setStatus(
      username === "admin" && password === "secret" ? "success" : "error"
    );
  useKeyboard((key) => {
    if (key.name === "tab")
      setFocused((prev) => (prev === "username" ? "password" : "username"));
  });
  return (
    <box style={{ border: true, padding: 2, flexDirection: "column", gap: 1 }}>
      <text fg="#FFFF00">Login Form</text>
      <box title="Username" style={{ border: true, width: 40, height: 3 }}>
        <input
          placeholder="Enter username..."
          onInput={setUsername}
          onSubmit={submit}
          focused={focused === "username"}
        />
      </box>
      <box title="Password" style={{ border: true, width: 40, height: 3 }}>
        <input
          placeholder="Enter password..."
          onInput={setPassword}
          onSubmit={submit}
          focused={focused === "password"}
        />
      </box>
      <text
        style={{
          fg:
            status === "success"
              ? "green"
              : status === "error"
              ? "red"
              : "#999",
        }}
      >
        {status.toUpperCase()}
      </text>
    </box>
  );
}
```

### Keyboard-driven list navigator

```tsx
// See Interactive: Keyboard-driven list
```

### Tabs switcher

```tsx
// See Interactive: Tabs
```

### System monitor bars (adapted)

```tsx
// Based on README: System Monitor Animation
// Path: opentui/packages/react/README.md
```

```tsx
function SystemBars() {
  type Stats = { cpu: number; memory: number; network: number; disk: number };
  const [stats, setStats] = useState<Stats>({
    cpu: 0,
    memory: 0,
    network: 0,
    disk: 0,
  });
  const timeline = useTimeline({ duration: 3000 });
  useEffect(() => {
    timeline.add(
      stats,
      {
        cpu: 85,
        memory: 70,
        network: 95,
        disk: 60,
        duration: 3000,
        ease: "linear",
        onUpdate: (v) => setStats({ ...v.targets[0] }),
      },
      0
    );
  }, []);
  const rows = [
    { name: "CPU", key: "cpu", color: "#6a5acd" },
    { name: "Memory", key: "memory", color: "#4682b4" },
    { name: "Network", key: "network", color: "#20b2aa" },
    { name: "Disk", key: "disk", color: "#daa520" },
  ] as const;
  return (
    <box title="System Monitor" style={{ margin: 1, padding: 1, border: true }}>
      {rows.map((r) => (
        <box key={r.key}>
          <box flexDirection="row" justifyContent="space-between">
            <text>{r.name}</text>
            <text>{Math.round(stats[r.key as keyof Stats])}%</text>
          </box>
          <box style={{ backgroundColor: "#333" }}>
            <box
              style={{
                width: `${stats[r.key as keyof Stats]}%`,
                height: 1,
                backgroundColor: r.color,
              }}
            />
          </box>
        </box>
      ))}
    </box>
  );
}
```

### Modal/dialog overlay

```tsx
function Modal({
  open,
  onClose,
  children,
}: {
  open: boolean;
  onClose: () => void;
  children: any;
}) {
  if (!open) return null;
  return (
    <box
      style={{
        position: "absolute",
        left: 0,
        top: 0,
        width: "100%",
        height: "100%",
        backgroundColor: "#00000088",
      }}
    >
      <box
        border
        title="Dialog"
        style={{ width: 40, height: 10, margin: "auto", padding: 1 }}
      >
        {children}
      </box>
      <text>Press ESC to close</text>
    </box>
  );
}
```

### Toast / status bar

```tsx
function StatusBar({ text }: { text: string }) {
  return (
    <box style={{ height: 1, backgroundColor: "#414868", paddingLeft: 1 }}>
      <text>{text}</text>
    </box>
  );
}
```

### ASCII font switcher (adapted)

```tsx
// Based on README: ASCII Font Component
// Path: opentui/packages/react/README.md
```

```tsx
function AsciiSwitcher() {
  const text = "ASCII";
  const [font, setFont] = useState<"block" | "shade" | "slick" | "tiny">(
    "tiny"
  );
  const { width, height } = measureText({ text, font });
  return (
    <box style={{ border: true, paddingLeft: 1, paddingRight: 1 }}>
      <box style={{ height: 8, border: true, marginBottom: 1 }}>
        <select
          focused
          onChange={(_, o) => setFont(o?.value)}
          showScrollIndicator
          options={[
            { name: "Tiny", value: "tiny" },
            { name: "Block", value: "block" },
            { name: "Slick", value: "slick" },
            { name: "Shade", value: "shade" },
          ]}
          style={{ flexGrow: 1 }}
        />
      </box>
      <ascii-font style={{ width, height }} text={text} font={font} />
    </box>
  );
}
```

### Counter with timer (adapted)

```tsx
// Based on README: Counter with Timer
// Path: opentui/packages/react/README.md
```

```tsx
function CounterTimer() {
  const [count, setCount] = useState(0);
  useEffect(() => {
    const id = setInterval(() => setCount((c) => c + 1), 1000);
    return () => clearInterval(id);
  }, []);
  return (
    <box title="Counter" style={{ padding: 1 }}>
      <text fg="#00FF00">{`Count: ${count}`}</text>
    </box>
  );
}
```

## Component extension (only when necessary)

When built-ins are insufficient, extend core renderables and register them.

```tsx
// Minimal example inspired by EXTEND.md and README
class ButtonRenderable extends BoxRenderable {
  private _label = "Button";
  constructor(ctx: RenderContext, options: BoxOptions & { label?: string }) {
    super(ctx, { border: true, minHeight: 3, ...options });
    if (options.label) this._label = options.label;
  }
  protected renderSelf(buffer: OptimizedBuffer): void {
    super.renderSelf(buffer);
    const cx = this.x + Math.floor(this.width / 2 - this._label.length / 2);
    const cy = this.y + Math.floor(this.height / 2);
    buffer.drawText(this._label, cx, cy, RGBA.fromInts(255, 255, 255, 255));
  }
  set label(v: string) {
    this._label = v;
    this.requestRender();
  }
}

declare module "@opentui/react" {
  interface OpenTUIComponents {
    consoleButton: typeof ButtonRenderable;
  }
}
extend({ consoleButton: ButtonRenderable });
```

Best practices

- Always call `requestRender()` in setters that affect visuals.
- Use the closest base class (`BoxRenderable`, `TextRenderable`, etc.).
- Keep names unique and declare module augmentation for TypeScript support.

## Intent → constructs mapping (AI cheatsheet)

- Scrollable list with highlight and arrows
  - Use `scrollbox` + `useKeyboard` + selected index state + colored `<span>` for highlight.
- Two-column layout with resizable main area
  - `box` row with left `width: 24` and right `flexGrow: 1`; use percent widths if needed.
- Settings panel with titled sections
  - Use `box border title` for each section; stack with `flexDirection: "column"` and `gap`.
- Tabs with content switching
  - `tab-select` to choose; render conditional content below.
- Animated progress/status
  - `useTimeline` to tween numeric value and map to `width: "X%"` bar.
- Form submit with validation
  - `input` controls, `focused` management, derived status `<text>` with colored `fg`.
- ASCII banner heading
  - `measureText` + `<ascii-font>` with computed `width/height`.

## Quality checklist for AI

- Maintain stable `key`s for mapped lists.
- Explicitly control focus (`focused={...}`) when multiple inputs exist.
- Keep layout shallow; prefer multiple small `box`es over deep nesting.
- Use `gap`, `padding`, `margin` consistently; avoid hardcoding too many absolute sizes.
- Lift animated values into state via `onUpdate` and update the minimum region.
- For long content, wrap in `scrollbox` and avoid rendering thousands of heavy children unnecessarily.
- Prefer `style` for visual grouping; direct props for simple flags.

## Appendix: sources referenced

- `opentui/packages/react/README.md`
- `opentui/packages/react/docs/EXTEND.md`
- `opentui/packages/react/examples/*.tsx`
