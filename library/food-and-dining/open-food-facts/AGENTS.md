# Open Food Facts Print Notes

- Keep every command read-only.
- Use documented JSON API endpoints, not HTML scraping.
- Prefer API v3 for product reads and API v2 for structured search.
- Store and purchase-place tags are discovery metadata, not live local inventory.
- `--store` is generic and may be used for Publix, Sprouts, Costco, Aldi, or any other valid Open Food Facts store tag.
- Keep requests bounded and preserve Open Food Facts data-quality caveats.
- Configure `OPEN_FOOD_FACTS_USER_AGENT` and `OPEN_FOOD_FACTS_CONTACT_EMAIL` for regular use.
- Run `go test ./...` before shipping changes.
