# Grammar — Staged Implementation Log

The parser is rewritten incrementally. Each stage appends one section below
with the COMPLETE grammar implemented at that stage. Sections are never
edited afterwards; the most recent stage section is the current state, and
the final stage section will be the full grammar.

## Stage 1 — Expressions only (2026-10-09)

Scope: a single expression, no statements, no identifiers. The parser entry
returns one `ast.Expr`. Nodes implemented: `Binary`, `Literal`, `Unary`,
`Grouping`.

```
expression     → equality ;
equality       → comparison ( ( "!=" | "==" ) comparison )* ;
comparison     → term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
term           → factor ( ( "-" | "+" ) factor )* ;
factor         → unary ( ( "/" | "*" ) unary )* ;
unary          → ( "!" | "-" ) unary
               | primary ;
primary        → NUMBER | STRING | "true" | "false" | "nil"
               | "(" expression ")" ;
```

### Design decisions

- Error handling: productions return `(ast.Expr, error)`. Errors are values
  created once at the point of failure with full context (token, line,
  message via `perror.ParseError`) and propagated unchanged — no
  `fmt.Errorf`/`%w` wrapping through the production chain, no panic/recover.
  Printing is not part of parsing: the CLI entry prints, tests assert the
  returned error.
- The scanner keeps its current immediate-print (`log.Logger.Error`) error
  style for now; it is considered suboptimal and will be migrated to the
  error-value style later.
