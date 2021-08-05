export function isStdLibRecord(
  value: unknown,
): value is Record<string, unknown> {
  return typeof value === "object" && value !== null;
}
