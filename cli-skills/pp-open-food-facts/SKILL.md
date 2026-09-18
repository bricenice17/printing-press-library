---
name: pp-open-food-facts
description: "Use Open Food Facts for product lookup, structured food search, retailer-aware discovery, nutrition, ingredients, allergens, additives, NOVA, Nutri-Score, and comparisons. Supports Publix, Sprouts, Costco, Aldi, and other OFF store tags via --store."
author: "Dhilip Subramanian; store-aware modifications by bricenice17"
license: "Apache-2.0"
argument-hint: "<command> [args]"
allowed-tools: "Read Bash"
metadata:
  openclaw:
    requires:
      bins:
        - open-food-facts-pp-cli
    install:
      - kind: go
        bins: [open-food-facts-pp-cli]
        module: github.com/bricenice17/printing-press-library/library/food-and-dining/open-food-facts/cmd/open-food-facts-pp-cli
---
<!-- GENERATED FILE — DO NOT EDIT.
     This file is a verbatim mirror of library/food-and-dining/open-food-facts/SKILL.md,
     regenerated post-merge by tools/generate-skills/. Hand-edits here are
     silently overwritten on the next regen. Edit the library/ source instead.
     See AGENTS.md "Generated artifacts: registry.json, cli-skills/". -->

# Open Food Facts — Store-aware Printing Press CLI

Use this skill when an agent needs Open Food Facts data for packaged-food
research, food-profile evaluation, barcode lookup, or retailer-aware product discovery.

## Install in Hermes

Install the customized binary from Mike's fork through the Printing Press installer:

```bash
npx -y @mvanhorn/printing-press-library install open-food-facts --cli-only \
  --registry-url https://raw.githubusercontent.com/bricenice17/printing-press-library/main/registry.json
```

Install this fork's focused skill into Hermes:

```bash
hermes skills install bricenice17/printing-press-library/cli-skills/pp-open-food-facts --force
```

Verify:

```bash
open-food-facts-pp-cli --version
open-food-facts-pp-cli doctor --agent
```

Direct Go fallback:

```bash
go install github.com/bricenice17/printing-press-library/library/food-and-dining/open-food-facts/cmd/open-food-facts-pp-cli@latest
```

No API key is required for read operations.

For regular use configure:

```bash
export OPEN_FOOD_FACTS_USER_AGENT="mike-food-intelligence/0.1"
export OPEN_FOOD_FACTS_CONTACT_EMAIL="YOUR_EMAIL"
```

## Store discovery

`--store` is generic and accepts an Open Food Facts store tag. Mike's preferred
stores are:

- Publix
- Sprouts
- Costco
- Aldi

Examples:

```bash
open-food-facts-pp-cli search --category "breakfast cereals" --country "united-states" --store "publix" --page-size 25 --agent
open-food-facts-pp-cli search --category "breakfast cereals" --country "united-states" --store "sprouts" --page-size 25 --agent
open-food-facts-pp-cli search --category "breakfast cereals" --country "united-states" --store "costco" --page-size 25 --agent
open-food-facts-pp-cli search --category "breakfast cereals" --country "united-states" --store "aldi" --page-size 25 --agent
```

For a request spanning several preferred stores, run one bounded search per store,
then deduplicate by barcode. Respect OFF's search rate limit.

`--purchase-place` may be used when a useful OFF purchase-place tag is known.

Never describe store tags as live inventory.

## Product inspection

```bash
open-food-facts-pp-cli product <barcode> --agent
open-food-facts-pp-cli nutrition <barcode> --agent
open-food-facts-pp-cli allergens <barcode> --agent
```

The allergen command also reports ingredient text, ingredient analysis, traces,
and additives.

## Mike Approved integration

This CLI is the **data adapter**, not the health-policy engine.

A separate Mike Approved skill should:
1. Discover candidates with this CLI.
2. Inspect ingredients/NOVA/Nutri-Score/additives.
3. Apply deterministic profile rules.
4. Reject forbidden ingredients.
5. Score/rank survivors.
6. Explain why each product passes or fails.

Open Food Facts data is community-contributed and may be incomplete or stale.
