# Project guidance

## Maintaining project notes

- Read the project notes below before starting work.
- Keep these notes current as project decisions and overall direction evolve.
- Record significant decisions, their brief rationale, and unresolved questions that affect future work.
- Distinguish agreed decisions from proposals and open questions.
- Update existing notes when decisions change, and remove obsolete guidance.
- Keep notes concise and focused on durable context rather than a log of routine tasks.
- Keep the README consistent with the project direction recorded here.

## Project notes

### Agreed direction

- Migrate `https://www.chopshopdance.org/` from Squarespace to a static site hosted on GitHub Pages.
- Use `export/Squarespace-Wordpress-Export-11-07-2026.xml` as the initial content source and retain it for reference during the migration.
- Keep editable site files in `www/` and shared HTML in `templates/`. Use `make dev` (`go tool modd`) to rebuild them into `dist/` on source changes and serve that output with `devd` and automatic browser reload. Keep watcher configuration in `modd.conf` rather than custom Go code.
- Track Go development tools with `tool` directives in `go.mod` and retain `go.sum` so contributors use pinned dependencies without installing tools globally.
- Keep local development accessible to non-engineers with Go installed. Document a direct Go command wherever Make is used, since Make may not be available.
- Build the site with directly editable HTML, CSS, and JavaScript, using Go’s standard `html/template` package to share the header and footer. `make build` generates static deployment files in `dist/`; no frontend framework or additional runtime is required. The faithful migration is complete. Redesign around consistent, mobile-first layouts while preserving the migrated content and routes.
- Keep all site styling in `www/styles.css`, with shared design variables and clearly organized component and responsive sections. Use descriptive site-owned classes instead of Squarespace classes, inline styles, or page-specific stylesheets.
- Preserve existing page routes, content, image descriptions, and photo credits. Keep the original export outside the deployed `dist/` directory.
- Use Tufte CSS as the basis for a coherent editorial design, adapted to system sans-serif fonts, 20px mobile body text, readable line lengths, and mobile-first layouts. Keep its MIT attribution in `www/licenses/tufte-css.txt`.
- Keep the palette neutral gray, black, and white. Text links are full black, with light text for navigation and actions on dark backgrounds.
- End each page’s main heading with a period.
- Use simple semantic HTML with shared article, profile, and archive layouts. The image logo appears at every size; navigation collapses below 900px. Review mobile and desktop rendering rather than reproducing the original Squarespace styling.
- Keep migration coverage, intentional differences, and launch questions current in `docs/migration.md`. Record downloaded asset sources in `export/asset-manifest.json`.
- Use clearly labeled email fallbacks until replacement contact and newsletter services are selected. Do not imply that an email draft or signup request has been submitted.
- Include the dance-film archive as public content. Michael removed its password and authorized migration of the page and any linked CHOP SHOP pages.

### Open questions

- Which refinements should follow Michael’s review of the new editorial design?
- Which services should handle newsletter signup and contact submissions?
- Which historical pages and dated application or ticket information should be updated before launch?
