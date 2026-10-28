# pr-autolabel-demo

Sandbox for testing automatic PR labels before proposing them to
[kubescape/kubescape](https://github.com/kubescape/kubescape).

Labels come from two sources:

- **PR title** (conventional commits): `feat` → `feature`, `fix` → `bug`,
  `docs` → `documentation`, `ci` → `ci/cd`, `test` → `Tests`,
  `refactor` → `Refactoring`, `chore`/`build` → `maintenance`.
  A scope containing `security` or `cve` adds `security`.
- **Changed files** ([.github/labeler.yml](.github/labeler.yml)).
