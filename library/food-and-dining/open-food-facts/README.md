# Open Food Facts CLI — bricenice17 store-aware build

This is a read-only Open Food Facts Printing Press CLI based on the upstream
`mvanhorn/printing-press-library` Open Food Facts print, with generic retailer
and purchase-place search support added.

It supports product lookup, structured search, nutrition summaries, ingredients,
allergens, additives, comparisons, NOVA, Nutri-Score, and agent-ready JSON.

## Setup

No API key is required for Open Food Facts read operations.

For regular use:

```bash
export OPEN_FOOD_FACTS_USER_AGENT="mike-food-intelligence/0.1"
export OPEN_FOOD_FACTS_CONTACT_EMAIL="YOUR_EMAIL"
```

## Install in Hermes

Install the customized binary from Mike's Printing Press catalog:

```bash
npx -y @mvanhorn/printing-press-library install open-food-facts --cli-only \
  --registry-url https://raw.githubusercontent.com/bricenice17/printing-press-library/main/registry.json
```

Then install the matching focused skill from this fork:

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

## Store-aware searches

The `--store` option is generic. It is **not hardcoded to Publix**.

Examples:

```bash
open-food-facts-pp-cli search --category "breakfast cereals" --country "united-states" --store "publix" --page-size 25 --agent

open-food-facts-pp-cli search --category "breakfast cereals" --country "united-states" --store "sprouts" --page-size 25 --agent

open-food-facts-pp-cli search --category "breakfast cereals" --country "united-states" --store "costco" --page-size 25 --agent

open-food-facts-pp-cli search --category "breakfast cereals" --country "united-states" --store "aldi" --page-size 25 --agent
```

You can also filter on a community-contributed purchase-place tag:

```bash
open-food-facts-pp-cli search --category "breakfast cereals" --purchase-place "tampa-florida" --agent
```

Returned product summaries include `stores` and `purchase_places` when OFF has them.

## Other commands

```bash
open-food-facts-pp-cli product 3017620422003 --agent
open-food-facts-pp-cli nutrition 3017620422003 --agent
open-food-facts-pp-cli allergens 3017620422003 --agent
open-food-facts-pp-cli compare 3017620422003 5449000000996 --agent
open-food-facts-pp-cli category "breakfast cereals" --page-size 5 --agent
open-food-facts-pp-cli sources --agent
open-food-facts-pp-cli doctor --agent
```

## Important caveat

Open Food Facts store and purchase-place tags are community-contributed.
A result tagged `publix`, `sprouts`, `costco`, or `aldi` means OFF associates
that product with that retailer. It does **not** prove that a specific nearby
location has the product in stock tonight.

Live inventory can be added later through a retailer/Instacart layer.
