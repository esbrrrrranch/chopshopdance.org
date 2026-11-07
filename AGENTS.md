# Project guidance

## Maintaining project notes

- Read the project notes below before starting work.
- Keep these notes current as project decisions and overall direction evolve.
- Log each significant decision in `PROJECT.md` (see below).
- Record significant decisions, their brief rationale, and unresolved questions that affect future work.
- Distinguish agreed decisions from proposals and open questions.
- Update existing notes when decisions change, and remove obsolete guidance.
- Keep notes concise and focused on durable context rather than a log of routine tasks.
- Keep the README consistent with the project direction recorded here.

## Decision log

- Maintain `PROJECT.md` at the repository root as a dated, append-only decision log to support long-term thinking about this work.
- Add an entry whenever a significant decision is made, changed, or reversed. Each entry has a date, the decision, brief rationale and alternatives considered, and status (agreed, proposed, superseded).
- Do not delete superseded entries; mark them superseded and link to the entry that replaces them. Track unresolved questions in a separate section and move them into the log once decided.
- Keep AGENTS.md notes as the current-state summary; `PROJECT.md` holds the history and reasoning.

## Project notes

### Agreed direction

- Migrate `https://www.chopshopdance.org/` to a static site hosted on GitHub Pages.
- The original WordPress XML export was used as the initial content source and has been removed from the repository.
- Keep editable site files in `www/` and shared HTML in `templates/`. Use `make dev` (`go tool modd`) to rebuild them into `dist/` on source changes and serve that output with `devd` and automatic browser reload. Keep watcher configuration in `modd.conf` rather than custom Go code.
- Track Go development tools with `tool` directives in `go.mod` and retain `go.sum` so contributors use pinned dependencies without installing tools globally.
- Keep local development accessible to non-engineers with Go installed. Document a direct Go command wherever Make is used, since Make may not be available.
- Build the site with directly editable HTML, CSS, and JavaScript, using Go’s standard `html/template` package to share the header and footer. `make build` generates static deployment files in `dist/`; no frontend framework or additional runtime is required. First reproduce the current site closely, then review design improvements with Michael.
- Keep all site styling in `www/styles.css`, with shared design variables and clearly organized component and responsive sections. Use descriptive site-owned classes instead of generic platform classes, inline styles, or page-specific stylesheets.
- Preserve existing page routes, content, image descriptions, and photo credits.
- Preserve the original mobile layout at widths up to 640px, including the text masthead, collapsible navigation, stacked columns, and hero typography. Compare against the live mobile site when changing responsive styles.
- Keep migration coverage, intentional differences, and launch questions current in the migration status section of `PROJECT.md`.
- Use clearly labeled email fallbacks until replacement contact and newsletter services are selected. Do not imply that an email draft or signup request has been submitted.
- Include the dance-film archive as public content. Michael removed its password and authorized migration of the page and any linked CHOP SHOP pages.

### Open questions

- Which typography and layout refinements should follow the initial faithful rebuild?
- Which services should handle newsletter signup and contact submissions?
- Which historical pages and dated application or ticket information should be updated before launch?
