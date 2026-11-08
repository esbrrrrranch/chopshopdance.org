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
- Keep the site files in `www/` and use `make dev` to serve them locally with `devd` and automatic browser reload.
- Track Go development tools with `tool` directives in `go.mod` and retain `go.sum` so contributors use pinned dependencies without installing tools globally.
- Keep local development accessible to non-engineers with Go installed. Document a direct Go command wherever Make is used, since Make may not be available.
- Build the site with directly editable static HTML, CSS, and JavaScript, without a framework or production build step. First reproduce the current site closely, then review design improvements with Michael.
- Keep all site styling in `www/styles.css`, with shared design variables and clearly organized component and responsive sections. Use descriptive site-owned classes instead of Squarespace classes, inline styles, or page-specific stylesheets.
- Preserve existing page routes, content, image descriptions, and photo credits. Keep the original export outside the deployed `www/` directory.
- Preserve the original mobile layout at widths up to 640px, including the text masthead, collapsible navigation, stacked columns, and hero typography. Compare against the live mobile site when changing responsive styles.
- Keep migration coverage, intentional differences, and launch questions current in `docs/migration.md`. Record downloaded asset sources in `export/asset-manifest.json`.
- Use clearly labeled email fallbacks until replacement contact and newsletter services are selected. Do not imply that an email draft or signup request has been submitted.
- Include the dance-film archive as public content. Michael removed its password and authorized migration of the page and any linked CHOP SHOP pages.

### Open questions

- Which typography and layout refinements should follow the initial faithful rebuild?
- Which services should handle newsletter signup and contact submissions?
- Which historical pages and dated application or ticket information should be updated before launch?
