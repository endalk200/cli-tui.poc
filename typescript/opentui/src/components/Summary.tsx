export function Summary({
  filesReviewed,
  suggestions,
}: {
  filesReviewed: number;
  suggestions: number;
}) {
  return (
    <text>
      Reviewed {filesReviewed} file{filesReviewed === 1 ? "" : "s"} ·{" "}
      {suggestions} suggestion
      {suggestions === 1 ? "" : "s"}
    </text>
  );
}
