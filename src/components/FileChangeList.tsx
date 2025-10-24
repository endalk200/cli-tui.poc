type FileChange = {
  filename: string;
  added: number;
  changed: number;
  removed: number;
};

export function FileChangeList({ items }: { items: FileChange[] }) {
  const rows = items.map((it) => (
    <text key={it.filename}>
      <strong>{it.filename}</strong> · <span fg="#00ff00">+{it.added}</span>{" "}
      <span fg="#ffff00">~{it.changed}</span>{" "}
      <span fg="#ff5555">-{it.removed}</span>
    </text>
  ));

  return <>{rows}</>;
}
