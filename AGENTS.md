# Project guidance

## Maintaining project notes

- Read the project notes below before starting work.
- Keep these notes current as project decisions and overall direction evolve.
- Record significant decisions and their brief rationale in the dated `PROJECT.md` log. Preserve historical entries and mark superseded decisions. Keep unresolved questions and current migration status in `docs/migration.md`.
- Distinguish agreed decisions from proposals and open questions.
- Update existing notes when decisions change, and remove obsolete guidance.
- Keep notes concise and focused on durable context rather than a log of routine tasks.
- Keep the README consistent with the project direction recorded here.

## Project notes

### Agreed direction

- Migrate `https://www.chopshopdance.org/` from Squarespace to a static site hosted on GitHub Pages.
- Use `export/Squarespace-Wordpress-Export-11-07-2026.xml` as the initial content source and retain it for reference during the migration.
- Keep editable site files in `www/` and shared HTML in `templates/`. Use `pnpm dev` to rebuild into `dist/` with Eleventy and serve with automatic browser reload. Keep build configuration in `eleventy.config.js`.
- Deploy only `dist/` to GitHub Pages with `.github/workflows/deploy.yml`, after formatting checks, regression tests, and a successful build. Restrict deployment to the default branch. Keep the custom error page in `www/404.html`, emitted as `dist/404.html`. Repository Pages settings and custom-domain DNS remain launch tasks.
- Pin development dependencies in `package.json` and retain `pnpm-lock.yaml`. Use pnpm for dependency management and scripts.
- Keep local development accessible to non-engineers with Node.js and pnpm installed. Go and Make are not required.
- Build directly editable HTML, CSS, and JavaScript with Eleventy and Nunjucks shared templates. `pnpm build` generates static deployment files in `dist/`; no browser framework or server runtime is required. Preserve source routes, including `/404.html`, and remove stale output after successful builds. The faithful migration is complete. Redesign around consistent, mobile-first layouts while preserving migrated content and routes.
- Format source JavaScript, HTML, and CSS with the pinned oxfmt dependency using `pnpm format` and `pnpm format:check`, including shared templates and unpublished drafts. Keep Nunjucks conditions inside quoted attribute values so the HTML parser can handle them.
- Link directly to local assets, including `/styles.css` and `/assets/site.js`, without cache-busting query parameters to avoid unrelated page changes during styling iterations.
- Keep all site styling in `www/styles.css`, with shared design variables and clearly organized component and responsive sections. Use descriptive site-owned classes instead of Squarespace classes, inline styles, or page-specific stylesheets.
- Preserve existing page routes, content, image descriptions, and photo credits. Keep the original export outside the deployed `dist/` directory.
- Use Tufte CSS as the basis for a coherent editorial design, adapted to system sans-serif fonts, 20px mobile body text, readable line lengths, and mobile-first layouts. Keep its MIT attribution in `www/licenses/tufte-css.txt`.
- Keep the palette neutral gray, black, and white. Text links are full black, with light text for navigation and actions on dark backgrounds.
- Do not transform or offset the header logo. Preserve its natural placement and the agreed `4px 0 6px` container padding.
- End each page’s main heading with a period.
- Keep page-heading and homepage photographs full width below their titles at every viewport size.
- Use simple semantic HTML with shared article, profile, and archive layouts. The image logo appears at every size; navigation collapses below 900px. Review mobile and desktop rendering rather than reproducing the original Squarespace styling.
- Keep migration coverage, intentional differences, and launch questions current in `docs/migration.md`. Record downloaded asset sources in `export/asset-manifest.json`.
- Use clearly labeled email fallbacks until replacement contact and newsletter services are selected. Do not imply that an email draft or signup request has been submitted.
- Include the dance-film archive as public content. Michael removed its password and authorized migration of the page and any linked CHOP SHOP pages.

- Keep tickets, master classes, and venues and directions unpublished in `drafts/` until approved for release. Only `www/` is built; preserve draft content and remove public links to draft routes. Historical master classes remain public.

### Open questions

- Which refinements should follow Michael’s review of the new editorial design?
- Which services should handle newsletter signup and contact submissions?
- Which historical pages and dated application or ticket information should be updated before launch?
