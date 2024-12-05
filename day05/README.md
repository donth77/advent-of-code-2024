```markdown
# --- Day 5: Print Queue ---

Satisfied with their search on Ceres, the squadron of scholars suggests subsequently scanning the stationery stacks of sub-basement 17.

The North Pole printing department is busier than ever this close to Christmas, and while The Historians continue their search of this historically significant facility, an Elf operating a very familiar printer beckons you over.

The Elf must recognize you, because they waste no time explaining that the new sleigh launch safety manual updates won't print correctly. Failure to update the safety manuals would be dire indeed, so you offer your services.

Safety protocols clearly indicate that new pages for the safety manuals must be printed in a very specific order. The notation `X|Y` means that if both page number `X` and page number `Y` are to be produced as part of an update, page number `X` must be printed at some point before page number `Y`.

The Elf has for you both the page ordering rules and the pages to produce in each update (your puzzle input), but can't figure out whether each update has the pages in the right order.

## Part 1: Determine Correctly-Ordered Updates

For example:

```
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13
```

```
75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
```

The first section specifies the page ordering rules, and the second section lists the updates to check. Only updates that obey the rules are valid. For instance:

1. Update `75,47,61,53,29` is valid because all rules are followed.
2. Update `75,97,47,61,53` is invalid because it violates `97|75`.
3. Update `61,13,29` is invalid because it violates `29|13`.

Correctly-ordered updates:
- `75,47,61,53,29`
- `97,61,53,29,13`
- `75,29,13`

**Middle page numbers** of these updates:
- 61 (middle of `75,47,61,53,29`)
- 53 (middle of `97,61,53,29,13`)
- 29 (middle of `75,29,13`)

**Sum of middle page numbers**: `61 + 53 + 29 = 143`.

---

## Part 2: Fix Incorrect Updates

For each of the incorrectly-ordered updates, reorder the pages to satisfy the rules.

Example:

1. `75,97,47,61,53` becomes `97,75,47,61,53` (to satisfy `97|75`).
2. `61,13,29` becomes `61,29,13` (to satisfy `29|13`).
3. `97,13,75,29,47` becomes `97,75,47,29,13` (to satisfy all relevant rules).

Reordered updates:
- `97,75,47,61,53`
- `61,29,13`
- `97,75,47,29,13`

**Middle page numbers** of these updates:
- 47 (middle of `97,75,47,61,53`)
- 29 (middle of `61,29,13`)
- 47 (middle of `97,75,47,29,13`)

**Sum of middle page numbers**: `47 + 29 + 47 = 123`.

```