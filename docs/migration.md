# Static site migration

## Direction

Reproduce the current Squarespace site closely before making design changes. The site consists of directly editable HTML in `www/`, shared CSS, and a small navigation and contact script. There is no site generator or production build step.

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

- All 45 site images, including the logo and hero photographs, are stored in `www/assets/images/`. Image descriptions and existing photo credits are retained. The CDN supplied optimized images, usually WebP, up to 2,000 pixels wide.
- The 2027 application PDF is stored in `www/s/`. The asset source URLs are recorded in `export/asset-manifest.json`.
- Fifteen YouTube embeds remain externally hosted and use HTTPS with the privacy-enhanced embed domain. Playback and availability depend on YouTube; every historical video has not been watched or independently verified.
- The application button retains the existing Google Forms destination. Donations retain the existing Fractured Atlas destination. No submissions or transactions were made during migration.
- The contact form prepares an email draft rather than submitting to Squarespace. It requires a configured email app and clearly states that the visitor must send the draft.
- Newsletter signup uses a labeled email link temporarily. The original signup markup reports that no storage destination is configured. A mailing-list provider must be chosen before restoring online signup.
- Adobe-hosted Proxima Nova and the original quote font have been replaced with system sans-serif and Georgia fallbacks. Exact typography requires licensed font files or an approved replacement.
- Squarespace analytics, cookie UI, platform scripts, and unused stylesheets are not included.

## Design and launch review

1. Review the homepage and one interior page together to settle typography, hero height, image crops, and mobile navigation before applying design changes throughout the site.
2. Choose a newsletter provider and decide whether contact should use a hosted form or keep email.
3. The dance-film archive is public. Its live source contains descriptions and credits but no film players; preserve that scope unless film access is explicitly added later.
4. Review dated content. The application deadline in the copied text is November 5, 2026; older tickets and classes pages also retain their original dates and external links.
5. Configure GitHub Pages to deploy only `www/` at the domain root. Root-relative links assume the custom domain or an account-level Pages site, not a repository subpath. The XML export and migration notes must remain outside the deployed directory.
6. Confirm external archive links and video availability before launch. Deployment and DNS changes have not been made.

## Verification

The first pass checked every local image, stylesheet, script, PDF, and page reference, compared all 140 live text blocks, and checked JavaScript syntax. Browser review covered the homepage at desktop, tablet, and mobile widths, the staff layout, and dropdown navigation. The existing development server was reused. These checks do not constitute a full accessibility audit or a complete external-link audit.

The dance-film follow-up verified all nine additional text blocks and the hero image, retained all 17 external artist and social links, and removed the temporary contact notice and its `noindex` directive.
