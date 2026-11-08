# CHOP SHOP Dance

Source repository for migrating [chopshopdance.org](https://www.chopshopdance.org/) from Squarespace to a static site hosted on GitHub Pages.

## Migration source

The original Squarespace content is stored in `export/Squarespace-Wordpress-Export-11-07-2026.xml`. This WordPress XML export is retained as source material for the migration.

## Project status

Editable site sources live in `www/`, shared HTML lives in `templates/`, and generated output goes to `dist/`. The first static migration includes 20 public pages, local images and the application PDF, archive videos, and the original navigation. GitHub Pages deployment is not configured yet.

See [migration notes](docs/migration.md) for page coverage, verification, and remaining decisions.

## Local development

Install Go 1.25 or newer, then run this command from the repository directory if Make is available:

```sh
make dev
```

With only Go installed, run the equivalent command:

```sh
go run ./cmd/site dev
```

The repository pins [devd](https://github.com/cortesi/devd) in `go.mod` and records dependency checksums in `go.sum`. Go downloads and builds the tool automatically on the first run, which requires internet access and may take a little longer. No separate `devd` installation is needed.

Open the local URL printed by `devd`. Changes to `www/` or `templates/` automatically rebuild the site and reload the browser. Template errors are printed in the terminal; the last successful build remains available until the error is corrected. Press `Ctrl+C` to stop the server.

## Editing the site

Edit `www/index.html` for the homepage or `www/<page>/index.html` for an interior page. All styling lives in `www/styles.css`, and navigation and email-draft behavior live in `www/assets/site.js`. Images are in `www/assets/images/`.

Start with the variables at the top of `styles.css` to change shared colors, fonts, content width, or gutters. The following sections cover document defaults, navigation, banners, content layout, media, forms, the footer, and responsive behavior. Use the existing component classes when editing pages; keep new styles in this same file.

Content uses `layout-row`, `layout-column`, and `layout-span-1` through `layout-span-12` for column widths, with `content-block` for spacing and `text-content` for prose. Columns stack at the 640px mobile breakpoint. Components such as `hero`, `photo`, `card`, `video`, and `archive-links` share their appearance across pages.

Edit `templates/header.html` or `templates/footer.html` to update shared content once. Pages include these files with `{{template "header.html" .}}` and `{{template "footer.html" .}}`. The header uses `.Path` to highlight the current page automatically. Redirect pages remain standalone HTML.

Save a file while `make dev` is running to see the change in your browser. If an older development server is running, stop it and restart with the new command.

## Building for deployment

Run `make build` or, with only Go installed:

```sh
go run ./cmd/site build
```

The Go standard library renders the HTML templates and copies assets into `dist/`. This generated directory is ignored by Git; edit source files instead. GitHub Pages deployment will use a GitHub Actions workflow to build and upload `dist/`; that workflow is not configured yet. The build removes obsolete output files, so do not store hand-edited files there.

## Next steps

1. Review visual fidelity, then make a shared design pass.
2. Choose contact and newsletter services.
3. Review dated content and external links.
4. Configure GitHub Pages to deploy only `dist/` at the custom domain root.
