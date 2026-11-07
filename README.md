# CHOP SHOP Dance

Source repository for migrating [chopshopdance.org](https://www.chopshopdance.org/) from Squarespace to a static site hosted on GitHub Pages.

## Migration source

The original Squarespace content is stored in `export/Squarespace-Wordpress-Export-11-07-2026.xml`. This WordPress XML export is retained as source material for the migration.

## Project status

The site root is `www/`. The first static migration includes 20 public pages, local images and the application PDF, archive videos, and the original navigation. GitHub Pages deployment is not configured yet.

See [migration notes](docs/migration.md) for page coverage, verification, and remaining decisions.

## Local development

Install Go 1.25 or newer, then run this command from the repository directory if Make is available:

```sh
make dev
```

With only Go installed, run the equivalent command:

```sh
go tool devd -l ./www
```

The repository pins [devd](https://github.com/cortesi/devd) in `go.mod` and records dependency checksums in `go.sum`. Go downloads and builds the tool automatically on the first run, which requires internet access and may take a little longer. No separate `devd` installation is needed.

Open the local URL printed by `devd`. Changes to files in `www/` automatically reload the browser. Press `Ctrl+C` to stop the server.

## Editing the site

Edit `www/index.html` for the homepage or `www/<page>/index.html` for an interior page. Shared styling lives in `www/assets/site.css`, and navigation and email-draft behavior live in `www/assets/site.js`. Images are in `www/assets/images/`.

Save a file while `make dev` is running to see the change in your browser. There is no production build step. Header and footer HTML is repeated on each page; update all pages when changing shared navigation or footer content.

## Next steps

1. Review visual fidelity, then make a shared design pass.
2. Choose contact and newsletter services.
3. Review dated content and external links.
4. Configure GitHub Pages to deploy only `www/` at the custom domain root.
