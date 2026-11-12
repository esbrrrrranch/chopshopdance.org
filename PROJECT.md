# Project decision log

Decision log for the CHOP SHOP Dance site. Preserve historical entries and append new decisions to the log. Current guidance lives in `AGENTS.md`; current migration status and launch questions live in `docs/migration.md`. The migration snapshot below is retained from the remote repository for reference.

Entry format: date, decision, rationale and alternatives, status (agreed, proposed, superseded).

## Log

### 2026-11-07: Migrate to a static site on GitHub Pages
- **Status:** superseded for tooling by the November 11, 2026 decision below; GitHub Pages hosting remains agreed.
- **Decision:** Move chopshopdance.org to a static site hosted on GitHub Pages, built from editable HTML, CSS, and JavaScript with Go's `html/template` sharing the header and footer.
- **Rationale:** No framework or runtime to maintain; non-engineers with Go installed can build and preview locally.

### 2026-11-07: Reproduce the current site before redesigning
- **Status:** completed; the November 11, 2026 reconciliation retains the subsequent editorial redesign.
- **Decision:** Rebuild faithfully first, including the original mobile layout up to 640px, then review design improvements with Michael.
- **Rationale:** Separates migration risk from design changes.

### 2026-11-08: Public dance-film archive
- **Status:** agreed
- **Decision:** Include the dance-film archive and linked CHOP SHOP pages as public content.
- **Rationale:** Michael removed its password and authorized migration.

### 2026-11-08: Remove the original export from the repository
- **Status:** superseded by the November 11, 2026 reconciliation below.
- **Decision:** Drop the `export/` directory. The asset source manifest was later dropped too.
- **Rationale:** Migration content is complete; the manifest keeps asset provenance.

### 2026-11-11: Use Eleventy, pnpm, and GitHub Actions

- **Status:** agreed.
- **Decision:** Replace Go templates and development tools with Eleventy and Nunjucks, manage dependencies with pnpm, and format JavaScript, HTML, and CSS with oxfmt. Deploy `dist/` through GitHub Actions after checks, including the custom `404.html`.
- **Rationale:** Consolidate static-site development and formatting in Node tooling while retaining editable HTML and existing routes.

### 2026-11-11: Reconcile local and remote repository histories

- **Status:** agreed.
- **Decision:** Merge both histories, retain the current local site and editorial redesign, and preserve this remote decision log. Retain the original XML export and asset manifest for migration reference. Keep current migration status in `docs/migration.md`.
- **Rationale:** Michael approved retaining the reviewed local site while preserving the remote history and decision log, without a force push.

## Open questions

See `docs/migration.md` for current design, service, content, and launch questions.

## Historical migration snapshot

The following snapshot predates the editorial redesign and tooling migration. It is preserved as historical context; `docs/migration.md` supersedes it for current status.

### Direction

Reproduce the current site closely before making design changes. The site consists of directly editable HTML in `www/`, shared CSS, and a small navigation and contact script. A small Go generator renders shared header and footer templates from `templates/` and copies assets into `dist/`. `make build` creates deployment output; `make dev` uses the pinned `modd` tool and `modd.conf` to rebuild changed sources, while `devd` serves and reloads the output.

The first pass was prepared on November 7, 2026. A WordPress XML export of the original site (since removed from the repository) supplied the page inventory and image descriptions. The live pages supplied the full text, layout, banners, quotes, buttons, and video embeds that the export omits. All 140 live text blocks were checked against the migrated pages.

### Page inventory

| Group | Paths | Status |
| --- | --- | --- |
| Main navigation | `/`, `/history/`, `/staff/`, `/sponsors-press/`, `/support/`, `/festival-artists/`, `/experience-dance-program/`, `/apply/`, `/contact/` | Migrated. |
| Additional public pages | `/scholarships/`, `/tickets/`, `/venues-directions/`, `/accommodation/`, `/master-classes/`, `/dance-journalism-workshop/` | Migrated with existing content, including historical dates. |
| Festival archives | `/2020-master-classes/`, `/2020-edp/`, `/2020-performing-artists/`, `/performing-artists-2019/` | Migrated. |
| Dance-film archive | `/dance-films/` | Migrated after Michael removed its password. Includes seven film descriptions and credits. Its content links only to external artist websites and social profiles; no additional CHOP SHOP pages were discovered. |
| Old links | `/home/`, `/festival/`, `/edp/` | Static redirects to the homepage, festival artists, and experience dance program. |
| Missing pages | `/404.html` | Custom not-found page for GitHub Pages. |

### Assets and integrations

- All 45 site images, including the logo and hero photographs, are stored in `www/assets/images/`. Image descriptions and existing photo credits are retained. The CDN supplied optimized images, usually WebP, up to 2,000 pixels wide.
- The 2027 application PDF is stored in `www/s/`.
- Fifteen YouTube embeds remain externally hosted and use HTTPS with the privacy-enhanced embed domain. Playback and availability depend on YouTube; every historical video has not been watched or independently verified.
- The application button retains the existing Google Forms destination. Donations retain the existing Fractured Atlas destination. No submissions or transactions were made during migration.
- The contact form prepares an email draft rather than submitting to a service. It requires a configured email app and clearly states that the visitor must send the draft.
- Newsletter signup uses a labeled email link temporarily. The original signup markup reports that no storage destination is configured. A mailing-list provider must be chosen before restoring online signup.
- Adobe-hosted Proxima Nova and the original quote font have been replaced with system sans-serif and Georgia fallbacks. Exact typography requires licensed font files or an approved replacement.
- Analytics, cookie UI, platform scripts, and unused stylesheets are not included.

### Design and launch review

1. Review the homepage and one interior page together to settle typography, hero height, image crops, and mobile navigation before applying design changes throughout the site.
2. Choose a newsletter provider and decide whether contact should use a hosted form or keep email.
3. The dance-film archive is public. Its live source contains descriptions and credits but no film players; preserve that scope unless film access is explicitly added later.
4. Review dated content. The application deadline in the copied text is November 5, 2026; older tickets and classes pages also retain their original dates and external links.
5. Configure GitHub Pages to deploy only `dist/` at the domain root. Root-relative links assume the custom domain or an account-level Pages site, not a repository subpath. The migration notes must remain outside the deployed directory.
6. Confirm external archive links and video availability before launch. Deployment and DNS changes have not been made.

### Verification

The first pass checked every local image, stylesheet, script, PDF, and page reference, compared all 140 live text blocks, and checked JavaScript syntax. Browser review covered the homepage at desktop, tablet, and mobile widths, the staff layout, and dropdown navigation. The existing development server was reused. These checks do not constitute a full accessibility audit or a complete external-link audit.

The dance-film follow-up verified all nine additional text blocks and the hero image, retained all 17 external artist and social links, and removed the temporary contact notice and its `noindex` directive.

### Mobile fidelity follow-up

Comparison with the live site revealed that the first pass missed its separate mobile header and changed hero typography and image proportions. The shared header now switches at 640px to a text masthead and hamburger button with expandable navigation groups. The mobile menu supports Escape, outside clicks, and keyboard focus. Navigation remains available without JavaScript.

The follow-up restores the original 62px mobile hero title, natural card image proportions, hidden spacers, and mobile block padding. The system-font fallback still produces different line wrapping from Proxima Nova. Shared asset URLs include a version query so browsers refresh the revised CSS and JavaScript.

Browser checks at 320px, 390px, 640px, 641px, and 1280px confirmed the header breakpoint, menu behavior, and absence of horizontal page overflow. Mobile navigation to an interior page and stacked staff content were also checked.

### Shared stylesheet

All pages now use `www/styles.css`. It contains shared design variables, component styles, and responsive rules. Legacy platform classes and editor metadata have been removed from the HTML, and inline presentation has moved into reusable classes. Navigation selectors are scoped to the shared navigation component so future content elements do not inherit menu styling.

Dropdown links share the main navigation typography: 18px, medium weight, uppercase, and .1em letter spacing on desktop and mobile. Their width accommodates the longer labels.

Footer social links appear below the attribution text and use a shared local SVG sprite for Instagram, Facebook, and email, with white graphics on black circles matching the original treatment. Each link retains an accessible label and a 44px touch target without an external icon dependency.

Staff profiles use an explicit 40% portrait and 60% biography grid within their existing content section. This corrects the imported nested-column proportions. At a 1280px viewport, the portrait and text widths match the original at 377px and 583px; both stack to 342px at a 390px viewport. Font substitution still affects text wrapping.

The cleanup preserves page text, links, image sources, and the restored mobile behavior. Before-and-after measurements of the homepage, staff page, and history page matched at 390px and 1280px widths.
