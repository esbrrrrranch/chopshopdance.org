import assert from "node:assert/strict";
import { execFile } from "node:child_process";
import {
  cp,
  mkdir,
  mkdtemp,
  readFile,
  readdir,
  rm,
  writeFile,
} from "node:fs/promises";
import os from "node:os";
import path from "node:path";
import test from "node:test";
import { promisify } from "node:util";

const execute = promisify(execFile);
const root = path.resolve(import.meta.dirname, "..");
const command = path.join(root, "node_modules/@11ty/eleventy/cmd.cjs");

const files = async (directory) => {
  const entries = await readdir(directory, {
    recursive: true,
    withFileTypes: true,
  });
  return entries
    .filter(
      (entry) =>
        entry.isFile() && ![".DS_Store", "Thumbs.db"].includes(entry.name),
    )
    .map((entry) =>
      path.relative(directory, path.join(entry.parentPath, entry.name)),
    )
    .sort();
};

test("build preserves routes, renders shared content, and excludes drafts", async (t) => {
  const fixture = await mkdtemp(path.join(os.tmpdir(), "chopshop-build-"));
  t.after(() => rm(fixture, { force: true, recursive: true }));
  await Promise.all(
    ["www", "templates", "drafts"].map((directory) =>
      cp(path.join(root, directory), path.join(fixture, directory), {
        recursive: true,
      }),
    ),
  );
  await cp(
    path.join(root, "eleventy.config.js"),
    path.join(fixture, "eleventy.config.mjs"),
  );
  const build = () =>
    execute(process.execPath, [command, "--config=eleventy.config.mjs"], {
      cwd: fixture,
    });
  await mkdir(path.join(fixture, "dist"));
  await writeFile(path.join(fixture, "dist/stale.html"), "obsolete page");
  await writeFile(path.join(fixture, "dist/stale.css"), "obsolete asset");
  await build();
  const input = path.join(fixture, "www");
  const output = path.join(fixture, "dist");
  assert.deepEqual(await files(output), await files(input));
  for (const filename of await files(input)) {
    const rendered = await readFile(path.join(output, filename));
    if (!filename.endsWith(".html")) {
      assert.deepEqual(rendered, await readFile(path.join(input, filename)));
      continue;
    }
    const html = rendered.toString();
    assert.doesNotMatch(html, /\{%|\{\{/);
    const source = await readFile(path.join(input, filename), "utf8");
    if (source.includes('include "header.html"')) {
      assert.equal(
        (html.match(/<header class="site-header">/g) || []).length,
        1,
      );
      assert.equal(
        (html.match(/<footer class="site-footer">/g) || []).length,
        1,
      );
    }
  }
  const staff = await readFile(path.join(output, "staff/index.html"), "utf8");
  assert.match(staff, /aria-current="page"\s+href="\/staff\/"/);
  assert.match(staff, /<span\s+aria-current="page">Staff<\/span>/);
  const missing = await readFile(path.join(output, "404.html"), "utf8");
  assert.doesNotMatch(missing, /<a\b[^>]*aria-current="page"/);
  const header = path.join(fixture, "templates/header.html");
  const original = await readFile(header, "utf8");
  await writeFile(header, "{% invalid %}");
  await assert.rejects(build());
  await writeFile(
    header,
    original.replace("Skip to content", "Skip to main content"),
  );
  await rm(path.join(input, "staff/index.html"));
  await rm(path.join(input, "styles.css"));
  await build();
  assert.deepEqual(await files(output), await files(input));
  assert.match(
    await readFile(path.join(output, "index.html"), "utf8"),
    /Skip to main content/,
  );
});
