# Static site migration

## Direction

The faithful Squarespace migration is complete. Michael has requested a coherent, mobile-first redesign using Tufte CSS as a base with system sans-serif typography. The site consists of directly editable HTML in `www/`, shared CSS, and a small navigation and contact script. A small Go generator renders shared header and footer templates from `templates/` and copies assets into `dist/`. `make build` creates deployment output; `make dev` uses the pinned `modd` tool and `modd.conf` to rebuild changed sources, while `devd` serves and reloads the output.

The first pass was prepared on November 7, 2026. The WordPress XML export supplied the page inventory and image descriptions. The live pages supplied the full text, layout, banners, quotes, buttons, and video embeds that the export omits. All 140 live text blocks were checked against the migrated pages.

## Page inventory

| Group | Paths | Status |
| --- | --- | --- |
| Main navigation | `/`, `/history/`, `/staff/`, `/sponsors-press/`, `/support/`, `/festival-artists/`, `/experience-dance-program/`, `/apply/`, `/contact/` | Migrated. |
| Additional public pages | `/scholarships/`, `/tickets/`, `/venues-directions/`, `/accommodation/`, `/master-classes/`, `/dance-journalism-workshop/` | Migrated with existing content, including historical dates. |
| Festival archives | `/2020-master-classes/`, `/2020-edp/`, `/2020-performing-artists/`, `/performing-artists-2019/` | Migrated. |
| Dance-film archive | `/dance-films/` | Migrated after Michael removed its password. Includes seven film descriptions and credits. Its content links only to external artist websites and social profiles; no additional CHOP SHOP pages were discovered. |
| Old links | `/home/`, `/festival/`, `/edp/` | Static redirects to the homepage, festival artists, and experience dance program. |
| Missing pages | `/404.html` | Custom not-found page for GitHub Pages. |

## Assets and integrations

- The 45 migrated images are stored in `www/assets/images/`. Image descriptions and existing photo credits are retained. The CDN supplied optimized images, usually WebP, up to 2,000 pixels wide.
- The shared header uses Michael’s revised `www/assets/images/chop-shop-logo.png`, supplied on November 11, 2026. Its original proportions are preserved, with a compact image masthead on mobile and a 288px display width on desktop.
- The 2027 application PDF is stored in `www/s/`. The asset source URLs are recorded in `export/asset-manifest.json`.
- Fifteen YouTube embeds remain externally hosted and use HTTPS with the privacy-enhanced embed domain. Playback and availability depend on YouTube; every historical video has not been watched or independently verified.
- The application button retains the existing Google Forms destination. Donations link to the Inlet Arts Zeffy donation form. No submissions or transactions were made during migration.
- The contact form prepares an email draft rather than submitting to Squarespace. It requires a configured email app and clearly states that the visitor must send the draft.
- Newsletter signup uses a labeled email link temporarily. The original signup markup reports that no storage destination is configured. A mailing-list provider must be chosen before restoring online signup.
- Typography uses system sans-serif fonts throughout, as requested for the redesign. No remote fonts are loaded.
- Squarespace analytics, cookie UI, platform scripts, and unused stylesheets are not included.

## Design and launch review

1. Review the redesigned homepage, article, staff, and archive layouts, including image crops and mobile navigation.
2. Choose a newsletter provider and decide whether contact should use a hosted form or keep email.
3. The dance-film archive is public. Its live source contains descriptions and credits but no film players; preserve that scope unless film access is explicitly added later.
4. Review dated content. The application deadline in the copied text is November 5, 2026; older tickets and classes pages also retain their original dates and external links.
5. Configure GitHub Pages to deploy only `dist/` at the domain root. Root-relative links assume the custom domain or an account-level Pages site, not a repository subpath. The XML export and migration notes must remain outside the deployed directory.
6. Confirm external archive links and video availability before launch. Deployment and DNS changes have not been made.

## Verification

The first pass checked every local image, stylesheet, script, PDF, and page reference, compared all 140 live text blocks, and checked JavaScript syntax. Browser review covered the homepage at desktop, tablet, and mobile widths, the staff layout, and dropdown navigation. The existing development server was reused. These checks do not constitute a full accessibility audit or a complete external-link audit.

The dance-film follow-up verified all nine additional text blocks and the hero image, retained all 17 external artist and social links, and removed the temporary contact notice and its `noindex` directive.

## Editorial redesign

The global `www/styles.css` adapts Tufte CSS’s typographic hierarchy, reading measure, figures, and margin column. The MIT license is included in `www/licenses/tufte-css.txt`. The adaptation replaces serif fonts with system sans-serif, uses a mobile-first grid, and adds the site’s navigation, forms, and homepage components.

Body text is 20px on mobile and 22px at desktop widths. Titles and photography are separate, so text remains readable without an image overlay. Article content has a limited reading width; longer pages have an expandable table of contents. Staff and artist biographies share one profile layout, and the archive pairs each year’s artists with its trailer. The imported column and spacer markup has been removed in favor of semantic headings, paragraphs, lists, figures, and sections.

A full-width black header blends with the logo on mobile and desktop. The mobile header is at least 88px tall, with a larger logo and space reserved for the menu button. On desktop, the logo targets 40% of the available header width. The logo preserves its aspect ratio, and is capped at 120px tall. Its container has 4px of top padding and sits flush with the header’s bottom edge, without a bottom border. The image shifts down 1px and clips inside its link so the white lettering meets the white page without a resampling gap. Navigation and dropdowns use light text, with a collapsible menu below 900px. The mobile toggle uses three thick lines, and expanded navigation appears above the logo so it remains aligned with the header’s bottom edge. Tickets, master classes, and venue information are now also available in the festival menu. Interior pages have shared breadcrumbs: current festival pages return to the festival artists overview, historical pages return to festival history, and scholarships also link through master classes. Footer social links remain below the attribution text, with white graphics on black circles and accessible labels.

The homepage has a new editorial headline and composition. Existing body copy, links, images, image descriptions, video embeds, and credits are preserved. Banner photo credits sit beside their photographs. Dates and application status are unchanged and still need the content review noted above.

The design replaces the original Squarespace mobile layout and desktop proportions intentionally. Browser checks cover all 21 content/error pages at 320px and 1280px, with additional checks of seven representative page types at 390px and 768px. Navigation, keyboard dismissal, table-of-contents links, local assets, and preserved content are checked separately. External forms and donations are not submitted during testing.

Dense articles use spaced sections with dark gray headings and subtle dividing rules. Course titles, dates, descriptions, and registration links are separate elements. Venue names, highlighted addresses, and smaller travel details have distinct levels. Existing text, dates, and destinations remain unchanged.

The palette uses neutral grays, black, and white. Text links are full black; navigation and filled actions retain light text on dark backgrounds.

All action buttons use a black background with white text, including class registration. Text expands within the button on hover and keyboard focus, with reduced-motion preferences respected.
