# Tech Stack

> This document follows **Spec-Driven Development**: every technology decision is recorded here before implementation begins. Code must not be written for a layer until the spec for that layer is approved.

---

## Development Phases

### Phase 1 — Frontend (current)

| Concern | Technology | Notes |
|---|---|---|
| UI Framework | **Angular** | Component-based SPA |
| Language | TypeScript | Strict mode enabled |
| Data layer | Local JSON files | Flat-file DB, no backend required |
| Styling | TBD | e.g. TailwindCSS or Angular Material |
| Build tool | Angular CLI | `ng build`, `ng serve` |

**Constraints**
- 
- 
- 
- 

---

## Spec-Driven Development Rules

1. **Spec first** — No new feature may be coded until a specification exists in `frontend/specification/`.
2. **Schema changes** — Any change to a JSON data model must be documented in the relevant spec file before `homePage.json` (or equivalent) is edited.
3. **API contract** — Phase 2 endpoints must have an OpenAPI 3.x spec approved before Go implementation starts.
4. **Review** — All spec files are version-controlled and reviewed via pull request.
