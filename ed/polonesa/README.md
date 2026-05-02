# Transforme da notação padrão para polonesa reversa

## Descrição

- Faça o caso simples sem parênteses
- Utilize os operadores `+`, `-`, `*`, `/` e `^`
- Precedência 1(menor): `+` e `-`
- Precedência 2: `*` e `/`
- Precedência 3(maior): `^`

## Drafts

<!-- links .cache/drafts -->
- cpp
  - [fn.hpp](../../.tko/cache/6421490caf670842e74ba5681e807eeaa53f9028/base/polonesa/.cache/drafts/cpp/fn.hpp)
  - [main.cpp](../../.tko/cache/6421490caf670842e74ba5681e807eeaa53f9028/base/polonesa/.cache/drafts/cpp/main.cpp)
<!-- links -->

## Testes

```bash

>>>>>>>>
3 + 4 * 2
========
3 4 2 * +
<<<<<<<<


>>>>>>>>
3 + 4 * 2 - 1
========
3 4 2 * + 1 -
<<<<<<<<

>>>>>>>>
3 + 4 * 2 - 1 / 2
========
3 4 2 * + 1 2 / -
<<<<<<<<

>>>>>>>>
1 + 2 * 3 ^ 4
========
1 2 3 4 ^ * +
<<<<<<<<

>>>>>>>>
1 + 2 * 3 ^ 4 - 5
========
1 2 3 4 ^ * + 5 -
<<<<<<<<

```
