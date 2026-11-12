import { readdir, rm } from "node:fs/promises";
import path from "node:path";

const listFiles = async (directory) => {
  const entries = await readdir(directory, { withFileTypes: true });
  const files = await Promise.all(
    entries.map(async (entry) => {
      const filename = path.join(directory, entry.name);
      return entry.isDirectory() ? listFiles(filename) : filename;
    }),
  );
  return files.flat();
};

export default (config) => {
  config.addPassthroughCopy("www", {
    filter: (filename) => !filename.endsWith(".html"),
  });
  config.addWatchTarget("templates/");
  config.setNunjucksEnvironmentOptions({
    autoescape: true,
    throwOnUndefined: true,
  });
  config.addGlobalData(
    "permalink",
    () =>
      ({ page }) =>
        path.relative("www", page.inputPath).split(path.sep).join("/"),
  );
  config.on("eleventy.after", async ({ dir }) => {
    const sources = new Set(
      (await listFiles(dir.input))
        .filter(
          (filename) =>
            ![".DS_Store", "Thumbs.db"].includes(path.basename(filename)),
        )
        .map((filename) => path.relative(dir.input, filename)),
    );
    // Remove stale output only after a successful build.
    await Promise.all(
      (await listFiles(dir.output))
        .filter((filename) => !sources.has(path.relative(dir.output, filename)))
        .map((filename) => rm(filename)),
    );
  });
  return {
    dir: { includes: "../templates", input: "www", output: "dist" },
    htmlTemplateEngine: "njk",
    templateFormats: ["html"],
  };
};
