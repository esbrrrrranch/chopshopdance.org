# CHOP SHOP Dance

Source repository for migrating [chopshopdance.org](https://www.chopshopdance.org/) from Squarespace to a static site hosted on GitHub Pages.

## Migration source

The original Squarespace content is stored in `export/Squarespace-Wordpress-Export-11-07-2026.xml`. This WordPress XML export is retained as source material for the migration.

## Project status

Editable site sources live in `www/`, shared HTML lives in `templates/`, and generated output goes to `dist/`. The first static migration preserves 20 content pages, local images and the application PDF, archive videos, and shared navigation. The current design adapts Tufte CSS to a mobile-first layout with system sans-serif typography. GitHub Pages deployment is not configured yet.

See [migration notes](docs/migration.md) for page coverage, verification, and remaining decisions.

## Local development

Install Go 1.25 or newer, then run this command from the repository directory if Make is available:

```sh
make dev
```

With only Go installed, run the equivalent command:

```sh
go tool modd
```

The repository pins [modd](https://github.com/cortesi/modd) and [devd](https://github.com/cortesi/devd) in `go.mod` and records dependency checksums in `go.sum`. Go downloads and builds the tools automatically on the first run, which requires internet access and may take a little longer. No separate tool installation is needed.

Open the local URL printed by `devd`. `modd.conf` watches `www/`, `templates/`, the Go generator, and module files. Changes automatically rebuild the site; `devd` watches `dist/` and reloads the browser. Template errors are printed in the terminal; the last successful build remains available until the error is corrected. Press `Ctrl+C` to stop the server.

## Editing the site

Edit `www/index.html` for the homepage or `www/<page>/index.html` for an interior page. All styling lives in `www/styles.css`, and navigation and email-draft behavior live in `www/assets/site.js`. Images are in `www/assets/images/`.

Start with the variables at the top of `styles.css` to change shared colors, fonts, content width, or gutters. The following sections cover document defaults, navigation, page headings, reading layouts, media, forms, the footer, and responsive behavior. Use the existing component classes when editing pages; keep new styles in this same file.

Use normal headings, paragraphs, lists, and figures inside an `article.prose`. A `reading-layout` pairs the article with a `page-aside` for navigation or festival information. Long pages use native `details` for an expandable table of contents, placed before the article on mobile. Staff and artist pages use `section.profile` with a photograph and `profile-body`; festival history uses `section.archive-entry` with `archive-year` and `archive-body`.

The layout starts with one column, 20px body text, and a compact menu. At 900px it expands to adjacent content, a margin column where useful, and 22px body text. Keep body copy within `--measure`, and use the shared `photo`, `video`, `button`, and `newsletter` components. Use `section.content-section` for major article sections, with semantic subheadings for courses and venues. `metadata` styles dates, `detail-label` and `detail` distinguish supporting information, and `venue-address` groups addresses and map links. Keep these hierarchy styles shared rather than applying one-off font sizes. Avoid spacer elements, inline styles, and nested presentation wrappers.

The stylesheet adapts [Tufte CSS](https://github.com/edwardtufte/tufte-css) rather than loading a second CSS file or its serif fonts. Its license is included at `www/licenses/tufte-css.txt`.

Edit `templates/header.html` or `templates/footer.html` to update shared content once. Pages include these files with `{{template "header.html" .}}` and `{{template "footer.html" .}}`. The header uses `.Path` to highlight the current page automatically. Interior pages call `breadcrumbs.html` with label/URL pairs, for example `{{template "breadcrumbs.html" (breadcrumbs "Festival" "/festival-artists/" "Master Classes" "")}}`. Home is included automatically; an empty URL marks the current page. Redirect pages remain standalone HTML.

Save a file while `make dev` is running to see the change in your browser. If an older development server is running, stop it and restart with the new command.

## Formatting HTML

To format all source HTML with pnpm installed, run:

```sh
pnpm dlx oxfmt@0.78.0 --write 'www/**/*.html' 'templates/**/*.html' 'drafts/**/*.html'
```

Use `--check` instead of `--write` to check formatting. Keep Go template conditions inside quoted attribute values, using Go raw strings for path comparisons, so oxfmt can parse the HTML. Formatting is optional for local development; serving and building still require only Go.

## Unpublished pages

Tickets, master classes, and venues and directions are saved in `drafts/`, outside the published `www/` tree. They are excluded from both local previews and `dist/`. The historical 2020 master classes page remains public.

To restore a page, move its directory from `drafts/` back to `www/`, review its dated content, and restore its link in `templates/header.html`. When restoring master classes, also restore the Master Classes parent breadcrumb on the scholarships page. Rebuild with `make build` or `go run ./cmd/site build`.

## Building for deployment

Run `make build` or, with only Go installed:

```sh
go run ./cmd/site build
```

The Go standard library renders the HTML templates and copies assets into `dist/`. This generated directory is ignored by Git; edit source files instead. GitHub Pages deployment will use a GitHub Actions workflow to build and upload `dist/`; that workflow is not configured yet. The build removes obsolete output files, so do not store hand-edited files there.

## Next steps

1. Review the new editorial design on phones and desktop.
2. Choose contact and newsletter services.
3. Review dated content and external links.
4. Configure GitHub Pages to deploy only `dist/` at the custom domain root.

The shared palette is gray, black, and white. Text links are black, with light text on dark navigation and action backgrounds.
