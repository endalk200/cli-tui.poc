type Finding = {
  location: string;
  finding: string;
  suggestion: string;
};

type FileFindings = {
  filename: string;
  findings: Finding[];
};

export function ReviewFindingsList({ items }: { items: FileFindings[] }) {
  return (
    <>
      {items.map((file) => (
        <>
          <text key={`h-${file.filename}`}>
            <strong>{file.filename}</strong>
          </text>
          {file.findings.map((f, i) => (
            <text key={`${file.filename}-${i}`}>
              <span fg="#7aa2f7">{f.location}</span>: {f.finding} →{" "}
              {f.suggestion}
            </text>
          ))}
        </>
      ))}
    </>
  );
}
