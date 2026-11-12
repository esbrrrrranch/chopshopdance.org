# CHOP SHOP Dance

Source repository for migrating [chopshopdance.org](https://www.chopshopdance.org/) from Squarespace to a static site hosted on GitHub Pages.

## Migration source

The original Squarespace content is stored in `export/Squarespace-Wordpress-Export-11-07-2026.xml`. This WordPress XML export is retained as source material for the migration.

## Project status

Editable site sources live in `www/`, shared HTML lives in `templates/`, and generated output goes to `dist/`. The first static migration preserves 20 content pages, local images and the application PDF, archive videos, and shared navigation. The current design adapts Tufte CSS to a mobile-first layout with system sans-serif typography. A GitHub Pages deployment workflow is included; repository Pages settings and DNS still need to be configured.

See [migration notes](docs/migration.md) for page coverage, verification, and remaining decisions.

## Local development

Install Node.js 22.18 or newer and pnpm, then run:

```sh
pnpm install
pnpm dev
```

Open the local URL printed by Eleventy. It watches `www/` and `templates/`, rebuilds changed sources, and reloads the browser. Template errors appear in the terminal; correct the error to resume rebuilding. Press `Ctrl+C` to stop the server.

Dependencies are pinned in `package.json` and `pnpm-lock.yaml`. Go and Make are no longer required.

## Editing the site

Edit `www/index.html` for the homepage or `www/<page>/index.html` for an interior page. All styling lives in `www/styles.css`, and navigation and email-draft behavior live in `www/assets/site.js`. Images are in `www/assets/images/`.

Start with the variables at the top of `styles.css` to change shared colors, fonts, content width, or gutters. The following sections cover document defaults, navigation, page headings, reading layouts, media, forms, the footer, and responsive behavior. Use the existing component classes when editing pages; keep new styles in this same file.

Use normal headings, paragraphs, lists, and figures inside an `article.prose`. A `reading-layout` pairs the article with a `page-aside` for navigation or festival information. Long pages use native `details` for an expandable table of contents, placed before the article on mobile. Staff and artist pages use `section.profile` with a photograph and `profile-body`; festival history uses `section.archive-entry` with `archive-year` and `archive-body`.

The layout starts with one column, 20px body text, and a compact menu. At 900px it expands to adjacent content, a margin column where useful, and 22px body text. Keep body copy within `--measure`, and use the shared `photo`, `video`, `button`, and `newsletter` components. Use `section.content-section` for major article sections, with semantic subheadings for courses and venues. `metadata` styles dates, `detail-label` and `detail` distinguish supporting information, and `venue-address` groups addresses and map links. Keep these hierarchy styles shared rather than applying one-off font sizes. Avoid spacer elements, inline styles, and nested presentation wrappers.

The stylesheet adapts [Tufte CSS](https://github.com/edwardtufte/tufte-css) rather than loading a second CSS file or its serif fonts. Its license is included at `www/licenses/tufte-css.txt`.

Edit `templates/header.html` or `templates/footer.html` to update shared content once. Pages use Nunjucks includes such as `{% include "header.html" %}`. The header uses `page.url` to highlight the current page automatically. Interior pages set a `breadcrumbs` array with `label` and `url` fields, then include `breadcrumbs.html`. Include Home first; an empty URL marks the current page. Redirect pages remain standalone HTML.

Save a file while `pnpm dev` is running to see the change in your browser. Stop any previous Go development server before switching.

## Formatting and checks

Oxfmt formats JavaScript, HTML, and CSS, including shared templates and unpublished drafts:

```sh
pnpm format
pnpm format:check
```

Generated output, dependencies, and the original export are excluded. Keep Nunjucks conditions inside quoted HTML attribute values so the HTML formatter can parse them.

Run the build regression checks with `pnpm test`. They verify output paths, asset copying, shared templates, navigation, draft exclusion, stale-file cleanup, and recovery after template errors.

## Unpublished pages

Tickets, master classes, and venues and directions are saved in `drafts/`, outside the published `www/` tree. They are excluded from both local previews and `dist/`. The historical 2020 master classes page remains public.

To restore a page, move its directory from `drafts/` back to `www/`, review its dated content, and restore its link in `templates/header.html`. When restoring master classes, also restore the Master Classes parent breadcrumb on the scholarships page. Rebuild with `pnpm build`.

## Building for deployment

Run:

```sh
pnpm build
```

Eleventy renders HTML through Nunjucks and copies assets into `dist/`. Configuration lives in `eleventy.config.js`; source paths are preserved, including `/404.html`. After successful builds, obsolete output files are removed. This generated directory is ignored by Git; edit source files instead.

`.github/workflows/deploy.yml` deploys pushes to the repository’s default branch and supports manual runs from that branch. It installs locked dependencies, checks formatting, runs regression tests, builds the site, verifies `dist/404.html`, and uploads only `dist/` to GitHub Pages. Runs on other branches skip deployment.

In the GitHub repository, select **Settings → Pages → Build and deployment → Source → GitHub Actions**. Configure the custom domain and DNS when ready to launch. The site uses root-relative URLs and expects the custom domain root or an account-level Pages site, rather than a repository subpath.

The custom error page lives in `www/404.html`. Eleventy emits it as `dist/404.html`, where GitHub Pages serves it for missing URLs.

## Next steps

1. Review the new editorial design on phones and desktop.
2. Choose contact and newsletter services.
3. Review dated content and external links.
4. Enable GitHub Actions as the Pages source and configure the custom domain and DNS.

The shared palette is gray, black, and white. Text links are black, with light text on dark navigation and action backgrounds.
